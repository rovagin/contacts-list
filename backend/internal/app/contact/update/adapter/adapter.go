package adapter

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type Adapter struct {
	collection Collection
	timeout    time.Duration
}

func New(collection Collection, timeout time.Duration) *Adapter {
	return &Adapter{
		collection: collection,
		timeout:    timeout,
	}
}

func (a *Adapter) Update(userID int, contactID int, fields map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	preparedFields := make(map[string]interface{})

	for k, v := range fields {
		preparedFields["contacts.$."+k] = v
	}

	_, err := a.collection.UpdateOne(ctx, bson.M{"_id": userID, "contacts.id": contactID},
		bson.M{"$set": preparedFields},
	)

	return err
}
