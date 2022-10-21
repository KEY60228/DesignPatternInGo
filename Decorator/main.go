package main

import "fmt"

func main() {
	b1 := NewStringDisplay("Hello, World.")
	b2 := NewUpDownBorder(b1, "-")
	b3 := NewSideBorder(b2, "*")

	b1.Show(b1)
	fmt.Println()
	b2.Show(b2)
	fmt.Println()
	b3.Show(b3)
	fmt.Println()

	b4 := NewFullBorder(NewUpDownBorder(NewSideBorder(NewUpDownBorder(NewSideBorder(NewStringDisplay("Hello, World"), "*"), "="), "|"), "/"))
	b4.Show(b4)
	fmt.Println()
}
