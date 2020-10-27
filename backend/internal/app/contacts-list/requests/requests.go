package requests

import (
	"contacts-list/api"
	"contacts-list/internal/app/contacts-list/usecase"
	"contacts-list/internal/pkg/errors"
	"contacts-list/internal/pkg/http/wrapper"
	"encoding/json"
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

	userID, err := sanitize(vars["user"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := r.interactor.Do(userID)
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

	formatted := toAPI(result)

	payloadBytes, err := json.Marshal(formatted)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fullResponse, err := wrapper.BuildResponse(0, payloadBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(fullResponse)
}

func sanitize(payload string) (int, error) {
	userID, err := strconv.Atoi(payload)
	if err != nil {
		return 0, errors.New("bad user id")
	}

	if userID < 0 {
		return 0, errors.New("bad user id")
	}

	return userID, nil
}

func toAPI(contacts usecase.Contacts) api.ContactsListResponse {
	result := make(api.ContactsListResponse, 0, len(contacts))

	for _, v := range contacts {
		result = append(result, api.ContactsListContact{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Phone:     v.Phone,
			Email:     v.Email,
			Note:      v.Note,
		})
	}

	return result
}
