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

// AliasAddress hosts alias addresses
// https://www.baruwa.com/docs/api/#alias-addresses
type AliasAddress struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address"`
	Enabled bool   `json:"enabled"`
}

// GetAliasAddress returns an alias address
// https://www.baruwa.com/docs/api/#retrieve-an-existing-alias-address
func (c *Client) GetAliasAddress(id int) (alias *AliasAddress, err error) {
	if id <= 0 {
		err = fmt.Errorf("The id param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("aliasaddresses/%d", id), alias)

	return
}

// CreateAliasAddress creates an alias address
// https://www.baruwa.com/docs/api/#create-an-alias-address
func (c *Client) CreateAliasAddress(userid int, alias *AliasAddress) (err error) {
	var v url.Values

	if userid <= 0 {
		err = fmt.Errorf("The userid param should be > 0")
		return
	}

	if alias == nil {
		err = fmt.Errorf("The alias param cannot be nil")
		return
	}

	if v, err = query.Values(alias); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("aliasaddresses/%d", userid), v, alias)

	return
}

// UpdateAliasAddress updates an alias address
// https://www.baruwa.com/docs/api/#update-an-alias-address
func (c *Client) UpdateAliasAddress(alias *AliasAddress) (err error) {
	var v url.Values

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

	err = c.put(fmt.Sprintf("aliasaddresses/%d", alias.ID), v, alias)

	return
}

// DeleteAliasAddress deletes an alias address
// https://www.baruwa.com/docs/api/#delete-an-alias-address
func (c *Client) DeleteAliasAddress(id int) (err error) {
	if id <= 0 {
		err = fmt.Errorf("The id param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("aliasaddresses/%d", id))

	return
}
