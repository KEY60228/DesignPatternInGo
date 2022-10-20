package main

type PrintProxy struct {
	name string
	real *Printer
}

func NewPrintProxy(name string) *PrintProxy {
	return &PrintProxy{
		name: name,
		real: nil,
	}
}

func (pp *PrintProxy) SetPrinterName(name string) {
	if pp.real != nil {
		pp.real.SetPrinterName(name)
	}
	pp.name = name
}

func (pp *PrintProxy) GetPrinterName() string {
	return pp.name
}

func (pp *PrintProxy) Print(str string) {
	pp.Realize()
	pp.real.Print(str)
}

func (pp *PrintProxy) Realize() {
	if pp.real == nil {
		pp.real = NewPrinter(pp.name)
	}
}
