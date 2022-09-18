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
	v := NewSizeVisitor()
	for _, e := range d.dir {
		e.Accept(v)
	}
	return v.totalSize
}

func (d *Directory) Add(e Entry) {
	d.dir = append(d.dir, e)
}

func (d *Directory) Accept(v Visitor) {
	v.VisitDir(d)
}
