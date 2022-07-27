// Package http is a wrapper around the standard net/http package, missing docs.
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type HttpClient struct {
	client  *http.Client
	baseURL string

	request  *http.Request
	response *http.Response
	Code     int
	result   bytes.Buffer

	err error
}

func NewHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		client:  &http.Client{},
		baseURL: baseURL,

		Code:   0,
		result: bytes.Buffer{},
	}
}

type RequestParams struct {
	Method string
	URL    string
	Body   io.Reader
	Query  url.Values
}

func (h *HttpClient) RequestWith(params RequestParams) *HttpClient {
	if params.Method == "" {
		params.Method = "GET"
	}

	var err error
	h.request, err = http.NewRequest(
		params.Method,
		fmt.Sprintf("%s%s", h.baseURL, params.URL),
		params.Body,
	)
	if err != nil {
		h.err = err
	}

	if params.Query != nil {
		h.request.URL.RawQuery = params.Query.Encode()
	}

	return h
}

func (h *HttpClient) WithHeaders(headers map[string]string) *HttpClient {
	for key, value := range headers {
		h.request.Header.Set(key, value)
	}

	return h
}

func (h *HttpClient) Do() *HttpClient {
	// early exit if there was an error in Params
	if h.err != nil {
		return h
	}

	var err error
	h.response, err = h.client.Do(h.request)
	if err != nil {
		h.err = err
		return h
	}

	h.Code = h.response.StatusCode
	h.err = h.unmarshal()

	return h
}

func (h *HttpClient) JSON(target interface{}) error {
	if h.err != nil {
		log.Println("early aborting JSON unmarshal")
		return h.err
	}
	defer h.result.Reset()
	return json.Unmarshal(h.result.Bytes(), target)
}

func (h *HttpClient) Buffer() (bytes.Buffer, error) {
	if h.err != nil {
		return bytes.Buffer{}, h.err
	}
	defer h.result.Reset()
	return h.result, nil
}

func (h *HttpClient) Response() (*http.Response, error) {
	if h.err != nil {
		return nil, h.err
	}

	return h.response, nil
}

func (h *HttpClient) unmarshal() error {
	if h.err != nil {
		return h.err
	}

	if h.response.Body == nil {
		_, h.err = h.result.WriteString("")
		return nil
	}

	bytes, err := io.ReadAll(h.response.Body)
	if err != nil {
		return err
	}

	defer h.response.Body.Close()

	_, err = h.result.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
