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

func organizationRelayShow(cmd *cli.Cmd) {
	var (
		id  *int
		err error
		b   []byte
		c   *api.Client
		r   *api.RelaySetting
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Relay setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if r, err = c.GetRelaySetting(*id); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(r); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationRelayCreate(cmd *cli.Cmd) {
	var (
		b                                                    []byte
		err                                                  error
		c                                                    *api.Client
		r                                                    *api.RelaySetting
		lowScore, highScore                                  api.LocalFloat64
		id, spamActions, highSpamActions, rateLimit          *int
		enabled, requireTLS, blockMacros                     *bool
		address, username, password1, password2, description *string
	)

	cmd.Spec = "--organization-id --address|--username [--password1][--password2][--description][--enable][--require-tls][--block-macros][--spam-actions][--high-spam-actions][--rate-limit][--low-score][--high-score]"

	id = cmd.Int(cli.IntOpt{
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
	password1 = cmd.String(cli.StringOpt{
		Name: "password1",
		Desc: "Password",
	})
	password2 = cmd.String(cli.StringOpt{
		Name: "password2",
		Desc: "Retype Password",
	})
	description = cmd.String(cli.StringOpt{
		Name: "description",
		Desc: "Description",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this relay setting",
	})
	requireTLS = cmd.Bool(cli.BoolOpt{
		Name: "require-tls",
		Desc: "Require TLS",
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "block-macros",
		Desc: "Enable blocking Attachments with Macros",
	})
	spamActions = cmd.Int(cli.IntOpt{
		Name:  "spam-actions",
		Desc:  "What to do with Probable spam",
		Value: 2,
	})
	highSpamActions = cmd.Int(cli.IntOpt{
		Name:  "high-spam-actions",
		Desc:  "What to do with Definite spam",
		Value: 2,
	})
	rateLimit = cmd.Int(cli.IntOpt{
		Name:  "rate-limit",
		Desc:  "Number of messages per 15 minutes",
		Value: 250,
	})
	cmd.Var(cli.VarOpt{
		Name:  "low-score",
		Desc:  "The score at which an email is considered to be suspected spam, 0.0 uses system defaults",
		Value: &lowScore,
	})
	cmd.Var(cli.VarOpt{
		Name:  "high-score",
		Desc:  "The score at which an email is considered to be definitely spam, 0.0 uses system defaults",
		Value: &highScore,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		r = &api.RelaySetting{
			Address:         *address,
			Username:        *username,
			Password1:       *password1,
			Password2:       *password2,
			Description:     *description,
			Enabled:         *enabled,
			RequireTLS:      *requireTLS,
			BlockMacros:     *blockMacros,
			SpamActions:     *spamActions,
			HighSpamActions: *highSpamActions,
			RateLimit:       *rateLimit,
			LowScore:        lowScore,
			HighScore:       highScore,
		}

		if err = c.CreateRelaySetting(*id, r); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(r); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func organizationRelayUpdate(cmd *cli.Cmd) {
	var (
		err                                                                         error
		c                                                                           *api.Client
		r                                                                           *api.RelaySetting
		lowScore, highScore                                                         api.LocalFloat64
		id, spamActions, highSpamActions, rateLimit                                 *int
		enabled, requireTLS, blockMacros                                            *bool
		address, username, password1, password2, description                        *string
		addressSet, usernameSet, password1Set, password2Set, descriptionSet         bool
		enabledSet, requireTLSSet, blockMacrosSet                                   bool
		spamActionsSet, highSpamActionsSet, rateLimitSet, lowScoreSet, highScoreSet bool
	)

	cmd.Spec = "--id --disable|--enable --disable-require-tls|--require-tls --disable-block-macros|--block-macros [--address][--username][--password1][--password2][--description][--spam-actions][--high-spam-actions][--rate-limit][--low-score][--high-score]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Relay setting ID",
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
	password1 = cmd.String(cli.StringOpt{
		Name:      "password1",
		Desc:      "Password",
		SetByUser: &password1Set,
	})
	password2 = cmd.String(cli.StringOpt{
		Name:      "password2",
		Desc:      "Retype Password",
		SetByUser: &password2Set,
	})
	description = cmd.String(cli.StringOpt{
		Name:      "description",
		Desc:      "Description",
		SetByUser: &descriptionSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this relay setting",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this relay setting",
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
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "disable-block-macros",
		Desc: "Disable blocking Attachments with Macros",
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name:      "block-macros",
		Desc:      "Enable blocking Attachments with Macros",
		SetByUser: &blockMacrosSet,
	})
	spamActions = cmd.Int(cli.IntOpt{
		Name:      "spam-actions",
		Desc:      "What to do with Probable spam",
		SetByUser: &spamActionsSet,
	})
	highSpamActions = cmd.Int(cli.IntOpt{
		Name:      "high-spam-actions",
		Desc:      "What to do with Definite spam",
		SetByUser: &highSpamActionsSet,
	})
	rateLimit = cmd.Int(cli.IntOpt{
		Name:      "rate-limit",
		Desc:      "Number of messages per 15 minutes",
		SetByUser: &rateLimitSet,
	})
	cmd.Var(cli.VarOpt{
		Name:      "low-score",
		Desc:      "The score at which an email is considered to be suspected spam, 0.0 uses system defaults",
		Value:     &lowScore,
		SetByUser: &lowScoreSet,
	})
	cmd.Var(cli.VarOpt{
		Name:      "high-score",
		Desc:      "The score at which an email is considered to be definitely spam, 0.0 uses system defaults",
		Value:     &highScore,
		SetByUser: &highScoreSet,
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if r, err = c.GetRelaySetting(*id); err != nil {
			log.Fatal(err)
		}

		*enabled = enabledSet
		*requireTLS = requireTLSSet
		*blockMacros = blockMacrosSet

		r.Enabled = *enabled
		r.RequireTLS = *requireTLS
		r.BlockMacros = *blockMacros
		if addressSet {
			r.Address = *address
		}
		if usernameSet {
			r.Username = *username
		}
		if password1Set {
			r.Password1 = *password1
		}
		if password2Set {
			r.Password2 = *password2
		}
		if descriptionSet {
			r.Description = *description
		}
		if spamActionsSet {
			r.SpamActions = *spamActions
		}
		if highSpamActionsSet {
			r.HighSpamActions = *highSpamActions
		}
		if rateLimitSet {
			r.RateLimit = *rateLimit
		}
		if lowScoreSet {
			r.LowScore = lowScore
		}
		if highScoreSet {
			r.HighScore = highScore
		}

		if err = c.UpdateRelaySetting(r); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The relay setting: %s[%s] has been updated\n", r.Address, r.Username)
	}
}

func organizationRelayDelete(cmd *cli.Cmd) {
	var (
		id  *int
		err error
		c   *api.Client
		r   *api.RelaySetting
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Relay setting ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if r, err = c.GetRelaySetting(*id); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteRelaySetting(r); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The relay setting: %s[%s] has been deleted\n", r.Address, r.Username)
	}
}
