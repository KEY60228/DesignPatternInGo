package main

import "fmt"

type ListVisitor struct {
	currentDir string
}

func NewListVisitor() *ListVisitor {
	return &ListVisitor{
		currentDir: "",
	}
}

func (lv *ListVisitor) VisitFile(f *File) {
	fmt.Printf("%s/%s\n", lv.currentDir, f.ToString(f))
}

func (lv *ListVisitor) VisitDir(d *Directory) {
	fmt.Printf("%s/%s\n", lv.currentDir, d.ToString(d))
	saveDir := lv.currentDir
	lv.currentDir = lv.currentDir + "/" + d.GetName()
	for _, e := range d.dir {
		e.Accept(lv)
	}
	lv.currentDir = saveDir
}
