// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package api Golang bindings for Baruwa REST API

Usage:

	import "github.com/baruwa-enterprise/baruwa-go/api"

Create a new Baruwa API client, then use the various methods on the client to
access different parts of the Baruwa REST API. For example:

	import (
		"fmt"
		"log"
		"os"

		"github.com/baruwa-enterprise/baruwa-go/api"
		prettyjson "github.com/hokaccha/go-prettyjson"
	)

	func main() {
		var (
			err                 error
			c                   *api.Client
			u                   *api.UserList
			opts                *api.ListOptions
			serverURL, apiToken string
		)

		serverURL = os.Getenv("BARUWA_API_SERVER")
		apiToken = os.Getenv("BARUWA_API_TOKEN")

		if c, err = api.New(serverURL, apiToken, nil); err != nil {
			log.Fatal(err)
		}

		// page through users
		for {
			if u, err = c.GetUsers(opts); err != nil {
				log.Fatal(err)
			}

			if len(u.Items) == 0 {
				fmt.Println()
				break
			}

			if b, err = prettyjson.Marshal(u); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\n", b)

			if u.Links.Pages.Next == "" {
				break
			}

			opts = &api.ListOptions{
				Page: u.Links.Pages.Next,
			}
		}
	}


Refer to the https://github.com/baruwa-enterprise/baruwactl for a full
application built using this api for further usage information.

*/
package api
