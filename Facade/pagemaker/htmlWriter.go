package pagemaker

import (
	"fmt"
	"os"
)

type HTMLWriter struct {
	*os.File
}

func NewHTMLWriter(f *os.File) *HTMLWriter {
	return &HTMLWriter{f}
}

func (w *HTMLWriter) Title(title string) {
	w.WriteString("<!DOCTYPE html>")
	w.WriteString("<html>")
	w.WriteString("<head>")
	w.WriteString("<title>" + title + "</title>")
	w.WriteString("</head>")
	w.WriteString("<body>")
	w.WriteString("\n")
	w.WriteString("<h1>" + title + "</h1>")
	w.WriteString("\n")
}

func (w *HTMLWriter) Paragraph(msg string) {
	w.WriteString("<p>" + msg + "</p>")
	w.WriteString("\n")
}

func (w *HTMLWriter) Link(href string, caption string) {
	w.Paragraph(fmt.Sprintf("<a href=\"%s\">%s</a>", href, caption))
}

func (w *HTMLWriter) MailTo(mailAddr string, username string) {
	w.Link("mailto:"+mailAddr, username)
}

func (w *HTMLWriter) Close() {
	w.WriteString("</body>")
	w.WriteString("</html>")
	w.WriteString("\n")
	w.File.Close()
}
