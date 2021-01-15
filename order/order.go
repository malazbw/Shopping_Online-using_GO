package customer

import (
	"context"
	"math/rand"
	"fmt"
	// "sync"
	"time"
	 "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type Order struct {
	orders         map[int32]map[string]int32
	ordersAvalability         map[int32]bool
	stock api.StockService
}

const (
	maxOrderId int32 = 987654321
)

func (o *Order) appendANewOrder(id int32, order map[string]int32) bool {
	if id != 0  {
		(*o.getOrderMap())[id] = order
		return true
	}
	return false
}
// getUserMap will return a pointer to the current user map in order to work in that. //
func (o *Order) getOrderMap() *map[int32]map[string]int32 {
	return &o.orders
}
func (u *Order) containsID(id int32) bool {
	_, inMap := (*u.getOrderMap())[id]
	return inMap
}

func (u *Order) getRandomOrderID(length int32) int32 {
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
		ordersAvalability:  make(map[int32]bool),
		stock:stock,

	
	}
}
func (o *Order) Place(ctx context.Context, request *api.OrderRequest, response *api.OrderResponse) error {
	logger.Infof("step1")
	orderCheckResponse, err := o.stock.ReserveOrder(context.Background(), &api.ReserveOrderRequest{
		Products: request.Products,
	})
	logger.Infof("step2")
	if err != nil {
		logger.Error(err)
	}else{
		response.Orderid  = o.getRandomOrderID(maxOrderId)
		if len(orderCheckResponse.Products) == 0 {
			o.ordersAvalability[response.Orderid] = true
			logger.Infof("order ",response.Orderid,"is not fully taken")
		}else{
			o.ordersAvalability[response.Orderid] = false
			logger.Infof("order ",response.Orderid,"is fully taken")
		}
		
		
	}
	return nil
}

