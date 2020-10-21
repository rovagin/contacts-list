package getbyname

import (
	"contacts-list/internal/app/contact/get-by-name/adapter"
	"contacts-list/internal/app/contact/get-by-name/requests"
	"contacts-list/internal/app/contact/get-by-name/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
