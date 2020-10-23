package create

import (
	"contacts-list/internal/app/contact/create/adapter"
	"contacts-list/internal/app/contact/create/requests"
	"contacts-list/internal/app/contact/create/usecase"
	"time"
)

func Setup(coll adapter.Collection, collTImeout time.Duration) *requests.Requests {
	repo := adapter.New(coll, collTImeout)

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
