// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"testing"

	"github.com/baruwa-enterprise/baruwa-go/api"
)

func TestGetClientError(t *testing.T) {
	var err error
	*serverURL = ""
	*apiToken = ""

	if _, err = GetClient(); err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != unsetVarsErr {
		t.Errorf("Expected %s got %s", unsetVarsErr, err.Error())
	}
}

func TestGetClientOK(t *testing.T) {
	var err error
	var c *api.Client
	token := "test-token"
	su := "https://baruwa.example.com"
	*serverURL = su
	*apiToken = token

	if c, err = GetClient(); err != nil {
		t.Fatalf("An error should not be returned")
	}
	if c.BaseURL.String() != su {
		t.Errorf("Expected %s got %s", su, c.BaseURL)
	}
}
