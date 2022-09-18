package main

type Visitor interface {
	VisitFile(*File)
	VisitDir(*Directory)
}
