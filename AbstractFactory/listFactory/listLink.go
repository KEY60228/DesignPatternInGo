package listFactory

import (
	"fmt"

	"DesignPatternInGo/AbstractFactory/factory"
)

type ListLink struct {
	*factory.Link
}

func NewListLink(caption string, url string) *ListLink {
	return &ListLink{factory.NewLink(caption, url)}
}

func (l *ListLink) MakeHTML() string {
	return fmt.Sprintf("\t<li><a href=\"%s\">%s</a></li>\n", l.URL, l.Caption)
}
