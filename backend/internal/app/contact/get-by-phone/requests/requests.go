package requests

import (
	"net/http"

	"contacts-list/api"
	"contacts-list/internal/app/contact/get-by-phone/usecase"
)

type Requests struct {
	interactor *usecase.Usecase
}

func New(interactor *usecase.Usecase) *Requests {
	return &Requests{
		interactor: interactor,
	}
}

// TODO: serve request
func (r *Requests) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}

func toAPI(contacts usecase.Contacts) api.Contacts {
	result := make(api.Contacts, 0, len(contacts))

	for _, v := range contacts {
		result = append(result, api.Contact{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Phone:     v.Phone,
			Email:     v.Email,
			Note:      v.Note,
		})
	}

	return result
}
