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

`ab -c 100 -n 100 'http://localhost:9090/mysql_max_conns'`

### go设置mysql连接池

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





