package adapter

import (
	"contacts-list/internal/app/contacts-list/usecase"
	"contacts-list/internal/pkg/mongo"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserContacts struct {
	UserID   int       `bson:"user_id"`
	Contacts []Contact `bson:"contacts"`
}

type Contact struct {
	ID        int    `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Phone     string `bson:"phone"`
	Email     string `bson:"email"`
	Note      string `bson:"note"`
}

type Collection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) mongo.Decoder
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

func (a *Adapter) Get(id int) (usecase.Contacts, error) {
	result := make(usecase.Contacts, 0)

	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	decoder := a.collection.FindOne(ctx, map[string]interface{}{"_id": id})

	userContacts := &UserContacts{Contacts: make([]Contact, 0)}

	err := decoder.Decode(userContacts)
	if err != nil {
		return nil, err
	}

	for _, v := range userContacts.Contacts {
		result = append(result, usecase.Contact{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Phone:     v.Phone,
			Email:     v.Email,
			Note:      v.Note,
		})
	}

	return result, nil
}
