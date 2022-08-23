package divFactory

import "DesignPatternInGo/AbstractFactory/factory"

type DivFactory struct{}

func NewDivFactory() *DivFactory {
	return &DivFactory{}
}

func (f *DivFactory) CreateLink(caption string, url string) factory.ILink {
	return NewDivLink(caption, url)
}

func (f *DivFactory) CreateTray(caption string) factory.ITray {
	return NewDivTray(caption)
}

func (f *DivFactory) CreatePage(title string, author string) factory.IPage {
	return NewDivPage(title, author)
}
