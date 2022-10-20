package main

import "fmt"

func main() {
	p := NewPrintProxy("Alice")
	fmt.Printf("名前は現在 %s です\n", p.GetPrinterName())
	p.SetPrinterName("Bob")
	fmt.Printf("名前は現在 %s です\n", p.GetPrinterName())
	p.Print("Hello World!")
}
