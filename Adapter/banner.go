package main

import "fmt"

type Banner struct {
	s string
}

func NewBanner(s string) *Banner {
	return &Banner{s: s}
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(" + b.s + ")")
}

func (b *Banner) ShowWithAster() {
	fmt.Println("*" + b.s + "*")
}
