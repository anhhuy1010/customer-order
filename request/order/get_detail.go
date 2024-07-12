package order

type (
	GetOrderRequestUri struct {
		OrderUuid string `uri:"uuid"`
	}
	GetOrderResponse struct {
		Uuid         string                 `json:"order_uuid"`
		CustomerName string                 `json:"customer_name"`
		OrderItem    []GetOrderItemResponse `json:"order_item"`
		Address      string                 `json:"address"`
		Phone        string                 `json:"phone"`
		Total        float64                `json:"total"`
	}
	GetOrderItemResponse struct {
		OrderItemUuid string  `json:"order_item_uuid" bson:"order_item_uuid"`
		OrderUuid     string  `json:"order_uuid" bson:"order_uuid"`
		ProductUuid   string  `json:"product_uuid" bson:"product_uuid"`
		ProductName   string  `json:"product_name" bson:"product_name"`
		ProductPrice  float64 `json:"product_price" bson:"product_price"`
		Quantity      int64   `json:"quantity" bson:"quantity"`
		ProductTotal  float64 `json:"product_total" bson:"prodtuct_total"`
	}
)
