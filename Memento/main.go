package main

import (
	"fmt"
	"time"
)

func main() {
	gamer := NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 100; i++ {
		fmt.Printf("=== %d\n", i)
		fmt.Printf("現状: %s", gamer)

		gamer.Bet()

		fmt.Printf("所持金は%d円になりました\n", gamer.GetMoney())

		if gamer.GetMoney() > memento.GetMoney() {
			fmt.Println("「だいぶ増えたし現在の状態を保存しておこう…」")
			memento = gamer.CreateMemento()
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			fmt.Println("「だいぶ減ったから前の状態まで戻そう！」")
			gamer.RestoreMemento(memento)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
