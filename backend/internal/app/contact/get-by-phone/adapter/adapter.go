package adapter

import "contacts-list/internal/app/contact/get-by-phone/usecase"

type Adapter struct {
	stub usecase.Contacts
}

func New() *Adapter {
	return &Adapter{
		stub: []usecase.Contact{
			{
				ID:        0,
				FirstName: "123",
				LastName:  "123",
				Phone:     "123",
				Email:     "123",
				Note:      "123",
			},
		},
	}
}

func (r *Adapter) Get(phone string) (usecase.Contacts, error) {
	return r.stub, nil
}
