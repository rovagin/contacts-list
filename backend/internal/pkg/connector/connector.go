package connector

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Connection interface {
	Name() string
	Connect(ctx context.Context) error
	Close() error
}

type Config struct {
	Timeout time.Duration `env:"CONNECTOR_TIMEOUT"`
}

func Connect(config Config, connections ...Connection) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	for _, conn := range connections {
		var loopConn = conn // https://golang.org/doc/faq#closures_and_goroutines

		g.Go(func() error {
			err := loopConn.Connect(ctx)
			if err != nil {
				return fmt.Errorf("unable to connect with %v: %v", loopConn.Name(), err)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func Close(connections ...Connection) error {
	var g errgroup.Group

	for _, conn := range connections {
		var loopConn = conn //https://golang.org/doc/faq#closures_and_goroutines

		g.Go(func() error {
			err := loopConn.Close()
			if err != nil {
				return fmt.Errorf("failed to close '%s': %v", loopConn.Name(), err)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
