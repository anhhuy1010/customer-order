package controllers

import (
	"fmt"
	"net/http"

	"github.com/anhhuy1010/customer-order/helpers/respond"
	"github.com/anhhuy1010/customer-order/models"
	request "github.com/anhhuy1010/customer-order/request/order"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderController struct {
}

func (orderCtl OrderController) Detail(c *gin.Context) {
	orderItemModel := new(models.OrderItem)
	orderModels := new(models.Orders)
	var reqUri request.GetOrderRequestUri
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.OrderUuid}
	order, err := orderModels.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, respond.ErrorCommon("Order not found!"))
		return
	}
	condition = bson.M{"order_uuid": reqUri.OrderUuid}
	orderItems, err := orderItemModel.Find(condition)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", orderItems)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, respond.ErrorCommon("Cart items not found!"))
		return
	}

	total := 0.0
	var itemm []request.GetOrderItemResponse
	for _, item := range orderItems {
		productTotal := item.ProductPrice * float64(item.Quantity)
		total += productTotal

		itemm = append(itemm, request.GetOrderItemResponse{
			ProductUuid:   item.ProductUuid,
			OrderUuid:     item.OrderUuid,
			OrderItemUuid: item.OrderUuid,
			ProductName:   item.ProductName,
			ProductPrice:  item.ProductPrice,
			Quantity:      item.Quantity,
			ProductTotal:  productTotal,
		})
	}
	response := request.GetOrderResponse{
		Uuid:         reqUri.OrderUuid,
		OrderItem:    itemm,
		CustomerName: order.Name,
		Address:      order.Address,
		Phone:        order.Phone,
		Total:        total,
	}
	c.JSON(http.StatusOK, respond.Success(response, "Successfully retrieved order details"))
}
