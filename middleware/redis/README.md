# redis为啥是单线程, 为什么要设置连接池？

### 关于为何是单线程
   首先redis单线程指的是网络请求模块使用了一个线程（所以不需考虑并发安全性），其他模块仍用了多个线程，
   其次redis主要工作是基于内存的命令读写，所以说cpu不是redis的瓶颈，redis的瓶颈最有可能是机器内存或者网络带宽。
   总体来说快速的原因如下：
      1>绝大部分请求是纯粹的内存操作（非常快速）
      2>采用单线程,避免了不必要的上下文切换和竞争条件
      3>非阻塞IO

>1.启动redis服务

```shell

docker run -d -ti \
  -p 6380:6380 \
  -v $PWD/conf/redis.conf:/conf/redis.conf \
  -v $PWD/data:/data \
  --restart always \
  --name my_redis \
  redis:latest \
  redis-server /conf/redis.conf

   docker logs e9037df83dba
   docker exec -it e9037df83dba redis-cli -p 6380
   docker ps -a
   docker stop  e9037df83dba &&  docker rm  e9037df83dba
   
```
`./redis-server --loglevel debug`

>2.启用10个conn
`tcpkali -c 10 127.0.0.1:6380'`

### go 设置redis连接池

``` go

//不使用连接池，直接获取连接
func getSingleConn() (c redis.Conn) {
	var err error
	c, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return
}

//conn.close 连接直接释放
func (c *conn) Close() error {
	c.mu.Lock()
	err := c.err
	if c.err == nil {
		c.err = errors.New("redigo: closed")
		err = c.conn.Close()
	}
	c.mu.Unlock()
	return err
}


// 建立连接池
pool = &redis.Pool{
		MaxIdle:     5,               //定义连接池中最大连接数（超过这个数会关闭老的链接，总会保持这个数）
		MaxActive:   20,              //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 5 * time.Second, //定义链接的超时时间，每次p.Get()的时候会检测这个连接是否超时（超时会关闭，并释放可用连接数）
		Wait:        true,            // 当可用连接数为0是，那么当wait=true,那么当调用p.Get()时，会阻塞等待，否则，返回nil.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr,
				redis.DialDatabase(db),  //默认是索引0，可以自己指定
				// redis.DialPassword(pwd), //默认是空，如果不设置密码则dial不通
			  redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second),
				redis.DialConnectTimeout(200*time.Millisecond))
			if err != nil {
				c.Close()
				log.Fatalf("redis.Dial error(%v)\n", err)
				return nil, err
			}
			if _, err := c.Do("AUTH", pwd); err != nil {
				c.Close()
				log.Fatalf("c.Do auth error(%v)\n", err)
				return nil, err
			}
      if _, err := c.Do("SELECT", db); err != nil { //选择db
				c.Close()
				return nil, err
			}
			return c, err
		},
		// 如果设置了给func,那么每次p.Get()的时候都会调用改方法来验证连接的可用性
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
  }
  
  //conn 关闭，回收到连接池
  func (ac *activeConn) Close() error {
      pc := ac.pc
      if pc == nil {
        return nil
      }
      ac.pc = nil

      if ac.state&internal.MultiState != 0 {
        pc.c.Send("DISCARD")
        ac.state &^= (internal.MultiState | internal.WatchState)
      } else if ac.state&internal.WatchState != 0 {
        pc.c.Send("UNWATCH")
        ac.state &^= internal.WatchState
      }
      if ac.state&internal.SubscribeState != 0 {
        pc.c.Send("UNSUBSCRIBE")
        pc.c.Send("PUNSUBSCRIBE")
        // To detect the end of the message stream, ask the server to echo
        // a sentinel value and read until we see that value.
        sentinelOnce.Do(initSentinel)
        pc.c.Send("ECHO", sentinel)
        pc.c.Flush()
        for {
          p, err := pc.c.Receive()
          if err != nil {
            break
          }
          if p, ok := p.([]byte); ok && bytes.Equal(p, sentinel) {
            ac.state &^= internal.SubscribeState
            break
          }
        }
      }
      pc.c.Do("")
      ac.p.put(pc, ac.state != 0 || pc.c.Err() != nil)
      return nil
}
```

> 1.不设置连接池  
如果启动5个连接，5个连接用完则释放，如下所示

``` shell
    3709:M 13 Dec 08:27:05.143 - Accepted 127.0.0.1:36552
    3709:M 13 Dec 08:27:05.143 - Accepted 127.0.0.1:36550
    3709:M 13 Dec 08:27:05.143 - Accepted 127.0.0.1:36554
    3709:M 13 Dec 08:27:05.144 - Accepted 127.0.0.1:36556
    3709:M 13 Dec 08:27:05.144 - Accepted 127.0.0.1:36558
    3709:M 13 Dec 08:27:05.144 - Client closed connection
    3709:M 13 Dec 08:27:05.144 - Client closed connection
    3709:M 13 Dec 08:27:05.144 - Client closed connection
    3709:M 13 Dec 08:27:05.145 - Client closed connection
    3709:M 13 Dec 08:27:05.145 - Client closed connection
    3709:M 13 Dec 08:31:26.134 - DB 0: 5 keys (0 volatile) in 8 slots HT.
    3709:M 13 Dec 08:31:26.135 - 0 clients connected (0 slaves), 831408 bytes in use
    3709:M 13 Dec 08:31:31.206 - DB 0: 5 keys (0 volatile) in 8 slots HT.
    3709:M 13 Dec 08:31:31.207 - 0 clients connected (0 slaves), 831408 bytes in use
    3709:M 13 Dec 08:31:36.287 - DB 0: 5 keys (0 volatile) in 8 slots HT.

```
   如果启用10000个连接，则会出现服务端连接枯竭，
   客户端得到：panic: dial tcp 127.0.0.1:6380: socket: too many open files
   原因是每次操作完成连接会被关闭，如果频繁操作会造成连接泄露

