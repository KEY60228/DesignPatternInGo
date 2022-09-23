package pagemaker

import (
	"fmt"
	"log"
	"os"
)

type PageMaker struct{}

func NewPageMaker() *PageMaker {
	return &PageMaker{}
}

func (pm *PageMaker) MakeWelcomePage(mailaddr string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mailprop := NewDatabase().LoadData("maildata")
	username := mailprop[mailaddr]

	writer := NewHTMLWriter(f)
	defer writer.Close()

	writer.Title(username + "'s web page")
	writer.Paragraph(fmt.Sprintf("Welcome to %s's web page!", username))
	writer.Paragraph("Nice to meet you!")
	writer.MailTo(mailaddr, username)
	fmt.Printf("%s is created for %s (%s)\n", filename, mailaddr, username)
}
