package update

import (
	"contacts-list/internal/app/contact/update/adapter"
	"contacts-list/internal/app/contact/update/requests"
	"contacts-list/internal/app/contact/update/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
