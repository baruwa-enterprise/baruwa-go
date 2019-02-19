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

func radiusShow(cmd *cli.Cmd) {
	var (
		c             *api.Client
		s             *api.RadiusSettings
		err           error
		b             []byte
		did, sid, rid *int
	)

	cmd.Spec = "--domain-id --settings-id --radius-settings-id"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	rid = cmd.Int(cli.IntOpt{
		Name: "radius-settings-id",
		Desc: "RADIUS setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetRadiusSettings(*did, *sid, *rid); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func radiusCreate(cmd *cli.Cmd) {
	var (
		c                 *api.Client
		s                 *api.RadiusSettings
		err               error
		b                 []byte
		did, sid, timeout *int
		secret            *string
	)

	cmd.Spec = "--domain-id --settings-id --secret [--timeout]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	secret = cmd.String(cli.StringOpt{
		Name: "secret",
		Desc: "Radius Secret",
	})
	timeout = cmd.Int(cli.IntOpt{
		Name: "timeout",
		Desc: "Radius Timeout",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		s = &api.RadiusSettings{
			Secret:  *secret,
			Timeout: *timeout,
		}

		if err = c.CreateRadiusSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func radiusUpdate(cmd *cli.Cmd) {
	var (
		c                      *api.Client
		s                      *api.RadiusSettings
		err                    error
		did, sid, rid, timeout *int
		secret                 *string
		secretSet, timeoutSet  bool
	)

	cmd.Spec = "--domain-id --settings-id --radius-settings-id [--secret][--timeout]"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	rid = cmd.Int(cli.IntOpt{
		Name: "radius-settings-id",
		Desc: "RADIUS setting ID",
	})
	secret = cmd.String(cli.StringOpt{
		Name:      "secret",
		Desc:      "Radius Secret",
		SetByUser: &secretSet,
	})
	timeout = cmd.Int(cli.IntOpt{
		Name:      "timeout",
		Desc:      "Radius Timeout",
		SetByUser: &timeoutSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetRadiusSettings(*did, *sid, *rid); err != nil {
			log.Fatal(err)
		}

		if secretSet {
			s.Secret = *secret
		}
		if timeoutSet {
			s.Timeout = *timeout
		}

		if err = c.UpdateRadiusSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The radius settings: %d have been updated\n", s.ID)
	}
}

func radiusDelete(cmd *cli.Cmd) {
	var (
		c             *api.Client
		s             *api.RadiusSettings
		err           error
		did, sid, rid *int
	)

	cmd.Spec = "--domain-id --settings-id --radius-settings-id"

	did = cmd.Int(cli.IntOpt{
		Name: "domain-id",
		Desc: "Domain ID",
	})
	sid = cmd.Int(cli.IntOpt{
		Name: "settings-id",
		Desc: "Authentication setting ID",
	})
	rid = cmd.Int(cli.IntOpt{
		Name: "radius-settings-id",
		Desc: "RADIUS setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if s, err = c.GetRadiusSettings(*did, *sid, *rid); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteRadiusSettings(*did, *sid, s); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The radius settings: %d have been deleted\n", s.ID)
	}
}
