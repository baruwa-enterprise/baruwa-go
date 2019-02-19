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

func organizationShow(cmd *cli.Cmd) {
	var (
		b   []byte
		id  *int
		err error
		c   *api.Client
		o   *api.Organization
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "User ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if o, err = c.GetOrganization(*id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(o); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationCreate(cmd *cli.Cmd) {
	var (
		b               []byte
		err             error
		f               *api.OrganizationForm
		o               *api.Organization
		c               *api.Client
		name            *string
		domains, admins *[]int
	)

	cmd.Spec = "--name [--domain...][--admin...]"

	name = cmd.String(cli.StringOpt{
		Name: "name",
		Desc: "Organization name",
	})
	domains = cmd.Ints(cli.IntsOpt{
		Name: "domain",
		Desc: "The domains to add to this organization",
	})
	admins = cmd.Ints(cli.IntsOpt{
		Name: "admin",
		Desc: "The admins who manage the domains under this organization",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		f = &api.OrganizationForm{
			Name:    *name,
			Domains: *domains,
			Admins:  *admins,
		}

		if o, err = c.CreateOrganization(f); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(o); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationUpdate(cmd *cli.Cmd) {
	var (
		id                             *int
		b                              []byte
		err                            error
		td                             []int
		f                              *api.OrganizationForm
		o                              *api.Organization
		c                              *api.Client
		name                           *string
		domains, admins                *[]int
		nameSet, domainsSet, adminsSet bool
	)

	cmd.Spec = "--id [--name][--domain...][--admin...]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "User ID",
	})
	name = cmd.String(cli.StringOpt{
		Name:      "name",
		Desc:      "Organization name",
		SetByUser: &nameSet,
	})
	domains = cmd.Ints(cli.IntsOpt{
		Name:      "domain",
		Desc:      "The domains to add to this organization",
		SetByUser: &domainsSet,
	})
	admins = cmd.Ints(cli.IntsOpt{
		Name:      "admin",
		Desc:      "The admins who manage the domains under this organization",
		SetByUser: &adminsSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if o, err = c.GetOrganization(*id); err != nil {
			log.Fatal(err)
		}

		for _, dm := range o.Domains {
			td = append(td, dm.ID)
		}

		f = &api.OrganizationForm{
			ID:      *id,
			Name:    o.Name,
			Domains: td,
		}
		if nameSet {
			f.Name = *name
		}
		if domainsSet {
			f.Domains = *domains
		}
		if adminsSet {
			f.Admins = *admins
		}

		if err = c.UpdateOrganization(f, o); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(o); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationDelete(cmd *cli.Cmd) {
	var (
		id  *int
		err error
		c   *api.Client
	)

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "User ID",
	})

	cmd.Spec = "--id"

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteOrganization(*id); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The organization: %d has been deleted\n", *id)
	}
}

func organizationsList(cmd *cli.Cmd) {
	var (
		page    *int
		b       []byte
		err     error
		pageSet bool
		c       *api.Client
		opts    *api.ListOptions
		o       *api.OrganizationList
	)

	cmd.Spec = "[--page]"

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
				Page: fmt.Sprintf(api.OrgListURL, *serverURL, api.APIVersion, *page),
			}
		}

		if o, err = c.GetOrganizations(opts); err != nil {
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
