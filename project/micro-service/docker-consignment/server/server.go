package main

import (
	"context"
	"log"
	"os"

	pb "gowhole/project/micro-service/docker-consignment/server/proto"

	"github.com/micro/go-micro"
)

//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment                                   // 获取仓库中所有的货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//
// 定义微服务
//
type service struct {
	repo Repository
}

//
// 实现 consignment.pb.go 中的 ShippingServiceHandler 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	// resp = &pb.Response{Created: true, Consignment: consignment}
	*resp = pb.Response{Created: true, Consignment: consignment}
	return nil
}

// 获取目前所有托运的货物
// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	resp.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	log.SetOutput(os.Stdout)

	server := micro.NewService(
		micro.Name("go.micro.srv.proto"),
		micro.Version("v1"),
	)

	// 解析命令行参数
	server.Init()
	pb.RegisterShippingServiceHandler(server.Server(), &service{Repository{}})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
