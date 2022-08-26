package factory

import (
	"fmt"
	"log"
	"os"
)

type IPage interface {
	Add(IItem)
	Output(string, IPage)
	MakeHTML() string
}

type Page struct {
	Title   string
	Author  string
	Content []IItem
}

func NewPage(title string, author string) *Page {
	return &Page{
		Title:   title,
		Author:  author,
		Content: []IItem{},
	}
}

func (p *Page) Add(item IItem) {
	p.Content = append(p.Content, item)
}

func (p *Page) Output(filename string, page IPage) {
	filename += ".html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = fmt.Fprint(f, page.MakeHTML())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%sを作成しました\n", filename)
}
