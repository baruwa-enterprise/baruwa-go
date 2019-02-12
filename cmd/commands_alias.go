// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cmd cmdline client for the Baruwa REST API
package cmd

import (
	"fmt"
	"log"

	"github.com/baruwa-enterprise/baruwa-go/api"
	prettyjson "github.com/hokaccha/go-prettyjson"
	cli "github.com/jawher/mow.cli"
)

func aliasShow(cmd *cli.Cmd) {
	var b []byte
	var id *int
	var err error
	var a *api.AliasAddress
	var c *api.Client

	cmd.Spec = "-id"
	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Alias address id",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if a, err = c.GetAliasAddress(*id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(a); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func aliasCreate(cmd *cli.Cmd) {
	var uid *int
	var b []byte
	var err error
	var email *string
	var enabled *bool
	var c *api.Client
	var a *api.AliasAddress

	cmd.Spec = "-u --alias-address --enabled"
	uid = cmd.Int(cli.IntOpt{
		Name: "u uid",
		Desc: "User ID",
	})
	email = cmd.String(cli.StringOpt{
		Name: "alias-address",
		Desc: "Alias Address",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enabled",
		Desc: "Enable or disable this alias",
	})

	cmd.Action = func() {
		a = &api.AliasAddress{
			Address: *email,
			Enabled: *enabled,
		}

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if err = c.CreateAliasAddress(*uid, a); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(a); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func aliasUpdate(cmd *cli.Cmd) {
	var b []byte
	var err error
	var aid *int
	var email *string
	var enabled *bool
	var c *api.Client
	var a *api.AliasAddress
	var enableSet, disableSet, emailSet bool

	aid = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Alias address ID",
	})
	email = cmd.String(cli.StringOpt{
		Name:      "alias-address",
		Desc:      "Alias Address",
		SetByUser: &emailSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this alias",
		SetByUser: &enableSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "disable",
		Desc:      "Disable this alias",
		SetByUser: &disableSet,
	})

	cmd.Spec = "--id [--alias-address] --enable|--disable"

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if a, err = c.GetAliasAddress(*aid); err != nil {
			log.Fatal(err)
		}

		if emailSet {
			a.Address = *email
		}
		if enableSet {
			*enabled = true
		}
		if disableSet {
			*enabled = false
		}
		a.Enabled = *enabled

		if err = c.UpdateAliasAddress(a); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(a); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func aliasDelete(cmd *cli.Cmd) {
	var err error
	var aid *int
	var c *api.Client
	var a *api.AliasAddress

	aid = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Alias address ID",
	})

	cmd.Spec = "--id"

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if a, err = c.GetAliasAddress(*aid); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteAliasAddress(a); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The alias address: %s has been deleted\n", a.Address)
	}
}
