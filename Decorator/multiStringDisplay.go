package main

import (
	"log"
	"strings"
)

type MultiStringDisplay struct {
	*AbstractDisplay
	strs   []string
	length int
}

func NewMultiStringDisplay() *MultiStringDisplay {
	return &MultiStringDisplay{
		AbstractDisplay: NewAbstractDisplay(),
		strs:            make([]string, 0),
		length:          0,
	}
}

func (d *MultiStringDisplay) GetColumns() int {
	return d.length
}

func (d *MultiStringDisplay) GetRows() int {
	return len(d.strs)
}

func (d *MultiStringDisplay) GetRowText(row int) string {
	if len(d.strs)-1 < row {
		log.Fatal("index out of bounds")
	}

	var s strings.Builder
	s.WriteString(d.strs[row])
	for i := 0; i < d.length-len(d.strs[row]); i++ {
		s.WriteString(" ")
	}
	return s.String()
}

func (d *MultiStringDisplay) Add(str string) {
	if d.length < len(str) {
		d.length = len(str)
	}
	d.strs = append(d.strs, str)
}
