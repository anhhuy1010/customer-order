package order

type (
	OrderRequest struct {
		CustomerName string `json:"customer_name"`
		Phone        int64  `json:"phone"`
		Address      string `json:"address"`
		CartUuid     string `json:"cart_uuid"`
	}
	OrderResponse struct {
		CustomerName string               `json:"customer_name"`
		Uuid         string               `json:"uuid"`
		CartUuid     string               `json:"cart_uuid"`
		Phone        int                  `json:"phone"`
		Address      string               `json:"address"`
		Items        []OrderItemsResponse `json:"items"`
		Total        int                  `json:"total"`
	}
	OrderItemsResponse struct {
		ProductUuid  string  `json:"product_uuid" bson:"product_uuid"`
		ProductName  string  `json:"product_name" bson:"product_name"`
		ProductPrice float64 `json:"product_price" bson:"product_price"`
		Quantity     int     `json:"quantity" bson:"quantity"`
		ProductTotal float64 `json:"product_total" bson:"product_total"`
	}
)
