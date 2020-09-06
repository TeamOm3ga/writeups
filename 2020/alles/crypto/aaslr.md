# crypto / Actual ASLR 1

## Question

> Prove that you can win 0xf times against the house.
> Then go to Vegas.

### Provided Files

- [`aaslr`](./aaslr)
- `Dockerfile`
- `ynetd`
- `flag1`: Dummy flag
- `flag2`: Dummy flag

## Solution

Playing around with the program, all we need to do is predict a random number 16 times.

```text
$ ncat --ssl 7b000000cf533960c57d134b.challenges.broker2.allesctf.net 1337
Welcome To The Actual-ASLR (AASLR) Demo
  1. throw dice
  2. create entry
  3. read entry
  4. take a guess

Select menu Item:
1
[>] Threw dice: 2
Welcome To The Actual-ASLR (AASLR) Demo
  1. throw dice
  2. create entry
  3. read entry
  4. take a guess

Select menu Item:
4
(1/16) guess next dice roll:
2
(2/16) guess next dice roll:
3
```

Analyzing `aaslr` in Ghidra, the function `raninit` starts up the random number generator.
`raninit` is called by `init_heap`.

```c
void init_heap(void)
{
  time_t t;

  HEAP = mmap((void *)0x0,0x10000,3,0x22,-1,0);
  t = time((time_t *)0x0);
  raninit(t);
  return;
}
```

`init_heap` initializes the random number generator with `time`, which is only second-precise.
If we start two connections within a second, they should generate the same numbers.
This can be done with the following Python script:

```py
#!/usr/bin/env python3
import subprocess

a = subprocess.Popen(["ncat", "--ssl", "7b000000cf533960c57d134b.challenges.broker2.allesctf.net", "1337"], stdin=subprocess.PIPE, stdout=subprocess.PIPE)
b = subprocess.Popen(["ncat", "--ssl", "7b000000cf533960c57d134b.challenges.broker2.allesctf.net", "1337"], stdin=subprocess.PIPE, stdout=subprocess.PIPE)

def readmenu(proc):
  for i in range(7):
    proc.stdout.readline()

def throw_dice(proc):
  proc.stdin.write(b"1\n")
  proc.stdin.flush()
  line = proc.stdout.readline()
  return int(line.split(b" ")[3])

readmenu(b)
b.stdin.write(b"4\n")
b.stdin.flush()
for i in range(16):
  readmenu(a)
  dice = throw_dice(a)
  b.stdout.readline()
  b.stdin.write(f"{dice}\n".encode())
  b.stdin.flush()

while True:
  print(b.stdout.readline())
```

Running this gets our flag.

```
$ python aaslr.py
b'ALLES{ILLEGAL_CARD_COUNTING!_BANNED}\n'
b'Welcome To The Actual-ASLR (AASLR) Demo\n'
b'  1. throw dice\n'
b'  2. create entry\n'
b'  3. read entry\n'
b'  4. take a guess\n'
b'\n'
b'Select menu Item:\n'
b'(1/16) guess next dice roll:\n'
```

### Flag

`ALLES{ILLEGAL_CARD_COUNTING!_BANNED}`
