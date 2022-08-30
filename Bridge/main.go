package main

import (
	"log"
	"os"
)

func main() {
	d1 := NewDisplay(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewCountDisplay(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))
	d4 := NewRandomDisplay(NewStringDisplayImpl("Random Display"))

	f, err := os.Open("string.txt")
	if err != nil {
		log.Fatal(err)
	}
	d5 := NewDisplay(NewFileDisplayImpl(f))

	d1.Do()
	d2.Do()
	d3.Do()
	d3.MultiDisplay(5)
	d4.RandomDo(100)
	d5.Do()
}
