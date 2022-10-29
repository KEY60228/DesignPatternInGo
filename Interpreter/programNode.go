package main

import "fmt"

type ProgramNode struct {
	CommandListNode Node
}

func NewProgramNode() *ProgramNode {
	return &ProgramNode{}
}

func (n *ProgramNode) Parse(condition *Condition) error {
	condition.SkipToken("program")
	n.CommandListNode = NewCommandListNode()
	n.CommandListNode.Parse(condition)
	return nil
}

func (n *ProgramNode) String() string {
	return fmt.Sprintf("[program %s]", n.CommandListNode)
}
