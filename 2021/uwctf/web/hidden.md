# web / hidden

> Points: 100

## Question

> What is the place of all hidden things? \
> http://chall.uwc.tf:2004

## Solution

Loading the page gives a blank website that says "There's nothing on this page :("

Since the challenge name is `hidden`, we check `/robots.txt`:

```text
User-Agent: *
Disallow: /ctf-key
```

Navigating to `/ctf-key` gives

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="shortcut icon" href="/images/logo.jpg" type="image/jpg" />
    <link rel="stylesheet" href="/styles/ctf-key.css">
    <title>CTF Key</title>
</head>
<body>
    <h1 id="main-heading">You came to the right place!</h1>
    <h4 id="ctf-key">Here's the key: 
        <!-- Empty -->
        <a id="ctf-final-key"></a>
        <!-- Clue - 2: does any other file reference #ctf-final-key? -->
    </h4>
</body>
</html>
```

If we access the `ctf-key.css` file, we find the key:
```css
#ctf-final-key {
    display: none;
    /*
    Here's the key:

    hidden_in_plain_sight_418bf76abfed4e04

    The flag format is: uwctf{the final key}
    Enter the key inside the braces, without any spaces
    */
}
```

### Flag

`uwctf{hidden_in_plain_sight_418bf76abfed4e04}`