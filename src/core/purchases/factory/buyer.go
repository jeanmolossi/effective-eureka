package factory

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"
)

type Buyer interface {
	Create(buyer map[string]interface{}) Buyer
	Build() (domain.Buyer, error)
}

type buyer struct {
	domain.Buyer
	err error
}

func NewBuyer() Buyer {
	return &buyer{}
}

func (b *buyer) Create(buyer map[string]interface{}) Buyer {
	if buyer == nil {
		b.err = errors.New("buyer is required")
		return b
	}

	var email string
	if buyerEmail, ok := buyer["email"].(string); ok {
		email = buyerEmail
	} else {
		b.err = errors.New("buyer email is not a string")
		return b
	}

	var name string
	if buyerName, ok := buyer["name"].(string); ok {
		name = buyerName
	} else {
		b.err = errors.New("buyer name is not a string")
		return b
	}

	b.Buyer = domain.NewBuyer(name, email)

	return b
}

func (b *buyer) Build() (domain.Buyer, error) {
	if b.err != nil {
		return nil, b.err
	}

	return b.Buyer, nil
}
