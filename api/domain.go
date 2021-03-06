// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Domain holds domains
type Domain struct {
	ID                int          `json:"id,omitempty" url:"id,omitempty"`
	Name              string       `json:"name" url:"name"`
	SiteURL           string       `json:"site_url" url:"site_url"`
	Enabled           bool         `json:"status" url:"status"`
	AcceptInbound     bool         `json:"accept_inbound" url:"accept_inbound"`
	DiscardMail       bool         `json:"discard_mail" url:"discard_mail"`
	SMTPCallout       bool         `json:"smtp_callout" url:"smtp_callout"`
	LdapCallout       bool         `json:"ldap_callout" url:"ldap_callout"`
	VirusChecks       bool         `json:"virus_checks" url:"virus_checks"`
	VirusChecksAtSMTP bool         `json:"virus_checks_at_smtp" url:"virus_checks_at_smtp"`
	BlockMacros       bool         `json:"block_macros" url:"block_macros"`
	SpamChecks        bool         `json:"spam_checks" url:"spam_checks"`
	SpamActions       int          `json:"spam_actions" url:"spam_actions"`
	HighspamActions   int          `json:"highspam_actions" url:"highspam_actions"`
	VirusActions      int          `json:"virus_actions" url:"virus_actions"`
	LowScore          LocalFloat64 `json:"low_score" url:"low_score"`
	HighScore         LocalFloat64 `json:"high_score" url:"high_score"`
	MessageSize       string       `json:"message_size" url:"message_size"`
	DeliveryMode      int          `json:"delivery_mode" url:"delivery_mode"`
	Language          string       `json:"language" url:"language"`
	Timezone          string       `json:"timezone" url:"timezone"`
	ReportEvery       int          `json:"report_every" url:"report_every"`
	Organizations     []int        `json:"organizations,omitempty" url:"organizations,omitempty"`
}

// DomainList holds domain smarthosts
type DomainList struct {
	Items []Domain `json:"items"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}

// GetDomains returns a DomainList object
// This contains a paginated list of domains and links
// to the neighbouring pages.
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#list-all-domains
func (c *Client) GetDomains(opts *ListOptions) (l *DomainList, err error) {
	l = &DomainList{}

	err = c.get("domains", opts, l)

	return
}

// GetDomain returns a domain
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-a-domain
func (c *Client) GetDomain(domainID int) (domain *Domain, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	domain = &Domain{}

	err = c.get(fmt.Sprintf("domains/%d", domainID), nil, domain)

	return
}

// GetDomainByName returns a domain
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-a-domain-by-name
func (c *Client) GetDomainByName(domainName string) (domain *Domain, err error) {
	if domainName == "" {
		err = fmt.Errorf(domainNameParamError)
		return
	}

	domain = &Domain{}

	err = c.get(fmt.Sprintf("domains/byname/%s", domainName), nil, domain)

	return
}

// CreateDomain creates a domain
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-a-new-domain
func (c *Client) CreateDomain(domain *Domain) (err error) {
	var v url.Values

	if domain == nil {
		err = fmt.Errorf(domainParamError)
		return
	}

	v, _ = query.Values(domain)

	err = c.post("domains", v, domain)

	return
}

// UpdateDomain updates a domain
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-a-domain
func (c *Client) UpdateDomain(domain *Domain) (err error) {
	var v url.Values

	if domain == nil {
		err = fmt.Errorf(domainParamError)
		return
	}

	if domain.ID <= 0 {
		err = fmt.Errorf(domainSIDError)
		return
	}

	v, _ = query.Values(domain)

	err = c.put(fmt.Sprintf("domains/%d", domain.ID), v, domain)

	return
}

// DeleteDomain deletes a domain
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-a-domain
func (c *Client) DeleteDomain(domainID int) (err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	err = c.delete(fmt.Sprintf("domains/%d", domainID), nil)

	return
}
