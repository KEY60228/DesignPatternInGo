package main

import "strings"

type UpDownBorder struct {
	*Border
	str string
}

func NewUpDownBorder(display Display, str string) *UpDownBorder {
	return &UpDownBorder{
		Border: NewBorder(display),
		str:    str,
	}
}

func (b *UpDownBorder) GetColumns() int {
	return b.Display.GetColumns()
}

func (b *UpDownBorder) GetRows() int {
	return 1 + b.Display.GetRows() + 1
}

func (b *UpDownBorder) GetRowText(row int) string {
	if row == 0 || row == b.Display.GetRows()+1 {
		return b.makeLine(b.str, b.Display.GetColumns())
	}
	return b.Display.GetRowText(row - 1)
}

func (b *UpDownBorder) makeLine(str string, count int) string {
	var s strings.Builder
	for i := 0; i < count; i++ {
		s.WriteString(str)
	}
	return s.String()
}
