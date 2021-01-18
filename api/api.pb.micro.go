// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Customer service

func NewCustomerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Customer service

type CustomerService interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...client.CallOption) (*CreateCustomerResponse, error)
}

type customerService struct {
	c    client.Client
	name string
}

func NewCustomerService(name string, c client.Client) CustomerService {
	return &customerService{
		c:    c,
		name: name,
	}
}

func (c *customerService) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...client.CallOption) (*CreateCustomerResponse, error) {
	req := c.c.NewRequest(c.name, "Customer.CreateCustomer", in)
	out := new(CreateCustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerHandler interface {
	CreateCustomer(context.Context, *CreateCustomerRequest, *CreateCustomerResponse) error
}

func RegisterCustomerHandler(s server.Server, hdlr CustomerHandler, opts ...server.HandlerOption) error {
	type customer interface {
		CreateCustomer(ctx context.Context, in *CreateCustomerRequest, out *CreateCustomerResponse) error
	}
	type Customer struct {
		customer
	}
	h := &customerHandler{hdlr}
	return s.Handle(s.NewHandler(&Customer{h}, opts...))
}

type customerHandler struct {
	CustomerHandler
}

func (h *customerHandler) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, out *CreateCustomerResponse) error {
	return h.CustomerHandler.CreateCustomer(ctx, in, out)
}

// Api Endpoints for Stock service

func NewStockEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Stock service

type StockService interface {
	Supply(ctx context.Context, in *SupplyRequest, opts ...client.CallOption) (*SupplyResponse, error)
	ReserveOrder(ctx context.Context, in *ReserveOrderRequest, opts ...client.CallOption) (*ReserveOrderResponse, error)
	ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...client.CallOption) (*ShipOrderResponse, error)
}

type stockService struct {
	c    client.Client
	name string
}

func NewStockService(name string, c client.Client) StockService {
	return &stockService{
		c:    c,
		name: name,
	}
}

