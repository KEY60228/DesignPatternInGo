package factory

type ILink interface {
	MakeHTML() string
}

type Link struct {
	*Item
	URL string
}

func NewLink(caption string, url string) *Link {
	return &Link{
		Item: NewItem(caption),
		URL:  url,
	}
}
