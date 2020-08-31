# reverse / XO

> Points: 475

## Question

> Sometimes the simpler the better !  
> nc xo.fword.wtf 5554  
> flag format : FwordCTF{}  
> Author: Semah BA

## Solution

We are given an executable, `task`. Let's see what happens when we run it:

```
$ ./task
input : 
ABC
3
input : 
ABCDE
5
input : 
NNNNN
0
```

It seems to be returning the length of our input, but not always. We need to disassemble it to find out what's going on.

The first challenge is finding the `main` function, your disassembler might not be able to locate it because the executable has had all of its symbols stripped.

Radare2 is able to locate `main` just fine:

```
$ r2 task
Warning: Cannot initialize dynamic strings
[0x00400a60]> s main
[0x00400bde]>
```

But in ghidra, we need to locate it ourselves. Luckily, the `entry` function will always be visible and we can find `main` from there:

```c
void entry(undefined8 param_1,undefined8 param_2,undefined8 param_3)

{
  undefined8 in_stack_00000000;

  FUN_00400fc0(FUN_00400bde,in_stack_00000000,&stack0x00000008,FUN_00401a20,FUN_00401ac0,param_3);
  do {
                    /* WARNING: Do nothing block with infinite loop */
  } while( true );
}
```

Here, `FUN_00400bde` is the address of our `main` function, rename it and we're good to start reverse engineering.

Continuing in ghidra, let's look at the decompiled code:

```c
void main(void)

{
  int iVar1;
  long lVar2;
  long lVar3;
  long lVar4;
  ulong uVar5;
  ulong uVar6;
  undefined8 uVar7;
  int local_40;
  
  FUN_00400b7d();
  lVar2 = FUN_004201f0(0x32);
  lVar3 = FUN_004201f0(0x32);
  lVar4 = FUN_00410bd0("flag.txt",&DAT_004ac8e8);
  if (lVar4 == 0) {
    FUN_0040fa60("Error while opening the file. Contact an admin!\n");
    FUN_0040ec20(1);
  }
  FUN_0040fd20(lVar4,&DAT_004ac929,lVar2);
  do {
    lVar4 = FUN_004201f0(0x32);
    FUN_00410cf0("input : ");
    FUN_0040fb40(&DAT_004ac929);
    uVar5 = thunk_FUN_004004ee(lVar2);
    uVar6 = thunk_FUN_004004ee(lVar3);
    if (uVar6 < uVar5) {
      iVar1 = thunk_FUN_004004ee(lVar3);
    }
    else {
      iVar1 = thunk_FUN_004004ee(lVar2);
    }
    local_40 = 0;
    while (local_40 < iVar1) {
      *(byte *)(lVar4 + local_40) = *(byte *)(lVar2 + local_40) ^ *(byte *)(lVar3 + local_40);
      local_40 = local_40 + 1;
    }
    uVar7 = thunk_FUN_004004ee(lVar4);
    FUN_0040f840(&DAT_004ac935,uVar7);
    FUN_00420ab0(lVar4);
  } while( true );
}
```

Remember, the binary is stripped, so we don't have any function names, not even for the statically linked libc. There's tools to automatically find these, but let's do it by hand instead. For most functions we can find the name just by looking at the context, for example:

```c
// we can see &DAT_004ac935 points to "%ld", so this is probably printf
FUN_0040f840(&DAT_004ac935,uVar7);
```

After renaming all the functions we can start to get an idea of what's happening.

```c
    while (index < length) {
      *(byte *)(tmp + index) = *(byte *)(flag + index) ^ *(byte *)(input + index);
      index = index + 1;
    }
    result = strlen(tmp);
```

The xor will return 0 if at that character our input matches the flag, so the printed `strlen` is actually the length of our string up to when the first character matches the flag. Now we can find the flag by trying all characters in every position, until we get a length value that indicates our input matches the flag:

```python
#!/usr/bin/env python3
import socket

sock = socket.create_connection(("xo.fword.wtf", 5554))
flag = b""
buffer = b""

def readline():
        global buffer
        while b"\n" not in buffer:
                buffer += sock.recv(1024)
        buffer = buffer.split(b"\n")
        result = buffer[0]
        buffer = b"\n".join(buffer[1:])
        return result

while True:
        for i in range(33, 127):
                readline()
                ch = bytes([i])
                sock.send(b"!" * len(flag) + ch + b"\n")
                length = readline()
                if int(length) == len(flag):
                        flag += ch
                        print(flag)
                        break
        else:
                break
```

### Flag

`NuL1_Byt35?15_IT_the_END?Why_i_c4nT_h4ndl3_That!`
