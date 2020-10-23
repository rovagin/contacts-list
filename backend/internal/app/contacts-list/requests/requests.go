package requests

import (
	"contacts-list/api"
	"contacts-list/internal/app/contacts-list/usecase"
	"contacts-list/internal/pkg/errors"
	"contacts-list/internal/pkg/http/wrapper"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request, err := wrapper.GetRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestPayload := &api.ContactsListRequest{}

	err = json.Unmarshal(request.Payload, requestPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate(requestPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := r.interactor.Do(requestPayload.UserID)
	if err != nil {
		code, payload := errors.ProcessError(err)
		fullResponse, err := wrapper.BuildResponse(request.RID, code, payload)
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

	fullResponse, err := wrapper.BuildResponse(request.RID, 0, payloadBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(fullResponse)
}

func validate(payload *api.ContactsListRequest) error {
	if payload.UserID < 0 {
		return errors.New("bad user id")
	}

	return nil
}

func toAPI(contacts usecase.Contacts) api.ContactsListResponse {
	result := make(api.ContactsListResponse, 0, len(contacts))

	for _, v := range contacts {
		result = append(result, api.Contact{
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
