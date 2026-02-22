# JWT - CHEATSHEET + MOST IMPORTANT LABS

- [jwt.io (usefull debugging)](https://www.jwt.io/)
- [jwt secret wordlist](https://github.com/wallarm/jwt-secrets/blob/master/jwt.secrets.list)
- jwt editor (burp extension)

## Key Terms

```bash
- JWT structure: Header (alg + typ) → Payload (claims/data) → Signature → <HEADER>.<PAYLOAD>.<SIGNATURE>
- Algorithms:
    > symmetric  -> uses the same secret to sign and verify (e.g. HS256)
    > asymmetric -> uses a key pair: private key to sign, public key to verify (e.g. RS256)

- jwk → JSON Web Key (public key embedded directly in the JWT header)
- jku → JWK Set URL (pointer to a JWKS hosted on a server you control)
- kid → Key ID (identifier used to select which key in a JWKS or key store should be used to verify the signature)
- PEM (Privacy-Enchanced Mail) A text-based container format that encodes binary cryptographic data (keys, certificates, CSRs) in Base64 with -----BEGIN/END----- delimiters

```

## Walkthrough - Most Important Labs

- [JWT auth bypass weak simetric alg signing key (brute force secret)](jwt-auth-bypass-weak-signing-key.md)
- [JWT auth bypass via JWK header injection](jwt-auth-bypass-jwk-header-injection.md)
- [JWT auth bypass via JKU header injection](jwt-auth-bypass-via-jku-header-injection.md)
- [JWT auth bypass via KID path traversal to /dev/null](jwt-auth-bypass-kid-path-traversal.md)
- [JWT auth bypass algorithm confusion + public key exposed](jwt-auth-bypass-algorithm-confusion.md)

<br>

## Methodology

```bash
- Detect algorithm
    > symmetric?  -> brute force ,  path traversal in kid
    > asymmetric? -> JWK , JKU
- Modify the payload value.
- Change the alg to none -> remove signature from the JWT.
- Weak signature? — try brute‑forcing the secret.
- Try injecting header parameters — jwk (embedded key) or jku (JWKS URL).
- Try path‑traversal on kid (e.g. point it to /dev/null).
- Try to change asymmetric algorithm to symmetric
```

```bash
hashcat -a 0 <jwt-token> <jwt-wordlist>
/.well-known
/jwks.json
```