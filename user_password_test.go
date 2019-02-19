// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"net/http"
	"testing"
)

func TestChangeUserPasswordError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.ChangeUserPassword(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != pwFormError {
		t.Errorf("Expected '%s' got '%s'", pwFormError, err)
	}
	f := &PasswordForm{}
	err = client.ChangeUserPassword(0, f)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
}

func TestChangeUserPasswordOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusAccepted, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	f := &PasswordForm{
		Password1: "password",
		Password2: "password",
	}
	err = client.ChangeUserPassword(1, f)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
}
