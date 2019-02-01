// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Domain holds domains
type Domain struct {
	ID                int     `json:"id,omitempty"`
	Name              string  `json:"name"`
	SiteURL           string  `json:"site_url"`
	Status            bool    `json:"status"`
	AcceptInbound     bool    `json:"accept_inbound"`
	DiscardMail       bool    `json:"discard_mail"`
	SMTPCallout       bool    `json:"smtp_callout"`
	LdapCallout       bool    `json:"ldap_callout"`
	VirusChecks       bool    `json:"virus_checks"`
	VirusChecksAtSMTP bool    `json:"virus_checks_at_smtp"`
	BlockMacros       bool    `json:"block_macros"`
	SpamChecks        bool    `json:"spam_checks"`
	SpamActions       int     `json:"spam_actions"`
	HighspamActions   int     `json:"highspam_actions"`
	VirusActions      int     `json:"virus_actions"`
	LowScore          float64 `json:"low_score"`
	HighScore         float64 `json:"high_score"`
	MessageSize       string  `json:"message_size"`
	DeliveryMode      int     `json:"delivery_mode"`
	Language          string  `json:"language"`
	Timezone          string  `json:"timezone"`
	ReportEvery       int     `json:"report_every"`
	Organizations     int     `json:"organizations,omitempty"`
}

// GetDomain returns a domain
// https://www.baruwa.com/docs/api/#retrieve-a-domain
func (c *Client) GetDomain(domainID int) (domain *Domain, err error) {
	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("domains/%d", domainID), nil, domain)

	return
}

// GetDomainByName returns a domain
// https://www.baruwa.com/docs/api/#retrieve-a-domain-by-name
func (c *Client) GetDomainByName(domainName string) (domain *Domain, err error) {
	if domainName == "" {
		err = fmt.Errorf("The domainName param is required")
		return
	}

	err = c.get(fmt.Sprintf("domains/byname/%s", domainName), nil, domain)

	return
}

// CreateDomain creates a domain
// https://www.baruwa.com/docs/api/#create-a-new-domain
func (c *Client) CreateDomain(domain *Domain) (err error) {
	var v url.Values

	if domain == nil {
		err = fmt.Errorf("The domain param cannot be nil")
		return
	}

	if v, err = query.Values(domain); err != nil {
		return
	}

	err = c.post("domains", v, domain)

	return
}

// UpdateDomain updates a domain
// https://www.baruwa.com/docs/api/#update-a-domain
func (c *Client) UpdateDomain(domain *Domain) (err error) {
	var v url.Values

	if domain == nil {
		err = fmt.Errorf("The domain param cannot be nil")
		return
	}

	if domain.ID <= 0 {
		err = fmt.Errorf("The domain.ID param should be > 0")
		return
	}

	if v, err = query.Values(domain); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("domains/%d", domain.ID), v, domain)

	return
}

// DeleteDomain deletes a domain
// https://www.baruwa.com/docs/api/#delete-a-domain
func (c *Client) DeleteDomain(domainID int) (err error) {
	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("domains/%d", domainID), nil)

	return
}

// ListDomains returns paged domain list
// https://www.baruwa.com/docs/api/#list-all-domains
func (c *Client) ListDomains() (err error) {
	return
}
