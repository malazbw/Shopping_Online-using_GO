package stock

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)

type StockHandlerService struct {
	prodcuts   map[string]int32
	
}


func CreateNewStockHandleInstance() *StockHandlerService {
	return &StockHandlerService{
		prodcuts:  make(map[string]int32),
	}
}


func (c *StockHandlerService) containsProdcut(name string, count int32) bool {
	
	if val, ok := c.prodcuts[name]; ok {
		if val >= count{
		return true
		}
	}
	return false
}


func (c *StockHandlerService) Supply(ctx context.Context, req *api.SupplyRequest, rsp *api.SupplyResponse) error {
	
	
	for key, value := range req.Products {
		c.prodcuts[key] = c.prodcuts[key] + value
	}
	rsp.State = "Hiho " 
	logger.Infof("zesss")
		
	return nil
}
func (c *StockHandlerService) ShipOrder(ctx context.Context, req *api.ShipOrderRequest, rsp *api.ShipOrderResponse) error {
	
	
	for key, value := range req.Products {
		c.prodcuts[key] = c.prodcuts[key] - value
	}
	rsp.State =  true 
	logger.Infof("shipping from stock")
		
	return nil
}



func (c *StockHandlerService) CheckOrder(ctx context.Context, req *api.OrderRequest, rsp *api.OrderCheckResponse) error {
	logger.Infof("step2") 
	var check = true
	for key, value := range req.Products {
		if (!c.containsProdcut(key, value)) {
			check = false
			break
		}
	}

	if check{
		rsp.State = "order is taken, you can pay now"
		logger.Infof("All items are available")
	}else {
		rsp.State = "order is not taken "
		logger.Infof("Not all items are available") 
	}
	
		
	return nil
}
