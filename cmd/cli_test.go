// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"os"
	"testing"
)

func TestNewCLI(t *testing.T) {
	token := "test-token"
	su := "https://baruwa.example.com"
	os.Setenv("BARUWA_API_TOKEN", token)
	os.Setenv("BARUWA_API_SERVER", su)
	_ = NewCLI()
	if *apiToken != token {
		t.Errorf("Expected %s got %s", token, *apiToken)
	}
	if *serverURL != su {
		t.Errorf("Expected %s got %s", su, *serverURL)
	}
	os.Unsetenv("BARUWA_API_TOKEN")
	os.Unsetenv("BARUWA_API_SERVER")
	_ = NewCLI()
	if *apiToken != "" {
		t.Errorf("Expected %s got %s", "", *apiToken)
	}
	if *serverURL != "" {
		t.Errorf("Expected %s got %s", "", *serverURL)
	}
}
