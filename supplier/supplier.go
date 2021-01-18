package supplier

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)

type Supplier struct {
	stock api.StockService
}



func New(stock api.StockService) *Supplier {
	return &Supplier{
		stock: stock,
	}
}

func (s *Supplier) Supply(ctx context.Context, req *api.SupplyRequest, rsp *api.SupplyResponse) error {
	
	logger.Info("Supply: ", req.Products)
	SupplyResponse, err := s.stock.Supply(context.Background(), &api.SupplyRequest{
		Products: req.Products,
	})
	if err != nil {
		logger.Error(err)
	}else{
		logger.Info(SupplyResponse.State)
		rsp.State = "Prodcuts are supplied"
	}
	return nil
}
