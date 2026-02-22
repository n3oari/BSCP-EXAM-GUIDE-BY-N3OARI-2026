# JWT authentication bypass via jku header injection

In this lab we will exploit the **jku** parameter in a JWT header. This parameter allows the server to dynamically fetch a public key from an external URL to verify the token’s signature. The vulnerability occurs when the server does not validate that the jku URL belongs to a trusted origin — an attacker can point jku to their own JWKS and cause the server to use an attacker-controlled public key to verify signatures.

<br>

- 1: Generate an RSA key pair (private + public).


![Screenshot1](../../04-Screenshots/jku1.png)


<br>

- 2: Convert the public key to JWK/JWKS and upload it to your Exploit Server (or serve it from a URL you control).
  
  
![Screenshot2](../../04-Screenshots/jku2.png)


<br>


- 3: Modify the JWT header: replace the kid with the kid from your key and add jku pointing to your exploit server. Sign the token with your private key and send the JWT to the server.


![Screenshot3](../../04-Screenshots/jku3.png)


```json

{
   "keys": [

     <public-key-JWT-key>
    
 ]
}

```

