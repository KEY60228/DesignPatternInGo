package main

import "math/rand"

type ProbStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func NewProbStrategy(seed int) *ProbStrategy {
	rand.Seed(int64(seed))
	return &ProbStrategy{
		prevHandValue:    0,
		currentHandValue: 0,
		history: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (s *ProbStrategy) NextHand() Hand {
	var sum int
	for _, i := range s.history {
		for _, j := range i {
			sum += j
		}
	}

	bet := rand.Intn(sum-1) + 1
	var handValue int
	if bet < s.history[s.currentHandValue][0] {
		handValue = 0
	} else if bet < s.history[s.currentHandValue][0]+s.history[s.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}

	s.prevHandValue = s.currentHandValue
	s.currentHandValue = handValue
	return Hand(handValue)
}

func (s *ProbStrategy) Study(win bool) {
	if win {
		s.history[s.prevHandValue][s.currentHandValue]++
	} else {
		s.history[s.prevHandValue][(s.currentHandValue+1)%3]++
		s.history[s.prevHandValue][(s.currentHandValue+2)%3]++
	}
}
