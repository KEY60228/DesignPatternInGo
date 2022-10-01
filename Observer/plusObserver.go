package main

import "fmt"

type PlusObserver struct{}

func NewPlusObserver() *PlusObserver {
	return &PlusObserver{}
}

func (o *PlusObserver) Update(generator NumberGenerator) {
	fmt.Print("PlusObserver: ")
	for i := 0; i < generator.GetNumber(); i++ {
		fmt.Print("+")
	}
	fmt.Println()
}
