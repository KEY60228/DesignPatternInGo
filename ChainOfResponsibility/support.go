package main

import "fmt"

type Support interface {
	SetNext(Support) Support
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

func (s *AbstractSupport) Do(support Support, trouble *Trouble) {
	if support.Resolve(trouble) {
		s.Done(support, trouble)
	} else if s.next != nil {
		s.next.Do(s.next, trouble)
	} else {
		s.Fail(trouble)
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
