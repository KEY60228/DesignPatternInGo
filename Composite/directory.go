package main

import "fmt"

type Directory struct {
	*AbstractEntry
	name    string
	entries []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		AbstractEntry: NewAbstractEntry(),
		name:          name,
		entries:       make([]Entry, 0),
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	var size int
	for _, entry := range d.entries {
		size += entry.GetSize()
	}
	return size
}

func (d *Directory) PrintList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, d.ToString(d))
	for _, entry := range d.entries {
		entry.PrintList(fmt.Sprintf("%s/%s", prefix, d.name))
	}
}

func (d *Directory) Add(entry Entry) {
	d.entries = append(d.entries, entry)
}
