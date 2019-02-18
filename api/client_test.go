// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestServer(code int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprint(w, body)
	}))
}

func getTestClient(endpoint string, opts *Options) (c *Client, e error) {
	c, e = New(endpoint, "test-token", opts)
	return
}

func getTestServerAndClient(code int, body string) (*httptest.Server, *Client, error) {
	server := getTestServer(code, body)
	client, err := getTestClient(server.URL, nil)
	return server, client, err
}

func TestNewErrors(t *testing.T) {
	c, e := New("", "", nil)
	if e == nil {
		t.Fatalf("An error should be returned as endpoint is required")
	}
	if e.Error() != endpointError {
		t.Errorf("Expected %s got %s", endpointError, e)
	}
	if c != nil {
		t.Errorf("Expected %v got %v", nil, c)
	}
	_, e = New("http://[fe80::%31]:8080/", "", nil)
	if e == nil {
		t.Fatalf("An error should be returned as endpoint is required")
	}
}

func TestNew(t *testing.T) {
	bu := "https://baruwa.example.com"
	c, e := New(bu, "", nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	if c.BaseURL.String() != bu {
		t.Errorf("Expected %s got %s", bu, c.BaseURL)
	}
}

func TestNewOpts(t *testing.T) {
	ua := "test-client"
	client := http.DefaultClient
	opts := &Options{
		UserAgent:  ua,
		HTTPClient: client,
	}
	bu := "https://baruwa.example.com"
	c, e := New(bu, "", opts)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	if c.BaseURL.String() != bu {
		t.Errorf("Expected %s got %s", bu, c.BaseURL)
	}
	if c.UserAgent != ua {
		t.Errorf("Expected %s got %s", ua, c.UserAgent)
	}
	if c.client != client {
		t.Errorf("Expected %v got %v", client, c.client)
	}
}

func TestRequestOptions(t *testing.T) {
	bu := "https://baruwa.example.com"
	c, e := New(bu, "test-token", nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	page := "https://baruwa.example.com/api/v1/users?page=2"
	opts := &ListOptions{
		Page: page + "&?test=1",
	}
	req, e := c.newRequest(http.MethodGet, apiPath("users"), opts, nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	if req.URL.String() != page {
		t.Errorf("Expected %s got %s", page, req.URL.String())
	}
	opts.Page = "https://b2.example.com/api/v1/users?page=2"
	req, e = c.newRequest(http.MethodGet, apiPath("users"), opts, nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	if req.URL.String() == page {
		t.Errorf("Expected %s got %s", "https://baruwa.example.com/api/v1/users", req.URL.String())
	}
}

func TestGetAccessTokenError(t *testing.T) {
	bu := "https://baruwa.example.com"
	c, e := New(bu, "", nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	if c.BaseURL.String() != bu {
		t.Errorf("Expected %s got %s", bu, c.BaseURL)
	}
	_, e = c.GetAccessToken("", "")
	if e == nil {
		t.Fatalf("An error should be returned")
	}
	if e.Error() != clientIDError {
		t.Errorf("Expected %s got %s", clientIDError, e)
	}
	_, e = c.GetAccessToken("test", "")
	if e == nil {
		t.Fatalf("An error should be returned")
	}
	if e.Error() != clientSecretError {
		t.Errorf("Expected %s got %s", clientSecretError, e)
	}
	server, client, e := getTestServerAndClient(http.StatusForbidden, ``)
	if e != nil {
		t.Fatalf("An error should not be returned: %s", e)
	}
	defer server.Close()
	_, e = client.GetAccessToken("test-id", "test-secret")
	if e == nil {
		t.Fatalf("An error should be returned")
	}
}

func TestGetAccessTokenOK(t *testing.T) {
	accessToken := "teUU1ApyURKgLpNctloeLpX7WAkCirgOYbTNwXQygqerOSEdxeeIRobTxNbR"
	data := fmt.Sprintf(`
	{
		"access_token": "%s",
		"token_type": "Bearer",
		"expires_in": 3600,
		"refresh_token": "Vcc0xyvhrTyoYIIiqs4LHhCoD4JYuvyodpGMXLPSirpM62KJD5qS6m6Zr0eT",
		"scope": "act-read act-create act-update act-delete dom-read dom-create dom-update dom-delete org-read org-create org-update org-delete sta-read"
	}
	`, accessToken)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	defer server.Close()
	token, err := client.GetAccessToken("test-id", "test-secret")
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if token.Token != accessToken {
		t.Errorf("Expected %s got %s", accessToken, token.Token)
	}
}

func TestApiPath(t *testing.T) {
	p := "users"
	expected := fmt.Sprintf("/api/%s/%s", APIVersion, p)
	g := apiPath(p)
	if g != expected {
		t.Errorf("Expected %s got %s", expected, g)
	}
}

func TestNewRequestErrors(t *testing.T) {
	tp := "../dir/"
	bu := "https://baruwa.example.com"
	c, e := New(bu, "test-token", nil)
	if e != nil {
		t.Fatalf("An error should not be returned")
	}
	_, e = c.newRequest(http.MethodGet, tp, nil, nil)
	if e == nil {
		t.Fatalf("An error should be returned")
	}
	e = c.get(tp, nil, nil)
	if e == nil {
		t.Fatalf("An error should be returned")
	}
}
