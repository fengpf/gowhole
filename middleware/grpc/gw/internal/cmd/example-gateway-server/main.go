/*
Command example-gateway-server is an example reverse-proxy implementation
whose HTTP handler is generated by grpc-gateway.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"gowhole/middleware/grpc/gw/internal/gateway"

	"github.com/golang/glog"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9999", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	swaggerDir = flag.String("swagger_dir", "examples/internal/proto/examplepb", "path to the directory which contains swagger definitions")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		SwaggerDir: *swaggerDir,
	}

	fmt.Printf("start grpc geteway server addr(%v)\n", opts.Addr)

	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
