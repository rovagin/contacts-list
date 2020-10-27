package api

type UpdateContactRequest struct {
	Fields map[string]interface{} `json:"fields"`
}
