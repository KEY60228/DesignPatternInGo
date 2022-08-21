package main

import "fmt"

func main() {
	fmt.Println("Start")
	singleton := NewSingleton()
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	if obj1 == obj2 {
		fmt.Println("同じインスタンス")
	} else {
		fmt.Println("異なるインスタンス")
	}
	fmt.Println("End")
}
