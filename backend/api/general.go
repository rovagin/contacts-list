package api

import "encoding/json"

type Request struct {
	RID     string          `json:"rid"`
	Payload json.RawMessage `json:"payload"`
}

type Response struct {
	RID     string          `json:"rid"`
	Code    int             `json:"code"`
	Payload json.RawMessage `json:"payload"`
}
