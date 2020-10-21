package adapter

type Adapter struct {
}

func New() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Update() error {
	return nil
}
