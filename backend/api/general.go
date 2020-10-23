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

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Note      string `json:"note"`
}
