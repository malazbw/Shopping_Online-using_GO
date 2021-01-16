package main

import (
	"context"
	"fmt"
	// "time"
	// "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
)
func main() {
	fmt.Println("Start Script")
	
	registry := etcdv3.NewRegistry()
	store := redis.NewStore()

	clientService := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	clientService.Init()

	supplier:=api.NewSupplierService("supplier", clientService.Client())
	var m map[string]int32
	m = make(map[string]int32)
	m["route"] = 66
	
	fmt.Println("Calling supply with product route %v", m["route"])
	rsp, err := supplier.Supply(context.TODO(), &api.SupplyRequest{Products: m})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received from supply %+v", rsp.State)
		}
	// time.Sleep(5 * time.Second)

		
	fmt.Println("Calling supply with product route %v", m["route"])
		rsp4, err4 := supplier.Supply(context.TODO(), &api.SupplyRequest{Products: m})
		if err4 != nil {
			fmt.Println(err4)
		} else {
			fmt.Println("Received from supply: %+v", rsp4.State)
		}	

}