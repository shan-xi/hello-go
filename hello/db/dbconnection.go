package db

import (
	"context"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	VisitCollection *mongo.Collection
	Ctx             = context.TODO()
)

func SetupMongoDBConnection(mongouri string) {
	logger := log.NewLogfmtLogger(os.Stderr)

	clientOptions := options.Client().ApplyURI(mongouri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		level.Error(logger).Log("msg", err)

	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		level.Error(logger).Log("msg", err)
	}
	db := client.Database("hello")
	VisitCollection = db.Collection("visit")
}
