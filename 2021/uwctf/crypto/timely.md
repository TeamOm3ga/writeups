# crypto / timely

> Points: 200

## Question

> Naive encryption and bad APIs will fall to a timely attack.
>
> `nc chall.uwc.tf 2000`

### Provided Files

- [`main.go`](./main.go)

## Solution

We can analyze `main.go` to find that we can list files
with `GET /list` and get them with `POST /get`.
The available files are:
```json
{"filenames":["app","flag.txt","key.txt","main.go"]}
```

When we analyzed the "encryption" in `main.go`,
we noticed that it is an XOR with the `key`
and dependent on the current timestamp in seconds.

But we have a copy of `main.go` in plaintext.
This means we can retrieve the key by XORing.
Then, we can use that to decrypt `flag.txt`,
provided it was also retrieved in the same second.
We can get and verify using `curl`:
```text
curl -d '{"filename":"flag.txt"}' http://chall.uwc.tf:2000/get -vv && curl -d '{"filename":"main.go"}' http://chall.uwc.tf:2000/get -vv
```
This gives two payloads with the same timestamp:
```json
// flag.txt
{"payload":"4tblGxBW2fi7WAqbICRR/1xy82yQmw+3dWpQQ60="}
// main.go
{"payload":"58DlBBdK3b2lZgCBWBFnoBwtsWvv0DHSOzpGR9eU...
```
Now, we can bytewise XOR everything together:
```py
encrypted_main = base64.b64decode("58DlBBdK3b2lZgCBWBFnoBwtsWvv0DHSOzpGR9eU")
plaintext_main = b'package main\n\nimport (\n\t"crypto/aes"\n\t"crypto/cypher"'
encrypted_flag = base64.b64decode("4tblGxBW2fi7WAqbICRR/1xy82yQmw+3dWpQQ60=")
print("".join([chr(a^b^c) for (a,b,c) in zip(encrypted_main, plaintext_main, encrypted_flag)]))
```

### Flag

`uwctf{aes_ctr?_2000s_c4ll3d}`
