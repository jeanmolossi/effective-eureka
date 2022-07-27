package handler

import "github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"

type HttpBuyer struct {
	Email string `json:"email" example:"olga.dev@email.com"`
	Name  string `json:"name" example:"Olga Doutel Frade"`
}

type HttpProduct struct {
	ID   uint32 `json:"id" example:"1564852"`
	Name string `json:"name" example:"Product02"`
}

type HttpPurchase struct {
	Status          domain.Status `json:"status" example:"COMPLETE"`
	Transaction     string        `json:"transaction" example:"HP17715690036011"`
	WarrantyIntTime int64         `json:"warranty_expire_date" example:"1625022000000"`
}

type HttpPurchasesItem struct {
	HttpBuyer    HttpBuyer    `json:"buyer"`
	HttpProduct  HttpProduct  `json:"product"`
	HttpPurchase HttpPurchase `json:"purchase"`
}

func NewHttpPurchasesItem(item domain.Item) *HttpPurchasesItem {
	return &HttpPurchasesItem{
		HttpBuyer{
			Email: item.Buyer().Email(),
			Name:  item.Buyer().Name(),
		},
		HttpProduct{
			ID:   item.Product().ID(),
			Name: item.Product().Name(),
		},
		HttpPurchase{
			Status:          item.Status(),
			Transaction:     item.Transaction(),
			WarrantyIntTime: item.WarranyExpireDate().Unix(),
		},
	}
}

type HttpPurchasesOk struct {
	Items []HttpPurchasesItem `json:"items"`
}

func NewHttpPurchaseOk(items []domain.Item) *HttpPurchasesOk {
	httpPurchasesItems := make([]HttpPurchasesItem, len(items))
	for i, item := range items {
		httpPurchasesItems[i] = *NewHttpPurchasesItem(item)
	}
	return &HttpPurchasesOk{
		Items: httpPurchasesItems,
	}
}
