  
### 使用默认文件启动
`docker-compose up -d`

### borker总数从1个扩展到4个，新启动的broker端口号随机
`docker-compose --scale kafka=4`

### 指定compose 文件启动和关闭
`docker-compose -f docker-compose-multi.yml up -d`

`docker-compose -f docker-compose-multi.yml stop`


### 查看启动的一个zookeeper和一个kafka容器：

`docker ps`

``` shell 
➜  concurrent git:(master) ✗ docker ps -a
CONTAINER ID        IMAGE                              COMMAND                  CREATED             STATUS                      PORTS                                                NAMES
696024154ff0        wurstmeister/kafka:2.11-0.11.0.3   "start-kafka.sh"         5 days ago          Up 4 days                   0.0.0.0:9094->9092/tcp                               kafka3
c35a5a8255c8        wurstmeister/kafka:2.11-0.11.0.3   "start-kafka.sh"         5 days ago          Up 4 days                   0.0.0.0:9093->9092/tcp                               kafka2
87be44627f17        wurstmeister/kafka:2.11-0.11.0.3   "start-kafka.sh"         5 days ago          Up 4 days                   0.0.0.0:9092->9092/tcp                               kafka1
9e0628d2b081        wurstmeister/zookeeper             "/bin/sh -c '/usr/sb…"   5 days ago          Up 4 days                   22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp   zookeeper

```

### 查看容器中的kafka版本号
`docker exec kafka1 find / -name \*kafka_\* | head -1 | grep -o '\kafka[^\n]*'`
```shell 
➜  concurrent git:(master) ✗ docker exec kafka1 find / -name \*kafka_\* | head -1 | grep -o '\kafka[^\n]*'
kafka_2.11-0.11.0.3
```

### 查看zookeeper版本：
`docker exec zookeeper pwd`
```shell
➜  concurrent git:(master) ✗ docker exec zookeeper pwd
/opt/zookeeper-3.4.9

```

### 创建一个topic，名为topic001，4个partition，副本因子2

```
docker exec kafka1 \
      kafka-topics.sh \
      --create --topic topic002 \
      --partitions 4 \
      --zookeeper zookeeper:2181 \
      --replication-factor 2

```

### 查看刚刚创建的topic
``` 
docker exec kafka1 \
   kafka-topics.sh --list \
   --zookeeper zookeeper:2181 \
   topic002

```

### 查看刚刚创建的topic的情况，borker和副本情况
```
docker exec kafka1  \
   kafka-topics.sh \
   --describe \
   --topic topic002 \
   --zookeeper zookeeper:2181

```

### 生产消息
-d :分离模式: 在后台运行
-i :即使没有附加也保持STDIN 打开
-t :分配一个伪终端

```
docker exec -it kafka1 \
   kafka-console-producer.sh \
   --topic topic002 \
   --broker-list kafka1:9092,kafka2:9092,kafka3:9092

```

### 消费消息
```
docker exec kafka1 \
   kafka-console-consumer.sh \
   --topic topic002 \
   --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092

```



#关闭kafka和zk集群
`docker-compose stop`

### 增加kafkacat 命令，主机里面发消息，接收消息，多个broker逗号隔开

注意：10.23.39.129 为本地ip {ipconfig getifaddr en0指令的结果}

`kafkacat -b 10.23.39.129:9092,10.23.39.129:9093,10.23.39.129:9094  -P -t topic001`
`kafkacat -b 10.23.39.129:9092,10.23.39.129:9093,10.23.39.129:9094  -C -t topic001`


###参考：

https://github.com/wurstmeister/kafka-docker/wiki/Connectivity
https://blog.csdn.net/boling_cavalry/article/details/85395080
https://www.jianshu.com/p/ac03f126980e


（1）使用默认的 ‘local’ driver 创建一个 volume
 docker volume create --name vol1
 docker volume inspect vol1

（2）使用这个 volume
 docker run -d -P --name web4 -v vol1:/volume training/webapp python app.p

（3）删除这个 volume
    可以使用 docker rm -v 命令在删除容器时删除该容器的卷
    docker run -d -P --name web5 -v /webapp training/webapp python app.py
    批量删除孤单 volumes
    从上面的介绍可以看出，使用 docker run -v 启动的容器被删除以后，在主机上会遗留下来孤单的卷。可以使用下面的简单方法来做清理：
    docker volume ls -qf dangling=true

