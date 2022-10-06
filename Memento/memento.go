package main

type Memento struct {
	money  int
	fruits []string
}

func NewMemento(money int) *Memento {
	return &Memento{
		money:  money,
		fruits: []string{},
	}
}

func (m *Memento) GetMoney() int {
	return m.money
}

func (m *Memento) AddFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) GetFruits() []string {
	return m.fruits
}
