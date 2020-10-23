package usecase

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Note      string
}

type Repository interface {
	Save(userID int, contact Contact) error
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
func (u *Usecase) Do(userID int, contact Contact) error {
	err := u.contactsRepo.Save(userID, contact)
	if err != nil {
		return err
	}

	return nil
}
