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
func New(supplier api.SupplierService, order api.OrderService, payment api.PaymentService) *Client {
	return &Client{
		supplier: supplier,
		order:    order,
		payment:  payment,
	}
}

func (c *Client) Interact() {
	logger.Infof("Szenario 3")
	var products map[string]int32
	products = make(map[string]int32)
	products["laptop"] = 1
	rsp, err := c.supplier.Supply(context.TODO(), &api.SupplyRequest{Products: products})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Received: %+v", rsp.State)
	}

	var order map[string]int32
	order = make(map[string]int32)
	order["laptop"] = 1
	order["watch"] = 1
	order["tv"] = 1
	//order:=api.NewOrderService("order", clientService.Client())
	orderRsp, err2 := c.order.Place(context.TODO(), &api.OrderRequest{Products: order})
	if err2 != nil {
		logger.Error(err2)
	} else {
		logger.Infof("OrderId: %+v", orderRsp.Orderid)
	}

	//payment:=api.NewPaymentService("payment", clientService.Client())
	rspCancel, err3 := c.order.Cancel(context.TODO(), &api.CancelRequest{Orderid: orderRsp.Orderid})
	if err3 != nil {
		logger.Error(err3)
	} else {
		logger.Infof("OrderId: %+v", rspCancel.Message)

	}
	var products2 map[string]int32
	products2 = make(map[string]int32)
	products2["watch"] = 1
	products2["tv"] = 1
	rsp4, err4 := c.supplier.Supply(context.TODO(), &api.SupplyRequest{Products: products2})
	if err4 != nil {
		logger.Error(err4)
	} else {
		logger.Infof("Received: %+v", rsp4.State)
	}

}
