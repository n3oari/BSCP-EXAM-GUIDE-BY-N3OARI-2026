
# Lab: Authentication: Broken brute-force protection (IP block)  
https://portswigger.net/web-security/authentication/password-based/lab-broken-bruteforce-protection-ip-block

In this lab we observed that the server **blocks the IP after the third failed login attempt**. A successful login resets the counter.

## Cheeatsheet proviedes by Portswigger:

users: https://portswigger.net/web-security/authentication/auth-lab-usernames

passwords: https://portswigger.net/web-security/authentication/auth-lab-passwords
## Goal

Create two payload lists that bypass this restriction by using a three-request pattern: two failing attempts followed by a legitimate successful login that resets the IP block.

## Payload structure
- **users.txt**: repeat the pattern `carlos`, `carlos`, `wiener`, ...
- **final_pass.txt**: for each block of three lines, the first two are guesses for `carlos`, and the third is the valid password for `wiener` (e.g., `peter`), then repeat.

## Example flow (Burp Intruder)
Three consecutive requests (one block):
```
username=carlos&password=<PASSWORD> -> failure 
username=carlos&password=<PASSWORD> -> failure
username=wiener&password=peter -> success (resets IP block)
```



