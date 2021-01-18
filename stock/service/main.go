package main

import (

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	st"blatt2-grp03/stock"
	
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	broker := nats.NewBroker()
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("stock"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
		
	)

	
	service.Init()

	

	if err :=  api.RegisterStockHandler(service.Server(),
	 	st.CreateNewStockHandleInstance(micro.NewEvent("log.stock", service.Client()))); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
