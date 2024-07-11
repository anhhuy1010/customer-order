package models

import (
	"github.com/anhhuy1010/customer-order/database"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson"
)

type OrderItem struct {
	Uuid         string  `json:"uuid" bson:"uuid"`
	OrderUuid    string  `json:"order_uuid" bson:"order_uuid"`
	ProductUuid  string  `json:"product_uuid" bson:"product_uuid"`
	ProductName  string  `json:"product_name" bson:"product_name"`
	ProductPrice float64 `json:"product_price" bson:"product_price"`
	Quantity     int     `json:"quantity" bson:"quantity"`
	ProductTotal float64 `json:"product_total" bson:"prodtuct_total"`
}

func (u *OrderItem) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("order_item")
}
