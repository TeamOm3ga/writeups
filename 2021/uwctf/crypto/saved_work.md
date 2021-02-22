# crypto / saved_work

> Points: 100

## Question

> I generated RSA parameters like so, saving some work:
> ```py
> start = random.randint(2 ** 2047, 2 ** 2048)
> p = gmpy2.next_prime(start)
> q = gmpy2.next_prime(p)
> ```
> Something tells me this wasn't a bright idea, though...

### Provided Files

- [`flag.enc`](./saved_work/flag.enc)
- [`pub.pem`](./saved_work/pub.pem)

## Solution

Notice that `q` is the next prime after `p`.
Searching for "consecutive prime RSA attack" on Google, we found a
[StackOverflow post](https://math.stackexchange.com/questions/71122/factoring-n-where-n-pq-and-p-and-q-are-consecutive-primes)
explaining how to derive `p` and `q` from `n`.

We can extract `n` and `e` from the `pub.pem` file:

```py
from Crypto.PublicKey import RSA
with open('pub.pem', 'r') as f:
    pub_key = RSA.import_key(f.read())
    n = pub_key.n
    e = pub_key.e
```

Now, we apply the algorithm to find `p` and `q` from the SO post.
We can also find `d` here; Mathematica is good at modular arithmetic.
```mathematica
m = Block[{$MaxExtraPrecision = 10000}, Ceiling[Sqrt[n]]];
diff = Sqrt[m^2 - n];
{p, q} = {m + diff, m - diff};
d = D /. Solve[e*D == 1, D, Modulus -> LCM[p - 1, q - 1]][[1]];
```

Now, we use `n`, `e`, and `d` to construct a private key:
```py
key = RSA.construct((n, e, d))
with open('priv.pem', 'wb') as f:
    f.write(key.export_key())
```
and decrypt the flag:
```sh
base64 -d flag.enc > flag.enc.bin
openssl rsautl -decrypt -inkey priv.pem -in flag.enc.bin
# uwctf{thats_too_close_dddeb91c2eaafe6b}
```

### Flag

`uwctf{thats_too_close_dddeb91c2eaafe6b}`
