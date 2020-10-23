package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Config struct {
	URI             string        `env:"MONGO_URI"`
	RequestsTimeout time.Duration `env:"MONGO_REQUESTS_TIMEOUT"`
}

type Connection struct {
	// Used in Connector to differentiate connects from other in case of connection error
	name string

	uri string

	*mongo.Database
}

func New(name string, config Config) *Connection {
	return &Connection{
		name: name,
		uri:  config.URI,
	}
}

func (c *Connection) Name() string {
	return c.name
}

func (c *Connection) Connect(ctx context.Context) (err error) {
	config, err := connstring.Parse(c.uri)
	if err != nil {
		return errors.Wrap(err, "connection string parse")
	}

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(c.uri),
	)
	if err != nil {
		return errors.Wrap(err, "connect")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "ping")
	}

	c.Database = client.Database(config.Database)

	return nil
}

func (c *Connection) Disconnect(ctx context.Context) error {
	return c.Client().Disconnect(ctx)
}

func (c *Connection) Close() error {
	return c.Disconnect(context.Background())
}

func (c *Connection) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	return &Collection{
		conn:       c,
		Collection: c.Database.Collection(name, opts...),
	}
}
