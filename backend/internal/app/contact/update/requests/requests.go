package requests

import (
	"contacts-list/api"
	"contacts-list/internal/app/contact/update/usecase"
	"contacts-list/internal/pkg/errors"
	"contacts-list/internal/pkg/http/wrapper"
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.Payload == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateReq := new(api.UpdateContactRequest)
	err = json.Unmarshal(request.Payload, updateReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate(updateReq)
	if err != nil {
		// TODO: there could be general response with hint to bad data
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.interactor.Do(updateReq.UserID, updateReq.ContactID, updateReq.Fields)
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

var updateFields = map[string]struct{}{
	"last_name":  {},
	"first_name": {},
	"phone":      {},
	"email":      {},
	"note":       {},
}

func validate(req *api.UpdateContactRequest) error {
	if req.UserID < 0 {
		return errors.New("bad user id")
	}

	if req.ContactID < 0 {
		return errors.New("bad contact id")

	}

	for k := range req.Fields {
		if _, ok := updateFields[k]; !ok {
			return errors.New("bad update field")
		}
	}

	return nil
}
