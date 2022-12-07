package service

import (
	"fmt"
	"hello/db"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type HelloService interface {
	SayHello(string) string
	GetVisit(string) (Visit, error)
	GetVisits() ([]Visit, error)
	DeleteVisit(string) error
}

func HelloServiceInstance(l log.Logger) HelloService {
	return &helloService{logger: l}
}

type helloService struct {
	logger log.Logger
}

func (h helloService) SayHello(s string) string {
	if s == "" {
		return "Hello World!"
	}
	v := Visit{s, time.Now()}
	result, err := CreateVisit(v)
	if err != nil {
		level.Error(h.logger).Log(err)
	}
	return fmt.Sprintf("Hello %s, visisted id=%v", s, result)
}

type Visit struct {
	Name      string    `bson:"name"`
	Timestamp time.Time `bson:"time"`
}

func CreateVisit(v Visit) (string, error) {
	result, err := db.VisitCollection.InsertOne(db.Ctx, v)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex()), err
}

func (h helloService) GetVisit(id string) (Visit, error) {
	var v Visit
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return v, err
	}
	err = db.VisitCollection.FindOne(db.Ctx, bson.M{"_id": objectId}).Decode(&v)
	if err != nil {
		return v, err
	}
	return v, nil
}

func (h helloService) GetVisits() ([]Visit, error) {
	var visit Visit
	var visits []Visit
	cursor, err := db.VisitCollection.Find(db.Ctx, bson.M{})
	if err != nil {
		defer cursor.Close(db.Ctx)
		return visits, err
	}
	for cursor.Next(db.Ctx) {
		err := cursor.Decode(&visit)
		if err != nil {
			return visits, err
		}
		visits = append(visits, visit)
	}
	return visits, nil
}

func (h helloService) DeleteVisit(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = db.VisitCollection.DeleteOne(db.Ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	return nil
}
