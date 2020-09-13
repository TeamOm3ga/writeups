# crypto / authy

## Question

> Check out this new storage application that your government has started! It's supposed to be pretty secure since everything is authenticated...
>
> `curl crypto.chal.csaw.io:5003`

### Provided Files

- [`handout.py`](./handout.py)

## Solution

### Background

The target is a plaintext note storage service.
Notes are submitted to `POST /new` with an `author` value and a `note` value.
This endpoint returns a base64 version of the note and a SHA1 hash of the metadata.

Notes can be retrieved from `POST /view` with the note `id` (the base64 string) and `integrity` (the SHA1 hash).
If the `id` sets `admin` and `access_sensitive` to true and has `entrynum` 7, then the flag will be printed.

The `id` data is stored and parsed using URL parameter notation (`param1=value1&param2=value&...`).
All parameters are concatenated together before parsing.
However, the code does not sanitize for `&` within a note.

```py
# handout.py L93-97
identifier = base64.b64decode(info["id"]).decode()
checksum = info["integrity"]

params = identifier.replace('&', ' ').split(" ")
note_dict = { param.split("=")[0]: param.split("=")[1]  for param in params }
```

This allows us to arbitrarily set values of `note_dict`, which is used to verify the `admin`, `access_sensitive`, and `entrynum` properties.

### Exploit

We can now craft a note with the required fields (and `entrynum=none` to have something to overwrite):

```text
$ curl -X POST \
  -F "author=1337" \
  -F "entrynum=none" \
  -F "note=AAAAAAAA&admin=True&access_sensitive=True&entrynum=7" \
  crypto.chal.csaw.io:5003/new

Successfully added YWRtaW49RmFsc2UmYWNjZXNzX3NlbnNpdGl2ZT1GYWxzZSZhdXRob3I9MTMzNyZlbnRyeW51bT03ODMmbm90ZT1BQUFBQUFBQSZhZG1pbj1UcnVlJmFjY2Vzc19zZW5zaXRpdmU9VHJ1ZSZlbnRyeW51bT03:1a51c1aa28c65fb763539c8055ae270b4c231a11
```

And access it, triggering the flag print statement:

```text
$ curl -X POST \
  -F "id=YWRtaW49RmFsc2UmYWNjZXNzX3NlbnNpdGl2ZT1GYWxzZSZhdXRob3I9MTMzNyZlbnRyeW51bT03ODMmbm90ZT1BQUFBQUFBQSZhZG1pbj1UcnVlJmFjY2Vzc19zZW5zaXRpdmU9VHJ1ZSZlbnRyeW51bT03"
  -F "integrity=1a51c1aa28c65fb763539c8055ae270b4c231a11" \
  crypto.chal.csaw.io:5003/view

Author: admin
Note: You disobeyed our rules, but here's the note: flag{h4ck_th3_h4sh}
```

### Flag

`flag{h4ck_th3_h4sh}`
