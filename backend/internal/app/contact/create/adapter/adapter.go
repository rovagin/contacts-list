package adapter

import (
	"contacts-list/internal/app/contact/create/usecase"
	"math/rand"
)

type Adapter struct {
}

func New() *Adapter {
	return &Adapter{}
}

// TODO: add actual DB request
func (a *Adapter) Save(contact usecase.Contact) (int, error) {
	return rand.Int(), nil
}
