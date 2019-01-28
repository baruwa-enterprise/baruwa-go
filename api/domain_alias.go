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

// AliasDomain hold alias domain entries
type AliasDomain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DomainAlias holds domain aliases
type DomainAlias struct {
	ID      int           `json:"id,omitempty"`
	Address string        `json:"address"`
	Enabled bool          `json:"enabled"`
	Domain  []AliasDomain `json:"domain,omitempty"`
}

// GetDomainAlias returns a domain alias
// https://www.baruwa.com/docs/api/#retrieve-domain-alias
func (c *Client) GetDomainAlias(domainid, aliasid int) (alias *DomainAlias, err error) {
	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if aliasid <= 0 {
		err = fmt.Errorf("The aliasid param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("domainaliases/%d/%d", domainid, aliasid), alias)

	return
}

// CreateDomainAlias creates a domain alias
// https://www.baruwa.com/docs/api/#create-a-domain-alias
func (c *Client) CreateDomainAlias(domainid int, alias *DomainAlias) (err error) {
	var v url.Values

	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if alias == nil {
		err = fmt.Errorf("The alias param cannot be nil")
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("domainaliases/%d", domainid), v, alias)

	return
}

// UpdateDomainAlias updates a domain alias
// https://www.baruwa.com/docs/api/#update-a-domain-alias
func (c *Client) UpdateDomainAlias(domainid int, alias *DomainAlias) (err error) {
	var v url.Values

	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if alias == nil {
		err = fmt.Errorf("The alias param cannot be nil")
		return
	}

	if alias.ID <= 0 {
		err = fmt.Errorf("The alias.ID param should be > 0")
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("domainaliases/%d/%d", domainid, alias.ID), v, alias)

	return
}

// DeleteDomainAlias deletes an domain alias
// https://www.baruwa.com/docs/api/#delete-a-domain-alias
func (c *Client) DeleteDomainAlias(domainid, aliasid int) (err error) {
	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if aliasid <= 0 {
		err = fmt.Errorf("The aliasid param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("domainaliases/%d/%d", domainid, aliasid))

	return
}
