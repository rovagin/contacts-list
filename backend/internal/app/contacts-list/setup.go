package contactslist

import (
	"contacts-list/internal/app/contacts-list/adapter"
	"contacts-list/internal/app/contacts-list/requests"
	"contacts-list/internal/app/contacts-list/usecase"
	"time"
)

func Setup(coll adapter.Collection, collTimeout time.Duration) *requests.Requests {
	repo := adapter.New(coll, collTimeout)

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
