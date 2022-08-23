package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "text" {
		textBuilder := NewTextBuilder("")
		director := NewDirector(textBuilder)
		director.construct()
		result := textBuilder.GetTextResult()
		fmt.Println(result)
		return
	}
	if flag.Arg(0) == "html" {
		htmlBuilder := NewHTMLBuilder()
		director := NewDirector(htmlBuilder)
		director.construct()
		result := htmlBuilder.GetHTMLResult()
		fmt.Println("HTMLファイル: " + result + "が作成されました")
		return
	}

	fmt.Println("usage: go run . [text|html]")
}
