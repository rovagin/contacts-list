package wrapper

import (
	"contacts-list/api"
	"encoding/json"
)

func GetRequest(body []byte) (*api.Request, error) {
	request := new(api.Request)

	err := json.Unmarshal(body, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func BuildResponse(code int, payload []byte) ([]byte, error) {
	result, err := json.Marshal(&api.Response{
		Code:    code,
		Payload: payload,
	})

	return result, err
}
