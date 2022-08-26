package listFactory

import (
	"strings"

	"DesignPatternInGo/AbstractFactory/factory"
)

type ListTray struct {
	*factory.Tray
}

func NewListTray(caption string) *ListTray {
	return &ListTray{factory.NewTray(caption)}
}

func (t *ListTray) MakeHTML() string {
	var s strings.Builder
	s.WriteString("<li>\n")
	s.WriteString(t.Caption)
	s.WriteString("\n<ul>\n")
	for _, item := range t.Items {
		s.WriteString(item.MakeHTML())
	}
	s.WriteString("</ul>\n")
	s.WriteString("</li>\n")
	return s.String()
}
