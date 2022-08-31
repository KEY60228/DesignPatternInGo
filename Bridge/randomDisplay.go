package main

import (
	"math/rand"
	"time"
)

type RandomDisplay struct {
	*Display
}

func NewRandomDisplay(impl DisplayImpl) *RandomDisplay {
	return &RandomDisplay{NewDisplay(impl)}
}

func (rd *RandomDisplay) RandomDo(times int) {
	rd.Open()
	for i := 0; i < times; i++ {
		rand.Seed(time.Now().UnixNano())
		if i == rand.Intn(10) {
			rd.Print()
		}
	}
	rd.Close()
}
