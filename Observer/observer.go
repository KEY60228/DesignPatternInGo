package main

type Observer interface {
	Update(NumberGenerator)
}
