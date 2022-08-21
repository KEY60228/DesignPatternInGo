package idcard

import "fmt"

type IDCard struct {
	id    int64
	owner string
}

func NewIDCard(id int64, owner string) *IDCard {
	fmt.Println(owner + "のカードを作ります")
	return &IDCard{
		id:    id,
		owner: owner,
	}
}

func (i *IDCard) Use() {
	fmt.Println(i.ToString() + "を使います")
}

func (i *IDCard) ToString() string {
	return fmt.Sprintf("[ID %d: %s]", i.id, i.owner)
}
