package main

type SpecialSupport struct {
	*AbstractSupport
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	return &SpecialSupport{
		AbstractSupport: NewAbstractSupport(name),
		number:          number,
	}
}

func (s *SpecialSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber() == s.number
}
