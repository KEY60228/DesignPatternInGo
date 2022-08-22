package main

import (
	"DesignPatternInGo/Prototype/framework"
)

func main() {
	manager := framework.NewManager()
	upen := NewUnderlinePen("-")
	mbox := NewMessageBox("*")
	sbox := NewMessageBox("/")

	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	p1 := manager.Create("strong message")
	p1.Use("Hello, World")

	p2 := manager.Create("warning box")
	p2.Use("Hello, World")

	p3 := manager.Create("slash box")
	p3.Use("Hello, World")
}
