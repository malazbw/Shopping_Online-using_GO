package stock

import (
	"blatt2-grp03/api"
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

type StockHandlerService struct {
	publisher        micro.Event
	prodcuts         map[string]int32
	reservedProducts map[string]int32
}

func CreateNewStockHandleInstance(publisher micro.Event) *StockHandlerService {
	return &StockHandlerService{
		prodcuts:         make(map[string]int32),
		reservedProducts: make(map[string]int32),
		publisher:        publisher,
	}
}

func (s *StockHandlerService) containsProdcut(name string, count int32) bool {

	if val, ok := s.prodcuts[name]; ok {
		if val >= count {
			return true
		}
	}
	return false
}

func (s *StockHandlerService) Supply(ctx context.Context, req *api.SupplyRequest, rsp *api.SupplyResponse) error {

	logger.Info("Before: Products in Stock", s.prodcuts)
	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] + value
	}
	rsp.State = "The Products are in stock"
	logger.Info("After: products in Stock: ", s.prodcuts)

	if err := s.publisher.Publish(context.Background(), &api.Event{
		Products: req.Products,
		Message: "New prodcuts in stock",
	}); err != nil {
		logger.Errorf("error while publishing: %+v", err)
	}

	return nil
}
func (s *StockHandlerService) ShipOrder(ctx context.Context, req *api.ShipOrderRequest, rsp *api.ShipOrderResponse) error {

	for key, value := range req.Products {
		s.prodcuts[key] = s.prodcuts[key] - value
		s.reservedProducts[key] = s.reservedProducts[key] - value
	}
	rsp.State = true
	logger.Info("Shipping ", req.Products, " from stock")
	logger.Info("Prodcuts in stock ", s.prodcuts)

	return nil
}

func (s *StockHandlerService) ReserveOrder(ctx context.Context, req *api.ReserveOrderRequest, rsp *api.ReserveOrderResponse) error {

	var order map[string]int32
	order = make(map[string]int32)
	var available = true
	logger.Info("order's products", req.Products)
	for key, value := range req.Products {
		if s.reservedProducts[key] < s.prodcuts[key] {
			order[key] = (s.reservedProducts[key] + value) - s.prodcuts[key]
		} else {
			order[key] = value
		}

		s.reservedProducts[key] = s.reservedProducts[key] + value
		if order[key] > 0 {
			available = false
		}

	}

	if available {

		logger.Info("All items are available")

	} else {

		logger.Info("Not all items are available")
	}
	rsp.State = available
	rsp.FaildItems = order
	return nil
}
