# boomer / olddos

## Question

> Running obsolete technology is awesome :D

## Solution

The website has a prompt to read any file.
Trying to get `flag_file.txt` fails with `forbidden!`, so it isn't that easy.
Accessing `main.go` by fetching `/show?path=main.go` reveals the code that performs the fetch.

```go
func read_file(base string, name string) string {
  if strings.Contains(strings.ToLower(name), "flag_file.txt") || strings.ContainsAny(name, "/") {
    return "forbidden!"
  }

  if content, err := ioutil.ReadFile(base + name); err == nil {
    return string(content)
  } else {
    return "error!"
  }
}
```

We have to find a filename that resolves to `flag_file.txt` without using the string `flag_file.txt`.
Based on the challenge name, we researched old DOS file naming conventions.
[8.3 filenames](https://en.wikipedia.org/wiki/8.3_filename) seem to be what we want.

In old versions of DOS/Windows, file names were limited to eight name characters and a three character extension.
Any more characters were truncated and replaced with `~X` for the `X`th file. In this case, `flag_file.txt` becomes `FLAG_F~1.TXT`.

Fetching `/show?path=FLAG_F~1.TXT` gives the flag.

```html
<html>
<body>
<form action="/show" method="get">
  Path:
  <input type="text" name="path" />
  <input type="submit" value="show" />
  <br>
</form>
<pre>
ALLES{legacy_FS_are_pr3tty_sc4ry}

</pre>
</body>
</html>
```

### Flag

`ALLES{legacy_FS_are_pr3tty_sc4ry}`
