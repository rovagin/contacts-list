package api

type CreateContactRequest struct {
	UserID  int     `json:"user_id"`
	Contact Contact `json:"contact"`
}
