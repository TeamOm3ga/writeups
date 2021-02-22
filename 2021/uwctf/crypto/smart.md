# crypto / smart

> Points: 500

## Question

> Now you're smart.
> I'm still wondering how my friend came up with that curve, though.
> Can you replicate that feat for this system which only likes curves of a particular size?
>
> `nc chall.uwc.tf 2002`

## Solution

Connecting to the server prints a message:
```text
Standard curves are so passé... Let's negotiate our own.
Give me p, a, b such that y^2 = x^3 + ax + b is an elliptic curve over Fp.
I will check that it's secure and send you P and Q = kP.
I don't think you can determine k. Prove me wrong and you'll get a flag!
p (384-bit prime): 
```
Essentially, we have to create a curve that would work with
the [notsmart](./notsmart) challenge.
After reading [another paper](http://www.monnerat.info/publications/anomalous.pdf)
(which, unfortunately, had no Sage code)
we found that it is pretty easy to generate these curves where the order equals the prime.

However, while searching for sample code that might help us implement the algorithm,
we found an entire [program](https://github.com/J08nY/ecgen) for it.
`ecgen` allows specification of the size of the prime, so that made it quite easy.

```json
// ./ecgen --anomalous --fp 384
[{
    "field": {
        "p": "0x9b0bb5d8d418b0aab36f6122bb8e7698782e6f22b55f8e54bb61a5364647f01ef7beea0eeca05080a3668a1833bfd6d9"
    },
    "a": "0x2fa5959f33210de060fb45c362f4b9ac571ec7b6e2dda458efe7673f8e7efd1b5e5ac06a6dc62ceec28fba7b27a9ab7a",
    "b": "0x626f8126623535dc7b158e8479bafa2752b985d03b5a6826cfe0e3808d3363cd1c25c2bec22534fdfa62a3e89d4701b0",
    "order": "0x9b0bb5d8d418b0aab36f6122bb8e7698782e6f22b55f8e54bb61a5364647f01ef7beea0eeca05080a3668a1833bfd6d9",
    "subgroups": [
        {
            "x": "0x18ca72e3128c3262c8e525aafe824d77e3705a68a0c1882cfcd27d99c2e85c3a42d53120b07745d87fc6468f09af7a01",
            "y": "0x3272899da856d713fc3123dfb3c90c4e0f357665c10ba6f6f0f13f65fbee68e109059e3953295a5e614831cae13e841a",
            "order": "0x9b0bb5d8d418b0aab36f6122bb8e7698782e6f22b55f8e54bb61a5364647f01ef7beea0eeca05080a3668a1833bfd6d9",
            "cofactor": "0x1",
            "points": [
                {
                    "x": "0x18ca72e3128c3262c8e525aafe824d77e3705a68a0c1882cfcd27d99c2e85c3a42d53120b07745d87fc6468f09af7a01",
                    "y": "0x3272899da856d713fc3123dfb3c90c4e0f357665c10ba6f6f0f13f65fbee68e109059e3953295a5e614831cae13e841a",
                    "order": "0x9b0bb5d8d418b0aab36f6122bb8e7698782e6f22b55f8e54bb61a5364647f01ef7beea0eeca05080a3668a1833bfd6d9"
                }
            ]
        }
    ]
}]
```
Using this and the same Sage code from `notsmart`, we negotiated a curve to get the flag:
```text
p (384-bit prime): 23863724010527267615392563341308477263697291340204146159335151977978529934678186881919081881996744109818265108207321
a: 7333515907157877475244809900796112790166948594758867769728677043526648835452128868248648059641421279290045621709690
b: 15150619997213698635641034892756724106870642580753672345177864460072323285888189053530809496985503174557201328439728
P: (8792262658743987381163815019237952513009044304501841544468182682458803257039874344202559062947311450322925137850275, 6270509291581139543041749977926479653260698310659666791597285835668714571053813110111624940693273788653801083259420)
Q: (11287257493245367241620332226821722753196875631481327269467872124299998879286449012103691627989149622686765223362133, 617505220228324653254494929328748953701200827361688833522846224344986925437766365608366437905491089909729523329475)
k: 13377567074704065378057795896310751071262043759406192850721723671243255326731795106061314673450570746315773090219750
```

### Flag

`uwctf{anomalous_moment_a5160cd470172a63}`
