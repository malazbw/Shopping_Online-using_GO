package main

import (
	// "time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	st"blatt2-grp03/stock"
	
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("stock"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
	)

	

	if err :=  api.RegisterStockHandler(service.Server(), st.CreateNewStockHandleInstance()); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
