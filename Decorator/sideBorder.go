package main

type SideBorder struct {
	*Border
	str string
}

func NewSideBorder(display Display, str string) *SideBorder {
	return &SideBorder{
		Border: NewBorder(display),
		str:    str,
	}
}

func (b *SideBorder) GetColumns() int {
	return 1 + b.Display.GetColumns() + 1
}

func (b *SideBorder) GetRows() int {
	return b.Display.GetRows()
}

func (b *SideBorder) GetRowText(row int) string {
	return b.str + b.Display.GetRowText(row) + b.str
}
