package requests

import (
	"contacts-list/internal/app/contact/update/usecase"
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

// TODO: serve request
func (r *Requests) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
