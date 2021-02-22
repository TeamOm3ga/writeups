# crypto / bythebook

> Points: 200

## Question

> I implemented RSA just as in my textbook!
> A cryptography text wouldn't let me down, would it..?

### Provided Files

- [`flag.enc`](./bythebook/flag.enc)
- [`key.pub`](./bythebook/key.pub)

## Solution

We're looking for another RSA implementation vulnerability.
Let's take a look at `n` and `e` from the public key:
```py
from Crypto.PublicKey import RSA
with open('key.pub', 'r') as f:
    pub_key = RSA.import_key(f.read())
    n = pub_key.n
    e = pub_key.e
```

We notice here that `e` is 3.
That's really small.
If we search for attacks, we find another
[StackOverflow post](https://crypto.stackexchange.com/questions/18301/textbook-rsa-with-exponent-e-3)
which explains that we can recover the plaintext by just taking the cube root.

Since Python can't handle that nicely, we go into Mathematica to do it,
then decode to ASCII:
```py
m = 980104003558357359783466286544512034491956481776073659699425644849051028010530706886282182800765
print(bytearray.fromhex(hex(m)[2:]).decode())
```

### Flag

`uwctf{textbook_mistake_efee1aff7d855eb5}`
