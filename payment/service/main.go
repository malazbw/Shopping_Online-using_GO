package main

import (
	// "time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	order"blatt2-grp03/order"
	
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	
	registry := etcdv3.NewRegistry()
	service := micro.NewService(	
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
	)

	


	if err := api.RegisterOrderHandler(service.Server(),
	order.New(api.NewStockService("stock", service.Client()))); err != nil {
	logger.Fatal(err)
}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
