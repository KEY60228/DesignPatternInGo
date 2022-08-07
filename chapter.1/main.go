package main

import (
	"fmt"
	"log"
)

func main() {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(NewBook("Bible"))
	bookShelf.AppendBook(NewBook("Cinderella"))
	bookShelf.AppendBook(NewBook("Daddy-Long-Legs"))

	itr := bookShelf.Iterator()
	for itr.HasNext() {
		book, err := itr.Next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(book.GetName())
	}
	fmt.Println()

	for _, book := range bookShelf.books {
		fmt.Println(book.GetName())
	}
}
