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

func organizationSMShow(cmd *cli.Cmd) {
	var (
		id, oid *int
		err     error
		b       []byte
		c       *api.Client
		s       *api.OrgSmartHost
	)

	cmd.Spec = "--id  --organization-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	oid = cmd.Int(cli.IntOpt{
		Name: "organization-id",
		Desc: "Organization ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetOrgSmartHost(*oid, *id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationSMCreate(cmd *cli.Cmd) {
	var (
		err                                      error
		b                                        []byte
		c                                        *api.Client
		s                                        *api.OrgSmartHost
		oid, port                                *int
		enabled, requireTLS                      *bool
		address, username, password, description *string
	)

	cmd.Spec = "--organization-id --address [--port] [--username] [--password] [--description] [--enable] [--require-tls]"

	oid = cmd.Int(cli.IntOpt{
		Name: "organization-id",
		Desc: "Organization ID",
	})
	address = cmd.String(cli.StringOpt{
		Name: "address",
		Desc: "Address",
	})
	username = cmd.String(cli.StringOpt{
		Name: "username",
		Desc: "Username",
	})
	password = cmd.String(cli.StringOpt{
		Name: "password",
		Desc: "Password",
	})
	description = cmd.String(cli.StringOpt{
		Name: "description",
		Desc: "Description",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this smarthost",
	})
	requireTLS = cmd.Bool(cli.BoolOpt{
		Name: "require-tls",
		Desc: "Require TLS",
	})
	port = cmd.Int(cli.IntOpt{
		Name:  "port",
		Desc:  "Port",
		Value: 25,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		s = &api.OrgSmartHost{
			Address:     *address,
			Username:    *username,
			Password:    *password,
			Description: *description,
			Enabled:     *enabled,
			RequireTLS:  *requireTLS,
			Port:        *port,
		}

		if err = c.CreateOrgSmartHost(*oid, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationSMUpdate(cmd *cli.Cmd) {
	var (
		err                                                  error
		c                                                    *api.Client
		s                                                    *api.OrgSmartHost
		id, oid, port                                        *int
		enabled, requireTLS                                  *bool
		address, username, password, description             *string
		addressSet, usernameSet, passwordSet, descriptionSet bool
		enabledSet, requireTLSSet, portSet                   bool
	)

	cmd.Spec = "--id  --organization-id --enable|--disable --require-tls|--disable-require-tls [--address][--port][--username][--password][--description]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	oid = cmd.Int(cli.IntOpt{
		Name: "organization-id",
		Desc: "Organization ID",
	})
	address = cmd.String(cli.StringOpt{
		Name:      "address",
		Desc:      "Address",
		SetByUser: &addressSet,
	})
	username = cmd.String(cli.StringOpt{
		Name:      "username",
		Desc:      "Username",
		SetByUser: &usernameSet,
	})
	password = cmd.String(cli.StringOpt{
		Name:      "password",
		Desc:      "Password",
		SetByUser: &passwordSet,
	})
	description = cmd.String(cli.StringOpt{
		Name:      "description",
		Desc:      "Description",
		SetByUser: &descriptionSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this smarthost",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this smarthost",
		SetByUser: &enabledSet,
	})
	requireTLS = cmd.Bool(cli.BoolOpt{
		Name: "disable-require-tls",
		Desc: "Disable Require TLS",
	})
	requireTLS = cmd.Bool(cli.BoolOpt{
		Name:      "require-tls",
		Desc:      "Require TLS",
		SetByUser: &requireTLSSet,
	})
	port = cmd.Int(cli.IntOpt{
		Name:      "port",
		Desc:      "Port",
		SetByUser: &portSet,
	})

	*enabled = enabledSet
	*requireTLS = requireTLSSet

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetOrgSmartHost(*oid, *id); err != nil {
			log.Fatal(err)
		}

		s.Enabled = *enabled
		s.RequireTLS = *requireTLS
		if addressSet {
			s.Address = *address
		}
		if usernameSet {
			s.Username = *username
		}
		if passwordSet {
			s.Password = *password
		}
		if descriptionSet {
			s.Description = *description
		}
		if portSet {
			s.Port = *port
		}

		if err = c.UpdateOrgSmartHost(*oid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The organization smarthost: %s has been updated\n", s.Address)
	}
}

func organizationSMDelete(cmd *cli.Cmd) {
	var (
		id, oid *int
		err     error
		c       *api.Client
		s       *api.OrgSmartHost
	)

	cmd.Spec = "--id  --organization-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	oid = cmd.Int(cli.IntOpt{
		Name: "organization-id",
		Desc: "Organization ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetOrgSmartHost(*oid, *id); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteOrgSmartHost(*oid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The organization smarthost: %s has been deleted\n", s.Address)
	}
}

func organizationSMList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		c        *api.Client
		opts     *api.ListOptions
		o        *api.OrgSmartHostList
	)

	cmd.Spec = "--id [--page]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Organization ID",
	})
	page = cmd.Int(cli.IntOpt{
		Name:      "page",
		Desc:      "Page number",
		SetByUser: &pageSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if pageSet {
			opts = &api.ListOptions{
				Page: fmt.Sprintf(api.OrgSMListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if o, err = c.GetOrgSmartHosts(*id, opts); err != nil {
			log.Fatal(err)
		}

		if len(o.Items) == 0 {
			fmt.Println()
			return
		}

		if b, err = prettyjson.Marshal(o); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
