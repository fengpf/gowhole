# protoc 编译器的 grpc 插件会处理 service 字段定义的 UserInfoService
# 使 service 能编码、解码 message
$ protoc -I . --go_out=plugins=grpc:. ./user.proto


go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc

go get -u github.com/micro/protobuf/proto
go get -u github.com/micro/protobuf/protoc-gen-go

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./proto/consignment.proto


protoc -I. --go_out=plugins=micro:$GOPATH/src/gowhole/project/micro-service/docker-consignment/server ./proto/consignment.proto


$ make run
# 在 Docker alpine 容器的 50001 端口上运行 consignment-service 服务
# 可添加 -d 参数将微服务放到后台运行
docker run -p 50051:50051 \
         -e MICRO_SERVER_ADDRESS=:50051 \
         -e MICRO_REGISTRY=mdns \
         consignment-service
2019/01/08 08:32:44 Listening on [::]:50051
2019/01/08 08:32:44 Broker Listening on [::]:35171
2019/01/08 08:32:44 Registering node: proto-f8adfb57-131f-11e9-9aac-0242ac110004


$ MICRO_REGISTRY=mdns  go run server.go
2019/01/08 16:38:19 Listening on [::]:64680
2019/01/08 16:38:19 Broker Listening on [::]:64681
2019/01/08 16:38:19 Registering node: proto-bfe03a7a-1320-11e9-8482-a860b63b3121