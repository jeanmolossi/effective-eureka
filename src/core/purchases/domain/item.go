package domain

import "time"

type item struct {
	buyer             Buyer
	product           Product
	status            Status
	transaction       string
	warranyExpireDate time.Time
}

func NewItem(
	buyer Buyer,
	product Product,
	status Status,
	transaction string,
	warranyExpireDate time.Time,
) Item {
	return &item{
		buyer,
		product,
		status,
		transaction,
		warranyExpireDate,
	}
}

func (i *item) Buyer() Buyer {
	return i.buyer
}

func (i *item) Product() Product {
	return i.product
}

func (i *item) Status() Status {
	return i.status
}

func (i *item) Transaction() string {
	return i.transaction
}

func (i *item) WarranyExpireDate() time.Time {
	return i.warranyExpireDate
}
