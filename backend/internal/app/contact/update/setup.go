package update

import (
	"contacts-list/internal/app/contact/update/adapter"
	"contacts-list/internal/app/contact/update/requests"
	"contacts-list/internal/app/contact/update/usecase"
	"time"
)

func Setup(coll adapter.Collection, collTimeout time.Duration) *requests.Requests {
	repo := adapter.New(coll, collTimeout)

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
