package factory

type ITray interface {
	Add(IItem)
	MakeHTML() string
}

type Tray struct {
	*Item
	Items []IItem
}

func NewTray(caption string) *Tray {
	return &Tray{
		Item:  NewItem(caption),
		Items: []IItem{},
	}
}

func (t *Tray) Add(item IItem) {
	t.Items = append(t.Items, item)
}
