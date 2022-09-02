package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatal("need more arguments")
	}

	seed1, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(fmt.Sprintf("%s is not int", flag.Arg(0)))
	}
	seed2, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(fmt.Sprintf("%s is not int", flag.Arg(1)))
	}

	player1 := NewPlayer("Taro", NewWinningStarategy(seed1))
	player2 := NewPlayer("Hana", NewProbStrategy(seed2))

	for i := 0; i < 10000; i++ {
		hand1 := player1.NextHand()
		hand2 := player2.NextHand()

		if hand1.IsStrongerThan(hand2) {
			player1.Win()
			player2.Lose()
		} else if hand1.IsWeakerThan(hand2) {
			player1.Lose()
			player2.Win()
		} else {
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println("Total Result")
	fmt.Println(player1.ToString())
	fmt.Println(player2.ToString())
}
