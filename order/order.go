package customer

import (
	"context"
	"math/rand"

	// "sync"
	"time"
	 "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type Order struct {
	orders         map[int32]map[string]int32
	availableOrders         map[int32]bool
	paidOrders        map[int32]bool
	stock api.StockService
	shipment api.ShipmentService
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

func New(stock api.StockService, shipment api.ShipmentService) *Order {
	return &Order{
		orders:  make(map[int32]map[string]int32),
		availableOrders:  make(map[int32]bool),
		paidOrders:  make(map[int32]bool),
		stock:stock,
		shipment:shipment,

	
	}
}
func (o *Order) Place(ctx context.Context, request *api.OrderRequest, response *api.OrderResponse) error {
	logger.Infof("step1")
	reserveOrderResponse, err := o.stock.ReserveOrder(context.Background(), &api.ReserveOrderRequest{
		Products: request.Products,
	})
	logger.Infof("step2")
	if err != nil {
		logger.Error(err)
	}else{
		response.Orderid  = o.getRandomOrderID(maxOrderId)
		o.orders[response.Orderid] = request.Products
		if (reserveOrderResponse.State) {
			o.availableOrders[response.Orderid] = true
			logger.Infof("order ",response.Orderid,"is  fully taken")
		}else{
			o.availableOrders[response.Orderid] = false
			logger.Infof("order ",response.Orderid,"is not fully taken")
		}
		
		
	}
	return nil
}


func (o *Order) InformPayment(ctx context.Context, request *api.InformPaymentRequest, response *api.InformPaymentResponse) error {
	logger.Infof("InformPayment",request.Orderid,":", o.availableOrders[request.Orderid])
	o.paidOrders[request.Orderid] = true
	if (o.availableOrders[request.Orderid]){
		_, err := o.shipment.Ship(context.Background(), &api.ShipmentRequest{
			Products: o.orders[request.Orderid],
		})
		if err != nil {
			logger.Error(err)
	}
	
	}
	return nil
}
