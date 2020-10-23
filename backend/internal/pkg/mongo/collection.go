package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	conn *Connection
	*mongo.Collection
}

func (c *Collection) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions,
) Decoder {
	return c.Collection.FindOne(ctx, filter, opts...)
}

type Decoder interface {
	Decode(interface{}) error
}

func (c *Collection) Database() *Connection {
	return c.conn
}

type Cursor interface {
	Decoder

	Next(context.Context) bool
	Close(context.Context) error
	Err() error
}

func (c *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	return c.Collection.Find(ctx, filter, opts...)
}

func (c *Collection) CountDocuments(
	ctx context.Context, filter interface{}, opts ...*options.CountOptions,
) (
	int64, error,
) {
	return c.Collection.CountDocuments(ctx, filter, opts...)
}