> 2.设置连接池  
  如果启动5个连接，5个连接用完不会被释放，如下所示

``` shell
    3709:M 13 Dec 08:54:16.342 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:54:16.342 - 0 clients connected (0 slaves), 2400664 bytes in use
    3709:M 13 Dec 08:54:17.231 - Accepted 127.0.0.1:41172
    3709:M 13 Dec 08:54:17.232 - Accepted 127.0.0.1:41170
    3709:M 13 Dec 08:54:17.232 - Accepted 127.0.0.1:41174
    3709:M 13 Dec 08:54:17.232 - Accepted 127.0.0.1:41176
    3709:M 13 Dec 08:54:17.233 - Accepted 127.0.0.1:41178
    3709:M 13 Dec 08:54:21.390 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:54:21.390 - 5 clients connected (0 slaves), 2487184 bytes in use
    3709:M 13 Dec 08:54:26.447 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:54:26.447 - 5 clients connected (0 slaves), 2487184 bytes in use
    3709:M 13 Dec 08:54:31.498 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:54:31.498 - 5 clients connected (0 slaves), 2487184 bytes in use
    3709:M 13 Dec 08:54:36.547 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:54:36.547 - 5 clients connected (0 slaves), 2487184 bytes in use

  #进程退出连接关闭
    3709:M 13 Dec 08:56:01.067 - Client closed connection
    3709:M 13 Dec 08:56:01.068 - Client closed connection
    3709:M 13 Dec 08:56:01.068 - Client closed connection
    3709:M 13 Dec 08:56:01.068 - Client closed connection
    3709:M 13 Dec 08:56:01.069 - Client closed connection

```

   如果服务端配置100个最大连接数，活跃连接数为20， 启用10000个客户端连接的时候，建立了20个连接

   ```shell 
    3709:M 13 Dec 08:59:33.186 - Accepted 127.0.0.1:41190
    3709:M 13 Dec 08:59:33.187 - Accepted 127.0.0.1:41192
    3709:M 13 Dec 08:59:33.187 - Accepted 127.0.0.1:41194
    3709:M 13 Dec 08:59:33.187 - Accepted 127.0.0.1:41196
    3709:M 13 Dec 08:59:33.187 - Accepted 127.0.0.1:41198
    3709:M 13 Dec 08:59:33.187 - Accepted 127.0.0.1:41200
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41202
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41204
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41206
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41208
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41210
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41212
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41214
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41216
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41218
    3709:M 13 Dec 08:59:33.188 - Accepted 127.0.0.1:41220
    3709:M 13 Dec 08:59:33.189 - Accepted 127.0.0.1:41222
    3709:M 13 Dec 08:59:33.189 - Accepted 127.0.0.1:41224
    3709:M 13 Dec 08:59:33.189 - Accepted 127.0.0.1:41226
    3709:M 13 Dec 08:59:33.189 - Accepted 127.0.0.1:41228
    3709:M 13 Dec 08:59:33.369 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:59:33.372 - 20 clients connected (0 slaves), 3402104 bytes in use
    3709:M 13 Dec 08:59:38.437 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
    3709:M 13 Dec 08:59:38.437 - 20 clients connected (0 slaves), 2746744 bytes in use
    3709:M 13 Dec 08:59:43.483 - DB 0: 10000 keys (0 volatile) in 16384 slots HT.
```


客户端依然可以正确完成请求
```shell 
    set success {"mid":9985,"from":0,"is_author":1,"timestamp":1544691573}
    set success {"mid":9986,"from":0,"is_author":2,"timestamp":1544691573}
    set success {"mid":9987,"from":0,"is_author":0,"timestamp":1544691573}
    set success {"mid":9988,"from":0,"is_author":1,"timestamp":1544691573}
    set success {"mid":9989,"from":0,"is_author":2,"timestamp":1544691573}
    set success {"mid":9990,"from":0,"is_author":0,"timestamp":1544691573}
    set success {"mid":9991,"from":0,"is_author":1,"timestamp":1544691573}
    set success {"mid":9992,"from":0,"is_author":2,"timestamp":1544691573}
    set success {"mid":9993,"from":0,"is_author":0,"timestamp":1544691573}
    set success {"mid":9994,"from":0,"is_author":1,"timestamp":1544691573}
    set success {"mid":9995,"from":0,"is_author":2,"timestamp":1544691573}
    set success {"mid":9996,"from":0,"is_author":0,"timestamp":1544691573}
    set success {"mid":9997,"from":0,"is_author":1,"timestamp":1544691573}
    set success {"mid":9998,"from":0,"is_author":2,"timestamp":1544691573}

```


