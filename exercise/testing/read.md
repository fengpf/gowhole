
### // 标志可以用来指定编译和测试的并发数

```

go test -p 1 ./...  

```

###    //将会执行 3 次，其中 GOMAXPROCS 值分别为 1，2，和 4

GOMAXPROCS=4 go test ./...

###  //-cpu 标志，搭配数据竞争的探测标志 -race

go test -cpu=1,2,4 ./... 

go test -cpu=1,2 -race ./...