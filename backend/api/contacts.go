package api

type Contacts []Contact

type Contact struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Note string `json:"note"`
}
