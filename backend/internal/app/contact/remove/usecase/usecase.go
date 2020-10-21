package usecase

type Repository interface {
	Remove(id int) error
}

type Usecase struct {
	contactsRepo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		contactsRepo: repo,
	}
}

func (u *Usecase) Do(id int) error {
	err := u.contactsRepo.Remove(id)
	if err != nil {
		return err
	}

	return nil
}
