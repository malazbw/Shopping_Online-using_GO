package client

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	"blatt2-grp03/api"
)

type Client struct {
	supplier api.SupplierService
	order    api.OrderService
	payment  api.PaymentService
}

//consructor
func New(supplier api.SupplierService, order api.OrderService, payment  api.PaymentService ) *Client {
	return &Client{
		supplier: supplier,
		order: order,
		payment: payment,
	}
}

func (c *Client) Interact() {

	logger.Info("Szenario 1: ")
	var product map[string]int32
	product = make(map[string]int32)
	product["watch"] = 2
	product["laptop"] = 5
	rsp, err := c.supplier.Supply(context.TODO(), &api.SupplyRequest{Products: product})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Received: %+v", rsp.State)
	}

	var m2 map[string]int32
	m2 = make(map[string]int32)
	m2["watch"] = 1
	//order:=api.NewOrderService("order", clientService.Client())
	orderRsp, err2 := c.order.Place(context.TODO(), &api.OrderRequest{Products: m2})
	if err2 != nil {
		logger.Error(err2)
	} else {
		logger.Info("Received OrderId: %+v", orderRsp.Orderid)
	}

	//payment:=api.NewPaymentService("payment", clientService.Client())
	paymentRsp, err3 := c.payment.Pay(context.TODO(), &api.PaymentRequest{Orderid: orderRsp.Orderid})
	if err3 != nil {
		logger.Error(err3)
	} else {
		logger.Info("OrderId: %+v", paymentRsp.Result)

	}

	rsp4, err4 := c.supplier.Supply(context.TODO(), &api.SupplyRequest{Products: product})
	if err4 != nil {
		logger.Error(err4)
	} else {
		logger.Info("Received: %+v", rsp4.State)
	}

}
