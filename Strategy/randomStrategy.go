package main

import "math/rand"

type RandomStrategy struct{}

func NewRandomStrategy(seed int) *RandomStrategy {
	rand.Seed(int64(seed))
	return &RandomStrategy{}
}

func (s *RandomStrategy) NextHand() Hand {
	return Hand(rand.Intn(2) + 1)
}

func (s *RandomStrategy) Study(bool) {}
