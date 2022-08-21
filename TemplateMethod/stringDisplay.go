package main

import "fmt"

type StringDisplay struct {
	s     string
	width int
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		s:     s,
		width: len(s),
	}
}

func (sd *StringDisplay) open() {
	sd.printLine()
}

func (sd *StringDisplay) print() {
	fmt.Println("|" + sd.s + "|")
}

func (sd *StringDisplay) close() {
	sd.printLine()
}

func (sd *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
