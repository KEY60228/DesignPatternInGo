package main

type Strategy interface {
	NextHand() Hand
	Study(bool)
}
