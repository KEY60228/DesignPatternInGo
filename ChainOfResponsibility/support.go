package main

import "fmt"

type Support interface {
	SetNext(Support) Support
	GetNext() Support
	Do(Support, *Trouble)
	Resolve(*Trouble) bool
	Done(Support, *Trouble)
	Fail(*Trouble)
	String() string
}

type AbstractSupport struct {
	name string
	next Support
}

func NewAbstractSupport(name string) *AbstractSupport {
	return &AbstractSupport{
		name: name,
		next: nil,
	}
}

func (s *AbstractSupport) SetNext(next Support) Support {
	s.next = next
	return next
}

func (s *AbstractSupport) GetNext() Support {
	return s.next
}

func (s *AbstractSupport) Do(support Support, trouble *Trouble) {
	next := support
	for {
		if next.Resolve(trouble) {
			s.Done(next, trouble)
			break
		}
		next = next.GetNext()
		if next == nil {
			s.Fail(trouble)
			break
		}
	}
}

func (s *AbstractSupport) Done(support Support, trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, support)
}

func (s *AbstractSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (s *AbstractSupport) String() string {
	return fmt.Sprintf("[%s]", s.name)
}
