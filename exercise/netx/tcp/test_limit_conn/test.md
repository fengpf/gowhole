
###测试最大连接数

`netstat -antp | grep 8080 |grep ESTABLISHED | wc -l`


`ulimit -n`

`cat /proc/sys/net/core/somaxconn`

`cat /proc/sys/net/ipv4/ip_local_port_range`