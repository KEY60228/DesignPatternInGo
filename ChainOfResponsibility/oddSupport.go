package main

type OddSupport struct {
	*AbstractSupport
}

func NewOddSupport(name string) *OddSupport {
	return &OddSupport{
		NewAbstractSupport(name),
	}
}

func (s *OddSupport) Resolve(trouble *Trouble) bool {
	return trouble.GetNumber()%2 == 1
}
