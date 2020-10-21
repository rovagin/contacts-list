package requests

import (
	"contacts-list/api"
	"contacts-list/internal/app/contacts-list/usecase"
	"contacts-list/internal/pkg/errors"
	"encoding/json"
	"net/http"
)

type Requests struct {
	interactor *usecase.Usecase
}

func New(interactor *usecase.Usecase) *Requests {
	return &Requests{
		interactor: interactor,
	}
}

func (r *Requests) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	result, err := r.interactor.Do(0)
	if err != nil {
		w.Write(errors.ProcessError(err))
	}

	formatted := toAPI(result)

	bytes, err := json.Marshal(formatted)
	if err != nil {
		w.Write(errors.ProcessError(err))
	}

	w.Write(bytes)
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
