# rev / ezbreezy

## Question

> This binary has nothing to hide!

### Provided Files

- [`app`](./app)

## Solution

After disassembling the `app` binary, it appears that it just opens a file `not_even_real.txt` and does nothing with it.
From Ghidra (with renamed variables):

```c
void read_file(void)
{
  int nextChar;
  FILE *__stream;

  __stream = fopen("not_even_real.txt","r");
  if (__stream != (FILE *)0x0) {
    while( true ) {
      nextChar = getc(__stream);
      if (nextChar == -1) break;
      putchar((uint)(nextChar != -1)); // will always be null bytes...
    }
    fclose(__stream);
  }
  return;
}
```

Upon further analysis of the binary, we notice a weirdly named section `.aj1ishudgqis`:

```text
$ objdump -h app
...
 24 .aj1ishudgqis 00000304  00000000008001a0  00000000008001a0  000031a0  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
```

Disassembling this reveals a repetitive group of assembly instructions forming what appears to be a massive nested conditional.

```text
                FUN_009001a0                                    XREF[2]:     001001d8(*),
                                                                            _elfSectionHeaders::00000650(*)  
009001a0 55              PUSH       RBP
009001a1 48 89 e5        MOV        RBP,RSP
009001a4 c7 45 fc        MOV        dword ptr [RBP + local_c],0x31436
    36 14 03 00
009001ab 81 7d fc        CMP        dword ptr [RBP + local_c],0x3f5a
    5a 3f 00 00
009001b2 75 0a           JNZ        LAB_009001be
009001b4 b8 01 00        MOV        EAX,0x1
    00 00
009001b9 e9 e4 02        JMP        LAB_009004a2
    00 00
                LAB_009001be                                    XREF[1]:     009001b2(j)  
009001be c6 45 fc 8e     MOV        byte ptr [RBP + local_c],0x8e
009001c2 c7 45 fc        MOV        dword ptr [RBP + local_c],0x93ca3
    a3 3c 09 00
009001c9 81 7d fc        CMP        dword ptr [RBP + local_c],0x1f40ec
    ec 40 1f 00
009001d0 75 0a           JNZ        LAB_009001dc
009001d2 b8 01 00        MOV        EAX,0x1
    00 00
009001d7 e9 c6 02        JMP        LAB_009004a2
    00 00
                LAB_009001dc                                    XREF[1]:     009001d0(j)  
009001dc c6 45 fd 94     MOV        byte ptr [RBP + local_c+0x1],0x94
009001e0 c7 45 fc        MOV        dword ptr [RBP + local_c],0x137ec9
    c9 7e 13 00
009001e7 81 7d fc        CMP        dword ptr [RBP + local_c],0x3fc47
    47 fc 03 00
009001ee 75 0a           JNZ        LAB_009001fa
009001f0 b8 01 00        MOV        EAX,0x1
    00 00
009001f5 e9 a8 02        JMP        LAB_009004a2
    00 00

...
```

Since the blocks all have the same structure, they could represent individual characters in the flag: a very strange sort of steganography.

Every block begins with a MOV instruction.
These each move a single byte into memory:
`8e 94 89 8f a3 9d 87 90 5c 9e 5b 87 9a 5b 8b 58 9e 5b 9a 5b 8c 87 95 5b a5`.
We noticed that these bytes are all in a relatively small range which feels somewhat ASCII-like.

The first byte, based on the flag format, must be `f` (hex `66`).
`8e` is `28` (decimal 40) more than `66`, so we subtract `28` from every byte.
This gives `66 6c 61 67 7b 75 5f 68 34 76 33 5f 72 33 63 30 76 33 72 33 64 5f 6d 33 7d`, which decodes to a valid flag.

### Flag

`flag{u_h4v3_r3c0v3r3d_m3}`
