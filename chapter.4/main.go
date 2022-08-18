package main

import (
	"DesignPatternInGo/chapter.4/framework"
	"DesignPatternInGo/chapter.4/idcard"
)

func main() {
	factory := framework.NewFactory(idcard.NewIDCardFactory())

	card1 := factory.Create("Hiroshi Yuki")
	card2 := factory.Create("Tomura")
	card3 := factory.Create("Hanako Sato")

	card1.Use()
	card2.Use()
	card3.Use()
}
