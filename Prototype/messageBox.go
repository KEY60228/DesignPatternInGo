package main

import (
	"fmt"

	"DesignPatternInGo/Prototype/framework"
)

type MessageBox struct {
	decochar string
}

func NewMessageBox(decochar string) *MessageBox {
	return &MessageBox{decochar: decochar}
}

func (mb *MessageBox) Use(s string) {
	decolen := 1 + len(s) + 1
	for i := 0; i < decolen; i++ {
		fmt.Print(mb.decochar)
	}
	fmt.Println()
	fmt.Println(mb.decochar + s + mb.decochar)
	for i := 0; i < decolen; i++ {
		fmt.Print(mb.decochar)
	}
	fmt.Println()
}

func (mb *MessageBox) CreateCopy() framework.Product {
	copy := *mb
	return &copy
}
