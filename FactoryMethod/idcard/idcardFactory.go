package idcard

import (
	"fmt"

	"DesignPatternInGo/FactoryMethod/framework"
)

type IDCardFactory struct {
	id      int64
	mapping map[int64]string
}

func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{
		id:      0,
		mapping: make(map[int64]string),
	}
}

func (idf *IDCardFactory) CreateProduct(owner string) framework.Product {
	idf.id++
	idf.mapping[idf.id] = owner
	return NewIDCard(idf.id, owner)
}

func (idf *IDCardFactory) RegisterProduct(product framework.Product) {
	fmt.Println(product.ToString() + "を登録しました")
}
