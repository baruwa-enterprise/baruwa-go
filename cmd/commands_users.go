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

func userShow(cmd *cli.Cmd) {
	var b []byte
	var uid *int
	var err error
	var u *api.User
	var c *api.Client

	cmd.Spec = "-u"
	uid = cmd.Int(cli.IntOpt{
		Name: "u uid",
		Desc: "User ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if u, err = c.GetUser(*uid); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func userCreate(cmd *cli.Cmd) {
	var b []byte
	var err error
	var f *api.UserForm
	var c *api.Client
	var u *api.User
	var accountType *int
	var domains, organizations *[]int
	var lowScore, highScore *api.LocalFloat64
	var enabled, sendReports, spamChecks, blockMacros *bool
	var username, firstname, lastname, password1, password2, email, timezone *string

	cmd.Spec = "--username --password1 --password2 --email --timezone --domain... [--firstname][--lastname][--accountType][--enabled][--sendReports][--spamChecks][--lowScore][--highScore][--blockMacros][--organization...]"

	username = cmd.String(cli.StringOpt{
		Name: "username",
		Desc: "The username for the account",
	})
	firstname = cmd.String(cli.StringOpt{
		Name: "firstname",
		Desc: "The first names",
	})
	lastname = cmd.String(cli.StringOpt{
		Name: "lastname",
		Desc: "The last names",
	})
	password1 = cmd.String(cli.StringOpt{
		Name: "password1",
		Desc: "The account Password",
	})
	password2 = cmd.String(cli.StringOpt{
		Name: "password2",
		Desc: "Retype account Password",
	})
	email = cmd.String(cli.StringOpt{
		Name: "email",
		Desc: "Email Address",
	})
	timezone = cmd.String(cli.StringOpt{
		Name: "timezone",
		Desc: "Users timezone, all dates and times will be displayed in this timezone",
	})
	accountType = cmd.Int(cli.IntOpt{
		Name:  "accountType",
		Desc:  "The account type",
		Value: 3,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enabled",
		Desc: "Enable or disable this account",
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name: "sendReports",
		Desc: "If enabled the user will receive reports from the system",
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:  "spamChecks",
		Desc:  "Enable or disable spam checking",
		Value: true,
	})
	cmd.VarOpt("lowScore", lowScore, "The score at which an email is considered to be suspected spam, 0.0 uses system defaults")
	cmd.VarOpt("highScore", highScore, "The score at which an email is considered to be definitely spam, 0.0 uses system defaults")
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "blockMacros",
		Desc: "Enable or disable blocking Attachments with Macros",
	})
	domains = cmd.Ints(cli.IntsOpt{
		Name: "domain",
		Desc: "The domains to which this user belongs",
	})
	organizations = cmd.Ints(cli.IntsOpt{
		Name: "organization",
		Desc: "The organizations to which this user belongs",
	})
	cmd.Action = func() {
		f = &api.UserForm{
			Username:      username,
			Firstname:     firstname,
			Lastname:      lastname,
			Password1:     password1,
			Password2:     password2,
			Email:         email,
			Timezone:      timezone,
			AccountType:   accountType,
			Enabled:       enabled,
			SendReport:    sendReports,
			SpamChecks:    spamChecks,
			LowScore:      lowScore,
			HighScore:     highScore,
			BlockMacros:   blockMacros,
			Domains:       *domains,
			Organizations: *organizations,
		}

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if u, err = c.CreateUser(f); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}

func userUpdate(cmd *cli.Cmd) {
	var uid *int
	var err error
	var u *api.User
	var c *api.Client
	var domains *[]int
	var f *api.UserForm
	var lowScore, highScore api.LocalFloat64
	var (
		enabled, sendReports, spamChecks, blockMacros                                                                                  *bool
		isSetfirstname, isSetlastname, isSetemail, isSettimezone, isSetEnable, isSetDisable, isEnableblockMacros, isDisableblockMacros bool
		isSetEnableReports, isSetDisableReports, isDisablespamChecks, isEnablespamChecks, isSetdomains, isSetlowScore, isSethighScore  bool
		username, firstname, lastname, email, timezone                                                                                 *string
	)

	cmd.Spec = "--uid --username --enable|--disable --enable-reports|--disable-reports --enable-spamChecks|--disable-spamChecks --enable-blockMacros|--disable-blockMacros [--email][--domain...][--timezone][--firstname][--lastname][--lowScore][--highScore]"

	uid = cmd.Int(cli.IntOpt{
		Name: "uid",
		Desc: "The user id for the account",
	})
	username = cmd.String(cli.StringOpt{
		Name: "username",
		Desc: "The username for the account",
	})
	firstname = cmd.String(cli.StringOpt{
		Name:      "firstname",
		Desc:      "The first names",
		SetByUser: &isSetfirstname,
	})
	lastname = cmd.String(cli.StringOpt{
		Name:      "lastname",
		Desc:      "The last names",
		SetByUser: &isSetlastname,
	})
	email = cmd.String(cli.StringOpt{
		Name:      "email",
		Desc:      "Email Address",
		SetByUser: &isSetemail,
	})
	timezone = cmd.String(cli.StringOpt{
		Name:      "timezone",
		Desc:      "Users timezone, all dates and times will be displayed in this timezone",
		SetByUser: &isSettimezone,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this account",
		SetByUser: &isSetEnable,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "disable",
		Desc:      "Disable this account",
		SetByUser: &isSetDisable,
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name:      "enable-reports",
		Desc:      "Enable reports for this account",
		SetByUser: &isSetEnableReports,
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name:      "disable-reports",
		Desc:      "Disable reports for this account",
		SetByUser: &isSetDisableReports,
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:      "enable-spamChecks",
		Desc:      "Enable spam checking",
		SetByUser: &isEnablespamChecks,
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:      "disable-spamChecks",
		Desc:      "Disable spam checking",
		SetByUser: &isDisablespamChecks,
	})
	cmd.Var(cli.VarOpt{
		Name:      "lowScore",
		Desc:      "The score at which an email is considered to be suspected spam, 0.0 uses system defaults",
		Value:     &lowScore,
		SetByUser: &isSetlowScore,
	})
	cmd.Var(cli.VarOpt{
		Name:      "highScore",
		Desc:      "The score at which an email is considered to be definitely spam, 0.0 uses system defaults",
		Value:     &highScore,
		SetByUser: &isSethighScore,
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name:      "enable-blockMacros",
		Desc:      "Enable or disable blocking Attachments with Macros",
		SetByUser: &isEnableblockMacros,
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name:      "disable-blockMacros",
		Desc:      "Enable or disable blocking Attachments with Macros",
		SetByUser: &isDisableblockMacros,
	})
	domains = cmd.Ints(cli.IntsOpt{
		Name:      "domain",
		Desc:      "The domains to which this user belongs",
		SetByUser: &isSetdomains,
	})
	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if u, err = c.GetUser(*uid); err != nil {
			log.Fatal(err)
		}

		var doms []int
		for _, dom := range u.Domains {
			doms = append(doms, dom.ID)
		}

		if isSetEnable {
			*enabled = true
		}
		if isSetDisable {
			*enabled = false
		}
		if isSetEnableReports {
			*sendReports = true
		}
		if isSetDisableReports {
			*sendReports = false
		}
		if isEnablespamChecks {
			*spamChecks = true
		}
		if isDisablespamChecks {
			*spamChecks = false
		}
		if isEnableblockMacros {
			*blockMacros = true
		}
		if isDisableblockMacros {
			*blockMacros = false
		}

		f = &api.UserForm{
			ID:          uid,
			Username:    &u.Username,
			Firstname:   &u.Firstname,
			Lastname:    &u.Lastname,
			Email:       &u.Email,
			Timezone:    &u.Timezone,
			Enabled:     enabled,
			SendReport:  sendReports,
			SpamChecks:  spamChecks,
			LowScore:    &u.LowScore,
			HighScore:   &u.HighScore,
			BlockMacros: blockMacros,
			Domains:     doms,
		}
		if isSetfirstname {
			f.Firstname = firstname
		}
		if isSetlastname {
			f.Lastname = lastname
		}
		if isSetemail {
			f.Email = email
		}
		if isSettimezone {
			f.Timezone = timezone
		}
		if isSetdomains {
			f.Domains = *domains
		}
		if isSetlowScore {
			f.LowScore = &lowScore
		}
		if isSethighScore {
			f.HighScore = &highScore
		}

		if err = c.UpdateUser(f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The user: %s has been updated\n", *username)
	}
}

func userDelete(cmd *cli.Cmd) {
	var uid *int

	uid = cmd.Int(cli.IntOpt{
		Name: "u uid",
		Desc: "The user id for the account",
	})

	cmd.Spec = "-u"

	cmd.Action = func() {
		var err error
		var c *api.Client

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteUser(*uid); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The user with id: %d has been deleted\n", *uid)
	}
}

func aliasShow(cmd *cli.Cmd) {
}

func aliasCreate(cmd *cli.Cmd) {
}

func aliasUpdate(cmd *cli.Cmd) {
}

func aliasDelete(cmd *cli.Cmd) {
}

func usersList(cmd *cli.Cmd) {
	cmd.Action = func() {
		var b []byte
		var err error
		var c *api.Client
		var u *api.UserList

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if u, err = c.GetUsers(nil); err != nil {
			log.Fatal(err)
		}

		if len(u.Items) == 0 {
			fmt.Println()
			return
		}

		if b, err = prettyjson.Marshal(u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
