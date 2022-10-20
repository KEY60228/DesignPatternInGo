package main

type Printable interface {
	SetPrinterName(string)
	GetPrinterName() string
	Print(string)
}
