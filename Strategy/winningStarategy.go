package main

import "math/rand"

type WinningStrategy struct {
	won      bool
	prevHand Hand
}

func NewWinningStarategy(seed int) *WinningStrategy {
	rand.Seed(int64(seed))
	return &WinningStrategy{
		won:      false,
		prevHand: 0,
	}
}

func (s *WinningStrategy) NextHand() Hand {
	if !s.won {
		s.prevHand = Hand(rand.Intn(2) + 1)
	}
	return s.prevHand
}

func (s *WinningStrategy) Study(win bool) {
	s.won = win
}
