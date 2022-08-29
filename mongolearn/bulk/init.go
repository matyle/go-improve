package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Decimal *decimal.Decimal

type text struct {
	Text string           `bson:"text"`
	Num  *decimal.Decimal `bson:"num"`
}

func Init() {
	uri := "mongodb://localhost:27018"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).
		SetRegistry(bson.NewRegistryBuilder().
			RegisterDecoder(reflect.TypeOf(decimal.Decimal{}), Decimal{}).
			RegisterEncoder(reflect.TypeOf(decimal.Decimal{}), Decimal{}).Build()))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin bulk
	coll := client.Database("insertDB").Collection("haikus")
	// models := []mongo.WriteModel{
	// 	// mongo.NewReplaceOneModel().SetFilter(bson.D{{"title", "Record of a Shriveled Datum"}}).
	// 	// 	SetReplacement(bson.D{{"title", "Dodging Greys"}, {"text", "When there're no matches, no longer need to panic. You can use upsert"}}).SetUpsert(true),
	// 	mongo.NewUpdateOneModel().SetFilter(bson.M{"title": "Dodging  test 223432rf"}).
	// 		SetUpdate(bson.M{
	// 			"$set":         bson.M{"symobl": "Dodge The Greys test"},
	// 			"$setOnInsert": bson.M{"dfsa": "Dodge The Greys test setInset"},
	// 		}).SetUpsert(true),
	// }
	// opts := options.BulkWrite().SetOrdered(true)

	// results, err := coll.BulkWrite(context.TODO(), models, opts)
	// // end bulk
	d := decimal.NewFromFloat(1.1)
	t := &text{
		Text: "decimal23",
		Num:  &d,
	}
	ptr := true
	results, err := coll.UpdateOne(
		context.TODO(),
		bson.M{"title": "Dodging Greys dafdfds 33432erqwe3eqwe23423f"},
		bson.M{
			"$set":         bson.M{"datanfdsb": "Dodging Greys testfdaddfd nochange sdfd"},
			"$setOnInsert": t,
		},
		&options.UpdateOptions{
			Upsert: &ptr,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of documents replaced or modified: %d", results.ModifiedCount)
}
