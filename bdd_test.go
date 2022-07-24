package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"

	"github.com/cucumber/godog"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

type apiFeature struct {
	baseURL string
	resp    *httptest.ResponseRecorder
}

var token string

func DoRegister() error {
	baseURL := os.Getenv("API_HOST")
	registerEndpoint := fmt.Sprintf("%s%s", baseURL, "/students/register")

	resp, err := http.Post(registerEndpoint, "application/json",
		strings.NewReader(`{"password": "123456789","username": "john@doe.com"}`),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		bytes, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(bytes), "already exists") {
			return errors.New("failed to register student")
		}
	}

	return nil
}

func DoLogin() error {
	baseURL := os.Getenv("API_HOST")
	loginEndpoint := fmt.Sprintf("%s%s", baseURL, "/auth/login")

	resp, err := http.Post(loginEndpoint, "application/json",
		strings.NewReader(`{"password": "123456789","username": "john@doe.com"}`),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to login student")
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type body struct {
		AccessToken string `json:"access_token"`
	}

	var b body
	err = json.Unmarshal(response, &b)
	if err != nil {
		return err
	}

	token = b.AccessToken

	return nil
}

func addTokenToRequest(req *http.Request) {
	req.Header.Add("Authorization", token)
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

	addTokenToRequest(req)

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

func (a *apiFeature) iRequestToWithPayload(method, endpoint string, payload *godog.DocString) (err error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", a.baseURL, endpoint),
		strings.NewReader(payload.Content))
	if err != nil {
		return
	}

	addTokenToRequest(req)
	req.Header.Add("Content-Type", "application/json")

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

func (a *apiFeature) thereAreAny(tableName string, data *godog.Table) error {
	var fields []string
	var marks []string

	head := data.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
		marks = append(marks, "?")
	}

	dbConn := shared.NewDbConnection()
	err := dbConn.Connect()
	if err != nil {
		return err
	}

	for i := 1; i < len(data.Rows); i++ {
		var vals []interface{}
		for _, cell := range data.Rows[i].Cells {
			vals = append(vals, cell.Value)
		}

		stmt := dbConn.DB().Exec(
			`INSERT INTO `+tableName+` (`+strings.Join(fields, ",")+`) VALUES (`+strings.Join(marks, ",")+`)`,
			vals...,
		)

		if err := stmt.Error; err != nil {
			return err
		}
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{
		baseURL: os.Getenv("API_HOST"),
	}

	err := DoRegister()
	if err != nil {
		panic(err)
	}

	// if has no token should login
	if token == "" {
		if err := DoLogin(); err != nil {
			panic(err)
		}
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	ctx.Step(`^I "(GET|POST|PUT|DELETE)" to "([^"]*)"$`, api.iRequestTo)
	ctx.Step(`^I "(POST|PUT)" to "([^"]*)" with:$`, api.iRequestToWithPayload)
	ctx.Step(`^the status code received should be (\d+)$`, api.theStatusCodeShouldBe)
	ctx.Step(`^the response received should match json:$`, api.theResponseMatchJSON)
	ctx.Step(`^there are "([^"]*)" with:$`, api.thereAreAny)
}
