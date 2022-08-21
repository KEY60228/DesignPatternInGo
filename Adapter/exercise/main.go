package main

import "fmt"

func main() {
	f := NewFileProperties()
	f.ReadFromFile("file.txt")
	f.SetValue("width", "1024")
	f.SetValue("height", "512")
	f.SetValue("depth", "32")
	f.WriteToFile("newfile.txt")
	fmt.Println("newfile.txt is created.")
}
