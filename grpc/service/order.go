package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/anhhuy1010/customer-order/grpc/proto/order"
	"github.com/anhhuy1010/customer-order/helpers/util"
	"github.com/anhhuy1010/customer-order/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
}

func NewOrderServer() pb.OrderServer {
	return &OrderService{}
}

func (s *OrderService) Create(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	//kiểm tra cart_uuid
	if req.CartUuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "CartUuid is required")
	}
	//kiểm tra phone
	if req.Phone == "" {
		return nil, status.Errorf(codes.InvalidArgument, "CartUuid is required")
	}
	//thêm vào database
	order := models.Orders{
		Uuid:      util.GenerateUUID(),
		Name:      req.CustomerName,
		Address:   req.Address,
		Phone:     req.Phone,
		Total:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	orderitem := models.OrderItem{
		Uuid:         util.GenerateUUID(),
		OrderUuid:    order.Uuid,
		ProductUuid:  "",
		ProductName:  "",
		ProductPrice: 0,
		Quantity:     0,
		ProductTotal: 0,
	}
	_, err := order.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = orderitem.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &pb.CreateOrderResponse{
		OrderUuid: order.Uuid,
	}, nil
}
