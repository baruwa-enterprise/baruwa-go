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

// AliasDomain hold alias domain entries
type AliasDomain struct {
	ID   int    `json:"id" url:"id"`
	Name string `json:"name" url:"name"`
}

// DomainAlias holds domain aliases
type DomainAlias struct {
	ID            int          `json:"id,omitempty" url:"id,omitempty"`
	Name          string       `json:"name" url:"name"`
	Enabled       bool         `json:"status" url:"status"`
	AcceptInbound bool         `json:"accept_inbound" url:"accept_inbound"`
	Domain        *AliasDomain `json:"domain,omitempty" url:"domain,omitempty"`
}

// DomainAliasForm holds domain aliases
type DomainAliasForm struct {
	ID            int    `json:"id,omitempty" url:"id,omitempty"`
	Name          string `json:"name" url:"name"`
	Enabled       bool   `json:"status" url:"status"`
	AcceptInbound bool   `json:"accept_inbound" url:"accept_inbound"`
	Domain        int    `json:"domain,omitempty" url:"domain,omitempty"`
}

// DomainAliasList holds domain smarthosts
type DomainAliasList struct {
	Items []DomainAlias `json:"items"`
	Links Links         `json:"links"`
	Meta  Meta          `json:"meta"`
}

// GetDomainAliases returns a DomainList object
// This contains a paginated list of domain aliases and links
// to the neighbouring pages.
// https://www.baruwa.com/docs/api/#domain-aliases
func (c *Client) GetDomainAliases(domainID int, opts *ListOptions) (l *DomainAliasList, err error) {
	l = &DomainAliasList{}

	err = c.get(fmt.Sprintf("domainaliases/%d", domainID), opts, l)

	return
}

// GetDomainAlias returns a domain alias
// https://www.baruwa.com/docs/api/#retrieve-domain-alias
func (c *Client) GetDomainAlias(domainID, aliasID int) (alias *DomainAlias, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if aliasID <= 0 {
		err = fmt.Errorf(aliasIDError)
		return
	}

	alias = &DomainAlias{}

	err = c.get(fmt.Sprintf("domainaliases/%d/%d", domainID, aliasID), nil, alias)

	return
}

// CreateDomainAlias creates a domain alias
// https://www.baruwa.com/docs/api/#create-a-domain-alias
func (c *Client) CreateDomainAlias(domainID int, form *DomainAliasForm) (alias *DomainAlias, err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if form == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	alias = &DomainAlias{}

	err = c.post(fmt.Sprintf("domainaliases/%d", domainID), v, alias)

	return
}

// UpdateDomainAlias updates a domain alias
// https://www.baruwa.com/docs/api/#update-a-domain-alias
func (c *Client) UpdateDomainAlias(domainID int, form *DomainAliasForm) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if form == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if form.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("domainaliases/%d/%d", domainID, form.ID), v, nil)

	return
}

// DeleteDomainAlias deletes an domain alias
// https://www.baruwa.com/docs/api/#delete-a-domain-alias
func (c *Client) DeleteDomainAlias(domainID int, form *DomainAliasForm) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if form == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if form.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("domainaliases/%d/%d", domainID, form.ID), v)

	return
}
