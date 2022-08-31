package main

import "fmt"

type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{
		str:   str,
		width: len(str),
	}
}

func (sd *StringDisplayImpl) RawOpen() {
	sd.printLine()
}

func (sd *StringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", sd.str)
}

func (sd *StringDisplayImpl) RawClose() {
	sd.printLine()
}

func (sd *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
