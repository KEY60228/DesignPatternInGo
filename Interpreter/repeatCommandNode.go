package main

import "fmt"

type RepeatCommandNode struct {
	num             int
	commandListNode Node
}

func NewRepeatCommandNode() *RepeatCommandNode {
	return &RepeatCommandNode{}
}

func (n *RepeatCommandNode) Parse(condition *Condition) error {
	condition.SkipToken("repeat")
	num, err := condition.CurrentNumber()
	if err != nil {
		return err
	}
	n.num = num
	condition.NextToken()
	n.commandListNode = NewCommandListNode()
	n.commandListNode.Parse(condition)
	return nil
}

func (n *RepeatCommandNode) String() string {
	return fmt.Sprintf("[repeat %d %s]", n.num, n.commandListNode)
}
