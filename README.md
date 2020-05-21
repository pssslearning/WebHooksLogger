# WebHooksLogger
Ejemplo de uso de Go Gin para disponer de un Web Hook Logger

## End-points habilitados

```go
// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showRootGET)
	router.POST("/", showRootPOST)

	// Handle the API samples route
	router.GET("/api/samples", showApiSamplesGET)
	router.POST("/api/samples", showApiSamplesPOST)

	// Handle the PAY-IN Webhook route
	router.POST("/webhook/event", showWebhookEventPOST)
}
```

## Compilación
```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/WebHooksLogger/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ go build -o webhooks-logger
```

## Preparación e instalación del servicio
```sh
ubuntu@ip-172-xx-yy-zz:~$ mkdir -pv webhooks-logger/log
mkdir: created directory 'webhooks-logger'
mkdir: created directory 'webhooks-logger/log'
ubuntu@ip-172-xx-yy-zz:~$ cd webhooks-logger/
ubuntu@ip-172-xx-yy-zz:~/webhooks-logger$ pwd
/home/ubuntu/webhooks-logger
```

> **NOTA**  
>   // By default it serves on :8080 unless a  
>	  // PORT environment variable was defined.  
>	  router.Run()  
>	  // router.Run(":3000") for a hard coded port  

> **NOTA: se dá por hecho que vía SFTP se sube el binario `webhooks-logger` a la carpeta `/home/ubuntu/webhooks-logger` 

```sh
ubuntu@ip-172-xx-yy-zz:~$ sudo vim /etc/systemd/system/webhooks-logger.service
ubuntu@ip-172-xx-yy-zz:~$ sudo cat /etc/systemd/system/webhooks-logger.service
[Unit]
Description=Webhooks Logger Service
After=syslog.target network.target

[Service]
Type=simple
User=ubuntu
Group=ubuntu
WorkingDirectory=/home/ubuntu/webhooks-logger
Environment="PORT=8030"

ExecStart=/home/ubuntu/webhooks-logger/webhooks-logger
ExecStop=/bin/kill -SIGTERM $MAINPID

[Install]
WantedBy=multi-user.target

```

## Prueba
```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/WebHooksLogger/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ export PORT=8030
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ ./webhooks-logger &
[1] 3728
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ sudo lsof -iTCP -sTCP:LISTEN -P -n | grep 8030
webhooks- 3728    devel1    6u  IPv6  35392      0t0  TCP *:8030 (LISTEN)

curl --verbose -H"Accept: application/json" -H"x-custom: custom-data" --data-ascii '{"mensaje":"Hola Mundo"}' --request POST "http://localhost:8030/webhook/event?type=pay-in&data=1234567890&data2=Hola%20Mundo&flag=true"


* Connected to localhost (::1) port 8030 (#0)
> POST /webhook/event?type=pay-in&data=1234567890&data2=Hola%20Mundo&flag=true HTTP/1.1
> Host: localhost:8030
> User-Agent: curl/7.64.0
> Accept: application/json
> x-custom: custom-data
> Content-Length: 28
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 28 out of 28 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Tue, 10 Mar 2020 07:27:10 GMT
< Content-Length: 18
< 
{"response":"OK"}
* Connection #0 to host localhost left intact


devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ ls -la log/
total 16
drwxr-xr-x 2 devel1 devel1 4096 mar 10 08:30 .
drwxr-xr-x 5 devel1 devel1 4096 mar 10 08:30 ..
-rw-r--r-- 1 devel1 devel1  532 mar 10 08:27 2020-03-10T08-27-10-01-00-POST-1649ddaf-0128-49ae-8bc3-cdabb2c2ab59.txt
-rw-r--r-- 1 devel1 devel1    0 mar 10 08:23 GIN-2020-03-10T08-23-25-01-00-error.log
-rw-r--r-- 1 devel1 devel1  771 mar 10 08:27 GIN-2020-03-10T08-23-25-01-00-request.log
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$

devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ cat log/GIN-2020-03-10T08-23-25-01-00-request.log
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.showRootGET (3 handlers)
[GIN-debug] POST   /                         --> main.showRootPOST (3 handlers)
[GIN-debug] POST   /webhook/event            --> main.showWebhookEventPOST (3 handlers)
[GIN-debug] Environment variable PORT="8030"
[GIN-debug] Listening and serving HTTP on :8030
[GIN] 2020/03/10 - 08:27:10 | 200 |     273.518µs |             ::1 | POST     /webhook/event?type=pay-in&data=1234567890&data2=Hola%20Mundo&flag=true

devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/WebHooksLogger$ cat log/2020-03-10T08-27-10-01-00-POST-1649ddaf-0128-49ae-8bc3-cdabb2c2ab59.txt 

---- URL & Query Params (begin) ------
/webhook/event?type=pay-in&data=1234567890&data2=Hola%20Mundo&flag=true
type: pay-in
data: 1234567890
data2: Hola Mundo
flag: true
---- URL & Query Params (end) --------
---- HEADERS (begin) -----------------
User-Agent: curl/7.64.0
Accept: application/json
X-Custom: custom-data
Content-Length: 28
Content-Type: application/x-www-form-urlencoded
---- HEADERS (end) -------------------
---- BODY    (begin) -----------------
{\"mensaje\":\"Hola Mundo\"}
---- BODY    (end) -------------------

```
