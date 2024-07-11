package models

import (
	"context"
	"log"
	"time"

	"github.com/anhhuy1010/customer-order/database"
	"github.com/anhhuy1010/customer-order/helpers/util"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson"

	"github.com/anhhuy1010/customer-order/constant"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Orders struct {
	Uuid         string    `json:"uuid" bson:"uuid"`
	Name         string    `json:"name" bson:"name"`
	Address      string    `json:"address" bson:"address"`
	Phone        int       `json:"phone" bson:"phone"`
	Total        float64   `json:"total" bson:"total"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	ProductUuid  string    `json:"product_uuid" bson:"product_uuid"`
	ProductName  string    `json:"product_name" bson:"product_name"`
	ProductPrice float64   `json:"product_price" bson:"product_price"`
	Quantity     int       `json:"quantity" bson:"quantity"`
	ProductTotal float64   `json:"product_total" bson:"prodtuct_total"`
}

func (u *Orders) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("orders")
}

func (u *Orders) Find(conditions map[string]interface{}, opts ...*options.FindOptions) ([]*Orders, error) {
	coll := u.Model()

	cursor, err := coll.Find(context.TODO(), conditions, opts...)
	if err != nil {
		return nil, err
	}

	var orders []*Orders
	for cursor.Next(context.TODO()) {
		var elem Orders
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return orders, nil
}

func (u *Orders) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*Users, error) {
	coll := u.Model()

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var users []*Users
	for cursor.Next(context.TODO()) {
		var elem Users
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("[Decode] PopularCuisine:", err)
			log.Println("-> #", elem.Uuid)
			continue
		}

		users = append(users, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return users, nil
}

func (u *Orders) Distinct(conditions map[string]interface{}, fieldName string, opts ...*options.DistinctOptions) ([]interface{}, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	values, err := coll.Distinct(context.TODO(), fieldName, conditions, opts...)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (u *Orders) FindOne(conditions map[string]interface{}) (*Orders, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *Orders) Insert() (interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (u *Orders) InsertMany(Orders []interface{}) ([]interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertMany(context.TODO(), Orders)
	if err != nil {
		return nil, err
	}

	return resp.InsertedIDs, nil
}

func (u *Orders) Update() (int64, error) {
	coll := u.Model()

	condition := make(map[string]interface{})
	condition["uuid"] = u.Uuid

	u.UpdatedAt = util.GetNowUTC()
	updateStr := make(map[string]interface{})
	updateStr["$set"] = u

	resp, err := coll.UpdateOne(context.TODO(), condition, updateStr)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Orders) UpdateByCondition(condition map[string]interface{}, data map[string]interface{}) (int64, error) {
	coll := u.Model()

	resp, err := coll.UpdateOne(context.TODO(), condition, data)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Orders) UpdateMany(conditions map[string]interface{}, updateData map[string]interface{}) (int64, error) {
	coll := u.Model()
	resp, err := coll.UpdateMany(context.TODO(), conditions, updateData)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *Orders) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
