package divFactory

import (
	"fmt"
	"strings"

	"DesignPatternInGo/AbstractFactory/factory"
)

type DivTray struct {
	*factory.Tray
}

func NewDivTray(caption string) *DivTray {
	return &DivTray{factory.NewTray(caption)}
}

func (t *DivTray) MakeHTML() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("<p><b>%s</b></p>\n", t.Caption))
	s.WriteString("<div class=\"TRAY\">\n")
	for _, item := range t.Items {
		s.WriteString(item.MakeHTML())
	}
	s.WriteString("</div>\n")
	return s.String()
}
