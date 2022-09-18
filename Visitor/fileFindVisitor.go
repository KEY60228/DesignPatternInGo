package main

import (
	"strings"
)

type FileFindVisitor struct {
	currentDir string
	targetExt  string
	targets    []*File
}

func NewFileFindVisitor(targetExt string) *FileFindVisitor {
	return &FileFindVisitor{
		currentDir: "",
		targetExt:  targetExt,
	}
}

func (ffv *FileFindVisitor) GetFoundFiles() []*File {
	return ffv.targets
}

func (ffv *FileFindVisitor) VisitFile(f *File) {
	filename := f.GetName()
	extIndex := strings.LastIndex(filename, ".")
	if extIndex >= 0 && filename[extIndex:] == ffv.targetExt {
		ffv.targets = append(ffv.targets, f)
	}
}

func (ffv *FileFindVisitor) VisitDir(d *Directory) {
	saveDir := ffv.currentDir
	ffv.currentDir = ffv.currentDir + "/" + d.GetName()
	for _, e := range d.dir {
		e.Accept(ffv)
	}
	ffv.currentDir = saveDir
}
