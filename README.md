#  README.MD  (MARCH 2026)

This repo contains each cheatsheet along with its respective most important labs.

Each topic includes its own cheatsheet/methodology, the most relevant labs, and useful resources.

<br>

>⭐This is the result of months of work — I hope you find it helpful. If you do, please leave a star. ⭐

<br>

> ❗This repository contains **my own** cheatsheets and methodologies for the exam — **not PortSwigger's**. You may consult PortSwigger's resources if you wish, but the files here are my personal notes.

> ❗I highly recommend creating **your own** cheatsheets. This repository is intended to help others, provide examples, show the overall organization, and demonstrate my work and methodologies.

> 💡🤓 This repository is a summary of all my Obsidian notes for the BSCP. I learned a lot while creating it because I focused on making it as clear and didactic as possible, fully internalizing all the concepts to be able to explain them correctly. 

<br>

The following are basic resources offered by PortSwigger for the exam:

- [usernames wordlist](https://portswigger.net/web-security/authentication/auth-lab-usernames)  
- [passwords wordlist](https://portswigger.net/web-security/authentication/auth-lab-passwords)  
- [delimiters wordlist (web cache deception)](https://portswigger.net/web-security/web-cache-deception/wcd-lab-delimiter-list)

<br>

The exam consists of **2 machines**, each with **3 phases**, and a duration of **4 hours**.
  - 1: Access any user account.
  - 2: Elevate privileges or compromising the administrator account.
  - 3: Exfiltrate contents of /home/carlos/secret and **submit solution**

 <br>

## Some utilities

- [Most Important HTTP Headers](/03-Extra/HTTP-HEADERS.md)
- [HTB machines I recommended for each topic](/03-Extra/HTB-machines.md)
- [REGEX explanation (util to interpreter code)](https://pythonium.net/regex)
- [js-beautify (util to interpreter code)](https://beautifier.io/)
- [URL validation bypass cheat sheet (SSRF)](https://portswigger.net/web-security/ssrf/url-validation-bypass-cheat-sheet)  

## Enumeration & Web Discovery

- [API testing / recon](01-All-Topics/API-testing/01-CHEATSHEET+LABS.md)
- [Obfuscation payloads (escape bypass)](03-Extra/obfuscating-payload.md) 

### PHASE 1 → Obtain Inicial User

- [XSS / DOM vulnerabilities](01-All-Topics/XSS-AND-DOM//01-CHEATSHEET+LABS.md)
- [Authentication / Brute-Force](01-All-Topics/Authentication/01-CHEATSHEET+LABS.md)
- [OAuth](01-All-Topics/OAuth/01-CHEATSHEET+LABS.md)
- [Host Header Injection](01-All-Topics/Host-Header-Injection/01-CHEATSHEET+LABS.md)
- [Web Cache Poisoning](01-All-Topics/Web-Cache-Poisoning/01-CHEATSHEET+LABS.md)
- [Web Cache Deception](01-All-Topics/Web-Cache-Deception/01-CHEATSHEET+LABS.md)
- [HTTP Request Smuggling](01-All-Topics/HTTP-Request-Smuggling/01-CHEATSHEET+LABS.md)


### PHASE 2 → Elevate Privileges 

- [SQL Injection](01-All-Topics/SQL-Injection/01-CHEATSHEET+LABS.md)
- [NoSQL Injection](01-All-Topics/NoSQL-Injection/01-CHEATSHEET+LABS.md)
- [Client Side Request Forgery (CSRF)](01-All-Topics/CSRF/01-CHEATSHEET+LABS.md) 
- [Cross-Origin Resource Sharing (CORS)](01-All-Topics/CORS/01-CHEATSHEET+LABS.md)
- [Prototype Pollution](01-All-Topics/Prototype-Pollution/01-CHEATSHEET+LABS.md)
- [Authentication -> Password Reset](01-All-Topics/Authentication/01-CHEATSHEET+LABS.md)
- [JWT](01-All-Topics/JWT/01-CHEATSHEET+LABS.md)
- [GraphQL API](01-All-Topics/GraphQL-API/01-CHEATSHEET+LABS.md)

### PHASE 3 → Exfiltrate Data

- [XML - XXE Injection](01-All-Topics/XML-XXE-Injection/01-CHEATSHEET+LABS.md)
- [Server Side Request Forgery (SSRF)](01-All-Topics/SSRF/01-CHEATSHEET+LABS.MD)
- [Path Traversal](/01-All-Topics/Path-Traversal/01-CHEATSHEET+LABS.md)
- [OS - Command Injection](/01-All-Topics/OS-Injection/01-CHEATSHEET+LABS.md)
- [Deserialization Insecure](01-All-Topics/Deserialization-Insecure/01-CHEATSHEET+LABS.md)
- [File Uploads (Web Shell)](01-All-Topics/File-Uploads/01-CHEATSHEET+LABS.MD)
- [Server Side Template Injection (SSTI)](01-All-Topics/SSTI/01-CHEATSHEET+LABS.md)


### OTHERS 

- [Race Condition](01-All-Topics/Race-Condition/01-CHEATSHEET+LABS.md)
- [Clickjacking](01-All-Topics/Clickjacking/01-CHEATSHEET+LABS.md)

<br>

<br>


> ⚠️ This is a reference — please, always verify and research on your own.

> ❗ Remember: You should probably chain multiple vulnerabilities.

<br>

<div align="center">
  
|         Category          | Stage 1 | Stage 2 | Stage 3 |
| :-----------------------: | :-----: | :-----: | :-----: |
|            XSS            |   🟢    |   🟢    |   🟡    |
|            DOM            |   🟢    |   🟢    |     🟡      |
|       SQL Injection       |     🔴    |   🟢    |   🟡    |         |
|      NoSQL Injection      |     🔴   |   🟢    |   🟡    |         |
|            CSRF           |   🟢    |   🟢    |   🔴      |
|            SSRF           |    🔴     |    🟡     |   🟢    |
|       Authentication      |   🟢    |   🟢    |     🔴    |
|           OAuth           |   🟢    |   🟢    |    🔴     |
|   OS Command Injection    |   🔴      |  🔴       |   🟢    |
|    Web Cache Poisoning    |   🟢    |   🟢    |    🔴     |
|    Web Cache Deception    |   🟢    |   🟢    |  🔴       |
|         File Upload       |   🔴    |   🔴    |   🟢    |
|   Host Header Injection   |   🟡    |    🟡   |    🟢      |
| Insecure Deserialization  |   🔴    |    🔴     |   🟢    |
|   HTTP Request Smuggling  |   🟢    |   🟢    |    🔴     |
|        API         |   🟢    |   🟢    |     🔴    |
|            CORS           |   🟢    |   🟢    |    🔴     |
|    Prototype Pollution    |     🟢     |   🟢    |  🟢        |
|             JWT           |   🟢    |   🟢    |     🔴    |
|   GraphQL - API Endpoints |         |   🟢    |      🔴   |
|         XML - XXE         |   🔴    |   🟡      |   🟢    |
|            SSTI           |     🔴    |   🔴      |   🟢    |
|   Broken Access Control   |   🔴    |   🟢    |   🔴      |
|       Path Traversal      |    🔴     |  🔴       |   🟢    |
|       Race Condition      |   🟢    |         |   🟢    |



</div>

<br>

## Contact me! 

[![Discord](https://img.shields.io/badge/Discord-n3oari-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/users/1078444408086728839)

<br>
