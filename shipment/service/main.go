package main

import (
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"blatt2-grp03/shipment"
	
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	
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

	
	service.Init(
		micro.Action(func(c *cli.Context) error {
			sleep := c.Int("sleep")
			if sleep > 0 {
				logger.Infof("sleeping %d seconds before startup", sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
			}

			return nil
		}),
	)


	


	if err := api.RegisterShipmentHandler(service.Server(),
	shipment.New(api.NewStockService("stock", service.Client()))); err != nil {
	logger.Fatal(err)
}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
