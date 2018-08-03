from 阮一峰的网络日志
http://www.ruanyifeng.com/blog/2018/02/nginx-docker.html

### 一、HTTP 服务

```
 docker container run \
  -d \
  -p 8080:80 \
  --rm \
  --name mynginx \
  nginx

```

    -d：在后台运行
    -p ：容器的80端口映射到127.0.0.2:8080
    --rm：容器停止运行后，自动删除容器文件
    --name：容器的名字为mynginx

    把这个容器终止，由于--rm参数的作用，容器文件会自动删除。

`docker container stop mynginx`

### 二、映射网页目录
网页文件都在容器里，没法直接修改，显然很不方便。下一步就是让网页文件所在的目录/usr/share/nginx/html映射到本地。

```
mkdir html

docker container run \
  -d \
  -p 8080:80 \
  --rm \
  --name mynginx \
  --volume "$PWD/html":/usr/share/nginx/html \
  nginx

```


### 三、拷贝配置

`docker container cp mynginx:/etc/nginx .`


### 四、映射配置目录

```
 docker container run \
  --rm \
  --name mynginx \
  --volume "$PWD/html":/usr/share/nginx/html \
  --volume "$PWD/conf":/etc/nginx \
  -p 8080:80 \
  -d \
  nginx

```

### 五、自签名证书
现在要为容器加入 HTTPS 支持，第一件事就是生成私钥和证书。正式的证书需要证书当局（CA）的签名，这里是为了测试，
搞一张自签名（self-signed）证书就可以了。
我参考的是 DigitalOcean 的教程。首先，确定你的机器安装了 OpenSSL
https://www.digitalocean.com/community/tutorials/how-to-create-a-self-signed-ssl-certificate-for-nginx-in-ubuntu-16-04

```
sudo openssl req \
  -x509 \
  -nodes \
  -days 365 \
  -newkey rsa:2048 \
  -keyout example.key \
  -out example.crt
  
  ```

  上面命令的各个参数含义如下
  ```
    req：处理证书签署请求。
    -x509：生成自签名证书。
    -nodes：跳过为证书设置密码的阶段，这样 Nginx 才可以直接打开证书。
    -days 365：证书有效期为一年。
    -newkey rsa:2048：生成一个新的私钥，采用的算法是2048位的 RSA。
    -keyout：新生成的私钥文件为当前目录下的example.key。
    -out：新生成的证书文件为当前目录下的example.crt。
  ```

  执行后，命令行会跳出一堆问题要你回答，比如你在哪个国家、你的 Email 等等。
  其中最重要的一个问题是 Common Name，正常情况下应该填入一个域名，这里可以填 127.0.0.1
  `Common Name (e.g. server FQDN or YOUR name) []:127.0.0.1`

 ```
 ➜  nginx mkdir conf/certs
 ➜  nginx  mv example.crt example.key conf/certs
 ```
 

```

docker container run \
  --rm \
  --name mynginx \
  --volume "$PWD/html":/usr/share/nginx/html \
  --volume "$PWD/conf":/etc/nginx \
  -p 8080:80 \
  -p 8081:443 \
  -d \
  nginx
  
  ```


`docker exec -it nginx /etc/init.d/nginx reload`