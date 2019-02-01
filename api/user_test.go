// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_User_NotFoundError(t *testing.T) {
	nf := "Not Found"
	server, client, err := getTestServerAndClient(http.StatusNotFound, fmt.Sprintf(`{"code":%d, "error":"%s"}`, http.StatusNotFound, nf))
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusNotFound {
			t.Errorf("Expected %d got %d", http.StatusNotFound, v.Code)
		}
		if v.Message != nf {
			t.Errorf("Expected '%s' got '%s'", nf, v.Message)
		}
	default:
		t.Errorf("Expected *ErrorResponse got %v", v)
	}
	if u != nil {
		t.Errorf("Expected %v got %v", nil, u)
	}
}

func Test_User_ServerError(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusInternalServerError, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusInternalServerError {
			t.Errorf("Expected %d got %d", http.StatusInternalServerError, v.Code)
		}
	default:
		t.Errorf("Expected error got %v", v)
	}
	if u != nil {
		t.Errorf("Expected %v got %v", nil, u)
	}
}

func Test_User_UnAuthError(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusUnauthorized, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusUnauthorized {
			t.Errorf("Expected %d got %d", http.StatusUnauthorized, v.Code)
		}
	default:
		t.Errorf("Expected error got %v", v)
	}
	if u != nil {
		t.Errorf("Expected %v got %v", nil, u)
	}
}

func Test_User_InvalidID(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusUnauthorized, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(0)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", err, userIDError)
	}
	if u != nil {
		t.Errorf("Expected %v got %v", nil, u)
	}
}
