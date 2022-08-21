package main

import "fmt"

type ByteDisplay struct {
	ch byte
}

func NewByteDisplay(ch byte) *ByteDisplay {
	return &ByteDisplay{ch: ch}
}

func (bd *ByteDisplay) open() {
	fmt.Print("<<")
}

func (bd *ByteDisplay) print() {
	fmt.Print(string(bd.ch))
}

func (bd *ByteDisplay) close() {
	fmt.Println(">>")
}
