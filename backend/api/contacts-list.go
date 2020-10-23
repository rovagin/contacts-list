package api

type ContactsListRequest struct {
	UserID int `json:"user_id"`
}

type ContactsListResponse []Contact
