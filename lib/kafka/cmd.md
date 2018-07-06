  
  
  使用docker-compose 启动集群
# docker-compose up -d

 启动1或者多个kafka 节点，比如3个
# docker-compose -f docker-compose-single-broker.yml up -d

# docker-compose up kafka=3


  使用docker 启动

1、启动zookeeper
# docker run -d --name zookeeper -p 2181 -t wurstmeister/zookeeper

2、启动kafka
# docker run --name kafka -e HOST_IP=localhost -e KAFKA_ADVERTISED_PORT=9092 -e KAFKA_BROKER_ID=1 -e ZK=zk -p 9092 --link zookeeper:zk -t wurstmeister/kafka

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
# $KAFKA_HOME/bin/kafka-console-producer.sh --topic test --broker-list a938097ce7c4:9092 

10、consumer接收消息
consumer订阅了topic test就可以接收上面的消息。命令行运行consumer将接收到的消息显示在终端：
# $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server a938097ce7c4:9092 --from-beginning --topic test 




docker run -d --name kafka --publish 9092:9092 \
--link zookeeper \
--env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
--env KAFKA_ADVERTISED_HOST_NAME=localhost\
--env KAFKA_ADVERTISED_PORT=9092 \
--volume /etc/localtime:/etc/localtime \
wurstmeister/kafka:latest
