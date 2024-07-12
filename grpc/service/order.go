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
	_, err := order.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, item := range req.OrderItem {
		orderItem := models.OrderItem{
			Uuid:         util.GenerateUUID(),
			OrderUuid:    order.Uuid,
			ProductUuid:  item.ProductUuid,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     item.Quantity,
			ProductTotal: item.ProductPrice * float64(item.Quantity),
		}
		_, err := orderItem.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return &pb.CreateOrderResponse{
		OrderUuid: order.Uuid,
	}, nil
}
