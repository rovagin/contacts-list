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

func (a *Adapter) Remove(userID int, contactID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	_, err := a.collection.UpdateOne(ctx,
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"contacts": bson.M{"id": contactID}}},
	)

	return err
}
