package idcard

import "fmt"

type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	fmt.Println(owner + "のカードを作ります")
	return &IDCard{
		owner: owner,
	}
}

func (i *IDCard) Use() {
	fmt.Println(i.ToString() + "を使います")
}

func (i *IDCard) ToString() string {
	return "[IDCard:" + i.owner + "]"
}
