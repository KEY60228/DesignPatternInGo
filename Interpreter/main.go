package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("program.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	commands := strings.Split(string(b), "\n")
	for _, command := range commands {
		fmt.Printf("command = \"%s\"\n", command)
		node := NewProgramNode()
		node.Parse(NewCondition(command))
		fmt.Printf("node = %s\n", node)
	}
}
