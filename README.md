# WebHooksLogger
Ejemplo de uso de Go Gin para disponer de un Web Hook Logger


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

ExecStart=/home/ubuntu/webhooks-logger/webhooks-logger
ExecStop=/bin/kill -SIGTERM $MAINPID

[Install]
WantedBy=multi-user.target

```

## Prueba
```sh
curl --verbose -H"Accept: application/json" -H"x-custom: custom-data"   --data-ascii "Hola Mundo" -X POST "http://ec2-aa-bb-ccc-ddd.eu-west-3.compute.amazonaws.com:8080/test/webhook/pay-in?data=1234567890&data2=Hola%20Mundo&flag=true"



ubuntu@ip-172-xx-yy-zz:~/webhooks-logger/log$ cat GIN-2019-11-08T18-47-01Z-request.log
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.showRootGET (3 handlers)
[GIN-debug] POST   /                         --> main.showRootPOST (3 handlers)
[GIN-debug] POST   /test/webhook/pay-in      --> main.showPayINWebhookPOST (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2019/11/08 - 18:52:11 | 200 |     242.393µs |   195.nn.nn.nn | POST     /test/webhook/pay-in?data=1234567890&data2=Hola%20Mundo&flag=true
ubuntu@ip-172-xx-yy-zz:~/webhooks-logger/log$ ls 2019-11-08T18-52-11Z-POST-b1d9b7a9-acbf-4e8f-8b6f-152e64578d8f.txt
2019-11-08T18-52-11Z-POST-b1d9b7a9-acbf-4e8f-8b6f-152e64578d8f.txt
ubuntu@ip-172-xx-yy-zz:~/webhooks-logger/log$ cat 2019-11-08T18-52-11Z-POST-b1d9b7a9-acbf-4e8f-8b6f-152e64578d8f.txt

---- URL & Query Params (begin) ------
/test/webhook/pay-in?data=1234567890&data2=Hola%20Mundo&flag=true
data: 1234567890
data2: Hola Mundo
flag: true
---- URL & Query Params (end) --------
---- HEADERS (begin) -----------------
User-Agent: curl/7.64.0
Accept: application/json
X-Custom: custom-data
Content-Length: 10
Content-Type: application/x-www-form-urlencoded
---- HEADERS (end) -------------------
---- BODY    (begin) -----------------
Hola Mundo
---- BODY    (end) -------------------
```