package remove

import (
	"contacts-list/internal/app/contact/remove/adapter"
	"contacts-list/internal/app/contact/remove/requests"
	"contacts-list/internal/app/contact/remove/usecase"
	"time"
)

func Setup(coll adapter.Collection, collTimeout time.Duration) *requests.Requests {
	repo := adapter.New(coll, collTimeout)

	u := usecase.New(repo)

	reqs := requests.New(u)

	return reqs
}
