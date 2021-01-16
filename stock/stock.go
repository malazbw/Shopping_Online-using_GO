package stock

import (
	"context"
	"fmt"
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
	
	fmt.Println("products before in Stock ", s.prodcuts)
	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] + value
	}
	rsp.State = "Hiho " 
	fmt.Println("products in Stock ", s.prodcuts)
		
	return nil
}
func (s *StockHandlerService) ShipOrder(ctx context.Context, req *api.ShipOrderRequest, rsp *api.ShipOrderResponse) error {
	
	
	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] - value
	}
	rsp.State =  true 
	logger.Infof("shipping from stock")
	fmt.Println(req.Products)
		
	return nil
}




func (s *StockHandlerService) ReserveOrder(ctx context.Context, req *api.ReserveOrderRequest, rsp *api.ReserveOrderResponse) error {
	
	var available = true
	fmt.Println("products in stock ", s.prodcuts)
	fmt.Println("products in req ", req.Products)
	for key, value := range req.Products {
		s.reservedProducts[key] = s.reservedProducts[key] + value
		if (s.reservedProducts[key] > s.prodcuts[key]){
			
			available = false
		}
		logger.Infof("All reservedProducts are available", s.reservedProducts[key] )
		logger.Infof("All prodcuts are available", s.prodcuts[key] )
	}
	rsp.State = available
	if available{
		logger.Infof("All items are available", )

	}else {
		
		logger.Infof(" all items Not are available") 
	}
		
	return nil
}