### docker hub
https://hub.docker.com/_/influxdb/

### docker 启动

`docker run --name=influxdb -d -p 8087:8087 influxdb:latest`

```

docker run -p 8087:8087 \
      -e INFLUXDB_DB=defaultdb \
      -e INFLUXDB_ADMIN_USER=admin \
      -e INFLUXDB_ADMIN_PASSWORD=admin \
      -e INFLUXDB_USER=user \
      -e INFLUXDB_USER_PASSWORD=user \
      -v influxdb:/var/lib/influxdb \
      influxdb:latest 

```

### influx
`docker exec -it c6111adcee24 influx`

### 显示数据库
`show databases`

### 新建数据库
`create database defaultdb`

### 删除数据库
`drop database defaultdb`

### 使用某个数据库
`use defaultdb`

### 显示所有表(在influxdb当中，并没有表（table）这个概念，取而代之的是MEASUREMENTS，MEASUREMENTS的功能与传统数据库中的表一致，因此我们也可以将MEASUREMENTS称为influxdb中的表)
`show measurements`

### 新建表influxdb中没有显式的新建表的语句，只能通过insert数据的方式来建立新表。
### task_user是表名,index=mid是tag，属于索引，value=111是field，这个可以随意写，随意定义。
`insert task_user,mid=123 state=1 1558610260`

### 查询表
```sql 
> select * from task_user
name: task_user
time       index mid state value
----       ----- --- ----- -----
1558608768 mid             111
1558610260       123 1
```

### 删除表
`drop measurement task_user`

### 修改和删除数据
influxdb属于时序数据库，没有提供修改和删除数据的方法。但是删除可以通过influxdb的数据保存策略（Retention Policies）来实现

### series操作 series表示这个表里面的数据，可以在图表上画成几条线，series主要通过tags排列组合算出来。
`show series from mem`

### http
curl -GET 'http://localhost:8086/query?pretty=true' --data-urlencode "db=defaultdb" --data-urlencode "q=SELECT value FROM task_user WHERE mid='123'"