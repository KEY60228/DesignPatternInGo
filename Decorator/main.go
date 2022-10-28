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

	b5 := NewMultiStringDisplay()
	b5.Add("Hi!")
	b5.Add("Good morning.")
	b5.Add("Good night!")
	b5.Show(b5)

	b6 := NewSideBorder(b5, "#")
	b6.Show(b6)

	b7 := NewFullBorder(b5)
	b7.Show(b7)
}
