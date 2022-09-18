package main

type File struct {
	*AbstractEntry
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		AbstractEntry: NewAbstractEntry(),
		name:          name,
		size:          size,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Accept(v Visitor) {
	v.VisitFile(f)
}
