package main

import (
	"github.com/JonasMuylaert/shippy-service-consignment/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
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
	srv.Handle(new(handler.ShippingService))

	//run server
	if err := srv.Run(); err != nil {
		logger.Fatalf("failed starting server: %v", err)
	}
}
