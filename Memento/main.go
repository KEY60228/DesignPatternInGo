package main

import (
	"fmt"
	"os"
	"time"
)

const filename = "game.dat"

func main() {
	gamer := &Gamer{}
	memento := &Memento{}

	_, err := os.Stat(filename)
	if err != nil {
		gamer = NewGamer(100)
		memento = gamer.CreateMemento()
	} else {
		memento = memento.LoadFromFile(filename)
		gamer.RestoreMemento(memento)
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("=== %d\n", i)
		fmt.Printf("現状: %s", gamer)

		gamer.Bet()

		fmt.Printf("所持金は%d円になりました\n", gamer.GetMoney())

		if gamer.GetMoney() > memento.GetMoney() {
			fmt.Println("「だいぶ増えたし現在の状態を保存しておこう…」")
			memento = gamer.CreateMemento()
			memento.SaveToFile(filename, memento)
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			fmt.Println("「だいぶ減ったから前の状態まで戻そう！」")
			gamer.RestoreMemento(memento)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
