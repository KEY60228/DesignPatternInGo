package main

import (
	"flag"
	"fmt"
	"log"

	"DesignPatternInGo/AbstractFactory/divFactory"
	framework "DesignPatternInGo/AbstractFactory/factory"
	"DesignPatternInGo/AbstractFactory/listFactory"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	classname := flag.Arg(1)

	var factory framework.IFactory
	switch classname {
	case "ListFactory":
		factory = listFactory.NewListFactory()
	case "DivFactory":
		factory = divFactory.NewDivFactory()
	default:
		log.Fatal("factory name is not match")
	}

	if factory == nil {
		log.Fatal("factory is null")
	}

	blogTray := factory.CreateTray("Blog Site")
	for i := 1; i <= 3; i++ {
		blog := factory.CreateLink(fmt.Sprintf("Blog %d", i), fmt.Sprintf("https://example.com/blog%d", i))
		blogTray.Add(blog)
	}

	news3 := factory.CreateTray("News 3")
	news3.Add(factory.CreateLink("News 3 (US)", "https://example.com/news3us"))
	news3.Add(factory.CreateLink("News 3 (Japan)", "https://example.com/news3jp"))

	newsTray := factory.CreateTray("News Site")
	for i := 1; i <= 2; i++ {
		news := factory.CreateLink(fmt.Sprintf("News %d", i), fmt.Sprintf("https://example.com/news%d", i))
		newsTray.Add(news)
	}
	newsTray.Add(news3)

	page := factory.CreatePage("Blog and News", "Hiroshi Yuki")
	page.Add(blogTray)
	page.Add(newsTray)

	page.Output(filename, page)
}
