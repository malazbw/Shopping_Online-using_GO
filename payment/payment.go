package payment

import (
	"context"

	"blatt2-grp03/api"

	"github.com/micro/go-micro/v2/logger"
)

type Payment struct {
	payments []int32

	order api.OrderService
}

func (p *Payment) remove(s []int32, i int32) []int32 {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func New(order api.OrderService) *Payment {
	return &Payment{
		order: order,
	}
}

func (p *Payment) Pay(ctx context.Context, request *api.PaymentRequest, response *api.PaymentResponse) error {
	logger.Infof("Zahlung erhalten für Bestellnummer: ", request.Orderid)
	response.Result = "Zahlung Erfolreich"
	p.payments = append(p.payments, request.Orderid)

	_, err := p.order.InformPayment(context.Background(), &api.InformPaymentRequest{
		Orderid: request.Orderid,
	})

	if err != nil {
		logger.Error(err)
	}

	return nil
}

func (p *Payment) Return(ctx context.Context, request *api.ReturnRequest, response *api.ReturnResponse) error {

	p.payments = p.remove(p.payments, request.Orderid)
	logger.Infof("Empfangen einer Retoure mit Ersatzlieferungswunsch.")

	return nil
}
