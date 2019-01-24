// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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
func (c *Client) GetDomain(id int) (domain *Domain, err error) {
	return
}

// GetDomainByName returns a domain
// https://www.baruwa.com/docs/api/#retrieve-a-domain-by-name
func (c *Client) GetDomainByName(name string) (domain *Domain, err error) {
	return
}

// CreateDomain creates a domain
// https://www.baruwa.com/docs/api/#create-a-new-domain
func (c *Client) CreateDomain(domain *Domain) (err error) {
	return
}

// UpdateDomain updates a domain
// https://www.baruwa.com/docs/api/#update-a-domain
func (c *Client) UpdateDomain(domain *Domain) (err error) {
	return
}

// DeleteDomain deletes a domain
// https://www.baruwa.com/docs/api/#delete-a-domain
func (c *Client) DeleteDomain(id int) (err error) {
	return
}

// ListDomains returns paged domain list
// https://www.baruwa.com/docs/api/#list-all-domains
func (c *Client) ListDomains() (err error) {
	return
}
