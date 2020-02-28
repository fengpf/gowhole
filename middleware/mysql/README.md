### 测试mysql最大连接数


``` sql

CREATE TABLE `test` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `a` int(11) NOT NULL,
  `b` int(11) NOT NULL,
  `c` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `a` (`a`,`b`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

set GLOBAL max_connections=256;--设置最大连接数

show variables like '%max_connections%';--mysql的最大连接数

show global status like 'Max_used_connections';--服务器响应的最大连接数

show processlist;--查看数据库的连接情况
show full processlist;--查看正在执行的完整SQL语句，完整显示

```

`ab -c 100 -n 100 'http://localhost:9000/mysql_max_conns'`

### go设置mysql连接池

关于连接池配置
>1.db.SetMaxIdleConns(n int) 设置连接池中的保持连接的最大连接数。
默认也是0，表示连接池不会保持释放连接池中的连接的连接状态：
即当连接释放回到连接池的时候，连接将会被关闭。这会导致连接再连接池中频繁的关闭和创建  

>2.db.SetMaxOpenConns(n int) 设置打开数据库的最大连接数。包含正在使用的连接和连接池的连接  
如果你的函数调用需要申请一个连接，并且连接池已经没有了连接或者连接数达到了最大连接数。
此时的函数调用将会被block，直到有可用的连接才会返回。
设置这个值可以避免并发太高导致连接mysql出现too many connections的错误。
该函数的默认设置是0，表示无限制   

>3.db.SetConnMaxLifetime(d time.Duration) 
设置连接可以被使用的最长有效时间，如果过期，连接将被拒绝   

``` go
    db.SetMaxOpenConns(256) //SetMaxOpenConns sets the maximum number of open connections to the database.
    db.SetMaxIdleConns(150)// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

```

> 1.所有请求使用db全局对象，不设置连接池  
   >>程序不会报错，响应变慢

> 2.所有请求使用db全局对象，设置连接池  
  >>程序不会报错，响应变快

> 3.每次请求创建新的db对象，不设置连接池  
  >>如果请求数量 > 最大连接数，则会出现 Error 1040: Too many connections

> 4.每次请求创建新的db对象，设置连接池  
  >>如果设置连接池的连接数 > 最大连接数，如果请求数量 >最大连接数，则会出现 Error 1040: Too many connections
  >>如果设置连接池的连接数 <= 最大连接数，请求数量的多少不会影响连接报错

### docker 启动

`docker pull mysql`

`docker run --name first-mysql -p 3306:3306 -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql`

run            运行一个容器
--name         后面是这个镜像的名称
-p 3306:3306   表示在这个容器中使用3306端口(第二个)映射到本机的端口号也为3306(第一个)
-d             表示使用守护进程运行，即服务挂在后台

-p 3306:3306：将容器的 3306 端口映射到主机的 3306 端口。

-v $PWD/conf:/etc/mysql/conf.d：将主机当前目录下的 conf/my.cnf 挂载到容器的 /etc/mysql/my.cnf。

-v $PWD/logs:/logs：将主机当前目录下的 logs 目录挂载到容器的 /logs。

-v $PWD/data:/var/lib/mysql ：将主机当前目录下的data目录挂载到容器的 /var/lib/mysql 。

-e MYSQL_ROOT_PASSWORD=123456：初始化 root 用户的密码。

`docker ps |grep mysql`

`docker exec -it b3f54a60b3ce /bin/bash`

docker 安装 mysql 8 版本

# docker 中下载 mysql
docker pull mysql

#启动
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Lzslov123! -d mysql

#进入容器
docker exec -it mysql bash

#登录mysql
mysql -u root -p
ALTER USER 'root'@'localhost' IDENTIFIED BY 'Lzslov123!';

#添加远程登录用户
CREATE USER 'liaozesong'@'%' IDENTIFIED WITH mysql_native_password BY 'Lzslov123!';
GRANT ALL PRIVILEGES ON *.* TO 'liaozesong'@'%';