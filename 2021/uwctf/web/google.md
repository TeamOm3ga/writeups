# web / google

> Points: 100

## Question

> You can get to Google by visiting http://chall.uwc.tf:2003

## Solution

Visiting the site redirects to Google.
We can analyze how this redirect works using curl:

```html
<!DOCTYPE HTML>

<html>
<head>
<title>Redirecting...</title>
<meta name="robots" content="no-index,no-follow">
<meta http-equiv="Refresh" content="0; url=https://google.com">
<!--uwctf{insecure_redirect_b0741313ee6d5696}-->
</head>
</html>
```

...and there's the flag.

### Flag

`uwctf{insecure_redirect_b0741313ee6d5696}`
