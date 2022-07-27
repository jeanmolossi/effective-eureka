package data

import "github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"

type RootModel struct {
	Items    []ItemModel `json:"items"`
	PageInfo struct {
		NextPageToken  string `json:"next_page_token"`
		ResultsPerPage uint8  `json:"results_per_page"`
		TotalResults   uint8  `json:"total_results"`
	} `json:"page_info"`
}

type ItemModel struct {
	Buyer struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"buyer"`
	Product struct {
		ID   uint32 `json:"id"`
		Name string `json:"name"`
	} `json:"product"`
	Purchase struct {
		Status          domain.Status `json:"status"`
		Transaction     string        `json:"transaction"`
		WarrantyIntTime int64         `json:"warranty_expire_date"`
	} `json:"purchase"`
}
