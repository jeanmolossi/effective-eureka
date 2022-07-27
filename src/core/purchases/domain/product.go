package domain

type product struct {
	id   uint32
	name string
}

func NewProduct(id uint32, name string) Product {
	return &product{id, name}
}

func (p *product) ID() uint32 {
	return p.id
}

func (p *product) Name() string {
	return p.name
}
