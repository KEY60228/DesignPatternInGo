package divFactory

import (
	"fmt"
	"strings"

	"DesignPatternInGo/AbstractFactory/factory"
)

type DivPage struct {
	*factory.Page
}

func NewDivPage(title string, author string) *DivPage {
	return &DivPage{factory.NewPage(title, author)}
}

func (p *DivPage) MakeHTML() string {
	var s strings.Builder
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html>\n")
	s.WriteString("<head>\n")
	s.WriteString(fmt.Sprintf("<title>%s</title>\n", p.Title))
	s.WriteString("<style>\n")
	s.WriteString("div.Tray { padding: 0.5em; margin-left: 5em; border: 1px solid black; }\n")
	s.WriteString("div.LINK { padding: 0.5em; background-color: lightgray; }\n")
	s.WriteString("</style>\n")
	s.WriteString("</head>\n")
	s.WriteString("<body>\n")
	s.WriteString(fmt.Sprintf("<h1>%s</h1>\n", p.Title))
	for _, item := range p.Content {
		s.WriteString(item.MakeHTML())
	}
	s.WriteString("<hr>\n")
	s.WriteString(fmt.Sprintf("<address>%s</address>\n", p.Author))
	s.WriteString("</body>\n")
	s.WriteString("</html>\n")
	return s.String()
}
