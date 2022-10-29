package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Condition struct {
	tokens    []string
	lastToken string
	index     int
}

func NewCondition(text string) *Condition {
	c := &Condition{
		tokens: strings.Split(text, " "),
		index:  0,
	}
	c.lastToken = c.NextToken()
	return c
}

func (c *Condition) NextToken() string {
	if c.index < len(c.tokens)-1 {
		c.index++
		c.lastToken = c.tokens[c.index]
	} else {
		c.lastToken = ""
	}
	return c.lastToken
}

func (c *Condition) CurrentToken() string {
	return c.lastToken
}

func (c *Condition) SkipToken(token string) error {
	if c.CurrentToken() == "" {
		return fmt.Errorf("error: '%s' is expected, but not more tokens is found", token)
	} else if c.CurrentToken() != token {
		return fmt.Errorf("error: '%s' is expected, but '%s' is found", token, c.CurrentToken())
	}
	c.NextToken()
	return nil
}

func (c *Condition) CurrentNumber() (int, error) {
	if c.CurrentToken() == "" {
		return 0, errors.New("error: not more token")
	}
	num, err := strconv.Atoi(c.CurrentToken())
	if err != nil {
		return 0, err
	}
	return num, nil
}
