// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cmd cmdline client for the Baruwa REST API
package cmd

import (
	"fmt"
	"log"

	"github.com/baruwa-enterprise/baruwa-go/api"
	prettyjson "github.com/hokaccha/go-prettyjson"
	cli "github.com/jawher/mow.cli"
)

func organizationShow(cmd *cli.Cmd) {
}

func organizationCreate(cmd *cli.Cmd) {
}

func organizationUpdate(cmd *cli.Cmd) {
}

func organizationDelete(cmd *cli.Cmd) {
}

func organizationsList(cmd *cli.Cmd) {
	cmd.Action = func() {
		var b []byte
		var err error
		var c *api.Client
		var u *api.OrganizationList

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if u, err = c.GetOrganizations(nil); err != nil {
			log.Fatal(err)
		}

		if len(u.Items) == 0 {
			fmt.Println()
			return
		}

		if b, err = prettyjson.Marshal(u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
