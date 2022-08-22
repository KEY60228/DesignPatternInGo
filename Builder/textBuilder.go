package main

import "strings"

type TextBuilder struct {
	s *strings.Builder
}

func NewTextBuilder(s string) *TextBuilder {
	return &TextBuilder{s: &strings.Builder{}}
}

func (b *TextBuilder) MakeTitle(title string) {
	b.s.WriteString("======================\n")
	b.s.WriteString("『")
	b.s.WriteString(title)
	b.s.WriteString("』\n\n")
}

func (b *TextBuilder) MakeString(str string) {
	b.s.WriteString("■")
	b.s.WriteString(str)
	b.s.WriteString("\n\n")
}

func (b *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		b.s.WriteString("\t・")
		b.s.WriteString(item)
		b.s.WriteString("\n")
	}
	b.s.WriteString("\n")
}

func (b *TextBuilder) Close() {
	b.s.WriteString("======================\n")
}

func (b *TextBuilder) GetTextResult() string {
	return b.s.String()
}
