package api

type RemoveContactRequest struct {
	UserID    int `json:"user_id"`
	ContactID int `json:"contact_id"`
}
