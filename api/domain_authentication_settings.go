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

// AuthServer holds an authentication server
type AuthServer struct {
	ID              int    `json:"id,omitempty"`
	Address         string `json:"address"`
	Protocol        int    `json:"protocol"`
	Port            int    `json:"port"`
	Enabled         bool   `json:"enabled"`
	SplitAddress    bool   `json:"split_address"`
	UserMapTemplate string `json:"user_map_template"`
}

// GetAuthServer returns an authentication server
// https://www.baruwa.com/docs/api/#retrieve-authentication-settings
func (c *Client) GetAuthServer(domainID, serverID int) (server *AuthServer, err error) {
	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("authservers/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateAuthServer creates an authentication server
// https://www.baruwa.com/docs/api/#create-authentication-settings
func (c *Client) CreateAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("authservers/%d", domainID), v, server)

	return
}

// UpdateAuthServer updates an authentication server
// https://www.baruwa.com/docs/api/#update-authentication-settings
func (c *Client) UpdateAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf("The server.ID param should be > 0")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("authservers/%d/%d", domainID, server.ID), v, server)

	return
}

// DeleteAuthServer deletes an authentication server
// https://www.baruwa.com/docs/api/#delete-authentication-settings
func (c *Client) DeleteAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf("The server.ID param should be > 0")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("authservers/%d/%d", domainID, server.ID), v)

	return
}
