package main

import "DesignPatternInGo/Facade/pagemaker"

func main() {
	pagemaker.NewPageMaker().MakeWelcomePage("hyuki@example.com", "welcome.html")
}
