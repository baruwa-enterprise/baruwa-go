// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"log"

	"github.com/baruwa-enterprise/baruwa-go/api"
	prettyjson "github.com/hokaccha/go-prettyjson"
	cli "github.com/jawher/mow.cli"
)

func ldapShow(cmd *cli.Cmd) {
	var (
		c             *api.Client
		s             *api.LDAPSettings
		err           error
		b             []byte
		did, sid, lid *int
	)

	cmd.Spec = "--domain-id --settings-id --ldap-settings-id"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	lid = cmd.Int(cli.IntOpt{
		Name: "ldap-settings-id",
		Desc: "LDAP setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetLDAPSettings(*did, *sid, *lid); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func ldapCreate(cmd *cli.Cmd) {
	var (
		c                                                              *api.Client
		s                                                              *api.LDAPSettings
		err                                                            error
		b                                                              []byte
		did, sid                                                       *int
		useTLS, useSearch                                              *bool
		basedn, nameattribute, emailattribute, binddn, bindpw          *string
		searchfilter, searchScope, emailSearchFilter, emailSearchScope *string
	)

	cmd.Spec = "--domain-id --settings-id --base-dn --name-attribute --email-attribute [--bind-dn][--bind-password][--search-filter][--search-scope][--email-search-filter][--email-search-scope][--use-tls][--use-search]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	basedn = cmd.String(cli.StringOpt{
		Name: "base-dn",
		Desc: "Base DN",
	})
	nameattribute = cmd.String(cli.StringOpt{
		Name:  "name-attribute",
		Desc:  "Name Attribute",
		Value: "uid",
	})
	emailattribute = cmd.String(cli.StringOpt{
		Name:  "email-attribute",
		Desc:  "Email Attribute",
		Value: "mail",
	})
	binddn = cmd.String(cli.StringOpt{
		Name: "bind-dn",
		Desc: "Bind DN",
	})
	bindpw = cmd.String(cli.StringOpt{
		Name: "bind-password",
		Desc: "Bind Password",
	})
	searchfilter = cmd.String(cli.StringOpt{
		Name: "search-filter",
		Desc: "Search Filter",
	})
	searchScope = cmd.String(cli.StringOpt{
		Name:  "search-scope",
		Desc:  "Search Scope",
		Value: "subtree",
	})
	emailSearchFilter = cmd.String(cli.StringOpt{
		Name: "email-search-filter",
		Desc: "Email Search Filter",
	})
	emailSearchScope = cmd.String(cli.StringOpt{
		Name:  "email-search-scope",
		Desc:  "Email Search Scope",
		Value: "subtree",
	})
	useTLS = cmd.Bool(cli.BoolOpt{
		Name: "use-tls",
		Desc: "Use TLS",
	})
	useSearch = cmd.Bool(cli.BoolOpt{
		Name: "use-search",
		Desc: "Use Search",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		s = &api.LDAPSettings{
			Basedn:            *basedn,
			NameAttribute:     *nameattribute,
			EmailAttribute:    *emailattribute,
			BindDN:            *binddn,
			BindPw:            *bindpw,
			SearchFilter:      *searchfilter,
			SearchScope:       *searchScope,
			EmailSearchFilter: *emailSearchFilter,
			EmailSearchScope:  *emailSearchScope,
			UseTLS:            *useTLS,
			UseSearch:         *useSearch,
		}

		if err = c.CreateLDAPSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func ldapUpdate(cmd *cli.Cmd) {
	var (
		c                                                              *api.Client
		s                                                              *api.LDAPSettings
		err                                                            error
		did, sid, lid                                                  *int
		useTLS, useSearch                                              *bool
		basedn, nameattribute, emailattribute, binddn, bindpw          *string
		searchfilter, searchScope, emailSearchFilter, emailSearchScope *string
		useTLSSet, useSearchSet, basednSet, nameattributeSet           bool
		emailattributeSet, binddnSet, bindpwSet, searchfilterSet       bool
		searchScopeSet, emailSearchFilterSet, emailSearchScopeSet      bool
	)

	cmd.Spec = "--domain-id --settings-id --ldap-settings-id --disable-use-tls|--use-tls --disable-use-search|--use-search [--base-dn][--name-attribute][--email-attribute][--bind-dn][--bind-password][--search-filter][--search-scope][--email-search-filter][--email-search-scope]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	lid = cmd.Int(cli.IntOpt{
		Name: "ldap-settings-id",
		Desc: "LDAP setting ID",
	})
	basedn = cmd.String(cli.StringOpt{
		Name:      "base-dn",
		Desc:      "Base DN",
		SetByUser: &basednSet,
	})
	nameattribute = cmd.String(cli.StringOpt{
		Name:      "name-attribute",
		Desc:      "Name Attribute",
		SetByUser: &nameattributeSet,
	})
	emailattribute = cmd.String(cli.StringOpt{
		Name:      "email-attribute",
		Desc:      "Email Attribute",
		SetByUser: &emailattributeSet,
	})
	binddn = cmd.String(cli.StringOpt{
		Name:      "bind-dn",
		Desc:      "Bind DN",
		SetByUser: &binddnSet,
	})
	bindpw = cmd.String(cli.StringOpt{
		Name:      "bind-password",
		Desc:      "Bind Password",
		SetByUser: &bindpwSet,
	})
	searchfilter = cmd.String(cli.StringOpt{
		Name:      "search-filter",
		Desc:      "Search Filter",
		SetByUser: &searchfilterSet,
	})
	searchScope = cmd.String(cli.StringOpt{
		Name:      "search-scope",
		Desc:      "Search Scope",
		SetByUser: &searchScopeSet,
	})
	emailSearchFilter = cmd.String(cli.StringOpt{
		Name:      "email-search-filter",
		Desc:      "Email Search Filter",
		SetByUser: &emailSearchFilterSet,
	})
	emailSearchScope = cmd.String(cli.StringOpt{
		Name:      "email-search-scope",
		Desc:      "Email Search Scope",
		SetByUser: &emailSearchScopeSet,
	})
	useTLS = cmd.Bool(cli.BoolOpt{
		Name: "disable-use-tls",
		Desc: "Disable use TLS",
	})
	useTLS = cmd.Bool(cli.BoolOpt{
		Name:      "use-tls",
		Desc:      "Use TLS",
		SetByUser: &useTLSSet,
	})
	useSearch = cmd.Bool(cli.BoolOpt{
		Name: "disable-use-search",
		Desc: "Disable use Search",
	})
	useSearch = cmd.Bool(cli.BoolOpt{
		Name:      "use-search",
		Desc:      "Use Search",
		SetByUser: &useSearchSet,
	})
	*useTLS = useTLSSet
	*useSearch = useSearchSet

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetLDAPSettings(*did, *sid, *lid); err != nil {
			log.Fatal(err)
		}

		if basednSet {
			s.Basedn = *basedn
		}
		if nameattributeSet {
			s.NameAttribute = *nameattribute
		}
		if emailattributeSet {
			s.EmailAttribute = *emailattribute
		}
		if binddnSet {
			s.BindDN = *binddn
		}
		if bindpwSet {
			s.BindPw = *bindpw
		}
		if searchfilterSet {
			s.SearchFilter = *searchfilter
		}
		if searchScopeSet {
			s.SearchScope = *searchScope
		}
		if emailSearchFilterSet {
			s.EmailSearchFilter = *emailSearchFilter
		}
		if emailSearchScopeSet {
			s.EmailSearchScope = *emailSearchScope
		}
		s.UseTLS = *useTLS
		s.UseSearch = *useSearch

		if err = c.UpdateLDAPSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The ldap settings: %d have been updated\n", s.ID)
	}
}

func ldapDelete(cmd *cli.Cmd) {
	var (
		c             *api.Client
		s             *api.LDAPSettings
		err           error
		did, sid, lid *int
	)

	cmd.Spec = "--domain-id --settings-id --ldap-settings-id"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	lid = cmd.Int(cli.IntOpt{
		Name: "ldap-settings-id",
		Desc: "LDAP setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetLDAPSettings(*did, *sid, *lid); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteLDAPSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The ldap settings: %d have been deleted\n", s.ID)
	}
}
