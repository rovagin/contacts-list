package usecase

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Note      string
}

type Contacts []Contact

type Repository interface {
	Get(name string) (Contacts, error)
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

func (u *Usecase) Do(name string) (Contacts, error) {
	result, err := u.contactsRepo.Get(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}
