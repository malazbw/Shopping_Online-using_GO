package main

import (

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"blatt2-grp03/shipment"
	"blatt2-grp03/misc"
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	service := micro.NewService(	
		micro.Name("shipment"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
		
	)

	
	service.Init()


	


	if err := api.RegisterShipmentHandler(service.Server(),
	shipment.New(api.NewStockService("stock", service.Client()))); err != nil {
	logger.Fatal(err)
}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
