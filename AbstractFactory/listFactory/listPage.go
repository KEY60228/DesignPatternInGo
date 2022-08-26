package listFactory

import (
	"fmt"
	"strings"

	"DesignPatternInGo/AbstractFactory/factory"
)

type ListPage struct {
	*factory.Page
}

func NewListPage(title string, author string) *ListPage {
	return &ListPage{factory.NewPage(title, author)}
}

func (p *ListPage) MakeHTML() string {
	var s strings.Builder
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html>\n")
	s.WriteString(fmt.Sprintf("<head><title>%s</title></head>\n", p.Title))
	s.WriteString("<body>\n")
	s.WriteString(fmt.Sprintf("<h1>%s</h1>\n", p.Title))
	s.WriteString("<ul>\n")
	for _, item := range p.Content {
		s.WriteString(item.MakeHTML())
	}
	s.WriteString("</ul>\n")
	s.WriteString("<hr>\n")
	s.WriteString(fmt.Sprintf("<address>%s</address>\n", p.Author))
	s.WriteString("</body>\n")
	s.WriteString("</html>\n")
	return s.String()
}
