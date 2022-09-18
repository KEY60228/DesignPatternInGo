package main

import "fmt"

type Entry interface {
	Element
	GetName() string
	GetSize() int
}

type AbstractEntry struct{}

func NewAbstractEntry() *AbstractEntry {
	return &AbstractEntry{}
}

func (ae *AbstractEntry) ToString(e Entry) string {
	return fmt.Sprintf("%s (%d)", e.GetName(), e.GetSize())
}
