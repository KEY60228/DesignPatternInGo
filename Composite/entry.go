package main

import "fmt"

type Entry interface {
	GetName() string
	GetSize() int
	PrintList(string)
	ToString(Entry) string
}

type AbstractEntry struct{}

func NewAbstractEntry() *AbstractEntry {
	return &AbstractEntry{}
}

func (e *AbstractEntry) ToString(entry Entry) string {
	return fmt.Sprintf("%s (%d)", entry.GetName(), entry.GetSize())
}
