package handler

import (
	"context"
	"sync"

	pb "github.com/JonasMuylaert/shippy-service-consignment/proto/consignment"
)

type ShippingService struct {
	Repo repository
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

func (s *ShippingService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return err
	}

	res = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

func (s *ShippingService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.Repo.GetAll()
	res = &pb.Response{Consignments: consignments}
	return nil
}
