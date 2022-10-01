package main

import (
	"math/rand"
	"time"
)

type RandomNumberGenerator struct {
	*AbstractNumberGenerator
	number int
}

func NewRandomNumberGenerator(number int) *RandomNumberGenerator {
	rand.Seed(time.Now().UnixNano())
	return &RandomNumberGenerator{
		AbstractNumberGenerator: NewAbstractNumberGenerator(),
		number:                  number,
	}
}

func (ng *RandomNumberGenerator) GetNumber() int {
	return ng.number
}

func (ng *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		ng.number = rand.Intn(50)
		ng.NotifyObservers(ng)
	}
}
