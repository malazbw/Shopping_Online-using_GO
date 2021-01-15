package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/misc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"blatt2-grp03/client"
	"blatt2-grp03/api"

)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()
	store := redis.NewStore()

	service := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	service.Init()

	client := client.New(api.NewSupplierService("supplier", service.Client()))

	client.Interact()
}
