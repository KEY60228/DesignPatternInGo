package main

import "log"

type StringDisplay struct {
	*AbstractDisplay
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		AbstractDisplay: NewAbstractDisplay(),
		str:             str,
	}
}

func (d *StringDisplay) GetColumns() int {
	return len(d.str)
}

func (d *StringDisplay) GetRows() int {
	return 1
}

func (d *StringDisplay) GetRowText(row int) string {
	if row != 0 {
		log.Fatal("index out of bounds")
	}
	return d.str
}
