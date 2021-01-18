package order

import (
	"context"

	// "sync"
	"blatt2-grp03/api"

	"github.com/micro/go-micro/v2/logger"
)

type OrderSubscriber struct {
	orderService *OrderService
}

func NewSubscriber(orderService *OrderService) *OrderSubscriber {
	return &OrderSubscriber{
		orderService:orderService,
	}
}
func (o *OrderSubscriber) Process(ctx context.Context, event *api.Event) error {
	logger.Infof("Received event msg: %+v", event.GetMessage())

	o.orderService.updateFaildOrders(event.Products)
	o.orderService.ShipIfPossible()
	

	return nil
}