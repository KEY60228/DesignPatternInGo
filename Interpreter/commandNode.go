package main

import "fmt"

type CommandNode struct {
	node Node
}

func NewCommandNode() *CommandNode {
	return &CommandNode{}
}

func (n *CommandNode) Parse(condition *Condition) error {
	if condition.CurrentToken() == "repeat" {
		n.node = NewRepeatCommandNode()
		n.node.Parse(condition)
	} else {
		n.node = NewPrimitiveCommandNode()
		n.node.Parse(condition)
	}
	return nil
}

func (n *CommandNode) String() string {
	return fmt.Sprintf("%s", n.node)
}
