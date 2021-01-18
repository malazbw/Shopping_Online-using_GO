package order

import (
	"context"

	"math/rand"

	// "sync"
	"blatt2-grp03/api"
	"time"

	"github.com/micro/go-micro/v2/logger"
)

type order struct {
	Key    int
	Option string
}

type OrderService struct {
	orders           map[int32]map[string]int32
	faildItemsOrders map[int32]map[string]int32
	paidOrders       map[int32]bool
	sentOrders       map[int32]bool
	stock            api.StockService
	shipment         api.ShipmentService
	payment          api.PaymentService
}

const (
	maxOrderId int32 = 987654321
)

func (o *OrderService) isOrderAvailabe(id int32) bool {
	var state = true
	for _, count := range o.faildItemsOrders[id] {
		if count > 0 {
			state = false
			break
		}
	}
	return state
}

func (o *OrderService) updateFaildOrders(items map[string]int32) error {
	for item, count := range items {

		var rest = count
		for id, _ := range o.faildItemsOrders {
			if rest == 0 {
				break
			}
			if o.faildItemsOrders[id][item] > 0 {
				if rest <= o.faildItemsOrders[id][item] {
					o.faildItemsOrders[id][item] -= rest
					rest = 0
				} else {
					rest -= o.faildItemsOrders[id][item]
					o.faildItemsOrders[id][item] = 0
				}
			}
		}

	}
	return nil
}

func (o *OrderService) appendANewOrder(id int32, order map[string]int32) bool {
	if id != 0 {
		(*o.getOrderMap())[id] = order
		return true
	}
	return false
}

// getUserMap will return a pointer to the current user map in order to work in that. //
func (o *OrderService) getOrderMap() *map[int32]map[string]int32 {
	return &o.orders
}
func (u *OrderService) containsID(id int32) bool {
	_, inMap := (*u.getOrderMap())[id]
	return inMap
}

func (u *OrderService) getRandomOrderID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !u.containsID(potantialID) {
			return potantialID
		}
	}
}

func New(stock api.StockService, shipment api.ShipmentService, payment api.PaymentService) *OrderService {
	return &OrderService{
		orders:           make(map[int32]map[string]int32),
		faildItemsOrders: make(map[int32]map[string]int32),
		paidOrders:       make(map[int32]bool),
		sentOrders:       make(map[int32]bool),
		stock:            stock,
		shipment:         shipment,
		payment:          payment,
	}
}
func (o *OrderService) Place(ctx context.Context, request *api.OrderRequest, response *api.OrderResponse) error {
	reserveOrderResponse, err := o.stock.ReserveOrder(context.Background(), &api.ReserveOrderRequest{
		Products: request.Products,
	})

	if err != nil {
		logger.Error(err)
	} else {
		var OrderID = o.getRandomOrderID(maxOrderId)

		response.Orderid = OrderID
		o.orders[OrderID] = request.Products
		o.paidOrders[OrderID] = false
		o.sentOrders[OrderID] = false
		//o.faildItemsOrders[response.Orderid] = request.Products
		var products map[string]int32
		products = make(map[string]int32)
		for item, _ := range request.Products {

			products[item] = 0
		}
		o.faildItemsOrders[OrderID] = products
		logger.Info("Order's items are reserved")
		logger.Info("Order id: ", OrderID)
		if reserveOrderResponse.State {

			logger.Info("All items are availabe in stock")
		} else {
			logger.Info("Not all items are availabe")
			for key, value := range reserveOrderResponse.FaildItems {
				o.faildItemsOrders[OrderID][key] = value

			}
		}

	}
	return nil
}

func (o *OrderService) ShipIfPossible() error {
	for id := range o.faildItemsOrders {
		if o.isOrderAvailabe(id) && o.paidOrders[id] {
			_, err := o.shipment.Ship(context.Background(), &api.ShipmentRequest{
				Products: o.orders[id],
			})
			if err != nil {
				logger.Error(err)
			}
		}
	}
	return nil
}

func (o *OrderService) InformPayment(ctx context.Context, request *api.InformPaymentRequest, response *api.InformPaymentResponse) error {
	logger.Info("InformPayment", request.Orderid, ":", o.orders[request.Orderid])
	o.paidOrders[request.Orderid] = true
	if o.isOrderAvailabe(request.Orderid) {
		_, err := o.shipment.Ship(context.Background(), &api.ShipmentRequest{
			Products: o.orders[request.Orderid],
		})
		if err != nil {
			logger.Error(err)
		} else {
			o.sentOrders[request.Orderid] = true
		}

	}
	return nil
}

func (o *OrderService) Cancel(ctx context.Context, request *api.CancelRequest, response *api.CancelResponse) error {
	var orderID = request.Orderid
	if _, ok := o.orders[orderID]; ok {
		if !o.sentOrders[orderID] {
			logger.Info("Empfangen einer Retour mit Ersatzlieferungswunsch für orderID", orderID)
			delete(o.sentOrders, orderID)
			delete(o.orders, orderID)
			if o.paidOrders[orderID] {
				_, err := o.payment.Return(context.Background(), &api.ReturnRequest{
					Orderid: orderID,
				})
				if err != nil {
					logger.Error(err)
				} else {
					logger.Info("payment", orderID, "is returned")
					response.Message = "order is cancelled payment is returned"
					delete(o.paidOrders, orderID)
				}
			} else {
				logger.Info("order", orderID, "is cancelled")
				response.Message = "order is cancelled"
			}

		} else {
			logger.Info("order", orderID, "is already shipped")
			response.Message = "order is already shipped "
		}
	} else {
		logger.Info("order not found")
		response.Message = "order not found"
	}
	return nil

}
