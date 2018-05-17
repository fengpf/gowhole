package main

import (
	"log"

	"github.com/tabalt/ipqueryd/src/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12101", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	iqc := pb.NewIpQueryClient(conn)
	ip := "1.1.8.1"
	r, err := iqc.Find(context.Background(), &pb.IpFindRequest{Ip: ip})
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	log.Printf("ip data: %s", r.Data)
}
