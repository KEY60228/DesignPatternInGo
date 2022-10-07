package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

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

func (m Memento) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Money  int      `json:"money"`
		Fruits []string `json:"fruits"`
	}{
		Money:  m.money,
		Fruits: m.fruits,
	})
	return v, err
}

func (m *Memento) UnmarshalJSON(b []byte) error {
	m2 := &struct {
		Money  int      `json:"money"`
		Fruits []string `json:"fruits"`
	}{}
	err := json.Unmarshal(b, m2)
	m.money = m2.Money
	m.fruits = m2.Fruits
	return err
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

func (m *Memento) SaveToFile(filename string, memento *Memento) bool {
	f, err := os.Create(filename)
	if err != nil {
		return false
	}
	defer f.Close()

	j, err := json.Marshal(memento)
	if err != nil {
		return false
	}

	if _, err := fmt.Fprint(f, string(j)); err != nil {
		return false
	}

	return true
}

func (m *Memento) LoadFromFile(filename string) *Memento {
	var memento Memento

	f, err := os.Open(filename)
	if err != nil {
		// return nil
		log.Fatal(err)
	}
	defer f.Close()

	j, err := io.ReadAll(f)
	if err != nil {
		// return nil
		log.Fatal(err)
	}

	if err := json.Unmarshal(j, &memento); err != nil {
		// return nil
		log.Fatal(err)
	}

	return &memento
}
