package shared

import "fmt"

type HttpMeta struct {
	Page         uint16 `json:"page" example:"2"`
	ItemsPerPage int    `json:"items_per_page" example:"1"`
	NextPage     string `json:"next_page,omitempty" example:"http://localhost:8080/resource?page=3&items_per_page=1"`
	PrevPage     string `json:"prev_page,omitempty" example:"http://localhost:8080/resource?page=1&items_per_page=1"`
}

func GetMeta(baseURL string, page uint16, itemsPerPage, itemsCount int) HttpMeta {
	prevPage := ""
	if page == 0 {
		page = 1
	}

	if itemsPerPage == 0 {
		itemsPerPage = 10
	}

	if page > 1 {
		prevPage = fmt.Sprintf("%s?page=%d&items_per_page=%d", baseURL, page-1, itemsPerPage)
	}

	nextPage := ""
	if itemsCount >= itemsPerPage {
		nextPage = fmt.Sprintf("%s?page=%d&items_per_page=%d", baseURL, page+1, itemsPerPage)
	}

	return HttpMeta{
		Page:         page,
		ItemsPerPage: itemsPerPage,
		NextPage:     nextPage,
		PrevPage:     prevPage,
	}
}
