package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/JonasMuylaert/shippy-service-consignment/handler"
	pb "github.com/JonasMuylaert/shippy-service-consignment/proto/consignment"
)

// func newServer() *shippingService {
// 	s := &shippingService{
// 		repo: &Repository{},
// 	}
// 	return s
// }

func main() {
	//register service with micro
	srv := service.New(
		service.Name("shippy.service.consignment"),
	)

	srv.Init()

	//register handler
	if err := pb.RegisterShippingServiceHandler(srv.Server(), &handler.ShippingService{
		Repo: &handler.Repository{},
	}); err != nil {
		logger.Errorf("Failed creating handler: %v", err)
	}

	//run server
	if err := srv.Run(); err != nil {
		logger.Fatalf("failed starting server: %v", err)
	}
}
