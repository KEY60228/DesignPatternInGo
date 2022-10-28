package main

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(int) string
}

type AbstractDisplay struct{}

func NewAbstractDisplay() *AbstractDisplay {
	return &AbstractDisplay{}
}

func (d *AbstractDisplay) Show(display Display) {
	for i := 0; i < display.GetRows(); i++ {
		fmt.Println(display.GetRowText(i))
	}
}
