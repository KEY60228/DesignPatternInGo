package main

import (
	"fmt"
	"time"
)

type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	heavyJob("Printer生成中")
	return &Printer{
		name: name,
	}
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) Print(str string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(str)
}

func heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	fmt.Println("完了")
}
