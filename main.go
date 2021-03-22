package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/JonasMuylaert/shippy-service-consignment/handler"
)

func main() {
	//register service with micro
	srv := service.New(
		service.Name("consignment"),
	)

	srv.Init()

	srv.Handle(handler.ShippingService{
		Repo: &handler.Repository{},
	})

	//run server
	if err := srv.Run(); err != nil {
		logger.Fatalf("failed starting server: %v", err)
	}
}
