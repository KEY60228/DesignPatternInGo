package main

import "errors"

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (bi *BookShelfIterator) HasNext() bool {
	return bi.index < bi.bookShelf.GetLength()
}

func (bi *BookShelfIterator) Next() (*Book, error) {
	if !bi.HasNext() {
		return nil, errors.New("no such element")
	}
	book := bi.bookShelf.GetBookAt(bi.index)
	bi.index++
	return book, nil
}
