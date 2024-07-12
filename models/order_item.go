package models

import (
	"context"
	"log"

	"github.com/anhhuy1010/customer-order/constant"
	"github.com/anhhuy1010/customer-order/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
)

type OrderItem struct {
	Uuid         string  `json:"uuid" bson:"uuid"`
	OrderUuid    string  `json:"order_uuid" bson:"order_uuid"`
	ProductUuid  string  `json:"product_uuid" bson:"product_uuid"`
	ProductName  string  `json:"product_name" bson:"product_name"`
	ProductPrice float64 `json:"product_price" bson:"product_price"`
	Quantity     int64   `json:"quantity" bson:"quantity"`
	ProductTotal float64 `json:"product_total" bson:"prodtuct_total"`
}

func (u *OrderItem) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("order_item")
}
func (u *OrderItem) Find(conditions map[string]interface{}, opts ...*options.FindOptions) ([]*OrderItem, error) {
	coll := u.Model()
	cursor, err := coll.Find(context.TODO(), conditions, opts...)
	if err != nil {
		return nil, err
	}

	var orderItem []*OrderItem
	for cursor.Next(context.TODO()) {
		var elem OrderItem
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		orderItem = append(orderItem, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return orderItem, nil
}

func (u *OrderItem) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*OrderItem, error) {
	coll := u.Model()

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var orderItem []*OrderItem
	for cursor.Next(context.TODO()) {
		var elem OrderItem
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("[Decode] PopularCuisine:", err)
			log.Println("-> #", elem.Uuid)
			continue
		}

		orderItem = append(orderItem, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return orderItem, nil
}

func (u *OrderItem) Distinct(conditions map[string]interface{}, fieldName string, opts ...*options.DistinctOptions) ([]interface{}, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	values, err := coll.Distinct(context.TODO(), fieldName, conditions, opts...)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (u *OrderItem) FindOne(conditions map[string]interface{}) (*OrderItem, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *OrderItem) Insert() (interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (u *OrderItem) InsertMany(OrderItem []interface{}) ([]interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertMany(context.TODO(), OrderItem)
	if err != nil {
		return nil, err
	}

	return resp.InsertedIDs, nil
}

func (u *OrderItem) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
