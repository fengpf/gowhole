

go run client/main.go 

go run server/main.go -port 50001 & > 50001.log
go run server/main.go -port 50002 & > 50002.log
go run server/main.go -port 50003 & > 50003.log



#查看etcd 链接数
# netstat -anltp | grep etcd | grep  ESTABLISHED | wc -l