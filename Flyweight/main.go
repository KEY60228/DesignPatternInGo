package main

import (
	"fmt"
	"log"
)

func main() {
	var str string
	fmt.Print("enter: ")
	fmt.Scan(&str)
	if len(str) < 1 {
		log.Fatal("enter some numbers")
	}
	factory := NewBigCharFactory()

	bs := NewBigString(factory, str, true)
	bs.Print()
}
