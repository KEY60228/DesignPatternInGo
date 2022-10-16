package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type BigChar struct {
	charname string
	fontdata string
}

func NewBigChar(charname string) *BigChar {
	filename := filepath.Join("text", fmt.Sprintf("big%s.txt", charname))

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return &BigChar{
		charname: charname,
		fontdata: string(b),
	}
}

func (c *BigChar) Print() {
	fmt.Println(c.fontdata)
}
