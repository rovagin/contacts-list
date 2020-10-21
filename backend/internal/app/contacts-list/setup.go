package contactslist

import (
	"contacts-list/internal/app/contacts-list/adapter"
	"contacts-list/internal/app/contacts-list/requests"
	"contacts-list/internal/app/contacts-list/usecase"
)

func Setup() *requests.Requests {
	repo := adapter.New()

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}