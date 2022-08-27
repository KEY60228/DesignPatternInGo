package main

type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{NewDisplay(impl)}
}

func (cd *CountDisplay) MultiDisplay(times int) {
	cd.Open()
	for i := 0; i < times; i++ {
		cd.Print()
	}
	cd.Close()
}
