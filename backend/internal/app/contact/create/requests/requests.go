package requests

import (
	"contacts-list/api"
	"contacts-list/internal/pkg/http/wrapper"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"contacts-list/internal/app/contact/create/usecase"
	"contacts-list/internal/pkg/errors"
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

	if request.Payload == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contactReq := new(api.CreateContactRequest)
	err = json.Unmarshal(request.Payload, contactReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateRequest(contactReq)
	if err != nil {
		// TODO: there could be general response with hint to bad data
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := r.interactor.Do(usecase.Contact{
		FirstName: contactReq.FirstName,
		LastName:  contactReq.LastName,
		Phone:     contactReq.Phone,
		Email:     contactReq.Email,
		Note:      contactReq.Note,
	})
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

	payloadBytes, err := json.Marshal(&api.CreateContactResponse{ID: result})
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

// TODO: Add some basic checks
func validateRequest(req *api.CreateContactRequest) error {
	if req.LastName == "" || req.FirstName == "" {
		return errors.New("bad last name")
	}

	return nil
}
