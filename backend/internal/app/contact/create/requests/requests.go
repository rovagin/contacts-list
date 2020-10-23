package requests

import (
	"contacts-list/api"
	"contacts-list/internal/pkg/http/wrapper"
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(err)
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

	err = validateRequest(contactReq.Contact)
	if err != nil {
		// TODO: there could be general response with hint to bad data
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.interactor.Do(contactReq.UserID, usecase.Contact{
		ID:        contactReq.Contact.ID,
		FirstName: contactReq.Contact.FirstName,
		LastName:  contactReq.Contact.LastName,
		Phone:     contactReq.Contact.Phone,
		Email:     contactReq.Contact.Email,
		Note:      contactReq.Contact.Note,
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

	w.WriteHeader(http.StatusOK)
}

// TODO: Add some basic checks
func validateRequest(contact api.CreateContact) error {
	if contact.FirstName == "" {
		return errors.New("bad first name")
	}

	if contact.LastName == "" {
		return errors.New("bad last name")
	}

	return nil
}
