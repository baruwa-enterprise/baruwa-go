// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"

	"github.com/baruwa-enterprise/baruwa-go/api"
)

const (
	unsetVarsErr = "Endpoint or token variables not set"
)

// GetClient returns a new Client
func GetClient() (c *api.Client, err error) {
	if *serverURL == "" || *apiToken == "" {
		err = fmt.Errorf(unsetVarsErr)
		return
	}
	c, err = api.New(*serverURL, *apiToken, nil)
	return
}
