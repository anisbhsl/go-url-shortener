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

2. Retrieve original URL [Access resource using shortURL]
```
/{shortURL}
```

- *Returns*
StatusCode: ```200```
Response:
```
{
    longURL:
    shortURL:
}
```




*This project is under development. Any suggestions and PRs are highly welcomed!*