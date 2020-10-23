package usecase

type Repository interface {
	Remove(userID, contactID int) error
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

func (u *Usecase) Do(userID, contactID int) error {
	err := u.contactsRepo.Remove(userID, contactID)
	if err != nil {
		return err
	}

	return nil
}
