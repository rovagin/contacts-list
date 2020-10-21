package remove

import (
	"contacts-list/internal/app/contact/remove/adapter"
	"contacts-list/internal/app/contact/remove/requests"
	"contacts-list/internal/app/contact/remove/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
