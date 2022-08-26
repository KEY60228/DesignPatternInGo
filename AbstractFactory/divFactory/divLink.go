package divFactory

import (
	"fmt"

	"DesignPatternInGo/AbstractFactory/factory"
)

type DivLink struct {
	*factory.Link
}

func NewDivLink(caption string, url string) *DivLink {
	return &DivLink{factory.NewLink(caption, url)}
}

func (l *DivLink) MakeHTML() string {
	return fmt.Sprintf("<div class=\"LINK\"><a href=\"%s\">%s</a></div>\n", l.URL, l.Caption)
}
