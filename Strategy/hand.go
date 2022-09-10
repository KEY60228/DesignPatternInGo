package main

type Hand int

const (
	Rock Hand = iota + 1
	Scissors
	Paper
)

func (m Hand) IsStrongerThan(y Hand) bool {
	return fight(m, y) == 1
}

func (m Hand) IsWeakerThan(y Hand) bool {
	return fight(m, y) == -1
}

func fight(a Hand, b Hand) int {
	if a == b {
		return 0
	}
	if (int(a)+1)%3 == int(b) {
		return 1
	}
	return -1
}
