package main

import "strings"

type FullBorder struct {
	*Border
}

func NewFullBorder(display Display) *FullBorder {
	return &FullBorder{
		Border: NewBorder(display),
	}
}

func (b *FullBorder) GetColumns() int {
	return 1 + b.Display.GetColumns() + 1
}

func (b *FullBorder) GetRows() int {
	return 1 + b.Display.GetRows() + 1
}

func (b *FullBorder) GetRowText(row int) string {
	if row == 0 {
		return "+" + b.makeLine("-", b.Display.GetColumns()) + "+"
	} else if row == b.Display.GetRows()+1 {
		return "+" + b.makeLine("-", b.Display.GetColumns()) + "+"
	} else {
		return "|" + b.Display.GetRowText(row-1) + "|"
	}
}

func (b *FullBorder) makeLine(str string, count int) string {
	var s strings.Builder
	for i := 0; i < count; i++ {
		s.WriteString(str)
	}
	return s.String()
}
