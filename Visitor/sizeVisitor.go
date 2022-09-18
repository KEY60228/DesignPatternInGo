package main

type SizeVisitor struct {
	currentDir string
	totalSize  int
}

func NewSizeVisitor() *SizeVisitor {
	return &SizeVisitor{
		currentDir: "",
		totalSize:  0,
	}
}

func (sv *SizeVisitor) GetTotalSize() int {
	return sv.totalSize
}

func (sv *SizeVisitor) VisitFile(f *File) {
	sv.totalSize += f.GetSize()
}

func (lv *SizeVisitor) VisitDir(d *Directory) {
	saveDir := lv.currentDir
	lv.currentDir = lv.currentDir + "/" + d.GetName()
	for _, e := range d.dir {
		e.Accept(lv)
	}
	lv.currentDir = saveDir
}
