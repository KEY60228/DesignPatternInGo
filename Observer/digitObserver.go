package main

import "fmt"

type DigitObserver struct{}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (o *DigitObserver) Update(generator NumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.GetNumber())
}
