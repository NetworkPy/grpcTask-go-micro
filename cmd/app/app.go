package main

import (
	"log"

	"github.com/NetworkPy/grpcTask/internal/point"
	"github.com/NetworkPy/grpcTask/internal/pointservice"
	micro "github.com/asim/go-micro/v3"
)

func main() {
	serv := micro.NewService(
		micro.Name("pointsservice"),
	)
	serv.Init()

	pointService := pointservice.NewPointService()

	point.RegisterPointserviceHandler(serv.Server(), pointService)

	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}
}
