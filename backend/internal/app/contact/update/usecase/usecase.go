package usecase

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Note      string
}

type Repository interface {
	Update() error
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

func (u *Usecase) Do() error {
	err := u.contactsRepo.Update()
	if err != nil {
		return err
	}

	return nil
}
