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

func domainSMShow(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		b       []byte
		c       *api.Client
		s       *api.DomainSmartHost
	)

	cmd.Spec = "--id  --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetDomainSmartHost(*did, *id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainSMCreate(cmd *cli.Cmd) {
	var (
		did, port                                *int
		err                                      error
		b                                        []byte
		c                                        *api.Client
		s                                        *api.DomainSmartHost
		enabled, requireTLS                      *bool
		address, username, password, description *string
	)

	cmd.Spec = "--domain-id --address [--port] [--username] [--password] [--description] [--enable] [--require-tls]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
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

		s = &api.DomainSmartHost{
			Address:     *address,
			Username:    *username,
			Password:    *password,
			Description: *description,
			Enabled:     *enabled,
			RequireTLS:  *requireTLS,
			Port:        *port,
		}

		if err = c.CreateDomainSmartHost(*did, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainSMUpdate(cmd *cli.Cmd) {
	var (
		err                                                  error
		c                                                    *api.Client
		s                                                    *api.DomainSmartHost
		id, did, port                                        *int
		enabled, requireTLS                                  *bool
		address, username, password, description             *string
		addressSet, usernameSet, passwordSet, descriptionSet bool
		enabledSet, requireTLSSet, portSet                   bool
	)

	cmd.Spec = "--id  --domain-id --enable|--disable --require-tls|--disable-require-tls [--address][--port][--username][--password][--description]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
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
	fmt.Println("ST", enabledSet, requireTLSSet)

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetDomainSmartHost(*did, *id); err != nil {
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

		if err = c.UpdateDomainSmartHost(*did, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain smarthost: %s has been updated\n", s.Address)
	}
}

func domainSMDelete(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		c       *api.Client
		s       *api.DomainSmartHost
	)

	cmd.Spec = "--id  --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Smarthost ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetDomainSmartHost(*did, *id); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteDomainSmartHost(*did, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain smarthost: %s has been deleted\n", s.Address)
	}
}

func domainSMList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		c        *api.Client
		opts     *api.ListOptions
		d        *api.DomainSmartHostList
	)

	cmd.Spec = "--id [--page]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain ID",
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
				Page: fmt.Sprintf(api.DSMListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if d, err = c.GetDomainSmartHosts(*id, opts); err != nil {
			log.Fatal(err)
		}

		if len(d.Items) == 0 {
			fmt.Println()
			return
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
