
```table下面是一些常用配置选项的说明：

--name：方便理解的节点名称，默认为 default，在集群中应该保持唯一
--data-dir：服务运行数据保存的路径，默认为 ${name}.etcd
--snapshot-count：指定有多少事务（transaction）被提交时，触发截取快照保存到磁盘
--heartbeat-interval：leader 多久发送一次心跳到 followers。默认值是 100ms
--eletion-timeout：重新投票的超时时间，如果follower在该时间间隔没有收到心跳包，会触发重新投票，默认为 1000 ms
--listen-peer-urls：和同伴通信的地址，比如 http://ip:2380，如果有多个，使用逗号分隔。需要所有节点都能够访问，所以不要使用 localhost
--advertise-client-urls：对外公告的该节点客户端监听地址，这个值会告诉集群中其他节点
--listen-client-urls：对外提供服务的地址：比如 http://ip:2379,http://127.0.0.1:2379，客户端会连接到这里和etcd交互
--initial-advertise-peer-urls：该节点同伴监听地址，这个值会告诉集群中其他节点
--initial-cluster：集群中所有节点的信息，格式为 node1=http://ip1:2380,node2=http://ip2:2380,…。需要注意的是，这里的 node1 是节点的--name指定的名字；后面的ip1:2380 是--initial-advertise-peer-urls 指定的值
--initial-cluster-state：新建集群的时候，这个值为 new；假如已经存在的集群，这个值为existing
--initial-cluster-token：创建集群的token，这个值每个集群保持唯一。这样的话，如果你要重新创建集群，即使配置和之前一样，也会再次生成新的集群和节点 uuid；否则会导致多个集群之间的冲突，造成未知的错误
```
　　所有以--init开头的配置都是在第一次启动etcd集群的时候才会用到，后续节点的重启会被忽略，如--initial-cluseter参数。所以当成功初始化了一个etcd集群以后，就不再需要这个参数或环境变量了。
　　如果服务已经运行过就要把修改 --initial-cluster-state 为existing
　　启用服务

### 我使用的etcd的api为v3版本。在使用命令时需要在前面加上ETCDCTL_API=3 
　如：查看集群成员
```
ETCDCTL_API=3 etcdctl member list
```

###查看集群状态
```
ETCDCTL_API=3  ./etcdctl  --endpoints 127.0.0.1:12380,127.0.0.1:22380,127.0.0.1:32380 endpoint status  --write-out="table" 
+-----------------+------------------+-----------+---------+-----------+------------+-----------+------------+--------------------+--------+
|    ENDPOINT     |        ID        |  VERSION  | DB SIZE | IS LEADER | IS LEARNER | RAFT TERM | RAFT INDEX | RAFT APPLIED INDEX | ERRORS |
+-----------------+------------------+-----------+---------+-----------+------------+-----------+------------+--------------------+--------+
| 127.0.0.1:12380 | 8211f1d0f64f3269 | 3.5.0-pre |   25 kB |      true |      false |         3 |         23 |                 23 |        |
| 127.0.0.1:22380 | 91bc3c398fb3c146 | 3.5.0-pre |   25 kB |     false |      false |         3 |         23 |                 23 |        |
| 127.0.0.1:32380 | fd422379fda50e48 | 3.5.0-pre |   25 kB |     false |      false |         3 |         23 |                 23 |        |
+-----------------+------------------+-----------+---------+-----------+------------+-----------+------------+--------------------+--------+
```

### 操作数据
   使用put和get命令可以保存和得到数据, del删除数据, 根据前缀查询
   
```
 ETCDCTL_API=3 etcdctl put test1 a
 ETCDCTL_API=3 etcdctl put test2 b
 ETCDCTL_API=3 etcdctl put test3 c
 ETCDCTL_API=3 etcdctl get --prefix test
```

###查询所有数据
 ```
   ETCDCTL_API=3 etcdctl get --from-key ""
 ```
 
 ### watch 监听
   watch 会监听key的变动 有变动时会在输出。这也正是服务发现需要使用的。
   我们监听 test键，然后对test执行修改和删除操作
```
ETCDCTL_API=3 etcdctl watch test
```

### lead 租约
  etcd可以为key设置超时时间，但与redis不同，etcd需要先创建lease，然后使用put命令加上参数–lease=<lease ID>
```
ETCDCTL_API=3 lease grant  ttl    创建lease，返回lease ID ttl秒
ETCDCTL_API=3 lease revoke  leaseId  删除lease，并删除所有关联的key
ETCDCTL_API=3 lease timetolive leaseId 取得lease的总时间和剩余时间
ETCDCTL_API=3 lease keep-alive leaseId     keep-alive会不间断的刷新lease时间，从而保证lease不会过期。
```

### 分布式锁
 使用lock命令后加锁名称 做分布式锁，如果没有显示释放锁，其他地方只能等待。
```
etcdctl --endpoints=$ENDPOINTS lock mutex1
# 在另一个终端输入
etcdctl --endpoints=$ENDPOINTS lock mutex1
```


▾ store/
  ▸ etcd1/data/
  ▸ etcd2/data/
  ▸ etcd3/data/
  docker-compose.yml
store/目录下的etcd1/data/, etcd2/data/和etcd1/data/用于存放存储数据, 避免docker重启之后数据丢失.

docker-compose-2.yml的内容如下:
