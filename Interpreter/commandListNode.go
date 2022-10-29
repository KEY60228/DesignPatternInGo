package main

import (
	"errors"
	"fmt"
)

type CommandListNode struct {
	list []Node
}

func NewCommandListNode() *CommandListNode {
	return &CommandListNode{
		list: []Node{},
	}
}

func (n *CommandListNode) Parse(condition *Condition) error {
	for {
		if condition.CurrentToken() == "" {
			return errors.New("error: Missing 'end'")
		} else if condition.CurrentToken() == "end" {
			condition.SkipToken("end")
			break
		} else {
			commandNode := NewCommandNode()
			commandNode.Parse(condition)
			n.list = append(n.list, commandNode)
		}
	}
	return nil
}

func (n *CommandListNode) String() string {
	return fmt.Sprintf("%s", n.list)
}
