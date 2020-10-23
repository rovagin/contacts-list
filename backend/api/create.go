package api

type CreateContactRequest struct {
	UserID  int           `json:"user_id"`
	Contact CreateContact `json:"contact"`
}

type CreateContact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Note      string `json:"note"`
}
