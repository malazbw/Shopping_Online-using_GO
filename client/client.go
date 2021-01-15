package client

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	"blatt2-grp03/api"
)

type Client struct {
	supplier api.SupplierService
	
}
func New(supplier api.SupplierService) *Client {
	return &Client{
		supplier: supplier,

	}
}
func (c *Client) Interact() {

	var m map[string]int32
	m = make(map[string]int32)
	m["route"] = 66
	rsp, err := c.supplier.Supply(context.TODO(), &api.SupplyRequest{Products: m})
		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Received: %+v", rsp.State)
		}


}
