package main

type NoSupport struct {
	*AbstractSupport
}

func NewNoSupport(name string) *NoSupport {
	return &NoSupport{
		NewAbstractSupport(name),
	}
}

func (s *NoSupport) Resolve(trouble *Trouble) bool {
	return false
}
