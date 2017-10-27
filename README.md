# Guessing Game Web Application

## HTTP header inspection using cURL 

To show the response of HTTP request to a specific server, you can use the "curl -i" command to include the HTTP response headers: 
```
$ curl -i http://localhost:8080/asddgasg/asdgasdg
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    13  100    13    0     0     13      0  0:00:01 --:--:--  0:00:01 13000HTTP/1.1 200 OK
Date: Fri, 27 Oct 2017 13:00:59 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8

Guessing Game
```

To see the HTTP request and response headers plus the response content, use "curl -v" command: 
```
$ curl -v http://localhost:8080/
* timeout on name lookup is not supported
*   Trying ::1...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Fri, 27 Oct 2017 13:01:04 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
<
{ [13 bytes data]
100    13  100    13    0     0     13      0  0:00:01 --:--:--  0:00:01 13000Guessing Game
* Connection #0 to host localhost left intact
```

The first output are some transaction's statistics about request time:
```
* timeout on name lookup is not supported
*   Trying ::1...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (::1) port 8080 (#0)
  ``` 


The first header section are the request's HTTP headers: 
```
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
```

The next header section are the response's HTTP headers coming from the server:
```
< HTTP/1.1 200 OK
< Date: Fri, 27 Oct 2017 13:01:04 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
```

And the last part are some response statistics plus the content: 
```
{ [13 bytes data]
100    13  100    13    0     0     13      0  0:00:01 --:--:--  0:00:01 13000Guessing Game
```