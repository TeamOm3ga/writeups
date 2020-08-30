# misc / Secret Array

> Points: 283

## Question

> `nc secretarray.fword.wtf 1337`

## Solution

Connecting to the server will provide instructions:

```text
I have a 1337 long array of secret positive integers. The only information I can provide is the sum of two elements. You can ask for that sum up to 1337 times by specifing two different indices in the array.

[!] - Your request should be in this format : "i j". In this case, I'll respond by arr[i]+arr[j]

[!] - Once you figure out my secret array, you should send a request in this format: "DONE arr[0] arr[1] ... arr[1336]"

[*] - Note 1: If you guessed my array before 1337 requests, you can directly send your DONE request.
[*] - Note 2: The DONE request doesn't count in the 1337 requests you are permitted to do.
[*] - Note 3: Once you submit a DONE request, the program will verify your array, give you the flag if it's a correct guess, then automatically exit.
```

Letting aₙ be the n-th value and Sₘₙ be the sum of the m-th and n-th elements,
we can determine the first three elements by fetching S₀₁, S₀₂, and S₁₂.
Notice that S₀₁+S₀₂-S₁₂=a₀+a₁+a₀+a₂-a₁-a₂=2a₀.
Once we have a₀, we can fetch the remaining elements with S₀ₙ-a₀=aₙ.

This algorithm is implemented (messily) in Python here:

```py
import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('secretarray.fword.wtf', 1337))

buffer = b""
def readline():
  global buffer
  global sock
  while b"\n" not in buffer:
    buffer += sock.recv(1024)
  buffer = buffer.split(b"\n")
  result = buffer[0].decode('ascii')
  buffer = b"\n".join(buffer[1:])
  print(result)
  return result
# get sum of the ith + jth element
def s(i, j):
  global sock
  req = f'{i} {j}'
  print(req)
  sock.send((req + '\n').encode('ascii'))
  return int(readline())

# read header
line = readline()
while line != 'START:':
  line = readline()
array = []
# use first three requests to get S(0,1) S(0,2) and S(1,2)
s01 = s(0,1)
s02 = s(0,2)
s12 = s(1,2)
# get 0th element from these algebraically
e0 = (s01 + s02 - s12) // 2
array = ['DONE', str(e0), str(s01 - e0), str(s02 - e0)]
# generate remaining elements
for i in range(3, 1337):
  s0i = s(0, i)
  array.append(str(s0i - e0))
done = ' '.join(array)
print(done)
sock.send((done + '\n').encode('ascii'))
readline()

```

Running this script (and waiting a bit) eventually gets us our flag:

```
Congratualtions! You guessed my secret array, here is your flag: FwordCTF{it_s_all_about_the_math}
```

### Flag

`FwordCTF{it_s_all_about_the_math}`
