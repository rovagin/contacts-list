package requests

import (
	"contacts-list/internal/app/contact/remove/usecase"
	"contacts-list/internal/pkg/errors"
	"contacts-list/internal/pkg/http/wrapper"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Requests struct {
	interactor *usecase.Usecase
}

func New(interactor *usecase.Usecase) *Requests {
	return &Requests{
		interactor: interactor,
	}
}

func (r *Requests) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	userID, err := sanitize("user id", vars["user"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contactID, err := sanitize("contact id", vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.interactor.Do(userID, contactID)
	if err != nil {
		code, payload := errors.ProcessError(err)
		fullResponse, err := wrapper.BuildResponse(code, payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(fullResponse)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func sanitize(field, payload string) (int, error) {
	userID, err := strconv.Atoi(payload)
	if err != nil {
		return 0, errors.Errorf("bad %s value", field)
	}

	if userID < 0 {
		return 0, errors.Errorf("bad %s value", field)
	}

	return userID, nil
}
