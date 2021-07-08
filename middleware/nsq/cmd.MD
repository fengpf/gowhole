
`./nsqlookupd`

`./nsqadmin --lookupd-http-address=127.0.0.1:4161`



#开一个nsqd实例
`./nsqd --lookupd-tcp-address=127.0.0.1:4160`

#开两个nsqd实例
`nsqd --lookupd-tcp-address=127.0.0.1:4160 -tcp-address=0.0.0.0:4152 -http-address=0.0.0.0:4153`