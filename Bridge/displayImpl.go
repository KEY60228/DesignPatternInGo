package main

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}
