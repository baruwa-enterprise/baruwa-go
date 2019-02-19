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

func domainAliasShow(cmd *cli.Cmd) {
	var (
		id, did *int
		err     error
		b       []byte
		d       *api.DomainAlias
		c       *api.Client
	)

	cmd.Spec = "--id  --did"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain Alias ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "did",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if d, err = c.GetDomainAlias(*did, *id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func domainAliasCreate(cmd *cli.Cmd) {
	var (
		b                      []byte
		err                    error
		d                      *api.DomainAlias
		c                      *api.Client
		f                      *api.DomainAliasForm
		did                    *int
		name                   *string
		enabled, acceptInbound *bool
	)

	cmd.Spec = "--domain-id --name [--enable][--accept-inbound]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	name = cmd.String(cli.StringOpt{
		Name: "name",
		Desc: "Domain Alias name",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this domain alias",
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name: "accept-inbound",
		Desc: "Enable accepting of inbound mail to this domain alias",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		f = &api.DomainAliasForm{
			Name:          *name,
			Enabled:       *enabled,
			AcceptInbound: *acceptInbound,
			Domain:        *did,
		}

		if d, err = c.CreateDomainAlias(*did, f); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainAliasUpdate(cmd *cli.Cmd) {
	var (
		err                                   error
		id, did                               *int
		d                                     *api.DomainAlias
		c                                     *api.Client
		f                                     *api.DomainAliasForm
		name                                  *string
		enabled, acceptInbound                *bool
		nameSet, enabledSet, acceptInboundSet bool
	)
	cmd.Spec = "--id --domain-id --enable|--disable --accept-inbound|--disable-accept-inbound [--name]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain Alias ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	name = cmd.String(cli.StringOpt{
		Name:      "name",
		Desc:      "Domain Alias name",
		SetByUser: &nameSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this domain alias",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this domain alias",
		SetByUser: &enabledSet,
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name: "disable-accept-inbound",
		Desc: "Disable accepting of inbound mail to this domain alias",
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name:      "accept-inbound",
		Desc:      "Enable accepting of inbound mail to this domain alias",
		SetByUser: &acceptInboundSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if d, err = c.GetDomainAlias(*did, *id); err != nil {
			log.Fatal(err)
		}

		*enabled = enabledSet
		*acceptInbound = acceptInboundSet

		f = &api.DomainAliasForm{
			ID:            *id,
			Domain:        *did,
			Name:          d.Name,
			AcceptInbound: *acceptInbound,
			Enabled:       *enabled,
		}

		if nameSet {
			f.Name = *name
		}

		if err = c.UpdateDomainAlias(*did, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain alias: %s has been updated\n", d.Name)
	}
}

func domainAliasDelete(cmd *cli.Cmd) {
	var (
		err     error
		id, did *int
		d       *api.DomainAlias
		c       *api.Client
		f       *api.DomainAliasForm
	)

	cmd.Spec = "--id --domain-id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain Alias ID",
	})
	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if d, err = c.GetDomainAlias(*did, *id); err != nil {
			log.Fatal(err)
		}

		f = &api.DomainAliasForm{
			ID:            *id,
			Domain:        *did,
			Name:          d.Name,
			AcceptInbound: d.AcceptInbound,
			Enabled:       d.Enabled,
		}

		if err = c.DeleteDomainAlias(*did, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain alias: %s has been deleted\n", d.Name)
	}
}

func domainAliasList(cmd *cli.Cmd) {
	var (
		id, page *int
		b        []byte
		err      error
		pageSet  bool
		opts     *api.ListOptions
		c        *api.Client
		d        *api.DomainAliasList
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
				Page: fmt.Sprintf(api.DAliasListURL, *serverURL, api.APIVersion, *id, *page),
			}
		}

		if d, err = c.GetDomainAliases(*id, opts); err != nil {
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
