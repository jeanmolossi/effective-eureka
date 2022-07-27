package data

import (
	"time"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"
)

func ModelToDomain(model *ItemModel) domain.Item {
	return domain.NewItem(
		domain.NewBuyer(
			model.Buyer.Name,
			model.Buyer.Email,
		),
		domain.NewProduct(
			model.Product.ID,
			model.Product.Name,
		),
		model.Purchase.Status,
		model.Purchase.Transaction,
		time.Unix(model.Purchase.WarrantyIntTime, 0),
	)
}
