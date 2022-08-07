package main

type Iterable interface {
	Iterator() *Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (*Book, error)
}
