# crypto / Randomness

> Points: 471

## Question

> XOR is simple , but if I choose random things, would it be more secure?
> Author: Semah BA

## Solution

We are given a script that outputs a sequence of "random" numbers:

```python
from Crypto.Util.number import *
from random import *

flag="TODO"
p=getPrime(64)
a=getrandbits(64)
b=getrandbits(64)
X=[]
X.append((a*getrandbits(64)+b)%p)
c=0
while c<len(flag):
	X.append((a*X[c]+b)%p)
	c+=1

output=[]

for i in range(len(flag)):
	output.append(ord(flag[i])^X[i])

print (output)

#output:[6680465291011788243L, 5100570103593250421L, 5906808313299165060L, 1965917782737693358L, 9056785591048864624L, 1829758495155458576L, 6790868899161600055L, 1596515234863242823L, 1542626304251881891L, 8104506805098882719L, 1007224930233032567L, 3734079115803760073L, 7849173324645439452L, 8732100672289854567L, 5175836768003400781L, 1424151033239111460L, 1199105222454059911L, 1664215650827157105L, 9008386209424299800L, 484211781780518254L, 2512932525834758909L, 270126439443651096L, 3183206577049996011L, 3279047721488346724L, 3454276445316959481L, 2818682432513461896L, 1198230090827197024L, 6998819122186572678L, 9203565046169681246L, 2238598386754583423L, 467098371562174956L, 5653529053698720276L, 2015452976526330232L, 2551998512666399199L, 7069788985925185031L, 5960242873564733830L, 8674335448210427234L, 8831855692621741517L, 6943582577462564728L, 2159276184039111694L, 8688468346396385461L, 440650407436900405L, 6995840816131325250L, 4637034747767556143L, 3074066864500201630L, 3089580429060692934L, 2636919931902761401L, 5048459994558771200L, 6575450200614822046L, 666932631675155892L, 3355067815387388102L, 3494943856508019168L, 3208598838604422062L, 1651654978658074504L, 1031697828323732832L, 3522460087077276636L, 6871524519121580258L, 6523448658792083486L, 127306226106122213L, 147467006327822722L, 3241736541061054362L, 8781435214433157730L, 7267936298215752831L, 3411059229428517472L, 6597995245035183751L, 1256684894889830824L, 6272257692365676430L, 303437276610446361L, 8730871523914292433L, 6472487383860532571L, 5022165523149187811L, 4462701447753878703L, 1590013093628585660L, 4874224067795612706L]
```

This script implements a [Linear Congruentual Generator](https://en.wikipedia.org/wiki/Linear_congruential_generator). The script then XORs the numbers from this LCG with the flag and prints the result of that.

To get the flag we need to find the values generated by the LCG and XOR them with the output. Because flags always start with `FwordCTF`, we can simply XOR this again with the output to get the first 8 values from the LCG back:

```python
#!/usr/bin/env python3
import functools

flag = "FwordCTF"
output = [...]

seeds = [ord(a) ^ b for a, b in zip(flag, output)]
print(seeds)
#[6680465291011788181, 5100570103593250306, 5906808313299165163, 1965917782737693404, 9056785591048864532, 1829758495155458643, 6790868899161600099, 1596515234863242753]
```

Next, we need to extract the parameters of the LCG to generate the rest of the sequence. We need to find the modulus, multiplier, and the increment, in that order. I'll explain the increment first though, because it's the simplest.

First let's define some variables. Remember, we already know `s(0..7)`.

```
m = modulus
a = multiplier
c = increment
s(n) = s(n-1) * a + c (mod m)
```

To find the increment, assuming we have the modulus and multiplier, is simple:

```
s(1) = s(0) * a + c (mod m)
c = s(1) - s(0) * a (mod m)
```

Note that we can find the increment with just 2 seeds! This would give us the increment if we know the modulus and multiplier. So let's work on the multiplier next, assuming we have only the modulus. We need 3 seeds this time:

```
s(1) = s(0) * a + c (mod m)
s(2) = s(1) * a + c (mod m)
s(2) - s(1) = (s(0) * a + c) - (s(1) * a + c) (mod m)
s(2) - s(1) = (s(0) * a) - (s(1) * a) (mod m)
s(2) - s(1) = a * (s(0) - s(1)) (mod m)
a = (s(2) - s(1)) / (s(0) - s(1)) (mod m)
```

We need to be careful about the division here, in modular arithmetic, division means multiplying by the [Modular multiplicative inverse](https://en.wikipedia.org/wiki/Modular_multiplicative_inverse):

```
a = (s(2) - s(1)) * inv(s(0) - s(1)) (mod m)
```

Great, once we have the modulus we can use this to find the multiplier, and then use the modulus and the multiplier to find the increment. So we just have one variable left.

Finding the modulus is much harder, but it's still possible. If we have a few multiples of some number `x`, we can find `x` by taking their [Greatest common divisor](https://en.wikipedia.org/wiki/Greatest_common_divisor), for example:

```
a = 1337 * rand()
b = 1337 * rand()
c = 1337 * rand()
gcd(a, b, c) = 1337
```

For modulus, a multiple of `m` is the same as a number [Congruent](https://mathworld.wolfram.com/Congruence.html) (but not equal) to 0 (mod `m`). So we just need to find some numbers congruent (but not equal) to 0 (mod `m`), and we can find the modulus by taking their `gcd`:

```
d(n) = s(n+1) - s(n) (mod m)
d(n) = (s(n) * a + c) - (s(n-1) * a + c) (mod m)
d(n) = (s(n) * a) - (s(n-1) * a) (mod m)
d(n) = a * (s(n) - (s(n-1))) (mod m)
d(n) = a * d(n-1) (mod m)
z(n) = d(n+2) * d(n) - d(n+1) * d(n+1) (mod m)
z(n) = a*a*d(n) * d(n) - a*d(n) * a*d(n) (mod m)
z(n) = 0 (mod m)
m = gcd(z(0), z(1), ...)
```

Let's write some code to do all of that:

```python
def egcd(a, b):
	if a == 0:
		return (b, 0, 1)
	g, x, y = egcd(b % a, a)
	return (g, y - b // a * x, x)
def inv(a, modulus): return egcd(a % modulus, modulus)[1] % modulus
def gcd(a, b): return egcd(a, b)[0]

diffs = [s1 - s0 for s0, s1 in zip(seeds, seeds[1:])]
zeroes = [d2 * d0 - d1 * d1 for d0, d1, d2 in zip(diffs, diffs[1:], diffs[2:])]
modulus = abs(functools.reduce(gcd, zeroes))
multiplier = (seeds[2] - seeds[1]) * inv(seeds[1] - seeds[0], modulus) % modulus
increment = (seeds[1] - seeds[0] * multiplier) % modulus
print(modulus, multiplier, increment)
#9444729917070668893 7762244320486225184 731234830430177597
```

Now that we have all the parameters needed to generate the key, we can decrypt the flag:

```python
flag = ""
seed = seeds[0]
for c in output:
	flag += chr(c ^ seed)
	seed = (seed * multiplier + increment) % modulus
print(flag)
#FwordCTF{LCG_easy_to_break!That_was_a_mistake_choosing_it_as_a_secure_way}
```

### Flag

`FwordCTF{LCG_easy_to_break!That_was_a_mistake_choosing_it_as_a_secure_way}`
