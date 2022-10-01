package main

type IncrementalNumberGenerator struct {
	*AbstractNumberGenerator
	number int
	limit  int
	inc    int
}

func NewIncrementalNumberGenerator(start int, limit int, inc int) *IncrementalNumberGenerator {
	return &IncrementalNumberGenerator{
		AbstractNumberGenerator: NewAbstractNumberGenerator(),
		number:                  start,
		limit:                   limit,
		inc:                     inc,
	}
}

func (ng *IncrementalNumberGenerator) GetNumber() int {
	return ng.number
}

func (ng *IncrementalNumberGenerator) Execute() {
	for ng.number <= ng.limit {
		ng.NotifyObservers(ng)
		ng.number += ng.inc
	}
}
