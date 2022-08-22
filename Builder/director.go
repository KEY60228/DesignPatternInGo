package main

type Director struct {
	Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) construct() {
	d.MakeTitle("Greeting")
	d.MakeString("一般的な挨拶")
	d.MakeItems([]string{"How are you?", "Hello.", "Hi."})
	d.MakeString("時間帯に応じた挨拶")
	d.MakeItems([]string{"Good morning.", "Good afternoon.", "Good evening."})
	d.Close()
}
