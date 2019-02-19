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

func domainDSShow(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		b       []byte
		c       *api.Client
		s       *api.DomainDeliveryServer
	)

	cmd.Spec = "--id  --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Delivery Server ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetDomainDeliveryServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainDSCreate(cmd *cli.Cmd) {
	var (
		b                                     []byte
		err                                   error
		c                                     *api.Client
		s                                     *api.DomainDeliveryServer
		f                                     *api.DomainDeliveryServerForm
		id, protocol, port                    *int
		address                               *string
		enabled, requireTLS, verificationOnly *bool
	)

	cmd.Spec = "--domain-id --address [--protocol][--port][--enable][--require-tls][--verification-only]"

	id = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	address = cmd.String(cli.StringOpt{
		Name: "address",
		Desc: "Address",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this delivery server",
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

		f = &api.DomainDeliveryServerForm{
			Address:          *address,
			Enabled:          *enabled,
			RequireTLS:       *requireTLS,
			VerificationOnly: *verificationOnly,
			Protocol:         *protocol,
			Port:             *port,
			Domain:           *id,
		}

		if s, err = c.CreateDomainDeliveryServer(*id, f); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainDSUpdate(cmd *cli.Cmd) {
	var (
		err                                                                              error
		id, did, protocol, port                                                          *int
		c                                                                                *api.Client
		s                                                                                *api.DomainDeliveryServer
		f                                                                                *api.DomainDeliveryServerForm
		address                                                                          *string
		enabled, requireTLS, verificationOnly                                            *bool
		addressSet, enabledSet, requireTLSSet, verificationOnlySet, protocolSet, portSet bool
	)

	cmd.Spec = "--id --domain-id --enable|--disable --require-tls|--disable-require-tls --verification-only|--disable-verification-only [--address][--protocol][--port]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Delivery server ID",
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
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this delivery server",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this delivery server",
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

		if s, err = c.GetDomainDeliveryServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		*enabled = enabledSet
		*requireTLS = requireTLSSet
		*verificationOnly = verificationOnlySet

		f = &api.DomainDeliveryServerForm{
			ID:               s.ID,
			Domain:           *did,
			Address:          s.Address,
			Enabled:          *enabled,
			RequireTLS:       *requireTLS,
			VerificationOnly: *verificationOnly,
			Protocol:         s.Protocol,
			Port:             s.Port,
		}

		if protocolSet {
			f.Protocol = *protocol
		}
		if portSet {
			f.Port = *port
		}
		if addressSet {
			f.Address = *address
		}

		if err = c.UpdateDomainDeliveryServer(*did, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain delivery server: %s has been updated\n", s.Address)
	}
}

func domainDSDelete(cmd *cli.Cmd) {
	var (
		err     error
		id, did *int
		c       *api.Client
		s       *api.DomainDeliveryServer
		f       *api.DomainDeliveryServerForm
	)

	cmd.Spec = "--id --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Delivery server ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetDomainDeliveryServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		f = &api.DomainDeliveryServerForm{
			ID:               s.ID,
			Domain:           *did,
			Address:          s.Address,
			Enabled:          s.Enabled,
			RequireTLS:       s.RequireTLS,
			VerificationOnly: s.VerificationOnly,
			Protocol:         s.Protocol,
			Port:             s.Port,
		}

		if err = c.DeleteDomainDeliveryServer(*did, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain delivery server: %s has been deleted\n", s.Address)
	}
}

func domainDSList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		c        *api.Client
		opts     *api.ListOptions
		d        *api.DomainDeliveryServerList
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
				Page: fmt.Sprintf(api.DDSListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if d, err = c.GetDomainDeliveryServers(*id, opts); err != nil {
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
