package main

func main() {
	d1 := NewAbstractDisplay(NewByteDisplay('H'))
	d2 := NewAbstractDisplay(NewStringDisplay("Hello, World."))

	d1.Display()
	d2.Display()
}
