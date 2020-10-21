package create

import (
	"contacts-list/internal/app/contact/create/adapter"
	"contacts-list/internal/app/contact/create/requests"
	"contacts-list/internal/app/contact/create/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
