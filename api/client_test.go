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
	c, e = New("http://[fe80::%31]:8080/", "", nil)
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

func TestApiPath(t *testing.T) {
	p := "users"
	expected := fmt.Sprintf("/api/%s/%s", APIVersion, p)
	g := apiPath(p)
	if g != expected {
		t.Errorf("Expected %s got %s", expected, g)
	}
}
