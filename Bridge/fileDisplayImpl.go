package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileDisplayImpl struct {
	str   string
	width int
}

func NewFileDisplayImpl(file *os.File) *FileDisplayImpl {
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return &FileDisplayImpl{
		str:   string(bytes),
		width: len(string(bytes)),
	}
}

func (fd *FileDisplayImpl) RawOpen() {
	fd.printLine()
}

func (fd *FileDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", fd.str)
}

func (fd *FileDisplayImpl) RawClose() {
	fd.printLine()
}

func (fd *FileDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < fd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
