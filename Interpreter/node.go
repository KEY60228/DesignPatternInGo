package main

type Node interface {
	Parse(*Condition) error
}
