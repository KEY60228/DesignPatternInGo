package main

type Displayable interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	Displayable
}

func NewAbstractDisplay(d Displayable) *AbstractDisplay {
	return &AbstractDisplay{d}
}

func (ad *AbstractDisplay) Display() {
	ad.open()
	for i := 0; i < 5; i++ {
		ad.print()
	}
	ad.close()
}
