// Copyright 2020 Mohammad Abdolirad
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"context"
	"github.com/atkrad/wait4x/pkg/checker"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var CreateRequestErr = checker.NewError("create the request error", "debug")
var DoCallErr = checker.NewError("do call error", "debug")
var KeyExistenceErr = checker.NewError("the key doesn't exist", "info")
var GetKeyErr = checker.NewError("get key", "debug")
var KeyValueExistenceErr = checker.NewError("the key and desired value doesn't exist", "info")

// Option configures an HTTP.
type Option func(s *HTTP)

// HTTP represents HTTP checker
type HTTP struct {
	address          string
	timeout          time.Duration
	expectBody       string
	expectStatusCode int
	*checker.LogAware
}

// New creates the HTTP checker
func New(address string, opts ...Option) checker.Checker {
	h := &HTTP{
		address:  address,
		timeout:  time.Second * 5,
		LogAware: &checker.LogAware{},
	}

	// apply the list of options to HTTP
	for _, opt := range opts {
		opt(h)
	}

	return h
}

// WithTimeout configures a time limit for requests made by the HTTP client
func WithTimeout(timeout time.Duration) Option {
	return func(h *HTTP) {
		h.timeout = timeout
	}
}

// WithExpectBody configures response body expectation
func WithExpectBody(body string) Option {
	return func(h *HTTP) {
		h.expectBody = body
	}
}

// WithExpectStatusCode configures response status code expectation
func WithExpectStatusCode(code int) Option {
	return func(h *HTTP) {
		h.expectStatusCode = code
	}
}

// Check checks HTTP connection
func (h *HTTP) Check(ctx context.Context) error {
	var httpClient = &http.Client{
		Timeout: h.timeout,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", h.address, nil)
	if err != nil {
		return CreateRequestErr.WithWrap(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return DoCallErr.WithWrap(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			h.Logger().Debug(err)
		}
	}()

	//if h.httpResponseCodeExpectation(h.expectStatusCode, resp) && h.httpResponseBodyExpectation(h.expectBody, resp) {
	//	return nil
	//}

	//return false

	return nil
}

func (h *HTTP) httpResponseCodeExpectation(expectStatusCode int, resp *http.Response) bool {
	if expectStatusCode == 0 {
		return true
	}

	h.Logger().InfoWithFields("Checking http response code expectation", map[string]interface{}{"actual": resp.StatusCode, "expect": expectStatusCode})

	return expectStatusCode == resp.StatusCode
}

func (h *HTTP) httpResponseBodyExpectation(expectBody string, resp *http.Response) bool {
	if expectBody == "" {
		return true
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.Logger().Fatal(err)
	}

	bodyString := string(bodyBytes)

	// TODO: Logging full body response in debug level.

	h.Logger().InfoWithFields("Checking http response body expectation", map[string]interface{}{"actual": h.truncateString(bodyString, 50), "expect": expectBody})

	matched, _ := regexp.MatchString(expectBody, bodyString)
	return matched
}

func (h *HTTP) truncateString(str string, num int) string {
	truncatedStr := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		truncatedStr = str[0:num] + "..."
	}

	return truncatedStr
}
