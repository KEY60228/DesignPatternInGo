package framework

type Product interface {
	Use(string)
	CreateCopy() Product
}
