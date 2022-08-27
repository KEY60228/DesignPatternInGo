package main

func main() {
	d1 := NewDisplay(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewCountDisplay(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))

	d1.Do()
	d2.Do()
	d3.Do()
	d3.MultiDisplay(5)
}
