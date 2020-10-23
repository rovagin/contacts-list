package requests

import (
	"contacts-list/api"
	"contacts-list/internal/app/contact/remove/usecase"
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

	requestPayload := &api.RemoveContactRequest{}

	err = json.Unmarshal(request.Payload, requestPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate(requestPayload)
	if err != nil {
		// TODO: there could be general response with hint to bad data
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.interactor.Do(requestPayload.UserID, requestPayload.ContactID)
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func validate(payload *api.RemoveContactRequest) error {
	if payload.UserID < 0 {
		return errors.New("bad user id")
	}

	if payload.ContactID < 0 {
		return errors.New("bad contact id")
	}

	return nil
}
