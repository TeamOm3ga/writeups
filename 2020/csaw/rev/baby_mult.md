# rev / baby_mult

## Question

> Welcome to reversing! Prove your worth and get the flag from this neat little program!

### Provided Files

- `program.txt`

## Solution

The provided `program.txt` is a list of numbers:

```text
85, 72, 137, 229, 72, 131, 236, 24, 72, 199, 69, 248, 79, 0, 0, 0, 72, 184, 21, 79, 231, 75, 1, 0, 0, 0, 72, 137, 69, 240, 72, 199, 69, 232, 4, 0, 0, 0, 72, 199, 69, 224, 3, 0, 0, 0, 72, 199, 69, 216, 19, 0, 0, 0, 72, 199, 69, 208, 21, 1, 0, 0, 72, 184, 97, 91, 100, 75, 207, 119, 0, 0, 72, 137, 69, 200, 72, 199, 69, 192, 2, 0, 0, 0, 72, 199, 69, 184, 17, 0, 0, 0, 72, 199, 69, 176, 193, 33, 0, 0, 72, 199, 69, 168, 233, 101, 34, 24, 72, 199, 69, 160, 51, 8, 0, 0, 72, 199, 69, 152, 171, 10, 0, 0, 72, 199, 69, 144, 173, 170, 141, 0, 72, 139, 69, 248, 72, 15, 175, 69, 240, 72, 137, 69, 136, 72, 139, 69, 232, 72, 15, 175, 69, 224, 72, 15, 175, 69, 216, 72, 15, 175, 69, 208, 72, 15, 175, 69, 200, 72, 137, 69, 128, 72, 139, 69, 192, 72, 15, 175, 69, 184, 72, 15, 175, 69, 176, 72, 15, 175, 69, 168, 72, 137, 133, 120, 255, 255, 255, 72, 139, 69, 160, 72, 15, 175, 69, 152, 72, 15, 175, 69, 144, 72, 137, 133, 112, 255, 255, 255, 184, 0, 0, 0, 0, 201
```

Since this is a "program", we can try interpreting these numbers as shellcode and disassembling them.
This approach is confirmed when we see `push rbp` and `mov rbp,rsp` at the top and `leave` at the end.

```asm
0:  55                      push   rbp
1:  48 89 e5                mov    rbp,rsp
4:  48 83 ec 18             sub    rsp,0x18
8:  48 c7 45 f8 4f 00 00    mov    QWORD PTR [rbp-0x8],0x4f
f:  00
10: 48 b8 15 4f e7 4b 01    movabs rax,0x14be74f15
17: 00 00 00
1a: 48 89 45 f0             mov    QWORD PTR [rbp-0x10],rax
1e: 48 c7 45 e8 04 00 00    mov    QWORD PTR [rbp-0x18],0x4
25: 00
26: 48 c7 45 e0 03 00 00    mov    QWORD PTR [rbp-0x20],0x3
2d: 00
2e: 48 c7 45 d8 13 00 00    mov    QWORD PTR [rbp-0x28],0x13
35: 00
36: 48 c7 45 d0 15 01 00    mov    QWORD PTR [rbp-0x30],0x115
3d: 00
3e: 48 b8 61 5b 64 4b cf    movabs rax,0x77cf4b645b61
45: 77 00 00
48: 48 89 45 c8             mov    QWORD PTR [rbp-0x38],rax
4c: 48 c7 45 c0 02 00 00    mov    QWORD PTR [rbp-0x40],0x2
53: 00
54: 48 c7 45 b8 11 00 00    mov    QWORD PTR [rbp-0x48],0x11
5b: 00
5c: 48 c7 45 b0 c1 21 00    mov    QWORD PTR [rbp-0x50],0x21c1
63: 00
64: 48 c7 45 a8 e9 65 22    mov    QWORD PTR [rbp-0x58],0x182265e9
6b: 18
6c: 48 c7 45 a0 33 08 00    mov    QWORD PTR [rbp-0x60],0x833
73: 00
74: 48 c7 45 98 ab 0a 00    mov    QWORD PTR [rbp-0x68],0xaab
7b: 00
7c: 48 c7 45 90 ad aa 8d    mov    QWORD PTR [rbp-0x70],0x8daaad
83: 00
84: 48 8b 45 f8             mov    rax,QWORD PTR [rbp-0x8]
88: 48 0f af 45 f0          imul   rax,QWORD PTR [rbp-0x10]
8d: 48 89 45 88             mov    QWORD PTR [rbp-0x78],rax
91: 48 8b 45 e8             mov    rax,QWORD PTR [rbp-0x18]
95: 48 0f af 45 e0          imul   rax,QWORD PTR [rbp-0x20]
9a: 48 0f af 45 d8          imul   rax,QWORD PTR [rbp-0x28]
9f: 48 0f af 45 d0          imul   rax,QWORD PTR [rbp-0x30]
a4: 48 0f af 45 c8          imul   rax,QWORD PTR [rbp-0x38]
a9: 48 89 45 80             mov    QWORD PTR [rbp-0x80],rax
ad: 48 8b 45 c0             mov    rax,QWORD PTR [rbp-0x40]
b1: 48 0f af 45 b8          imul   rax,QWORD PTR [rbp-0x48]
b6: 48 0f af 45 b0          imul   rax,QWORD PTR [rbp-0x50]
bb: 48 0f af 45 a8          imul   rax,QWORD PTR [rbp-0x58]
c0: 48 89 85 78 ff ff ff    mov    QWORD PTR [rbp-0x88],rax
c7: 48 8b 45 a0             mov    rax,QWORD PTR [rbp-0x60]
cb: 48 0f af 45 98          imul   rax,QWORD PTR [rbp-0x68]
d0: 48 0f af 45 90          imul   rax,QWORD PTR [rbp-0x70]
d5: 48 89 85 70 ff ff ff    mov    QWORD PTR [rbp-0x90],rax
dc: b8 00 00 00 00          mov    eax,0x0
e1: c9                      leave
```

Walking through this program, we see that the results of four sets of multiplications (`imul`, also challenge name) are stored into memory.
As hex, the four numbers concatenated are `666c61677b73757033725f76346c31645f7072306772346d7d`.
Decoded as ASCII, we get our flag.

### Flag

`flag{sup3r_v4l1d_pr0gr4m}`
