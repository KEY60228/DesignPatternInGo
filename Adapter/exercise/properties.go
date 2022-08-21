package main

import (
	"fmt"
	"io"
	"strings"
)

type Properties map[string]string

func NewProperties() Properties {
	return make(Properties)
}

func (p Properties) Load(reader io.Reader) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	s := string(bytes)

	// とりあえず"\n", "="のみ
	properties := strings.Split(s, "\n")
	for _, property := range properties {
		keyValue := strings.Split(property, "=")
		p[keyValue[0]] = keyValue[1]
	}

	return nil
}

func (p Properties) Store(writer io.Writer, comments string) error {
	_, err := fmt.Fprint(writer, comments, "\n")
	if err != nil {
		return err
	}
	for key, value := range p {
		_, err := fmt.Fprintf(writer, "%s=%s\n", key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
