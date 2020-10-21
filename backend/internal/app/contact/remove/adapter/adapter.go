package adapter

type Adapter struct {
}

func New() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Remove(id int) error {
	return nil
}
