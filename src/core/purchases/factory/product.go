package factory

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"
)

type Product interface {
	Create(product map[string]interface{}) Product
	Build() (domain.Product, error)
}

type product struct {
	domain.Product

	err error
}

func NewProduct() Product {
	return &product{}
}

func (p *product) Create(product map[string]interface{}) Product {
	if product == nil {
		p.err = errors.New("product is required")
		return p
	}

	var id uint32
	if productID, ok := product["id"].(uint32); ok {
		id = productID
	} else {
		p.err = errors.New("product id is not a uint32")
		return p
	}

	var name string
	if productName, ok := product["name"].(string); ok {
		name = productName
	} else {
		p.err = errors.New("product name is not a string")
		return p
	}

	p.Product = domain.NewProduct(id, name)

	return p
}

func (p *product) Build() (domain.Product, error) {
	if p.err != nil {
		return nil, p.err
	}

	return p.Product, nil
}
