package main

import (
	"fmt"
	"time"
)

type GraphObserver struct{}

func NewGraphObserver() *GraphObserver {
	return &GraphObserver{}
}

func (o *GraphObserver) Update(generator NumberGenerator) {
	fmt.Print("GraphObserver:")
	count := generator.GetNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	time.Sleep(100 * time.Millisecond)
}
