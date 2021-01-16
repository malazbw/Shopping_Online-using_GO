package payment

import (
	"context"
	
	 "github.com/micro/go-micro/v2/logger"
	"blatt2-grp03/api"
)


type Payment struct {

	order api.OrderService
}



func New(order api.OrderService) *Payment {
	return &Payment{
		
 	order:order,

	}
}
func (o *Payment) Pay(ctx context.Context, request *api.PaymentRequest, response *api.PaymentResponse) error {
	logger.Infof("step1")
	o.order.InformPayment(context.Background(), &api.InformPaymentRequest{
		Orderid: request.Orderid,
	})
	logger.Infof("Zahlung erhalten für Bestellnummer ", request.Orderid )
	return nil
}
