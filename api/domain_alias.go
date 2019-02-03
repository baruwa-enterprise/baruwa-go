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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DomainAlias holds domain aliases
type DomainAlias struct {
	ID            int          `json:"id,omitempty"`
	Address       string       `json:"address"`
	Enabled       bool         `json:"status"`
	AcceptInbound bool         `json:"accept_inbound"`
	Domain        *AliasDomain `json:"domain,omitempty"`
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
func (c *Client) CreateDomainAlias(domainID int, alias *DomainAlias) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("domainaliases/%d", domainID), v, alias)

	return
}

// UpdateDomainAlias updates a domain alias
// https://www.baruwa.com/docs/api/#update-a-domain-alias
func (c *Client) UpdateDomainAlias(domainID int, alias *DomainAlias) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if alias.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("domainaliases/%d/%d", domainID, alias.ID), v, nil)

	return
}

// DeleteDomainAlias deletes an domain alias
// https://www.baruwa.com/docs/api/#delete-a-domain-alias
func (c *Client) DeleteDomainAlias(domainID int, alias *DomainAlias) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if alias.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("domainaliases/%d/%d", domainID, alias.ID), v)

	return
}
