# crypto / modus_operandi

## Question

> Can't play CSAW without your favorite block cipher!
>
> `nc crypto.chal.csaw.io 5001`

## Solution

Connecting to the server provides a game of sorts where, given a known plaintext, you have to determine if the block cipher is ECB or CBC:

```text
$ nc crypto.chal.csaw.io 5001
Hello! For each plaintext you enter, find out if the block cipher used is ECB or CBC. Enter "ECB" or "CBC" to get the flag!
Enter plaintext:
AAAAAAA
Ciphertext is:  e7836ec8b4abe96fa9229df0aba08be6
ECB or CBC?
ECB
Enter plaintext:
AAAAAAA
Ciphertext is:  2cfcdac6629cad4b970570822e824603
ECB or CBC?
ECB
```

We can distinguish ECB from CBC by looking for cyclic patterns.
Given a long enough constant plaintext, the ciphertext in ECB will repeat itself.
Checking for this can be done in Python fairly simply:

```py
#!/usr/bin/env python
import socket

def brute():
    count = 0
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect(('crypto.chal.csaw.io', 5001))
        s.recv(2048)

        while True:
            s.send(b'A'*256+b'\n')

            inp = s.recv(2048)
            if inp == b'':
                raise ConnectionResetError
            else:
                inp = inp.split()[2]
            if inp[:128] == inp[128:256]:
                s.send(b'ECB\n')
            else:
                s.send(b'CBC\n')

            out = s.recv(1024).decode()
            if out == '':
                    raise ConnectionResetError
            elif 'Enter' not in out:
                print(out)
                return out
            else:
                print(f'[*] Solved Problem {count + 1}')
                count += 1

def main():
    try:
        print(brute())
    except ConnectionResetError:
        print('[!] Server reset the connection or sent no response')
        main()
    except Exception as e:
        print(e)

if __name__ == '__main__':
    main()
```

This code always seems to stop working after problem 175.
After trying a few times, we realized that it was probably deterministic: problem 175 was the end on purpose.
We also noticed the sequence of ECB and CBC seemed to be constant:

```
ECB CBC CBC ECB ECB CBC CBC
ECB ECB CBC CBC ECB CBC CBC
ECB ECB ECB CBC CBC ECB ECB
ECB ECB CBC ECB CBC CBC ECB
ECB CBC CBC CBC ECB CBC CBC
CBC CBC ECB CBC CBC ECB CBC
ECB ECB ECB CBC ECB CBC ECB
CBC ECB ECB ECB ECB CBC CBC
ECB CBC ECB ECB ECB ECB CBC
ECB ECB CBC ECB CBC CBC CBC
CBC CBC ECB CBC CBC CBC ECB
ECB CBC ECB ECB CBC CBC ECB
ECB CBC ECB CBC ECB CBC ECB
ECB ECB ECB ECB ECB ECB CBC
CBC ECB CBC CBC ECB ECB ECB
CBC CBC ECB CBC CBC ECB ECB
ECB CBC CBC CBC CBC ECB ECB
CBC ECB CBC ECB CBC CBC CBC
CBC CBC ECB CBC CBC CBC ECB
ECB CBC CBC ECB CBC ECB CBC
ECB CBC ECB CBC ECB CBC CBC
ECB ECB ECB CBC CBC ECB CBC
CBC ECB CBC ECB CBC CBC ECB
ECB CBC ECB ECB CBC ECB ECB
ECB CBC CBC CBC CBC CBC ECB
```

Interpreting this as binary ASCII with ECB as 0 and CBC as 1, we found the flag.

### Flag

`flag{ECB_re@lly_sUck$}`
