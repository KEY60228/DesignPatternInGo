package main

import "fmt"

func main() {
	b1 := NewStringDisplay("Hello, World.")
	b2 := NewSideBorder(b1, "#")
	b3 := NewFullBorder(b2)

	b1.AbstractDisplay.Show(b1)
	fmt.Println()
	b2.AbstractDisplay.Show(b2)
	fmt.Println()
	b3.AbstractDisplay.Show(b3)
	fmt.Println()

	b4 := NewSideBorder(NewFullBorder(NewFullBorder(NewSideBorder(NewFullBorder(NewStringDisplay("Hello, Wolrd")), "*"))), "/")
	b4.AbstractDisplay.Show(b4)
}
