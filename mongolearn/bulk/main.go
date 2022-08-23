package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://localhost:27018"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
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
	models := []mongo.WriteModel{
		// mongo.NewReplaceOneModel().SetFilter(bson.D{{"title", "Record of a Shriveled Datum"}}).
		// 	SetReplacement(bson.D{{"title", "Dodging Greys"}, {"text", "When there're no matches, no longer need to panic. You can use upsert"}}).SetUpsert(true),
		mongo.NewUpdateOneModel().SetFilter(bson.M{"title": "Dodging  test 2"}).
			SetUpdate(bson.M{
				"$set":         bson.M{"title": "Dodge The Greys test"},
				"$setOnInsert": bson.M{"symbol": "Dodge The Greys test setInset"},
			}).SetUpsert(true),
	}
	opts := options.BulkWrite().SetOrdered(true)

	results, err := coll.BulkWrite(context.TODO(), models, opts)
	// end bulk

	if err != nil {
		panic(err)
	}

	// When you run this file for the first time, it should print:
	// Number of documents replaced or modified: 2
	fmt.Printf("Number of documents replaced or modified: %d", results.ModifiedCount)
}
