package main

import (
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	payment"blatt2-grp03/payment"
	
	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	logger.Infof("BEFORE REGISTERY")

	registry := etcdv3.NewRegistry()
	service := micro.NewService(	
		micro.Name("payment"),
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

	logger.Infof("AFTER REGISTERY")



	if err := api.RegisterPaymentHandler(service.Server(),
	payment.New(api.NewOrderService("order", service.Client()))); err != nil {
		logger.Infof("NEW ORDER")

	logger.Fatal(err)
}

	if err := service.Run(); err != nil {
		logger.Infof("SERVICE RUN")
		logger.Fatal(err)
	}
}
