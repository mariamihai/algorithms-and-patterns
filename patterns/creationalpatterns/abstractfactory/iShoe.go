package abstractfactory

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) GetSize() int {
	return s.size
}
