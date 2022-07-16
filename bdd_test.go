package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"

	"github.com/cucumber/godog"
)

type apiFeature struct {
	baseURL string
	resp    *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", a.baseURL, endpoint),
		nil)
	if err != nil {
		return
	}

	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	ok := func(w http.ResponseWriter, data []byte) {
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, "%s", string(data))
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	a.resp.WriteHeader(res.StatusCode)
	if err != nil {
		fmt.Fprintf(a.resp, "Error: %s", err)
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(a.resp, "Error: %s", err)
		return
	}

	ok(a.resp, data)

	return
}

func (a *apiFeature) theStatusCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected status code %d, got %d", code, a.resp.Code)
	}

	return nil
}

func (a *apiFeature) theResponseMatchJSON(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{
		baseURL: os.Getenv("API_HOST"),
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	ctx.Step(`^I "(GET|POST|PUT|DELETE)" to "([^"]*)"$`, api.iRequestTo)
	ctx.Step(`^the status code received should be (\d+)$`, api.theStatusCodeShouldBe)
	ctx.Step(`^the response received should match json:$`, api.theResponseMatchJSON)

}
