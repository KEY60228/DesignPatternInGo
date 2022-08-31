package main

type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{impl: impl}
}

func (d *Display) Open() {
	d.impl.RawOpen()
}

func (d *Display) Print() {
	d.impl.RawPrint()
}

func (d *Display) Close() {
	d.impl.RawClose()
}

func (d *Display) Do() {
	d.Open()
	d.Print()
	d.Close()
}
