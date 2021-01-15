package payment

import (
	"context"
	"math/rand"
	
	// "sync"
	"time"
	 "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type Payment struct {
	orders         map[int32]map[string]int32
	stock api.StockService
}

const (
	maxOrderId int32 = 987654321
)

func (o *Payment) appendANewOrder(id int32, order map[string]int32) bool {
	if id != 0  {
		(*o.getOrderMap())[id] = order
		return true
	}
	return false
}
// getUserMap will return a pointer to the current user map in order to work in that. //
func (o *Payment) getOrderMap() *map[int32]map[string]int32 {
	return &o.orders
}
func (u *Payment) containsID(id int32) bool {
	_, inMap := (*u.getOrderMap())[id]
	return inMap
}

func (u *Payment) getRandomOrderID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !u.containsID(potantialID) {
			return potantialID
		}
	}
}

func New(stock api.StockService) *Order {
	return &Order{
		orders:  make(map[int32]map[string]int32),
		stock:stock,

	
	}
}
func (o *Payment) Pay(ctx context.Context, request *api.PaymentRequest, response *api.PaymentRequest) error {
	logger.Infof("step1")
	orderCheckResponse, err := o.stock.CheckOrder(context.Background(), &api.OrderRequest{
		Products: request.Products,
	})
	logger.Infof("step2")
	if err != nil {
		logger.Error(err)
	}else{
		response.Invoiceid  = o.getRandomOrderID(maxOrderId)
		logger.Infof(orderCheckResponse.State)
	}
	return nil
}
