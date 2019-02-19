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

func userShow(cmd *cli.Cmd) {
	var (
		b   []byte
		uid *int
		err error
		u   *api.User
		c   *api.Client
	)

	cmd.Spec = "--id"

	uid = cmd.Int(cli.IntOpt{
		Name: "id",
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
	var (
		b                                                                    []byte
		err                                                                  error
		f                                                                    *api.UserForm
		c                                                                    *api.Client
		u                                                                    *api.User
		accountType                                                          *int
		domains, organizations                                               *[]int
		lowScore, highScore                                                  api.LocalFloat64
		enabled, sendReports, spamChecks, blockMacros                        *bool
		username, firstname, lastname, password1, password2, email, timezone *string
	)

	cmd.Spec = "--username --password1 --password2 --email --timezone --domain... [--firstname][--lastname][--account-type][--enable][--send-reports][--spam-checks][--low-score][--high-score][--block-macros][--organization...]"

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
		Name:  "account-type",
		Desc:  "The account type",
		Value: 3,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this account",
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name: "send-reports",
		Desc: "If enabled the user will receive reports from the system",
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:  "spam-checks",
		Desc:  "Enable spam checking",
		Value: true,
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
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "block-macros",
		Desc: "Enable blocking Attachments with Macros",
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
			LowScore:      &lowScore,
			HighScore:     &highScore,
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
	var (
		uid                                                                                        *int
		err                                                                                        error
		u                                                                                          *api.User
		c                                                                                          *api.Client
		domains                                                                                    *[]int
		f                                                                                          *api.UserForm
		lowScore, highScore                                                                        api.LocalFloat64
		enabled, sendReports, spamChecks, blockMacros                                              *bool
		isSetfirstname, isSetlastname, isSetemail, isSettimezone, isSetEnable, isEnableblockMacros bool
		isSetEnableReports, isEnablespamChecks, isSetdomains, isSetlowScore, isSethighScore        bool
		username, firstname, lastname, email, timezone                                             *string
	)

	cmd.Spec = "--id --username --enable|--disable --enable-reports|--disable-reports --enable-spam-checks|--disable-spam-checks --enable-block-macros|--disable-block-macros [--email][--domain...][--timezone][--firstname][--lastname][--low-score][--high-score]"

	uid = cmd.Int(cli.IntOpt{
		Name: "id",
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
		Name: "disable",
		Desc: "Disable this account",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this account",
		SetByUser: &isSetEnable,
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name: "disable-reports",
		Desc: "Disable reports for this account",
	})
	sendReports = cmd.Bool(cli.BoolOpt{
		Name:      "enable-reports",
		Desc:      "Enable reports for this account",
		SetByUser: &isSetEnableReports,
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name: "disable-spam-checks",
		Desc: "Disable spam checking",
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:      "enable-spam-checks",
		Desc:      "Enable spam checking",
		SetByUser: &isEnablespamChecks,
	})
	cmd.Var(cli.VarOpt{
		Name:      "low-score",
		Desc:      "The score at which an email is considered to be suspected spam, 0.0 uses system defaults",
		Value:     &lowScore,
		SetByUser: &isSetlowScore,
	})
	cmd.Var(cli.VarOpt{
		Name:      "high-score",
		Desc:      "The score at which an email is considered to be definitely spam, 0.0 uses system defaults",
		Value:     &highScore,
		SetByUser: &isSethighScore,
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "disable-block-macros",
		Desc: "Enable or disable blocking Attachments with Macros",
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name:      "enable-block-macros",
		Desc:      "Enable or disable blocking Attachments with Macros",
		SetByUser: &isEnableblockMacros,
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

		*enabled = isSetEnable
		*sendReports = isSetEnableReports
		*spamChecks = isEnablespamChecks
		*blockMacros = isEnableblockMacros

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
		Name: "id",
		Desc: "The user id for the account",
	})

	cmd.Spec = "--id"

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

func usersList(cmd *cli.Cmd) {
	var (
		page    *int
		b       []byte
		err     error
		pageSet bool
		c       *api.Client
		u       *api.UserList
		opts    *api.ListOptions
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
				Page: fmt.Sprintf(api.UserListURL, *serverURL, api.APIVersion, *page),
			}
		}

		if u, err = c.GetUsers(opts); err != nil {
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
