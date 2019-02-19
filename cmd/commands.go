// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	cli "github.com/jawher/mow.cli"
)

// RegisterCommands registers all CLI commands
func (c *CLI) RegisterCommands() {
	// user
	c.Command("user", "manage user accounts", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of a user account", userShow)
		cmd.Command("create", "create a new user account", userCreate)
		cmd.Command("update", "update a user account", userUpdate)
		cmd.Command("delete", "delete a user account", userDelete)
		cmd.Command("alias", "manage user alias addresses", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of an alias address", aliasShow)
			cmd.Command("create", "create a new alias address", aliasCreate)
			cmd.Command("update", "update an alias address", aliasUpdate)
			cmd.Command("delete", "delete an alias address", aliasDelete)
		})
	})
	// users
	c.Command("users", "list user accounts", usersList)
	// domain
	c.Command("domain", "manage domains", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of a domain", domainShow)
		cmd.Command("create", "create a new domain", domainCreate)
		cmd.Command("update", "update a domain", domainUpdate)
		cmd.Command("delete", "delete a domain", domainDelete)
		cmd.Command("alias", "manage alias domains", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a domain alias", domainAliasShow)
			cmd.Command("create", "create a new domain alias", domainAliasCreate)
			cmd.Command("update", "update a domain alias", domainAliasUpdate)
			cmd.Command("delete", "delete a domain alias", domainAliasDelete)
		})
		cmd.Command("aliases", "list domain aliases", domainAliasList)
		cmd.Command("deliveryserver", "manage domain delivery servers", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a domain delivery server", domainDSShow)
			cmd.Command("create", "create a new domain delivery server", domainDSCreate)
			cmd.Command("update", "update a domain delivery server", domainDSUpdate)
			cmd.Command("delete", "delete a domain delivery server", domainDSDelete)
		})
		cmd.Command("deliveryservers", "list domain delivery servers", domainDSList)
		cmd.Command("userdeliveryserver", "manage user delivery servers", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a user delivery server", userDSShow)
			cmd.Command("create", "create a new user delivery server", userDSCreate)
			cmd.Command("update", "update a user delivery server", userDSUpdate)
			cmd.Command("delete", "delete a user delivery server", userDSDelete)
		})
		cmd.Command("userdeliveryservers", "list user delivery servers", userDSList)
		cmd.Command("authsetting", "manage authentication settings", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of an authentication setting", domainASShow)
			cmd.Command("create", "create a new authentication setting", domainASCreate)
			cmd.Command("update", "update a authentication setting", domainASUpdate)
			cmd.Command("delete", "delete a authentication setting", domainASDelete)
			cmd.Command("ldapsetting", "manage ldap settings", func(cmd *cli.Cmd) {
				cmd.Command("show", "show detailed information of an ldap setting", ldapShow)
				cmd.Command("create", "create a new ldap setting", ldapCreate)
				cmd.Command("update", "update a ldap setting", ldapUpdate)
				cmd.Command("delete", "delete a ldap setting", ldapDelete)
			})
			cmd.Command("radiussetting", "manage radius settings", func(cmd *cli.Cmd) {
				cmd.Command("show", "show detailed information of an radius setting", radiusShow)
				cmd.Command("create", "create a new radius setting", radiusCreate)
				cmd.Command("update", "update a radius setting", radiusUpdate)
				cmd.Command("delete", "delete a radius setting", radiusDelete)
			})
		})
		cmd.Command("authsettings", "list authentication settings", domainASList)
		cmd.Command("smarthost", "manage smarthosts", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a smarthost", domainSMShow)
			cmd.Command("create", "create a new smarthost", domainSMCreate)
			cmd.Command("update", "update a smarthost", domainSMUpdate)
			cmd.Command("delete", "delete a smarthost", domainSMDelete)
		})
		cmd.Command("smarthosts", "list smarthosts", domainSMList)
	})
	// domains
	c.Command("domains", "list domains", domainsList)
	// organization
	c.Command("organization", "manage organizations", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of an organization", organizationShow)
		cmd.Command("create", "create a new organization", organizationCreate)
		cmd.Command("update", "update a organization", organizationUpdate)
		cmd.Command("delete", "delete a organization", organizationDelete)
		cmd.Command("smarthost", "manage smarthosts", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a smarthost", organizationSMShow)
			cmd.Command("create", "create a new smarthost", organizationSMCreate)
			cmd.Command("update", "update a smarthost", organizationSMUpdate)
			cmd.Command("delete", "delete a smarthost", organizationSMDelete)
		})
		cmd.Command("smarthosts", "list smarthosts", organizationSMList)
		cmd.Command("fallbackserver", "manage fallback servers", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a fallback server", organizationFSShow)
			cmd.Command("create", "create a new fallback server", organizationFSCreate)
			cmd.Command("update", "update a fallback server", organizationFSUpdate)
			cmd.Command("delete", "delete a fallback server", organizationFSDelete)
		})
		cmd.Command("fallbackservers", "list fallback servers", organizationFSList)
		cmd.Command("relaysetting", "manage relay settings", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a relay setting", organizationRelayShow)
			cmd.Command("create", "create a new relay setting", organizationRelayCreate)
			cmd.Command("update", "update a relay setting", organizationRelayUpdate)
			cmd.Command("delete", "delete a relay setting", organizationRelayDelete)
		})
	})
	// organizations
	c.Command("organizations", "list organizations", organizationsList)
	// systemstatus
	c.Command("systemstatus", "show system status", systemStatus)
}
