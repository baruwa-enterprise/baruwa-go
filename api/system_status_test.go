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

func TestGetSystemStatus(t *testing.T) {
	data := `
	{
		"inbound": 0,
		"status": true,
		"total": {
		  "spam": 0,
		  "highspam": 0,
		  "lowspam": 0,
		  "infected": 0,
		  "clean": 16,
		  "total": 16,
		  "virii": 0
		},
		"outbound": 0
	  }
	`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	_, err = client.GetSystemStatus()
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
