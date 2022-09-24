package main

type NumberGenerator interface {
	AddObserver(Observer)
	DeleteObserver(Observer)
	NotifyObservers(NumberGenerator)
	GetNumber() int
	Execute()
}

type AbstractNumberGenerator struct {
	observers []Observer
}

func NewAbstractNumberGenerator() *AbstractNumberGenerator {
	return &AbstractNumberGenerator{
		observers: make([]Observer, 0),
	}
}

func (ng *AbstractNumberGenerator) AddObserver(observer Observer) {
	ng.observers = append(ng.observers, observer)
}

func (ng *AbstractNumberGenerator) DeleteObserver(observer Observer) {
	os := make([]Observer, 0, len(ng.observers))
	for _, o := range ng.observers {
		if o != observer {
			os = append(os, o)
		}
	}
	ng.observers = os
}

func (ng *AbstractNumberGenerator) NotifyObservers(g NumberGenerator) {
	for _, o := range ng.observers {
		o.Update(g)
	}
}
