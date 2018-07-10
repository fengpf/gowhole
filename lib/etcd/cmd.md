
#### 1.1解压安装

##### 解压
tar zxvf etcd-v3.2.11-linux-amd64.tar.gz 
cd etcd-v3.2.11-linux-amd64

##### ETCD版本
etcd --version

##### 客户端接口版本
etcdctl --version

##### API3的要这样
ETCDCTL_API=3 etcdctl version

##### 启动也很简单
./etcd

##### 试试
`ETCDCTL_API=3 etcdctl --endpoints=localhost:2379 put foo bar`
`ETCDCTL_API=3 etcdctl --endpoints=localhost:2379 get foo`

### 1.2 源码安装
`go get -u -v https://github.com/coreos/etcd`
`./build`
启动:
`./etcd`
### 1.3 docker安装
拉镜像:
`docker pull quay.io/coreos/etcd`
启动:
`docker run -it --rm -p 2379:2379 -p 2380:2380 --name etcd quay.io/coreos/etcd`
查询:
`docker exec -it etcd etcdctl member list`

### 启动详细说明

#### 2.1单机启动
`./etcd --name my-etcd-1  --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380  --initial-cluster my-etcd-1=http://0.0.0.0:2380`

#### 2.2集群启动
     我们用三个端口2380,2381,2382来模拟集群(这三个是成员之间通信),2379,2389,2399是给客户端连接的.公网IP是: 172.16.13.90, 如果在本机模拟集群, 可以将172.16.13.90改为0.0.0.0。
     带advertise参数是广播参数: 如--listen-client-urls和--advertise-client-urls, 前者是Etcd端监听客户端的url,后者是Etcd客户端请求的url, 
     两者端口是相同的, 只不过后者一般为公网IP, 暴露给外部使用.

```

./etcd --name my-etcd-1
       --listen-client-urls http://0.0.0.0:2379
       --advertise-client-urls http://172.16.13.90:2379
       --listen-peer-urls http://0.0.0.0:2380
       --initial-advertise-peer-urls http://172.16.13.90:2380  
       --initial-cluster-token etcd-cluster-test 
       --initial-cluster-state new 
       --initial-cluster my-etcd-1=http://172.16.13.90:2380,my-etcd-2=http://172.16.13.90:2381,my-etcd-3=http://172.16.13.90:2382


./etcd --name my-etcd-2  
       --listen-client-urls http://0.0.0.0:2389 
       --advertise-client-urls http://172.16.13.90:2389 
       --listen-peer-urls http://0.0.0.0:2381 
       --initial-advertise-peer-urls http://172.16.13.90:2381  
       --initial-cluster-token etcd-cluster-test 
       --initial-cluster-state new 
       --initial-cluster my-etcd-1=http://172.16.13.90:2380,my-etcd-2=http://172.16.13.90:2381,my-etcd-3=http://172.16.13.90:2382

./etcd --name my-etcd-3  
       --listen-client-urls http://0.0.0.0:2399 
       --advertise-client-urls http://172.16.13.90:2399 
       --listen-peer-urls http://0.0.0.0:2382 
       --initial-advertise-peer-urls http://172.16.13.90:2382  
       --initial-cluster-token etcd-cluster-test 
       --initial-cluster-state new --initial-cluster my-etcd-1=http://172.16.13.90:2380,my-etcd-2=http://172.16.13.90:2381,my-etcd-3=http://172.16.13.90:2382

```

查看成员:

`etcdctl member list`
使用时需要指定endpoints(默认本地端口2379), 集群时数据会迅速同步:

`ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2389 put foo xx`
`ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get foo`

### 2.3参数说明

| 参数	| 使用说明 |
|:-----:|:-----:|
| --name etcd0	| 本member的名字
|--initial-advertise-peer-urls http://192.168.2.55:2380 | 其他member使用，其他member通过该地址与本member交互信息.一定要保证从其他member能可访问该地址。静态配置方式下，该参数的value一定要同时在 --initial-cluster参数中存在。memberID的生成受--initial-cluster-token和--initial-advertise-peer-urls影响。 
|--listen-peer-urls http://0.0.0.0:2380	|  本member侧使用，用于监听其他member发送信息的地址。ip为全0代表监听本member侧所有接口
|--listen-client-urls http://0.0.0.0:2379 | 本member侧使用，用于监听etcd客户发送信息的地址。ip为全0代表监听本member侧所有接口
|--advertise-client-urls http://192.168.2.55:2379	| etcd客户使用，客户通过该地址与本member交互信息。一定要保证从客户侧能可访问该地址
|--initial-cluster-token etcd-cluster-2	| 用于区分不同集群。本地如有多个集群要设为不同。
|--initial-cluster etcd0=http://192.168.2.55:2380,etcd1=http://192.168.2.54:2380,etcd2=http://192.168.2.56:2380	| 本member侧使用。描述集群中所有节点的信息，本member根据此信息去联系其他member。memberID的生成受--initial-cluster-token和--initial-advertise-peer-urls影响。
| --initial-cluster-state new | 	用于指示本次是否为新建集群。有两个取值new和existing。如果填为existing，则该member启动时会尝试与其他member交互。集群初次建立时，要填为new，经尝试最后一个节点填existing也正常，其他节点不能填为existing。集群运行过程中，一个member故障后恢复时填为existing，经尝试填为new也正常。
| -data-dir	|  指定节点的数据存储目录，这些数据包括节点ID，集群ID，集群初始化配置，Snapshot文件，若未指定-wal-dir，还会存储WAL文件；如果不指定会用缺省目录。
| -discovery http://192.168.1.163:20003/v2/keys/discovery/78b12ad7-2c1d-40db-9416-3727baf686cb	|  用于自发现模式下，指定第三方etcd上key地址，要建立的集群各member都会向其注册自己的地址。

三.使用详细说明
ETCD API有两种, 一种是3, 一种是2, 默认为2, 我们主要用3:

API3:

```

root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=3 etcdctl put mykey "this is awesome"
 OK
root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=3 etcdctl get mykey
mykey
this is awesome

```

API2是这样的(可以不加前缀)

```

root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=2 etcdctl set /local/dd d
d
root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=2 etcdctl get /local/dd
d
root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=2 etcdctl set /local/dd d
d
root@ubuntu-xenial:~/code/src/github.com/coreos/etcd/bin$ ETCDCTL_API=2 etcdctl get /local/dd
d

```