func (c *stockService) Supply(ctx context.Context, in *SupplyRequest, opts ...client.CallOption) (*SupplyResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.Supply", in)
	out := new(SupplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockService) ReserveOrder(ctx context.Context, in *ReserveOrderRequest, opts ...client.CallOption) (*ReserveOrderResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.ReserveOrder", in)
	out := new(ReserveOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockService) ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...client.CallOption) (*ShipOrderResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.ShipOrder", in)
	out := new(ShipOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stock service

type StockHandler interface {
	Supply(context.Context, *SupplyRequest, *SupplyResponse) error
	ReserveOrder(context.Context, *ReserveOrderRequest, *ReserveOrderResponse) error
	ShipOrder(context.Context, *ShipOrderRequest, *ShipOrderResponse) error
}

func RegisterStockHandler(s server.Server, hdlr StockHandler, opts ...server.HandlerOption) error {
	type stock interface {
		Supply(ctx context.Context, in *SupplyRequest, out *SupplyResponse) error
		ReserveOrder(ctx context.Context, in *ReserveOrderRequest, out *ReserveOrderResponse) error
		ShipOrder(ctx context.Context, in *ShipOrderRequest, out *ShipOrderResponse) error
	}
	type Stock struct {
		stock
	}
	h := &stockHandler{hdlr}
	return s.Handle(s.NewHandler(&Stock{h}, opts...))
}

type stockHandler struct {
	StockHandler
}

func (h *stockHandler) Supply(ctx context.Context, in *SupplyRequest, out *SupplyResponse) error {
	return h.StockHandler.Supply(ctx, in, out)
}

func (h *stockHandler) ReserveOrder(ctx context.Context, in *ReserveOrderRequest, out *ReserveOrderResponse) error {
	return h.StockHandler.ReserveOrder(ctx, in, out)
}

func (h *stockHandler) ShipOrder(ctx context.Context, in *ShipOrderRequest, out *ShipOrderResponse) error {
	return h.StockHandler.ShipOrder(ctx, in, out)
}

// Api Endpoints for Supplier service

func NewSupplierEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Supplier service

type SupplierService interface {
	Supply(ctx context.Context, in *SupplyRequest, opts ...client.CallOption) (*SupplyResponse, error)
}

type supplierService struct {
	c    client.Client
	name string
}

func NewSupplierService(name string, c client.Client) SupplierService {
	return &supplierService{
		c:    c,
		name: name,
	}
}

func (c *supplierService) Supply(ctx context.Context, in *SupplyRequest, opts ...client.CallOption) (*SupplyResponse, error) {
	req := c.c.NewRequest(c.name, "Supplier.Supply", in)
	out := new(SupplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Supplier service

type SupplierHandler interface {
	Supply(context.Context, *SupplyRequest, *SupplyResponse) error
}

func RegisterSupplierHandler(s server.Server, hdlr SupplierHandler, opts ...server.HandlerOption) error {
	type supplier interface {
		Supply(ctx context.Context, in *SupplyRequest, out *SupplyResponse) error
	}
	type Supplier struct {
		supplier
	}
	h := &supplierHandler{hdlr}
	return s.Handle(s.NewHandler(&Supplier{h}, opts...))
}

type supplierHandler struct {
	SupplierHandler
}

func (h *supplierHandler) Supply(ctx context.Context, in *SupplyRequest, out *SupplyResponse) error {
	return h.SupplierHandler.Supply(ctx, in, out)
}

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	Place(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	InformPayment(ctx context.Context, in *InformPaymentRequest, opts ...client.CallOption) (*InformPaymentResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...client.CallOption) (*CancelResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Place(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Place", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) InformPayment(ctx context.Context, in *InformPaymentRequest, opts ...client.CallOption) (*InformPaymentResponse, error) {
	req := c.c.NewRequest(c.name, "Order.InformPayment", in)
	out := new(InformPaymentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Cancel(ctx context.Context, in *CancelRequest, opts ...client.CallOption) (*CancelResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Cancel", in)
	out := new(CancelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	Place(context.Context, *OrderRequest, *OrderResponse) error
	InformPayment(context.Context, *InformPaymentRequest, *InformPaymentResponse) error
	Cancel(context.Context, *CancelRequest, *CancelResponse) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		Place(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		InformPayment(ctx context.Context, in *InformPaymentRequest, out *InformPaymentResponse) error
		Cancel(ctx context.Context, in *CancelRequest, out *CancelResponse) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) Place(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderHandler.Place(ctx, in, out)
}

func (h *orderHandler) InformPayment(ctx context.Context, in *InformPaymentRequest, out *InformPaymentResponse) error {
	return h.OrderHandler.InformPayment(ctx, in, out)
}

func (h *orderHandler) Cancel(ctx context.Context, in *CancelRequest, out *CancelResponse) error {
	return h.OrderHandler.Cancel(ctx, in, out)
}

// Api Endpoints for Payment service

func NewPaymentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Payment service

type PaymentService interface {
	Pay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error)
	Return(ctx context.Context, in *ReturnRequest, opts ...client.CallOption) (*ReturnResponse, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Pay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.Pay", in)
	out := new(PaymentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Return(ctx context.Context, in *ReturnRequest, opts ...client.CallOption) (*ReturnResponse, error) {
	req := c.c.NewRequest(c.name, "Payment.Return", in)
	out := new(ReturnResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payment service

type PaymentHandler interface {
	Pay(context.Context, *PaymentRequest, *PaymentResponse) error
	Return(context.Context, *ReturnRequest, *ReturnResponse) error
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
		Pay(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error
		Return(ctx context.Context, in *ReturnRequest, out *ReturnResponse) error
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

func (h *paymentHandler) Pay(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error {
	return h.PaymentHandler.Pay(ctx, in, out)
}

func (h *paymentHandler) Return(ctx context.Context, in *ReturnRequest, out *ReturnResponse) error {
	return h.PaymentHandler.Return(ctx, in, out)
}

// Api Endpoints for Shipment service

func NewShipmentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Shipment service

type ShipmentService interface {
	Ship(ctx context.Context, in *ShipmentRequest, opts ...client.CallOption) (*ShipmentResponse, error)
}

type shipmentService struct {
	c    client.Client
	name string
}

func NewShipmentService(name string, c client.Client) ShipmentService {
	return &shipmentService{
		c:    c,
		name: name,
	}
}

func (c *shipmentService) Ship(ctx context.Context, in *ShipmentRequest, opts ...client.CallOption) (*ShipmentResponse, error) {
	req := c.c.NewRequest(c.name, "Shipment.Ship", in)
	out := new(ShipmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shipment service

type ShipmentHandler interface {
	Ship(context.Context, *ShipmentRequest, *ShipmentResponse) error
}

func RegisterShipmentHandler(s server.Server, hdlr ShipmentHandler, opts ...server.HandlerOption) error {
	type shipment interface {
		Ship(ctx context.Context, in *ShipmentRequest, out *ShipmentResponse) error
	}
	type Shipment struct {
		shipment
	}
	h := &shipmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Shipment{h}, opts...))
}

type shipmentHandler struct {
	ShipmentHandler
}

func (h *shipmentHandler) Ship(ctx context.Context, in *ShipmentRequest, out *ShipmentResponse) error {
	return h.ShipmentHandler.Ship(ctx, in, out)
}
