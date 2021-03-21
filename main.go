package main

import (
	"context"
	"sync"

	pb "github.com/JonasMuylaert/shippy-service-consignment/proto/consignment"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

const (
	PORT = ":50051"
)

type shippingServiceServer struct {
	repo repository
}

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

func (s *shippingServiceServer) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *shippingServiceServer) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func newServer() *shippingServiceServer {
	s := &shippingServiceServer{
		repo: &Repository{},
	}
	return s
}

func main() {
	//register service with micro
	srv := service.New(
		service.Name("shippy.service.consignment"),
	)

	srv.Init()

	//register handler
	if err := srv.Handle(newServer()); err != nil {
		logger.Fatalf("failed creating a logger: %v", err)
	}

	//run server
	if err := srv.Run(); err != nil {
		logger.Fatalf("failed starting server: %v", err)
	}
}
