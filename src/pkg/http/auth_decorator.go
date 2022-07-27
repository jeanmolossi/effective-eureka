// TODO: should be refactored to better dev experience
package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type HttpClientWithAuth struct {
	client      *HttpClient
	accessToken string
}

func NewHttpClientWithAuth(baseURL string) *HttpClientWithAuth {
	return &HttpClientWithAuth{
		NewHttpClient(baseURL),
		getAccessToken(),
	}
}

func (h *HttpClientWithAuth) RequestWith(params RequestParams) *HttpClientWithAuth {
	if h.accessToken == "" {
		err := h.authBeforeRequest()
		if err != nil {
			return h
		}
	}

	h.client = h.client.RequestWith(params).WithHeaders(
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", h.accessToken),
		},
	)

	return h
}

func (h *HttpClientWithAuth) Do() *HttpClient {
	h.client.Do()

	if h.client.Code == 401 {
		log.Println("revalidating token")
		h.authBeforeRequest()
		return h.client.Do()
	}

	return h.client
}

func (h *HttpClientWithAuth) authBeforeRequest() error {
	query := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {os.Getenv("HOTMART_CLIENT_ID")},
		"client_secret": {os.Getenv("HOTMART_CLIENT_SECRET")},
	}

	client := NewHttpClient(os.Getenv("HOTMART_AUTH_API"))
	client = client.RequestWith(RequestParams{
		Method: "POST",
		URL:    "/security/oauth/token",
	}).WithHeaders(map[string]string{
		"Authorization": os.Getenv("HOTMART_BASIC_TOKEN"),
	})
	client.request.URL.RawQuery = query.Encode()
	res := map[string]interface{}{}
	err := client.Do().JSON(&res)
	if err != nil {
		return err
	}

	h.saveSession(res["access_token"])

	return nil
}

func (h *HttpClientWithAuth) saveSession(token interface{}) error {
	file, err := os.Create(os.TempDir() + "/session.secret")
	if err != nil {
		return err
	}

	if strToken, ok := token.(string); ok {
		file.WriteString(strToken)
		h.accessToken = strToken
	}

	return nil
}

func getAccessToken() string {
	file, err := os.Open(os.TempDir() + "/session.secret")
	if err != nil {
		log.Println(err)
		return ""
	}

	accessToken, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(accessToken)
}
