package main

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf(maxsize int) *BookShelf {
	return &BookShelf{
		books: make([]*Book, maxsize),
		last:  0,
	}
}

func (b *BookShelf) GetBookAt(index int) *Book {
	return b.books[index]
}

func (b *BookShelf) AppendBook(book *Book) {
	b.books[b.last] = book
	b.last++
}

func (b *BookShelf) GetLength() int {
	return b.last
}

func (b *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(b)
}
