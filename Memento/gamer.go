package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fruit int

const (
	Apple Fruit = iota + 1
	Grape
	Banana
	Orange
)

func (f Fruit) String() string {
	m := map[Fruit]string{
		Apple:  "リンゴ",
		Grape:  "ぶどう",
		Banana: "バナナ",
		Orange: "みかん",
	}
	return m[f]
}

type Gamer struct {
	money  int
	fruits []string
}

func NewGamer(money int) *Gamer {
	rand.Seed(time.Now().UnixNano())
	return &Gamer{
		money:  money,
		fruits: []string{},
	}
}

func (g *Gamer) GetMoney() int {
	return g.money
}

func (g *Gamer) Bet() {
	dice := rand.Intn(6) + 1
	switch dice {
	case 1:
		g.money += 100
		fmt.Println("所持金が増えました")
	case 2:
		g.money /= 2
		fmt.Println("所持金が半分になりました")
	case 6:
		f := g.GetFruit()
		g.fruits = append(g.fruits, f)
		fmt.Printf("フルーツ %s をもらいました\n", f)
	default:
		fmt.Println("何も起きませんでした")
	}
}

func (g *Gamer) CreateMemento() *Memento {
	m := NewMemento(g.money)
	for _, f := range g.fruits {
		m.AddFruit(f)
	}
	return m
}

func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.money
	g.fruits = memento.fruits
}

func (g Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %s]\n", g.money, g.fruits)
}

func (g *Gamer) GetFruit() string {
	f := Fruit(rand.Intn(4) + 1)
	return f.String()
}
