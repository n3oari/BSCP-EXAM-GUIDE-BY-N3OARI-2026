# API Testing - CHEATSHEET + MOST IMPORTANT LABS

## Key Terms

```
- Burp Extension -> Content-Type Converter
```

## Walkthrough - Most Important Labs

[Finding-and-exploiting-an-unused-API-endpoint](Finding-and-exploiting-an-unused-API-endpoint.md)

[Exploiting-a-mass-assignament-vulnerability-(hidden-parameter-fields)](Exploiting-a-mass-assignament-vulnerability-(hidden-parameter-fields).md)

## Methodology

- 1: Site map -> Scan -> Deep Scan
- 2: Site map -> Engagement tools -> Discovery Content
- 3: Use the diccionary from [botes juan](https://github.com/botesjuan) [burp-labs-wordlists](https://github.com/botesjuan/Burp-Suite-Certified-Practitioner-Exam-Study/blob/main/wordlists/burp-labs-wordlist.txt)[api-wordlist-repo-by-chrislockard](https://github.com/chrislockard/api_wordlist) 

```bash
wfuzz -c --hc=404 -w burp-labs-wordlist.txt https://<LAB>/FUZZ 
```





```
- Recon ->
	  > Search in official api documentacion (different versions too)
	  > /api + Sitemap
	  > Site map -> Engagement tools -> discover content
	  > Burp Intruder ->
	   	  > Add HTTP verbs list in critical endpoints
	   	  > /api/<wordlist> 
	   	  > /api/<version>/<wordlist/
	  > Content-Type: json/xml? -> Content-Type-Converter -> see changes
```

