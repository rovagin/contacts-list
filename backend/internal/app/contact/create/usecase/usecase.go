package usecase

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Note      string
}

type Repository interface {
	Save(contact Contact) (int, error)
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

func (u *Usecase) Do(contact Contact) (int, error) {
	result, err := u.contactsRepo.Save(contact)
	if err != nil {
		return 0, err
	}

	return result, nil
}
