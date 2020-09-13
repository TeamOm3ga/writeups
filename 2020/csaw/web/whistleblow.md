# web / whistleblow

## Question

> One of your coworkers in the cloud security department sent you an urgent email, probably about some privacy concerns for your company.

### Provided Files

- `letter`

### Hints

> Hint 1: Presigning is always better than postsigning \
> Hint 2: Isn't one of the pieces you find a folder? \
> ~~Hint 3: Look for flag.txt!~~ (released after we solved)

## Solution

We are provided with a plaintext `letter`:

```text
Hey Fellow Coworker,

Heard you were coming into the Sacremento office today. I have some sensitive information for you to read out about company stored at ad586b62e3b5921bd86fe2efa4919208 once you are settled in. Make sure you're a valid user!
Don't read it all yet since they might be watching. Be sure to read it once you are back in Columbus.

Act quickly! All of this stuff will disappear a week from 19:53:23 on September 9th 2020.

- Totally Loyal Coworker
```

The only obvious information from this letter that we get is the 128-bit hex string `ad586b62e3b5921bd86fe2efa4919208`.
This is a web challenge, so there should be a website somewhere corresponding to this string.
The prompt indicated that we were looking for a cloud service provider.
We tried looking for GitHub commits, gist IDs, or a matching IPv6 address, to no avail.

When the first hint appeared, searching for `presigning` leads to a bunch of results about Amazon S3.
`ad586b62e3b5921bd86fe2efa4919208` is a valid S3 bucket.
Knowing that we have to be `a valid user` to access the bucket, we thought we would have to find creds somehow.
In actuality, we did not need _special_ creds, only to be logged in as any AWS user:

```text
$ aws s3 ls s3://ad586b62e3b5921bd86fe2efa4919208
                           PRE 06745a2d-18cd-477a-b045-481af29337c7/
                           PRE 23b699f9-bc97-4704-9013-305fce7c8360/
                           PRE 24f0f220-1c69-42e8-8c10-cc1d8b8d2a30/
                           PRE 286ef40f-cee0-4325-a65e-44d3e99b5498/
                           PRE 2967f8c2-4651-4710-bcfb-2f0f70ecea5c/
                           PRE 2b4bb8f9-559e-41ed-9f34-88b67e3021c2/
                           PRE 32ff8884-6eb9-4bc5-8108-0e84a761fe2c/
                           PRE 3a00dd08-541a-4c9f-b85e-ade6839aa4c0/
                           PRE 465d332a-dd23-459b-a475-26273b4de01c/
                           PRE 64c83ba4-8a37-4db8-b039-11d62d19a136/
                           PRE 6c748996-e05a-408a-8ed8-925bf01be752/
                           PRE 7092a3ec-8b3a-4f24-bdbd-23124af06a41/
                           PRE 84874ee9-cee1-4d6b-9d7a-24a9e4f470c8/
                           PRE 95e94188-4dd1-42d8-a627-b5a7ded71372/
                           PRE a50eb136-de5f-4bb6-94ef-e1ee89c26b05/
                           PRE b2896abb-92e7-4f76-9d8a-5df55b86cfd3/
                           PRE c05abd3c-444a-4dc3-9edc-bb22293e1e0f/
                           PRE c172e521-e50d-4e30-864b-f12d72f8bf7a/
                           PRE c9bf9d72-8f62-4233-9cd6-1a0f8805b0af/
                           PRE ff4ad932-5828-496b-abdc-6281600309c6/
```

That's a lot of folders. There are even more files in them:

```text
$ aws s3 ls s3://ad586b62e3b5921bd86fe2efa4919208 --recursive | wc -l
200
```

After syncing down the files, it looks like most files have 30 random lowercase letters in them.
We can use `find` and `grep` to filter those out:

```text
$ find . -iname "*.txt" -exec grep "" -H {} \; | grep -E -v ":[a-z]{30}"
./3a00dd08-541a-4c9f-b85e-ade6839aa4c0/3fa52aaa-78ed-4261-8bcc-04fc0b817395/4bcd2707-48db-4c04-9ec7-df522de2ccd7.txt:s3://super-top-secret-dont-look
./6c748996-e05a-408a-8ed8-925bf01be752/c1fe922c-aec8-4908-a97d-398029d39236/77010958-c8ed-4a7b-802a-f189d0f76ec0.txt:3560cef4b02815e7c5f95f1351c1146c8eeeb7ae0aff0adc5c1
06f6488db5b6b
./7092a3ec-8b3a-4f24-bdbd-23124af06a41/7db7f9b0-ab6a-4605-9fc1-1cc8ba7877a1/1b56b43a-7525-429a-9777-02602b52dc1e.txt:.sorry/.for/.nothing/
./c9bf9d72-8f62-4233-9cd6-1a0f8805b0af/acbad485-dd20-4295-99fa-f45e3d5bdb45/1eaddd5d-fe24-4deb-8e6e-5463f395fa03.txt:AKIAQHTF3NZUTQBCUQCK
```

Going back to "presigning", we detour into how S3 authenticates requests.
Typically, a request is signed by the secret key of the user making the request.
However, as an alternative, Amazon allows "presigning", where the authorized user creates a special link which can read a file for a certain amount of time.
This time maxes out at one week, the time specified in the `letter` file.

With three pieces of information from the text files:
an S3 bucket (`s3://super-top-secret-dont-look`),
an AWS credential (`AKIAQHTF3NZUTQBCUQCK`),
and a presigned signature (`3560cef4b02815e7c5f95f1351c1146c8eeeb7ae0aff0adc5c1`),
along with the expiry time (one week) and timestamp (`20200909T195323Z`) from the letter,
we can read a file with a carefully crafted request.

What file? We are only given a folder, `.sorry/.for/.nothing/`.
We tried using the ListObjectsV2 API, `/?list-type=2&prefix=.sorry/.for/.nothing/`, but that did not work.
Eventually, after solving [flask_caching](./flask_caching.md) by reading `/flag.txt`, we realized other challenges probably shared a similar setup.

Performing a presigned GET request to `/.sorry/.for/.nothing/flag.txt` gets the flag.

### Flag

`flag{pwn3d_th3_buck3ts}`
