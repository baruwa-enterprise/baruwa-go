// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/jawher/mow.cli"
)

var (
	apiToken  *string
	serverURL *string
)

// CLI struct for main
type CLI struct {
	*cli.Cli
}

// NewCLI initializes new command line interface
func NewCLI() *CLI {
	c := &CLI{cli.App("baruwa", "A cmdline client for the Baruwa REST API.")}
	c.Spec = "-k -s"

	apiToken = c.String(cli.StringOpt{
		Name:      "k api-token",
		Desc:      "Baruwa API OAUTH Token",
		EnvVar:    "BARUWA_API_TOKEN",
		HideValue: true,
	})
	serverURL = c.String(cli.StringOpt{
		Name:      "s server-url",
		Desc:      "Baruwa server url",
		EnvVar:    "BARUWA_API_SERVER",
		HideValue: true,
	})

	return c
}
