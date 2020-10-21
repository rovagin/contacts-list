package api

type CreateContactRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Note      string `json:"note"`
}

type CreateContactResponse struct {
	ID int `json:"id"`
}
