/*
Command example-grpc-server is an example grpc server
to be called by example-gateway-server.
*/
package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"

	"gowhole/middleware/grpc/gw/internal/server"
)

var (
	addr    = flag.String("addr", ":9999", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	ctx := context.Background()

	fmt.Println("start grpc server")

	if err := server.Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}

}
