package api

type ContactsListRequest struct {
	UserID int `json:"user_id"`
}

type ContactsListResponse []ContactsListContact

type ContactsListContact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Note      string `json:"note"`
}
