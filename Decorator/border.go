package main

type Border struct {
	Display
	*AbstractDisplay
}

func NewBorder(display Display) *Border {
	return &Border{
		Display:         display,
		AbstractDisplay: NewAbstractDisplay(),
	}
}
