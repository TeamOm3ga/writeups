# rev / cablefish

> Points: 100

## Question

> We have managed to capture a TLS session with this one website and steal their RSA key.
> Can you extract anything useful from this?

### Provided Files

- [`key.priv`](./key.priv)
- [`capture.pcap`](./capture.pcap)

## Solution

We searched Wireshark documentation for how to add a TLS key.
Right-clicking on a TLS packet and going to Protocol Preferences > RSA Keys List, we added the given `key.priv`.

We can now see a PDF file `2021-02.pdf` sent through TLS.
Opening it gives a message
```text
Well done. The flag is uwctf{use_ecdh_0c4ad8ce02dfe9aa}
```

### Flag

`uwctf{use_ecdh_0c4ad8ce02dfe9aa}`
