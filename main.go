// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main cmdline client for the Baruwa REST API
package main

import (
	"github.com/baruwa-enterprise/baruwa-go/api"
	"github.com/integrii/flaggy"
)

var (
	user                 *flaggy.Subcommand
	userGet              *flaggy.Subcommand
	userCreate           *flaggy.Subcommand
	userUpdate           *flaggy.Subcommand
	userDelete           *flaggy.Subcommand
	domain               *flaggy.Subcommand
	domainGet            *flaggy.Subcommand
	domainCreate         *flaggy.Subcommand
	domainUpdate         *flaggy.Subcommand
	domainDelete         *flaggy.Subcommand
	domainAlias          *flaggy.Subcommand
	domainAliasGet       *flaggy.Subcommand
	domainAliasCreate    *flaggy.Subcommand
	domainAliasUpdate    *flaggy.Subcommand
	domainAliasDelete    *flaggy.Subcommand
	domainDS             *flaggy.Subcommand
	domainDSGet          *flaggy.Subcommand
	domainDSCreate       *flaggy.Subcommand
	domainDSUpdate       *flaggy.Subcommand
	domainDSDelete       *flaggy.Subcommand
	userDS               *flaggy.Subcommand
	userDSGet            *flaggy.Subcommand
	userDSCreate         *flaggy.Subcommand
	userDSUpdate         *flaggy.Subcommand
	userDSDelete         *flaggy.Subcommand
	domainAS             *flaggy.Subcommand
	domainASGet          *flaggy.Subcommand
	domainASCreate       *flaggy.Subcommand
	domainASUpdate       *flaggy.Subcommand
	domainASDelete       *flaggy.Subcommand
	domainSM             *flaggy.Subcommand
	domainSMGet          *flaggy.Subcommand
	domainSMCreate       *flaggy.Subcommand
	domainSMUpdate       *flaggy.Subcommand
	domainSMDelete       *flaggy.Subcommand
	ldap                 *flaggy.Subcommand
	ldapGet              *flaggy.Subcommand
	ldapCreate           *flaggy.Subcommand
	ldapUpdate           *flaggy.Subcommand
	ldapDelete           *flaggy.Subcommand
	radius               *flaggy.Subcommand
	radiusGet            *flaggy.Subcommand
	radiusCreate         *flaggy.Subcommand
	radiusUpdate         *flaggy.Subcommand
	radiusDelete         *flaggy.Subcommand
	organization         *flaggy.Subcommand
	organizationGet      *flaggy.Subcommand
	organizationCreate   *flaggy.Subcommand
	organizationUpdate   *flaggy.Subcommand
	organizationDelete   *flaggy.Subcommand
	relay                *flaggy.Subcommand
	relayGet             *flaggy.Subcommand
	relayCreate          *flaggy.Subcommand
	relayUpdate          *flaggy.Subcommand
	relayDelete          *flaggy.Subcommand
	fallbackServer       *flaggy.Subcommand
	fallbackServerGet    *flaggy.Subcommand
	fallbackServerCreate *flaggy.Subcommand
	fallbackServerUpdate *flaggy.Subcommand
	fallbackServerDelete *flaggy.Subcommand
	organizationSM       *flaggy.Subcommand
	organizationSMGet    *flaggy.Subcommand
	organizationSMCreate *flaggy.Subcommand
	organizationSMUpdate *flaggy.Subcommand
	organizationSMDelete *flaggy.Subcommand
	organizations        *flaggy.Subcommand
	system               *flaggy.Subcommand
	users                *flaggy.Subcommand
	domains              *flaggy.Subcommand
	domainAliases        *flaggy.Subcommand
	userDSs              *flaggy.Subcommand
	domainASs            *flaggy.Subcommand
	domainDSs            *flaggy.Subcommand
	domainSMs            *flaggy.Subcommand
	fallbackServers      *flaggy.Subcommand
	organizationSMs      *flaggy.Subcommand
)

