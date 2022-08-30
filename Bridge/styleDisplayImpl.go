package main

import "fmt"

type StyleDisplayImpl struct {
	start string
	style string
	end   string
	count int
}

func NewStyleDisplay(start string, style string, end string) *StyleDisplayImpl {
	return &StyleDisplayImpl{
		start: start,
		style: style,
		end:   end,
		count: 0,
	}
}

func (sd *StyleDisplayImpl) RawOpen() {
	fmt.Print(sd.start)
}

func (sd *StyleDisplayImpl) RawPrint() {
	for i := 0; i < sd.count; i++ {
		fmt.Print(sd.style)
	}
	sd.count++
}

func (sd *StyleDisplayImpl) RawClose() {
	fmt.Println(sd.end)
}
