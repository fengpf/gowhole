  
  
  使用docker-compose 启动集群
# docker-compose up -d

 启动1或者多个kafka 节点，比如3个
# docker-compose -f docker-compose-single-broker.yml up -d

# docker-compose up kafka=3


  使用docker 启动

1、启动zookeeper
# docker run -d --name zookeeper -p 2181:2181 -t wurstmeister/zookeeper

2、启动kafka

docker run --rm --name kafka --link zookeeper:zk \
   -e KAFKA_PORT=9092 \
   -e LANG="en_US.UTF-8" \
   -e KAFKA_ADVERTISED_HOST_NAME=10.23.39.129 \
   -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
   -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
   -p 9092:9092 \
   -t wurstmeister/kafka

3、进入容器
   执行docker ps，找到kafka的CONTAINER ID，进入容器内部。
# docker exec -it ${CONTAINER ID} /bin/bash 


4、 启动zookeeper服务
   kafka自带了zookeeper，可以直接使用它建立一个单节点的zookeeper 实例。也可以自己安装配置zookeeper
 #  bin/zookeeper-server-start.sh config/zookeeper.properties

5、 启动kafka server
 #  bin/kafka-server-start.sh config/server.properties

6、 创建topic
    topic 名字为test，1个分块和1个副本
# $KAFKA_HOME/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic test

7、 查看topic列表
# $KAFKA_HOME/bin/kafka-topics.sh --list --zookeeper zookeeper:2181

8、 查看topic 详细信息
# $KAFKA_HOME/bin/kafka-topics.sh --zookeeper zookeeper:2181 --describe --topic test

9、producer发送消息
   kafka的命令行客户端可以将文件或标准输入作为消息来发送到kafka集群。默认情况下每一行都是独立一个消息。发送消息时要运行producer。按ctrl+c退出消息发送。
# $KAFKA_HOME/bin/kafka-console-producer.sh --topic test --broker-list 5d9a407caf6d:9092 

10、consumer接收消息
consumer订阅了topic test就可以接收上面的消息。命令行运行consumer将接收到的消息显示在终端：
# $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server 5d9a407caf6d:9092 --from-beginning --topic test 

11、增加kafkacat 命令，主机里面发消息，接收消息
本地
# kafkacat -b 10.23.39.129:9092 -P -t test
# kafkacat -b 10.23.39.129:9092 -C -t test

服务器
# kafkacat -b 10.163.6.225:32784 -P -t test
# kafkacat -b 10.163.6.225:32784 -C -t test

# kafkacat -b 10.163.6.225:32773,10.163.6.225:32774,10.163.6.225:32775 -P -t test
# kafkacat -b 10.163.6.225:32773,10.163.6.225:32774,10.163.6.225:32775 -C -t test



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

