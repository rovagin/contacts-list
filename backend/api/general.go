package api

import "encoding/json"

type Request struct {
	Payload json.RawMessage `json:"payload"`
}

type Response struct {
	Code    int             `json:"code"`
	Payload json.RawMessage `json:"payload"`
}
