package getbyphone

import (
	"contacts-list/internal/app/contact/get-by-phone/adapter"
	"contacts-list/internal/app/contact/get-by-phone/requests"
	"contacts-list/internal/app/contact/get-by-phone/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
