package main

type Directory struct {
	*AbstractEntry
	name string
	dir  []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		AbstractEntry: NewAbstractEntry(),
		name:          name,
		dir:           []Entry{},
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	var size int
	for _, e := range d.dir {
		size += e.GetSize()
	}
	return size
}

func (d *Directory) Add(e Entry) {
	d.dir = append(d.dir, e)
}

func (d *Directory) Accept(v Visitor) {
	v.VisitDir(d)
}
