# rev / pdf

> Points: 100

## Question

> This PDF has more to it than meets the eye.

### Provided Files

- [`PDFCHALL.pdf`](./PDFCHALL.pdf)

## Solution

Considering most PDF metadata is stored in plaintext,
we started by checking to see if the flag could be found
by grepping. It could.

```text
$ strings PDFCHALL.pdf | grep uwctf
<photoshop:CaptionWriter>uwctf{m3t4d4t4134k_4f2b5d9cacedc697}</photoshop:CaptionWriter>
```

### Flag

`uwctf{m3t4d4t4134k_4f2b5d9cacedc697}`
