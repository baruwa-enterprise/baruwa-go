// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main cmdline client for the Baruwa REST API
package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("baruwa", "A cmdline client for the Baruwa REST API.")
	token  = app.Flag("api-token", "OAUTH Token.").OverrideDefaultFromEnvar("BARUWA_API_TOKEN").Required().Short('k').String()
	server = app.Flag("server-url", "Baruwa server url without the path").OverrideDefaultFromEnvar("BARUWA_API_HOST").Required().Short('s').URL()
	// users       = app.Command("users", "Account management")
	// usersAction = users.Flag("action", "Action to perform").Required().Short('a').HintOptions("list", "get", "create", "update", "delete").Enum("list", "get", "create", "update", "delete")
	// usersList                 = users.Command("list", "List accounts")
	// usersGet                  = users.Command("get", "Retrieve an existing Account")
	// usersCreate               = users.Command("create", "Create a new Account")
	// usersUpdate               = users.Command("update", "Update an Account")
	// usersDelete               = users.Command("delete", "Delete an Account")
	// passwords                 = app.Command("passwords", "Password management")
	// passwordsUpdate           = passwords.Command("update", "Change account password")
	// addresses                 = app.Command("addresses", "Alias Address management")
	// addressesGet              = addresses.Command("get", "Retrieve an existing Alias Address")
	// addressesCreate           = addresses.Command("create", "Create an Alias Address")
	// addressesUpdate           = addresses.Command("update", "Update an Alias Address")
	// addressesDelete           = addresses.Command("delete", "Delete an Alias Address")
	// domains                   = app.Command("domains", "Domain management")
	// domainsList               = domains.Command("list", "List Domains")
	// domainsGet                = domains.Command("get", "Retrieve a Domain")
	// domainsCreate             = domains.Command("create", "Create a new Domain")
	// domainsUpdate             = domains.Command("update", "Update a Domain")
	// domainsDelete             = domains.Command("delete", "Delete a Domain")
	// domainAliases             = app.Command("domainaliases", "Domain Alias management")
	// domainAliasesList         = domainAliases.Command("list", "Listing Domain Aliases")
	// domainAliasesGet          = domainAliases.Command("get", "Retrieve a Domain Alias")
	// domainAliasesCreate       = domainAliases.Command("create", "Create a Domain Alias")
	// domainAliasesUpdate       = domainAliases.Command("update", "Update a Domain Alias")
	// domainAliasesDelete       = domainAliases.Command("delete", "Delete a Domain Alias")
	// deliveryServers           = app.Command("deliveryservers", "Delivery server management")
	// deliveryServersList       = deliveryServers.Command("list", "List Delivery servers")
	// deliveryServersGet        = deliveryServers.Command("get", "Retrieve a Delivery server")
	// deliveryServersCreate     = deliveryServers.Command("create", "Create a Delivery server")
	// deliveryServersUpdate     = deliveryServers.Command("update", "Update a Delivery server")
	// deliveryServersDelete     = deliveryServers.Command("delete", "Delete a Delivery server")
	// userDeliveryServers       = app.Command("userdeliveryservers", "User Delivery server management")
	// userDeliveryServersList   = userDeliveryServers.Command("list", "List User Delivery servers")
	// userDeliveryServersGet    = userDeliveryServers.Command("get", "Retrieve a User Delivery server")
	// userDeliveryServersCreate = userDeliveryServers.Command("create", "Create a User Delivery server")
	// userDeliveryServersUpdate = userDeliveryServers.Command("update", "Update a User Delivery server")
	// userDeliveryServersDelete = userDeliveryServers.Command("delete", "Delete a User Delivery server")
)

func addSubCommand(app *kingpin.Application, name string, description string) {
	c := app.Command(name, description).Action(func(c *kingpin.ParseContext) error {
		fmt.Printf("Would have run command %s => %s.\n", name, description)
		return nil
	})
	c.Flag("nop-flag", "Example of a flag with no options").Bool()
}

func main() {
	addSubCommand(app, "users", "Account management")
	addSubCommand(app, "domains", "Domain management")
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
