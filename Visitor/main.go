package main

import "fmt"

func main() {
	fmt.Println("Making root entries...")

	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))
	rootDir.Accept(NewListVisitor())

	fmt.Println()
	fmt.Println("Making user entries...")

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	usrDir.Add(yuki)
	usrDir.Add(hanako)
	usrDir.Add(tomura)
	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.go", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))
	rootDir.Accept(NewListVisitor())

	fmt.Println()
}
