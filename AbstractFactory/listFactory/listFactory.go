package listFactory

import "DesignPatternInGo/AbstractFactory/factory"

type ListFactory struct{}

func NewListFactory() *ListFactory {
	return &ListFactory{}
}

func (f *ListFactory) CreateLink(caption string, url string) factory.ILink {
	return NewListLink(caption, url)
}

func (f *ListFactory) CreateTray(caption string) factory.ITray {
	return NewListTray(caption)
}

func (f *ListFactory) CreatePage(title string, author string) factory.IPage {
	return NewListPage(title, author)
}
