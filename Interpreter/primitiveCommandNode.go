package main

import (
	"errors"
	"fmt"
)

type PrimitiveCommandNode struct {
	name string
}

func NewPrimitiveCommandNode() *PrimitiveCommandNode {
	return &PrimitiveCommandNode{}
}

func (n *PrimitiveCommandNode) Parse(condition *Condition) error {
	n.name = condition.CurrentToken()
	if n.name == "" {
		return errors.New("error: missing <primitive command>")
	} else if n.name != "go" && n.name != "right" && n.name != "left" {
		return fmt.Errorf("error: unknown <primitive command>: '%s'", n.name)
	}
	condition.SkipToken(n.name)
	return nil
}

func (n *PrimitiveCommandNode) String() string {
	return n.name
}
