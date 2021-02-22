# pwn / puppy

> Points: 150

## Question

> Spot the puppy accidentally destroyed his Puppy Linux image and now he can't boot.
> He needs you to save his files!

### Provided Files

- `fossapup64-9.5.iso`

## Solution

Trying to mount the ISO file gives an error.
Let's check if it's even an ISO file:

```text
$ file fossapup64-9.5.iso
fossapup64-9.5.iso: RAR archive data, v5
```
Well, let's extract it then.
```text
$ unrar l fossapup64-9.5.so

UNRAR 6.00 freeware      Copyright (c) 1993-2020 Alexander Roshal

Archive: fossapup64-9.5.iso
Details: RAR 5

 Attributes      Size     Date    Time   Name
----------- ---------  ---------- -----  ----
    ..A....     24576  2019-11-17 15:59  isolinux.bin
    ..A.... 136105984  2020-07-26 08:30  puppy_fossapup64_9.5.sfs
    ..A....     13808  2020-09-17 06:13  README.txt
    ..A....   6575264  2020-07-26 08:28  vmlinuz
    ..A....  27963392  2020-07-26 08:32  zdrv_fossapup64_9.5.sfs
    ..A....      2048  2020-09-17 06:15  boot/boot.catalog
    ..A....   9437184  2019-11-25 11:55  boot/efi.img
    ..A....      5004  2019-06-25 05:11  boot/grub/font.pf2
    ..A....    331557  2019-10-27 21:36  boot/grub/grldr
    ..A....      1598  2020-09-17 06:15  boot/grub/grub.cfg
    ..A....      1223  2020-09-17 06:15  boot/grub/loopback.cfg
    ..A....      1973  2020-09-17 06:15  boot/grub/menu.lst
    ..A....      3380  2020-09-17 06:15  boot/grub/menu_phelp.lst
    ..A....     38568  2019-10-27 14:43  boot/isolinux/chain.c32
    ..A....     21792  2019-08-10 15:12  boot/isolinux/isohybrid
    ..A....     14184  2019-08-19 09:53  boot/isolinux/isohybrid.pl
    ..A....     22856  2019-08-19 09:53  boot/isolinux/isohybrid64
    ..A....     24576  2019-10-27 14:43  boot/isolinux/isolinux.bin
    ..A....        93  2019-10-29 05:57  boot/isolinux/isolinux.cfg
    ..A....    358952  2020-06-23 07:33  boot/splash.jpg
    ..A....    541649  2020-09-17 06:15  boot/splash.png
    ..A....      1691  2019-11-02 12:05  Windows_Installer/readme.html
    ..A....      1292  2019-11-02 12:05  Windows_Installer/readme.txt
    ..A.... 239190016  2020-09-17 06:15  adrv_fossapup64_9.5.sfs
    ..A....  41168896  2020-09-17 06:15  fdrv_fossapup64_9.5.sfs
    ..A....      1598  2020-09-17 06:15  grub.cfg
    ..A....   1372644  2020-09-17 06:13  initrd.gz
    ...D...         0  2020-09-17 06:15  boot/grub
    ...D...         0  2019-11-17 16:17  boot/isolinux
    ...D...         0  2020-09-17 06:15  boot
    ...D...         0  2020-11-07 18:31  Windows_Installer
----------- ---------  ---------- -----  ----
            463225798                    31
```
The `.sfs` files are `squashfs` filesystem images.
One of them probably has the flag.

Spoiler: only one of them has stuff we care about in it.

Now, we can mount `puppy_fossapup64_9.5.sfs`:
```text
$ sudo mount puppy_fossapup64_9.5.sfs /mnt/test -t squashfs
```
Inside here, we find `/home/spot/README.txt`:
```text
"spot" is a restricted user of Puppy.
(you could also add finn, rover, rex!)

/home/spot is spot's home directory.

To change from root to user spot, type this:

# su spot
# cd ~

Or, do this:

# su --login spot

Note that spot has been setup to not require a password.

You can confirm that you are indeed now spot:

# whoami

When you have finished being spot, type this:

# exit

-------------------------

User "spot" is currently used in Puppy by DidiWiki.
See the script /usr/sbin/didiwiki-gui, which is run from
the window manager menu, in "Information managers" submenu.

uwctf{puppiesarecute_690b5d89d409ffcb}
```

### Flag

`uwctf{puppiesarecute_690b5d89d409ffcb}`
