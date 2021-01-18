package main

import (
	"blatt2-grp03/misc"
	order "blatt2-grp03/order"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"

	"blatt2-grp03/api"
)

/*
Main Function to start a new users service.
*/
func main() {

	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()
	service := micro.NewService(
		micro.Name("order"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()
	s := order.New(api.NewStockService("stock", service.Client()), api.NewShipmentService("shipment", service.Client()), api.NewPaymentService("payment", service.Client()))

	if err := api.RegisterOrderHandler(service.Server(),
		s); err != nil {
		panic(err)
	}

	if err := micro.RegisterSubscriber("log.*", service.Server(),
		order.NewSubscriber(s)); err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {

		logger.Fatal(err)
	}
}
