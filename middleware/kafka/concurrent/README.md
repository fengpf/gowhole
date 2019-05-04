  
### 指定compose 文件启动和关闭
`docker-compose -f docker-compose-multi.yml up -d`
`docker-compose -f docker-compose-multi.yml stop`

### 使用docker-compose 启动

`docker-compose up -d --scale kafka=4`

``` shell ➜  concurrent git:(master) ✗ docker-compose up -d
Creating network "concurrent_default" with the default driver
Pulling kafka (wurstmeister/kafka:2.11-0.11.0.3)...
2.11-0.11.0.3: Pulling from wurstmeister/kafka
bdf0201b3a05: Already exists
9e12771959ad: Already exists
ae272eb2814b: Already exists
f059dafc9b73: Pull complete
a6e2c9da29de: Pull complete
396a7b97c1dc: Pull complete
Digest: sha256:df6dbfde3828a7ec7b3ecfc73241d22fd3ba4f76f1b4c6417a9cf8d101ea14b0
Status: Downloaded newer image for wurstmeister/kafka:2.11-0.11.0.3
Creating concurrent_kafka_1     ... done
Creating concurrent_zookeeper_1 ... done

```

### 查看启动的一个zookeeper和一个kafka容器：

`docker ps`

``` shell 
➜  concurrent git:(master) ✗ docker ps
CONTAINER ID        IMAGE                              COMMAND                  CREATED             STATUS              PORTS                                                NAMES
6c4edb53b69f        wurstmeister/zookeeper             "/bin/sh -c '/usr/sb…"   9 seconds ago       Up 15 seconds       22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp   concurrent_zookeeper_1
194e3af8463a        wurstmeister/kafka:2.11-0.11.0.3   "start-kafka.sh"         9 seconds ago       Up 14 seconds       0.0.0.0:32768->9092/tcp                              concurrent_kafka_1
```

### 查看容器中的kafka版本号
`docker exec concurrent_kafka_1 find / -name \*kafka_\* | head -1 | grep -o '\kafka[^\n]*'`
```shell 
➜  concurrent git:(master) ✗ docker exec concurrent_kafka_1 find / -name \*kafka_\* | head -1 | grep -o '\kafka[^\n]*'
  kafka_2.11-0.11.0.3
```

### 查看zookeeper版本：
`docker exec concurrent_zookeeper_1 pwd`
```shell
➜  concurrent git:(master) ✗ docker exec concurrent_zookeeper_1 pwd
/opt/zookeeper-3.4.9
```

### borker总数从1个扩展到4个
`docker-compose --scale kafka=4`

```shell
➜  concurrent git:(master) ✗ docker-compose scale kafka=4
WARNING: The scale command is deprecated. Use the up command with the --scale flag instead.
Starting concurrent_kafka_1 ... done
Creating concurrent_kafka_2 ... done
Creating concurrent_kafka_3 ... done
Creating concurrent_kafka_4 ... done
```

### 创建一个topic，名为topic001，4个partition，副本因子2

```
docker exec concurrent_kafka_1 \
      kafka-topics.sh \
      --create --topic topic001 \
      --partitions 4 \
      --zookeeper zookeeper:2181 \
      --replication-factor 2

```

### 查看刚刚创建的topic
``` 
docker exec concurrent_kafka_1 \
   kafka-topics.sh --list \
   --zookeeper zookeeper:2181 \
   topic001

```

### 查看刚刚创建的topic的情况，borker和副本情况
```
docker exec concurrent_kafka_1 \
   kafka-topics.sh \
   --describe \
   --topic topic001 \
   --zookeeper zookeeper:2181
```

### 消费消息
```
docker exec concurrent_kafka_1\
   kafka-console-consumer.sh \
   --topic topic001 \
   --bootstrap-server concurrent_kafka_1:9092,concurrent_kafka_2:9092,concurrent_kafka_3:9092,concurrent_kafka_4:9092
```

### 生产消息

```
docker exec -it concurrent_kafka_1 \
   kafka-console-producer.sh \
   --topic topic001 \
   --broker-list concurrent_kafka_1:9092,concurrent_kafka_2:9092,concurrent_kafka_3:9092,concurrent_kafka_4:9092

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

