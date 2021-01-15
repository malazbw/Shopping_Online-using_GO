package main

import (
	"context"
	"fmt"

	// "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
)
func main() {
	fmt.Println("Start Test Skript")
	
	registry := etcdv3.NewRegistry()
	store := redis.NewStore()

	clientService := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	clientService.Init()



		var m2 map[string]int32
	m2 = make(map[string]int32)
	m2["route"] = 66
		order:=api.NewOrderService("order", clientService.Client())
		rsp2, err := order.Place(context.TODO(), &api.OrderRequest{Products: m2})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Received: %+v", rsp2.Invoiceid)
			}

}