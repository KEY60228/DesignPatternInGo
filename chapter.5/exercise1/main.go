package main

import "fmt"

func main() {
	tm := NewTicketMaker()
	tm1 := tm.GetInstance()
	tm2 := tm.GetInstance()
	fmt.Println(tm1.GetNextTicketNumber())
	fmt.Println(tm2.GetNextTicketNumber())
}
