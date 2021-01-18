package shipment

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type Shipment struct {
	stock api.StockService
}


func New(stock api.StockService) *Shipment {
	return &Shipment{
		stock:stock,
	}
}
func (o *Shipment) Ship(ctx context.Context, request *api.ShipmentRequest, response *api.ShipmentResponse) error {
	_, err := o.stock.ShipOrder(context.Background(), &api.ShipOrderRequest{
		Products: request.Products,
	})
	
	if err != nil {
		logger.Error(err)
	}else{
		
		logger.Infof("products in shipment ", request.Products)
	}
	return nil
}
