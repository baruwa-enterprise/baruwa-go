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

func domainASShow(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		b       []byte
		c       *api.Client
		s       *api.AuthServer
	)

	cmd.Spec = "--id  --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Authentication setting ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetAuthServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainASCreate(cmd *cli.Cmd) {
	var (
		err                   error
		b                     []byte
		c                     *api.Client
		s                     *api.AuthServer
		did, protocol, port   *int
		address, usm          *string
		enabled, splitAddress *bool
	)

	cmd.Spec = "--domain-id --address --port --protocol [--user-map-template][--enable][--split-address]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	address = cmd.String(cli.StringOpt{
		Name: "address",
		Desc: "Address",
	})
	usm = cmd.String(cli.StringOpt{
		Name: "user-map-template",
		Desc: "User Map Template",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this authentication setting",
	})
	splitAddress = cmd.Bool(cli.BoolOpt{
		Name: "split-address",
		Desc: "Split the address",
	})
	port = cmd.Int(cli.IntOpt{
		Name: "port",
		Desc: "Port",
	})
	protocol = cmd.Int(cli.IntOpt{
		Name: "protocol",
		Desc: "Protocol",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		s = &api.AuthServer{
			Address:         *address,
			Port:            *port,
			UserMapTemplate: *usm,
			Protocol:        *protocol,
			Enabled:         *enabled,
			SplitAddress:    *splitAddress,
		}

		if err = c.CreateAuthServer(*did, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainASUpdate(cmd *cli.Cmd) {
	var (
		err                                                                   error
		c                                                                     *api.Client
		s                                                                     *api.AuthServer
		id, did, protocol, port                                               *int
		address, usm                                                          *string
		enabled, splitAddress                                                 *bool
		addressSet, usmSet, enabledSet, splitAddressSet, protocolSet, portSet bool
	)

	cmd.Spec = "--id --domain-id --enable|--disable --split-address|--disable-split-address [--address][--port][--protocol][--user-map-template]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Authentication setting ID",
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
	usm = cmd.String(cli.StringOpt{
		Name:      "user-map-template",
		Desc:      "User Map Template",
		SetByUser: &usmSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this authentication setting",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this authentication setting",
		SetByUser: &enabledSet,
	})
	splitAddress = cmd.Bool(cli.BoolOpt{
		Name: "disable-split-address",
		Desc: "Disable split the address",
	})
	splitAddress = cmd.Bool(cli.BoolOpt{
		Name:      "split-address",
		Desc:      "Split the address",
		SetByUser: &splitAddressSet,
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

		if s, err = c.GetAuthServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		*enabled = enabledSet
		*splitAddress = splitAddressSet

		s.Enabled = *enabled
		s.SplitAddress = *splitAddress

		if addressSet {
			s.Address = *address
		}
		if usmSet {
			s.UserMapTemplate = *usm
		}
		if protocolSet {
			s.Protocol = *protocol
		}
		if portSet {
			s.Port = *port
		}

		if err = c.UpdateAuthServer(*did, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The authentication setting: %s has been updated\n", s.Address)
	}
}

func domainASDelete(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		c       *api.Client
		s       *api.AuthServer
	)

	cmd.Spec = "--id --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Authentication setting ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetAuthServer(*did, *id); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteAuthServer(*did, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The authentication setting: %s has been deleted\n", s.Address)
	}
}

func domainASList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		c        *api.Client
		opts     *api.ListOptions
		d        *api.AuthServerList
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
				Page: fmt.Sprintf(api.DASListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if d, err = c.GetAuthServers(*id, opts); err != nil {
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
