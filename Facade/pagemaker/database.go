package pagemaker

import (
	"io"
	"log"
	"os"
	"strings"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) LoadData(dbname string) map[string]string {
	filename := dbname + ".txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	rows := strings.Split(string(body), "\n")
	for _, row := range rows {
		col := strings.Split(row, "=")
		m[col[0]] = col[1]
	}

	return m
}
