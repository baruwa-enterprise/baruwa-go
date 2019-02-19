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

func organizationFSShow(cmd *cli.Cmd) {
	var (
		id  *int
		err error
		b   []byte
		c   *api.Client
		s   *api.FallBackServer
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Fallback Server ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetFallBackServer(*id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationFSCreate(cmd *cli.Cmd) {
	var (
		b                                     []byte
		err                                   error
		c                                     *api.Client
		s                                     *api.FallBackServer
		oid, protocol, port                   *int
		address                               *string
		enabled, requireTLS, verificationOnly *bool
	)

	cmd.Spec = "--organization-id --address [--protocol][--port][--enable][--require-tls][--verification-only]"

	oid = cmd.Int(cli.IntOpt{
		Name: "organization-id",
		Desc: "Organization ID",
	})
	address = cmd.String(cli.StringOpt{
		Name: "address",
		Desc: "Address",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this fallback server",
	})
	requireTLS = cmd.Bool(cli.BoolOpt{
		Name: "require-tls",
		Desc: "Require TLS",
	})
	verificationOnly = cmd.Bool(cli.BoolOpt{
		Name: "verification-only",
		Desc: "Verification only",
	})
	protocol = cmd.Int(cli.IntOpt{
		Name:  "protocol",
		Desc:  "Protocol",
		Value: 1,
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

		s = &api.FallBackServer{
			Address:          *address,
			Enabled:          *enabled,
			RequireTLS:       *requireTLS,
			VerificationOnly: *verificationOnly,
			Protocol:         *protocol,
			Port:             *port,
		}

		if err = c.CreateFallBackServer(*oid, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationFSUpdate(cmd *cli.Cmd) {
	var (
		err                                                                              error
		id, protocol, port                                                               *int
		c                                                                                *api.Client
		s                                                                                *api.FallBackServer
		address                                                                          *string
		enabled, requireTLS, verificationOnly                                            *bool
		addressSet, enabledSet, requireTLSSet, verificationOnlySet, protocolSet, portSet bool
	)

	cmd.Spec = "--id --enable|--disable --require-tls|--disable-require-tls --verification-only|--disable-verification-only [--address][--protocol][--port]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Fallback Server ID",
	})
	address = cmd.String(cli.StringOpt{
		Name:      "address",
		Desc:      "Address",
		SetByUser: &addressSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this fallback server",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this fallback server",
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
	verificationOnly = cmd.Bool(cli.BoolOpt{
		Name: "disable-verification-only",
		Desc: "Disable Verification only",
	})
	verificationOnly = cmd.Bool(cli.BoolOpt{
		Name:      "verification-only",
		Desc:      "Verification only",
		SetByUser: &verificationOnlySet,
	})
	protocol = cmd.Int(cli.IntOpt{
		Name:      "protocol",
		Desc:      "Protocol",
		SetByUser: &protocolSet,
	})
	port = cmd.Int(cli.IntOpt{
		Name:      "port",
		Desc:      "Port",
		SetByUser: &portSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetFallBackServer(*id); err != nil {
			log.Fatal(err)
		}

		*enabled = enabledSet
		*requireTLS = requireTLSSet
		*verificationOnly = verificationOnlySet

		s.Enabled = *enabled
		s.RequireTLS = *requireTLS
		s.VerificationOnly = *verificationOnly

		if protocolSet {
			s.Protocol = *protocol
		}
		if portSet {
			s.Port = *port
		}
		if addressSet {
			s.Address = *address
		}

		if err = c.UpdateFallBackServer(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The fallback server: %s has been updated\n", s.Address)
	}
}

func organizationFSDelete(cmd *cli.Cmd) {
	var (
		err error
		id  *int
		c   *api.Client
		s   *api.FallBackServer
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Fallback Server ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetFallBackServer(*id); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteFallBackServer(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The fallback server: %s has been deleted\n", s.Address)
	}
}

func organizationFSList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		c        *api.Client
		opts     *api.ListOptions
		o        *api.FallBackServerList
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
				Page: fmt.Sprintf(api.OrgFSListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if o, err = c.GetFallBackServers(*id, opts); err != nil {
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
