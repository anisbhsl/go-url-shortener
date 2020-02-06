**go-url-shortener**

A simple URL shortener API written in Go. This URL shortening service uses Redis KV store to store URLs and their shortened hash. 

**API Description**

Endpoints:

1. Shorten your URL
```
/shorten?url=<URLtoShorten>
```
- *Returns*:
StatusCode: ```200```
Response:
```
{
    longURL:
    shortURL:
}
```

2. Use short URL to redirect to original URL
```
/{shortURL}
```

*Redirects to original URL*

---------------
Example:

*1. URL Shorten Request*
```
/shorten?url=https://www.facebook.com
```
*Response*
```
{
longURL: "https://www.facebook.com",
shortURL: "BpLnfg"
}
```

*2. Using short url 
```
/BpLnfg
```
*Redirects to* ```https://www.facebook.com```



*This project is under development. Any suggestions or PRs are highly welcomed!*
