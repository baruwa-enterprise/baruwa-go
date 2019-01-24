// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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
func (c *Client) GetDomainAlias(id int) (alias *DomainAlias, err error) {
	return
}

// CreateDomainAlias creates a domain alias
// https://www.baruwa.com/docs/api/#create-a-domain-alias
func (c *Client) CreateDomainAlias(domain *DomainAlias) (err error) {
	return
}

// UpdateDomainAlias updates a domain alias
// https://www.baruwa.com/docs/api/#update-a-domain-alias
func (c *Client) UpdateDomainAlias(domain *DomainAlias) (err error) {
	return
}

// DeleteDomainAlias deletes an domain alias
// https://www.baruwa.com/docs/api/#delete-a-domain-alias
func (c *Client) DeleteDomainAlias(id int) (err error) {
	return
}
