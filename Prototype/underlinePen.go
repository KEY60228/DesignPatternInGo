package main

import (
	"fmt"

	"DesignPatternInGo/Prototype/framework"
)

type UnderlinePen struct {
	ulchar string
}

func NewUnderlinePen(ulchar string) *UnderlinePen {
	return &UnderlinePen{ulchar: ulchar}
}

func (up *UnderlinePen) Use(s string) {
	ulen := len(s)
	fmt.Println(s)
	for i := 0; i < ulen; i++ {
		fmt.Print(up.ulchar)
	}
	fmt.Println()
}

func (up *UnderlinePen) CreateCopy() framework.Product {
	copy := *up
	return &copy
}
