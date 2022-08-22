package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type HTMLBuilder struct {
	filename string
	s        *strings.Builder
}

func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{
		filename: "untitled.html",
		s:        &strings.Builder{},
	}
}

func (b *HTMLBuilder) MakeTitle(title string) {
	b.filename = title + ".html"
	b.s.WriteString("<!DOCTYPE html>\n")
	b.s.WriteString("<html>\n")
	b.s.WriteString("<head><tilte>")
	b.s.WriteString(title)
	b.s.WriteString("</title></head>\n")
	b.s.WriteString("<body>\n")
	b.s.WriteString("\t<h1>")
	b.s.WriteString(title)
	b.s.WriteString("</h1>\n\n")
}

func (b *HTMLBuilder) MakeString(str string) {
	b.s.WriteString("\t<p>")
	b.s.WriteString(str)
	b.s.WriteString("</p>\n\n")
}

func (b *HTMLBuilder) MakeItems(items []string) {
	b.s.WriteString("\t<ul>\n")
	for _, item := range items {
		b.s.WriteString("\t\t<li>")
		b.s.WriteString(item)
		b.s.WriteString("</li>\n")
	}
	b.s.WriteString("\t</ul>\n\n")
}

func (b *HTMLBuilder) Close() {
	b.s.WriteString("</body>\n")
	b.s.WriteString("</html>\n")

	f, err := os.Create(b.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = fmt.Fprint(f, b.s)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *HTMLBuilder) GetHTMLResult() string {
	return b.filename
}
