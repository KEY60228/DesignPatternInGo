package main

type PrintBanner struct {
	*Banner
}

func NewPrintBanner(s string) *PrintBanner {
	return &PrintBanner{NewBanner(s)}
}

func (pb *PrintBanner) PrintWeak() {
	pb.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.ShowWithAster()
}
