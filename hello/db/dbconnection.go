package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Client

func GetMongoDB(mongouri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongouri)
	mongoDB, err := mongo.Connect(context.Background(), clientOptions)
	fmt.Println(mongoDB)
	if err != nil {
		fmt.Println(err)
	}
	return mongoDB
}
