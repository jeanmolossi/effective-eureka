package data

import (
	"errors"
	"net/url"
	"os"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/domain"
	shared "github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/jeanmolossi/effective-eureka/src/pkg/http"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
)

type GetPurchases interface {
	Get(email, transaction string) ([]domain.Item, error)
}

type getPurchases struct {
	httpClient *http.HttpClientWithAuth

	logger logger.Logger
}

func NewGetPurchases() GetPurchases {
	return &getPurchases{
		httpClient: http.NewHttpClientWithAuth(
			os.Getenv("HOTMART_API"),
		),

		logger: logger.NewLogger(),
	}
}

func (g *getPurchases) Get(email, transaction string) ([]domain.Item, error) {
	g.logger.Infoln("getting purchases from email", email, "and transaction", transaction)

	if os.Getenv("ENVIRONMENT") == "development" {
		email = os.Getenv("HOTMART_SANDBOX_EMAIL")
		transaction = os.Getenv("HOTMART_SANDBOX_TID")
	}

	client := g.httpClient.RequestWith(http.RequestParams{
		URL: "/payments/api/v1/sales/history",
		Query: url.Values{
			"buyer_email": {email},
			"transaction": {transaction},
		},
	}).Do()

	var items RootModel
	if err := client.JSON(&items); err != nil {
		return nil, err
	}

	if items.PageInfo.TotalResults == 0 {
		return nil, shared.NewNotFoundErr(
			errors.New("no purchases found"),
		)
	}

	domainItems := make([]domain.Item, len(items.Items))
	for i, item := range items.Items {
		domainItems[i] = ModelToDomain(&item)
	}

	return domainItems, nil
}
