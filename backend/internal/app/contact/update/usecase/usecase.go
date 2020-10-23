package usecase

type Repository interface {
	Update(userID int, contactID int, fields map[string]interface{}) error
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

// TODO: add check for phone duplicate
func (u *Usecase) Do(userID int, contactID int, fields map[string]interface{}) error {
	err := u.contactsRepo.Update(userID, contactID, fields)
	if err != nil {
		return err
	}

	return nil
}
