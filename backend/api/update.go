package api

type UpdateContactRequest struct {
	UserID    int                    `json:"user_id"`
	ContactID int                    `json:"contact_id"`
	Fields    map[string]interface{} `json:"fields"`
}
