package main

import (
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/vesose/example-micro/api"
	"gitlab.lrz.devss/semester/ob-20ws/blatt2/blatt2-grp03/customer"
)

/*
Main Function to start a new users service.
*/
func main() {


	service := micro.NewService(
		micro.Name("customer"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
	)

	err1 := proto.RegisterUsersHandler(service.Server(), newUserService)

	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}
}