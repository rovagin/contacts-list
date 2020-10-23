package adapter

import (
	"contacts-list/internal/app/contact/create/usecase"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	ID        int    `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Phone     string `bson:"phone"`
	Email     string `bson:"email"`
	Note      string `bson:"note"`
}

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

func (a *Adapter) Save(userID int, contact usecase.Contact) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	_, err := a.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$addToSet": bson.M{"contacts": &Contact{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
		Note:      contact.Note,
	}}})

	return err
}
