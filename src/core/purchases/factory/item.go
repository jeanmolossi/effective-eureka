package factory

import (
	"errors"
	"time"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"
)

type Item interface {
	Create(item map[string]interface{}) Item
	WithBuyer(item map[string]interface{}) Item
	WithProduct(item map[string]interface{}) Item
	Build() (domain.Item, error)
}

type item struct {
	domain.Item
	err error

	buyer   Buyer
	product Product
}

func NewItem() Item {
	return &item{
		buyer:   NewBuyer(),
		product: NewProduct(),
	}
}

func (i *item) Create(item map[string]interface{}) Item {
	if item == nil {
		i.err = errors.New("item is required")
		return i
	}

	if item["purchase"] == nil {
		i.err = errors.New("item purchase data not found")
		return i
	}

	var purchase map[string]interface{}
	if itemPurchase, ok := item["purchase"].(map[string]interface{}); !ok {
		i.err = errors.New("item purchase data is not a map")
		return i
	} else {
		purchase = itemPurchase
	}

	var status string
	if purchase["status"] == nil {
		i.err = errors.New("item purchase status not found")
		return i
	} else {
		status = purchase["status"].(string)
	}

	var transaction string
	if purchase["transaction"] == nil {
		i.err = errors.New("item purchase transaction not found")
		return i
	} else {
		transaction = purchase["transaction"].(string)
	}

	var warrantyExpires uint64
	if purchase["warranty_expires"] == nil {
		i.err = errors.New("item purchase warranty_expires not found")
		return i
	} else {
		warrantyExpires = purchase["warranty_expires"].(uint64)
	}

	warrantyInTime := time.Unix(int64(warrantyExpires), 0)

	i.Item = domain.NewItem(
		nil, nil, // Buyer and Product not set at moment
		domain.Status(status),
		transaction,
		warrantyInTime,
	)

	return i
}

func (i *item) WithBuyer(buyerMap map[string]interface{}) Item {
	if buyerMap == nil {
		i.err = errors.New("item is required")
		return i
	}

	buyer, err := i.buyer.Create(buyerMap).Build()
	if err != nil {
		i.err = err
		return i
	}

	i.Item = domain.NewItem(
		buyer, nil, // Product not set at moment
		i.Item.Status(),
		i.Item.Transaction(),
		i.Item.WarranyExpireDate(),
	)

	return i
}

func (i *item) WithProduct(item map[string]interface{}) Item {
	if item == nil {
		i.err = errors.New("item is required")
		return i
	}

	if item["product"] == nil {
		i.err = errors.New("item product data not found")
		return i
	}

	var productMap map[string]interface{}
	if itemProduct, ok := item["product"].(map[string]interface{}); !ok {
		i.err = errors.New("item product data is not a map")
		return i
	} else {
		productMap = itemProduct
	}

	product, err := i.product.Create(productMap).Build()
	if err != nil {
		i.err = err
		return i
	}

	i.Item = domain.NewItem(
		i.Item.Buyer(), product, // Product not set at moment
		i.Item.Status(),
		i.Item.Transaction(),
		i.Item.WarranyExpireDate(),
	)

	return i
}

func (i *item) Build() (domain.Item, error) {
	if i.err != nil {
		return nil, i.err
	}

	return i.Item, nil
}
