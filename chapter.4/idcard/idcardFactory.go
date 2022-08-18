package idcard

import (
	"DesignPatternInGo/chapter.4/framework"
	"fmt"
)

type IDCardFactory struct {
}

func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{}
}

func (idf *IDCardFactory) CreateProduct(owner string) framework.Product {
	return NewIDCard(owner)
}

func (idf *IDCardFactory) RegisterProduct(product framework.Product) {
	fmt.Println(product.ToString() + "を登録しました")
}
