package adapter

import "time"

type Collection interface {
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

func (a *Adapter) Remove(id int) error {

	return nil
}
