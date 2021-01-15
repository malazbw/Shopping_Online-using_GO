package stock

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)

type StockHandlerService struct {
	prodcuts   map[string]int32
	reservedProducts   map[string]int32
	
}


func CreateNewStockHandleInstance() *StockHandlerService {
	return &StockHandlerService{
		prodcuts:  make(map[string]int32),
		reservedProducts:  make(map[string]int32),
	}
}


func (s *StockHandlerService) containsProdcut(name string, count int32) bool {
	
	if val, ok := s.prodcuts[name]; ok {
		if val >= count{
		return true
		}
	}
	return false
}


func (s *StockHandlerService) Supply(ctx context.Context, req *api.SupplyRequest, rsp *api.SupplyResponse) error {
	
	
	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] + value
	}
	rsp.State = "Hiho " 
	logger.Infof("zesss")
		
	return nil
}
func (s *StockHandlerService) ShipOrder(ctx context.Context, req *api.ShipOrderRequest, rsp *api.ShipOrderResponse) error {
	
	
	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] - value
	}
	rsp.State =  true 
	logger.Infof("shipping from stock")
		
	return nil
}




func (s *StockHandlerService) ReserveOrder(ctx context.Context, req *api.ReserveOrderRequest, rsp *api.ReserveOrderResponse) error {
	
	var available = true
	for key, value := range req.Products {
		s.reservedProducts[key] = s.reservedProducts[key] + value
		if (s.reservedProducts[key] > s.prodcuts[key]){
			available = false
		}
	}
	if available{
		rsp.State = "order is taken, you can pay now"
		logger.Infof("All items are available")
	}else {
		rsp.State = "order is not taken "
		logger.Infof("Not all items are available") 
	}
		
	return nil
}