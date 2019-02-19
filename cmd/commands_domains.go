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

func domainShow(cmd *cli.Cmd) {
	var (
		id    *int
		b     []byte
		err   error
		name  *string
		idSet bool
		c     *api.Client
		d     *api.Domain
	)

	id = cmd.Int(cli.IntOpt{
		Name:      "id",
		Desc:      "Domain ID",
		SetByUser: &idSet,
	})
	name = cmd.String(cli.StringOpt{
		Name: "name",
		Desc: "Domain name",
	})

	cmd.Spec = "--id|--name"

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if idSet {
			if d, err = c.GetDomain(*id); err != nil {
				log.Fatal(err)
			}
		} else {
			if d, err = c.GetDomainByName(*name); err != nil {
				log.Fatal(err)
			}
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainCreate(cmd *cli.Cmd) {
	var (
		b                                                                     []byte
		err                                                                   error
		c                                                                     *api.Client
		d                                                                     *api.Domain
		lowScore, highScore                                                   api.LocalFloat64
		spamActions, highSpamActions, virusActions, deliveryMode, reportEvery *int
		organizations                                                         *[]int
		name, siteURL, messageSize, language, timeZone                        *string
		enabled, acceptInbound, discardMail, smtpCallout, ldapCallout         *bool
		virusChecks, virusChecksAtSMTP, blockMacros, spamChecks               *bool
	)

	cmd.Spec = "--name --site-url [--message-size][--language][--timezone][--enable][--accept-inbound][--discard-mail][--smtp-callout][--ldap-callout][--virus-checks][--virus-checks-at-smtp][--block-macros][--spam-checks][--spam-actions][--high-spam-actions][--virus-actions][--delivery-mode][--report-frequency][--low-score][--high-score][--organization...]"

	name = cmd.String(cli.StringOpt{
		Name: "name",
		Desc: "Domain name",
	})
	siteURL = cmd.String(cli.StringOpt{
		Name: "site-url",
		Desc: "Site url",
	})
	messageSize = cmd.String(cli.StringOpt{
		Name:  "message-size",
		Desc:  "The maximum message size for email sent to and from this domain",
		Value: "0",
	})
	language = cmd.String(cli.StringOpt{
		Name:  "language",
		Desc:  "The default language for users under this domain",
		Value: "en",
	})
	timeZone = cmd.String(cli.StringOpt{
		Name:  "timezone",
		Desc:  "The default language for users under this domain",
		Value: "UTC",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "enable",
		Desc: "Enable this domain",
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name: "accept-inbound",
		Desc: "Enable accepting of inbound mail to this domain",
	})
	discardMail = cmd.Bool(cli.BoolOpt{
		Name: "discard-mail",
		Desc: "Enable discarding of all mail sent to this domain",
	})
	smtpCallout = cmd.Bool(cli.BoolOpt{
		Name: "smtp-callout",
		Desc: "Enable SMTP callout based recipient verification",
	})
	ldapCallout = cmd.Bool(cli.BoolOpt{
		Name: "ldap-callout",
		Desc: "Enable LDAP email address verification for this domain",
	})
	virusChecks = cmd.Bool(cli.BoolOpt{
		Name: "virus-checks",
		Desc: "Enable virus checks for this domain",
	})
	virusChecksAtSMTP = cmd.Bool(cli.BoolOpt{
		Name: "virus-checks-at-smtp",
		Desc: "Run Virus Checks at SMTP time",
	})
	blockMacros = cmd.Bool(cli.BoolOpt{
		Name: "block-macros",
		Desc: "Enable blocking Attachments with Macros",
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name: "spam-checks",
		Desc: "Enable spam checks for this domain",
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
	virusActions = cmd.Int(cli.IntOpt{
		Name:  "virus-actions",
		Desc:  "What to do with messages that match virus signatures",
		Value: 2,
	})
	deliveryMode = cmd.Int(cli.IntOpt{
		Name:  "delivery-mode",
		Desc:  "How messages should be delivered when multiple servers are set",
		Value: 1,
	})
	reportEvery = cmd.Int(cli.IntOpt{
		Name:  "report-frequency",
		Desc:  "How often PDF reports should be sent",
		Value: 3,
	})
	organizations = cmd.Ints(cli.IntsOpt{
		Name: "organization",
		Desc: "The organizations that own this domain",
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

		d = &api.Domain{
			Name:              *name,
			SiteURL:           *siteURL,
			MessageSize:       *messageSize,
			Language:          *language,
			Timezone:          *timeZone,
			Enabled:           *enabled,
			AcceptInbound:     *acceptInbound,
			DiscardMail:       *discardMail,
			SMTPCallout:       *smtpCallout,
			LdapCallout:       *ldapCallout,
			VirusChecks:       *virusChecks,
			VirusChecksAtSMTP: *virusChecksAtSMTP,
			BlockMacros:       *blockMacros,
			SpamChecks:        *spamChecks,
			SpamActions:       *spamActions,
			HighspamActions:   *highSpamActions,
			VirusActions:      *virusActions,
			DeliveryMode:      *deliveryMode,
			ReportEvery:       *reportEvery,
			Organizations:     *organizations,
			LowScore:          lowScore,
			HighScore:         highScore,
		}

		if err = c.CreateDomain(d); err != nil {
			log.Fatal(err)
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", b)
	}
}

func domainUpdate(cmd *cli.Cmd) {
	var (
		d                                                                                        *api.Domain
		err                                                                                      error
		c                                                                                        *api.Client
		organizations                                                                            *[]int
		lowScore, highScore                                                                      api.LocalFloat64
		id, spamActions, highSpamActions, virusActions, deliveryMode, reportEvery                *int
		name, siteURL, messageSize, language, timeZone                                           *string
		enabled, acceptInbound, discardMail, smtpCallout, ldapCallout                            *bool
		virusChecks, virusChecksAtSMTP, blockMacros, spamChecks                                  *bool
		nameSet, siteURLSet, messageSizeSet, languageSet, timeZoneSet, lowScoreSet, highScoreSet bool
		enabledSet, acceptInboundSet, discardMailSet, smtpCalloutSet, ldapCalloutSet             bool
		virusChecksSet, virusChecksAtSMTPSet, blockMacrosSet, spamChecksSet, spamActionsSet      bool
		highSpamActionsSet, virusActionsSet, deliveryModeSet, reportEverySet, organizationsSet   bool
	)

	cmd.Spec = "--id --enable|--disable --accept-inbound|--disable-inbound --discard-mail|--disable-discard-mail --smtp-callout|--disable-smtp-callout --ldap-callout|--disable-ldap-callout --virus-checks|--disable-virus-checks --virus-checks|--disable-virus-checks --virus-checks-at-smtp|--disable-virus-checks-at-smtp --block-macros|--disable-block-macros --spam-checks|--disable-spam-checks [--name][--site-url][--message-size][--language][--timezone][--spam-actions][--high-spam-actions][--virus-actions][--delivery-mode][--report-frequency][--low-score][--high-score][--organization...]"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain ID",
	})
	name = cmd.String(cli.StringOpt{
		Name:      "name",
		Desc:      "Domain name",
		SetByUser: &nameSet,
	})
	siteURL = cmd.String(cli.StringOpt{
		Name:      "site-url",
		Desc:      "Site url",
		SetByUser: &siteURLSet,
	})
	messageSize = cmd.String(cli.StringOpt{
		Name:      "message-size",
		Desc:      "The maximum message size for email sent to and from this domain",
		Value:     "0",
		SetByUser: &messageSizeSet,
	})
	language = cmd.String(cli.StringOpt{
		Name:      "language",
		Desc:      "The default language for users under this domain",
		Value:     "en",
		SetByUser: &languageSet,
	})
	timeZone = cmd.String(cli.StringOpt{
		Name:      "timezone",
		Desc:      "The default language for users under this domain",
		Value:     "UTC",
		SetByUser: &timeZoneSet,
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name: "disable",
		Desc: "Disable this domain",
	})
	enabled = cmd.Bool(cli.BoolOpt{
		Name:      "enable",
		Desc:      "Enable this domain",
		SetByUser: &enabledSet,
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name: "disable-inbound",
		Desc: "Disable accepting of inbound mail to this domain",
	})
	acceptInbound = cmd.Bool(cli.BoolOpt{
		Name:      "accept-inbound",
		Desc:      "Enable accepting of inbound mail to this domain",
		SetByUser: &acceptInboundSet,
	})
	discardMail = cmd.Bool(cli.BoolOpt{
		Name: "disable-discard-mail",
		Desc: "Disable discarding of all mail sent to this domain",
	})
	discardMail = cmd.Bool(cli.BoolOpt{
		Name:      "discard-mail",
		Desc:      "Enable discarding of all mail sent to this domain",
		SetByUser: &discardMailSet,
	})
	smtpCallout = cmd.Bool(cli.BoolOpt{
		Name: "disable-smtp-callout",
		Desc: "Disable SMTP callout based recipient verification",
	})
	smtpCallout = cmd.Bool(cli.BoolOpt{
		Name:      "smtp-callout",
		Desc:      "Enable SMTP callout based recipient verification",
		SetByUser: &smtpCalloutSet,
	})
	ldapCallout = cmd.Bool(cli.BoolOpt{
		Name: "disable-ldap-callout",
		Desc: "Disable LDAP email address verification for this domain",
	})
	ldapCallout = cmd.Bool(cli.BoolOpt{
		Name:      "ldap-callout",
		Desc:      "Enable LDAP email address verification for this domain",
		SetByUser: &ldapCalloutSet,
	})
	virusChecks = cmd.Bool(cli.BoolOpt{
		Name: "disable-virus-checks",
		Desc: "Disable virus checks for this domain",
	})
	virusChecks = cmd.Bool(cli.BoolOpt{
		Name:      "virus-checks",
		Desc:      "Enable virus checks for this domain",
		SetByUser: &virusChecksSet,
	})
	virusChecksAtSMTP = cmd.Bool(cli.BoolOpt{
		Name: "disable-virus-checks-at-smtp",
		Desc: "Disable run Virus Checks at SMTP time",
	})
	virusChecksAtSMTP = cmd.Bool(cli.BoolOpt{
		Name:      "virus-checks-at-smtp",
		Desc:      "Run Virus Checks at SMTP time",
		SetByUser: &virusChecksAtSMTPSet,
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
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name: "disable-spam-checks",
		Desc: "Disable spam checks for this domain",
	})
	spamChecks = cmd.Bool(cli.BoolOpt{
		Name:      "spam-checks",
		Desc:      "Enable spam checks for this domain",
		SetByUser: &spamChecksSet,
	})
	spamActions = cmd.Int(cli.IntOpt{
		Name:      "spam-actions",
		Desc:      "What to do with Probable spam",
		Value:     2,
		SetByUser: &spamActionsSet,
	})
	highSpamActions = cmd.Int(cli.IntOpt{
		Name:      "high-spam-actions",
		Desc:      "What to do with Definite spam",
		Value:     2,
		SetByUser: &highSpamActionsSet,
	})
	virusActions = cmd.Int(cli.IntOpt{
		Name:      "virus-actions",
		Desc:      "What to do with messages that match virus signatures",
		Value:     2,
		SetByUser: &virusActionsSet,
	})
	deliveryMode = cmd.Int(cli.IntOpt{
		Name:      "delivery-mode",
		Desc:      "How messages should be delivered when multiple servers are set",
		Value:     1,
		SetByUser: &deliveryModeSet,
	})
	reportEvery = cmd.Int(cli.IntOpt{
		Name:      "report-frequency",
		Desc:      "How often PDF reports should be sent",
		Value:     3,
		SetByUser: &reportEverySet,
	})
	organizations = cmd.Ints(cli.IntsOpt{
		Name:      "organization",
		Desc:      "The organizations that own this domain",
		SetByUser: &organizationsSet,
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

		if d, err = c.GetDomain(*id); err != nil {
			log.Fatal(err)
		}

		if nameSet {
			d.Name = *name
		}
		if siteURLSet {
			d.SiteURL = *siteURL
		}
		if messageSizeSet {
			d.MessageSize = *messageSize
		}
		if languageSet {
			d.Language = *language
		}
		if timeZoneSet {
			d.Timezone = *timeZone
		}
		if lowScoreSet {
			d.LowScore = lowScore
		}
		if highScoreSet {
			d.HighScore = highScore
		}
		if spamActionsSet {
			d.SpamActions = *spamActions
		}
		if highSpamActionsSet {
			d.HighspamActions = *highSpamActions
		}
		if virusActionsSet {
			d.VirusActions = *virusActions
		}
		if deliveryModeSet {
			d.DeliveryMode = *deliveryMode
		}
		if reportEverySet {
			d.ReportEvery = *reportEvery
		}
		if organizationsSet {
			d.Organizations = *organizations
		}
		*enabled = enabledSet
		*acceptInbound = acceptInboundSet
		*discardMail = discardMailSet
		*smtpCallout = smtpCalloutSet
		*ldapCallout = ldapCalloutSet
		*virusChecks = virusChecksSet
		*virusChecksAtSMTP = virusChecksAtSMTPSet
		*blockMacros = blockMacrosSet
		*spamChecks = spamChecksSet

		d.Enabled = *enabled
		d.AcceptInbound = *acceptInbound
		d.DiscardMail = *discardMail
		d.SMTPCallout = *smtpCallout
		d.LdapCallout = *ldapCallout
		d.VirusChecks = *virusChecks
		d.VirusChecksAtSMTP = *virusChecksAtSMTP
		d.BlockMacros = *blockMacros
		d.SpamChecks = *spamChecks

		if err = c.UpdateDomain(d); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain: %s has been updated\n", d.Name)
	}
}

func domainDelete(cmd *cli.Cmd) {
	var (
		id  *int
		err error
		c   *api.Client
	)

	cmd.Spec = "--id"

	id = cmd.Int(cli.IntOpt{
		Name: "id",
		Desc: "Domain ID",
	})

	cmd.Action = func() {
		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if err = c.DeleteDomain(*id); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The domain: %d has been deleted\n", *id)
	}
}

func domainsList(cmd *cli.Cmd) {
	var (
		page    *int
		b       []byte
		err     error
		pageSet bool
		c       *api.Client
		d       *api.DomainList
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
				Page: fmt.Sprintf(api.DomainListURL, *serverURL, api.APIVersion, *page),
			}
		}

		if d, err = c.GetDomains(opts); err != nil {
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
