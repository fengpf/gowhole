### 服务端与客户端版本对应关系：
#### 老版本
wurstmeister/kafka:2.11-0.11.0.3
github.com/Shopify/sarama v1.22.0
#### 新版本
wurstmeister/kafka:2.12-2.0.1
github.com/Shopify/sarama v1.29.1

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

### 进去容器操作kafka 命令

``` 1.进入docker：docker exec -it kakfa /bin/bash

2.进入docker中kafka的文件目录：cd opt/kafka_2.12-2.2.1

3.查看当前主题的信息：bin/kafka-topics.sh --zookeeper IP:2181 --describe --topic mykafka2
  bin/kafka-topics.sh --zookeeper zookeeper.concurrent_default:2181 --describe --topic test1

4.修改主题分区：bin/kafka-topics.sh --zookeeper IP:2181 -alter --partitions 4 --topic mykafka2

5.查看修改后的分区信息：bin/kafka-topics.sh --zookeeper IP:2181 --describe --topic mykafka2
 bin/kafka-topics.sh --zookeeper zookeeper.concurrent_default:2181 --describe --topic test1
 
6.退出虚拟环境：exit
```

### kafka查看topic和消息内容命令
```
 /opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```
### 创建topic 1个副本 3个分区
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-topics.sh --create --zookeeper zoo1:2181 --replication-factor 1 --partitions 3 --topic test
```

### 查看topic 信息
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-topics.sh --describe --zookeeper localhost:2181 --topic test 
```

### 查看kafka data log 文件
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-run-class.sh kafka.tools.DumpLogSegments --files ./00000000000000000000.log   --print-data-log
```
### 查看kafka index 文件

# /kafka/kafka-logs-e0af89e74b5b/test2-1 这个路径就是log.dirs，test2-1 是topic-partitionId

/opt/kafka_2.11-0.11.0.3/bin/kafka-run-class.sh kafka.tools.DumpLogSegments --files  ./00000000000000000000.index

### 删除topic
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-topics.sh --zookeeper zoo1:2181 --topic test --delete 
```
### 查看topic 列表
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-topics.sh --list --zookeeper zoo1:2181
```
### 增加topic的partition数 
/opt/kafka_2.11-0.11.0.3/bin/kafka-topics.sh --zookeeper localhost:2181 --alter --topic test --partitions 5 

### 生产消息 
/opt/kafka_2.11-0.11.0.3/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test 

### 消费消息 
 	1) 从头开始 
   /opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning 
 	2) 从尾部开始 
 	/opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --offset latest 
 	3) 指定分区 
 	/opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --offset latest --partition 1 
 	4) 取指定个数 
 	/opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --offset latest --partition 1 --max-messages 1 
 	5) 新消费者（ver>=0.9） 
 	/opt/kafka_2.11-0.11.0.3/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --new-consumer --from-beginning --consumer.config config/consumer.properties 
### 查看有哪些消费者Group 
```
/opt/kafka_2.11-0.11.0.3/bin/kafka-consumer-groups.sh --new-consumer --bootstrap-server 127.0.0.1:9092 --list 
```

### 进入容器
`docker exec -it de3dc9d04fe0 /bin/bash`
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

注意：10.23.9.38 为本地ip {ipconfig getifaddr en0指令的结果}

`kafkacat -b 10.23.9.38:9092,10.23.9.38:9093,10.23.9.38:9094  -P -t topic001`
`kafkacat -b 10.23.9.38:9092,10.23.9.38:9093,10.23.9.38:9094  -C -t topic001`


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

