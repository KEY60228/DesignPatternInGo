package framework

type IFactory interface {
	CreateProduct(string) Product
	RegisterProduct(Product)
}

type Factory struct {
	IFactory
}

func NewFactory(factory IFactory) *Factory {
	return &Factory{factory}
}

func (f *Factory) Create(owner string) Product {
	p := f.CreateProduct(owner)
	f.RegisterProduct(p)
	return p
}