func init() {
	flaggy.SetName("baruwa")
	flaggy.SetDescription("A cmdline client for the Baruwa REST API.")
	flaggy.SetVersion(api.Version)
	user = flaggy.NewSubcommand("user")
	user.Description = "Account management"
	userGet = flaggy.NewSubcommand("show")
	userGet.Description = "Show account details"
	userCreate = flaggy.NewSubcommand("create")
	userCreate.Description = "Create account"
	userUpdate = flaggy.NewSubcommand("update")
	userUpdate.Description = "Update account"
	userDelete = flaggy.NewSubcommand("delete")
	userDelete.Description = "Delete account"

	users = flaggy.NewSubcommand("users")
	users.Description = "List Accounts"

	domain = flaggy.NewSubcommand("domain")
	domain.Description = "Domain management"
	domainGet = flaggy.NewSubcommand("show")
	domainGet.Description = "Show domain details"
	domainCreate = flaggy.NewSubcommand("create")
	domainCreate.Description = "Create domain"
	domainUpdate = flaggy.NewSubcommand("update")
	domainUpdate.Description = "Update domain"
	domainDelete = flaggy.NewSubcommand("delete")
	domainDelete.Description = "Delete domain"

	domainAliases = flaggy.NewSubcommand("aliases")
	domainAliases.Description = "List domain aliases"

	domainAlias = flaggy.NewSubcommand("alias")
	domainAlias.Description = "Domain alias management"
	domainAliasGet = flaggy.NewSubcommand("show")
	domainAliasGet.Description = "Show domain alias details"
	domainAliasCreate = flaggy.NewSubcommand("create")
	domainAliasCreate.Description = "Create alias domain"
	domainAliasUpdate = flaggy.NewSubcommand("update")
	domainAliasUpdate.Description = "Update alias domain"
	domainAliasDelete = flaggy.NewSubcommand("delete")
	domainAliasDelete.Description = "Delete alias domain"

	domainDSs = flaggy.NewSubcommand("deliverservers")
	domainDSs.Description = "List domain delivery servers"

	domainDS = flaggy.NewSubcommand("deliverserver")
	domainDS.Description = "Delivery server management"
	domainDSGet = flaggy.NewSubcommand("show")
	domainDSGet.Description = "Show delivery server details"
	domainDSCreate = flaggy.NewSubcommand("create")
	domainDSCreate.Description = "Create delivery server"
	domainDSUpdate = flaggy.NewSubcommand("update")
	domainDSUpdate.Description = "Update delivery server"
	domainDSDelete = flaggy.NewSubcommand("delete")
	domainDSDelete.Description = "Delete delivery server"

	userDSs = flaggy.NewSubcommand("userdeliverservers")
	userDSs.Description = "List user delivery servers"

	userDS = flaggy.NewSubcommand("userdeliverserver")
	userDS.Description = "User delivery server management"
	userDSGet = flaggy.NewSubcommand("show")
	userDSGet.Description = "Show user delivery server details"
	userDSCreate = flaggy.NewSubcommand("create")
	userDSCreate.Description = "Create user delivery server"
	userDSUpdate = flaggy.NewSubcommand("update")
	userDSUpdate.Description = "Update user delivery server"
	userDSDelete = flaggy.NewSubcommand("delete")
	userDSDelete.Description = "Delete user delivery server"

	domains = flaggy.NewSubcommand("domains")
	domains.Description = "List domains"

	domainASs = flaggy.NewSubcommand("authsettings")
	domainASs.Description = "List authentication settings"

	domainAS = flaggy.NewSubcommand("authsetting")
	domainAS.Description = "Authentication settings management"
	domainASGet = flaggy.NewSubcommand("show")
	domainASGet.Description = "Show authentication settings details"
	domainASCreate = flaggy.NewSubcommand("create")
	domainASCreate.Description = "Create authentication settings"
	domainASUpdate = flaggy.NewSubcommand("update")
	domainASUpdate.Description = "Update authentication settings"
	domainASDelete = flaggy.NewSubcommand("delete")
	domainASDelete.Description = "Delete authentication settings"

	domainSMs = flaggy.NewSubcommand("smarthosts")
	domainSMs.Description = "List smarthosts"

	domainSM = flaggy.NewSubcommand("smarthost")
	domainSM.Description = "Smarthost management"
	domainSMGet = flaggy.NewSubcommand("show")
	domainSMGet.Description = "Show smarthost details"
	domainSMCreate = flaggy.NewSubcommand("create")
	domainSMCreate.Description = "Create smarthost"
	domainSMUpdate = flaggy.NewSubcommand("update")
	domainSMUpdate.Description = "Update smarthost"
	domainSMDelete = flaggy.NewSubcommand("delete")
	domainSMDelete.Description = "Delete smarthost"

	ldap = flaggy.NewSubcommand("ldap")
	ldap.Description = "AD/LDAP settings management"
	ldapGet = flaggy.NewSubcommand("show")
	ldapGet.Description = "Show AD/LDAP settings"
	ldapCreate = flaggy.NewSubcommand("create")
	ldapCreate.Description = "Create AD/LDAP settings"
	ldapUpdate = flaggy.NewSubcommand("update")
	ldapUpdate.Description = "Update AD/LDAP settings"
	ldapDelete = flaggy.NewSubcommand("delete")
	ldapDelete.Description = "Delete AD/LDAP settings"

	radius = flaggy.NewSubcommand("radius")
	radius.Description = "RADIUS settings management"
	radiusGet = flaggy.NewSubcommand("show")
	radiusGet.Description = "Show RADIUS settings"
	radiusCreate = flaggy.NewSubcommand("create")
	radiusCreate.Description = "Create RADIUS settings"
	radiusUpdate = flaggy.NewSubcommand("update")
	radiusUpdate.Description = "Update RADIUS settings"
	radiusDelete = flaggy.NewSubcommand("delete")
	radiusDelete.Description = "Delete RADIUS settings"

	organization = flaggy.NewSubcommand("organization")
	organization.Description = "Organization management"
	organizationGet = flaggy.NewSubcommand("show")
	organizationGet.Description = "Show organization details"
	organizationCreate = flaggy.NewSubcommand("create")
	organizationCreate.Description = "Create organization"
	organizationUpdate = flaggy.NewSubcommand("update")
	organizationUpdate.Description = "Update organization"
	organizationDelete = flaggy.NewSubcommand("delete")
	organizationDelete.Description = "Delete organization"

	relay = flaggy.NewSubcommand("relay")
	relay.Description = "Relay management"
	relayGet = flaggy.NewSubcommand("show")
	relayGet.Description = "Show relay details"
	relayCreate = flaggy.NewSubcommand("create")
	relayCreate.Description = "Create relay"
	relayUpdate = flaggy.NewSubcommand("update")
	relayUpdate.Description = "Update relay"
	relayDelete = flaggy.NewSubcommand("delete")
	relayDelete.Description = "Delete relay"

	fallbackServers = flaggy.NewSubcommand("fallbackservers")
	fallbackServers.Description = "List Fallback Servers"

	fallbackServer = flaggy.NewSubcommand("fallbackserver")
	fallbackServer.Description = "Fallback Server management"
	fallbackServerGet = flaggy.NewSubcommand("show")
	fallbackServerGet.Description = "Show fallback server details"
	fallbackServerCreate = flaggy.NewSubcommand("create")
	fallbackServerCreate.Description = "Create fallback server"
	fallbackServerUpdate = flaggy.NewSubcommand("update")
	fallbackServerUpdate.Description = "Update fallback server"
	fallbackServerDelete = flaggy.NewSubcommand("delete")
	fallbackServerDelete.Description = "Delete fallback server"

	organizationSMs = flaggy.NewSubcommand("smarthosts")
	organizationSMs.Description = "List smarthosts"

	organizationSM = flaggy.NewSubcommand("smarthost")
	organizationSM.Description = "Smarthost management"
	organizationSMGet = flaggy.NewSubcommand("show")
	organizationSMGet.Description = "Show smarthost details"
	organizationSMCreate = flaggy.NewSubcommand("create")
	organizationSMCreate.Description = "Create smarthost"
	organizationSMUpdate = flaggy.NewSubcommand("update")
	organizationSMUpdate.Description = "Update smarthost"
	organizationSMDelete = flaggy.NewSubcommand("delete")
	organizationSMDelete.Description = "Delete smarthost"

	organizations = flaggy.NewSubcommand("organizations")
	organizations.Description = "List Organizations"

	system = flaggy.NewSubcommand("system")
	system.Description = "System Status"

	user.AttachSubcommand(userGet, 1)
	user.AttachSubcommand(userCreate, 1)
	user.AttachSubcommand(userUpdate, 1)
	user.AttachSubcommand(userDelete, 1)
	flaggy.AttachSubcommand(user, 1)
	flaggy.AttachSubcommand(users, 1)
	domain.AttachSubcommand(domainGet, 1)
	domain.AttachSubcommand(domainCreate, 1)
	domain.AttachSubcommand(domainUpdate, 1)
	domain.AttachSubcommand(domainDelete, 1)
	domain.AttachSubcommand(domainAliases, 1)
	domain.AttachSubcommand(domainAlias, 1)
	domainAlias.AttachSubcommand(domainAliasGet, 1)
	domainAlias.AttachSubcommand(domainAliasCreate, 1)
	domainAlias.AttachSubcommand(domainAliasUpdate, 1)
	domainAlias.AttachSubcommand(domainAliasDelete, 1)
	domain.AttachSubcommand(domainDSs, 1)
	domain.AttachSubcommand(domainDS, 1)
	domain.AttachSubcommand(userDSs, 1)
	domain.AttachSubcommand(userDS, 1)
	domain.AttachSubcommand(domainASs, 1)
	domain.AttachSubcommand(domainAS, 1)
	domain.AttachSubcommand(domainSMs, 1)
	domain.AttachSubcommand(domainSM, 1)
	domainDS.AttachSubcommand(domainDSGet, 1)
	domainDS.AttachSubcommand(domainDSCreate, 1)
	domainDS.AttachSubcommand(domainDSUpdate, 1)
	domainDS.AttachSubcommand(domainDSDelete, 1)
	userDS.AttachSubcommand(userDSGet, 1)
	userDS.AttachSubcommand(userDSCreate, 1)
	userDS.AttachSubcommand(userDSUpdate, 1)
	userDS.AttachSubcommand(userDSDelete, 1)
	domainAS.AttachSubcommand(domainASGet, 1)
	domainAS.AttachSubcommand(domainASCreate, 1)
	domainAS.AttachSubcommand(domainASUpdate, 1)
	domainAS.AttachSubcommand(domainASDelete, 1)
	domainAS.AttachSubcommand(ldap, 1)
	domainAS.AttachSubcommand(radius, 1)
	domainSM.AttachSubcommand(domainSMGet, 1)
	domainSM.AttachSubcommand(domainSMCreate, 1)
	domainSM.AttachSubcommand(domainSMUpdate, 1)
	domainSM.AttachSubcommand(domainSMDelete, 1)
	ldap.AttachSubcommand(ldapGet, 1)
	ldap.AttachSubcommand(ldapCreate, 1)
	ldap.AttachSubcommand(ldapUpdate, 1)
	ldap.AttachSubcommand(ldapDelete, 1)
	radius.AttachSubcommand(radiusGet, 1)
	radius.AttachSubcommand(radiusCreate, 1)
	radius.AttachSubcommand(radiusUpdate, 1)
	radius.AttachSubcommand(radiusDelete, 1)
	flaggy.AttachSubcommand(domain, 1)
	flaggy.AttachSubcommand(domains, 1)
	flaggy.AttachSubcommand(organization, 1)
	organization.AttachSubcommand(organizationGet, 1)
	organization.AttachSubcommand(organizationCreate, 1)
	organization.AttachSubcommand(organizationUpdate, 1)
	organization.AttachSubcommand(organizationDelete, 1)
	organization.AttachSubcommand(relay, 1)
	organization.AttachSubcommand(fallbackServers, 1)
	organization.AttachSubcommand(fallbackServer, 1)
	organization.AttachSubcommand(organizationSMs, 1)
	organization.AttachSubcommand(organizationSM, 1)
	relay.AttachSubcommand(relayGet, 1)
	relay.AttachSubcommand(relayCreate, 1)
	relay.AttachSubcommand(relayUpdate, 1)
	relay.AttachSubcommand(relayDelete, 1)
	fallbackServer.AttachSubcommand(fallbackServerGet, 1)
	fallbackServer.AttachSubcommand(fallbackServerCreate, 1)
	fallbackServer.AttachSubcommand(fallbackServerUpdate, 1)
	fallbackServer.AttachSubcommand(fallbackServerDelete, 1)
	organizationSM.AttachSubcommand(organizationSMGet, 1)
	organizationSM.AttachSubcommand(organizationSMCreate, 1)
	organizationSM.AttachSubcommand(organizationSMUpdate, 1)
	organizationSM.AttachSubcommand(organizationSMDelete, 1)
	flaggy.AttachSubcommand(organizations, 1)
	flaggy.AttachSubcommand(system, 1)
	flaggy.Parse()
}

func main() {
}
