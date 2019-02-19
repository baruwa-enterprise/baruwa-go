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
	ID              int    `json:"id,omitempty" url:"id,omitempty"`
	Address         string `json:"address" url:"address"`
	Port            int    `json:"port" url:"port"`
	Protocol        int    `json:"protocol" url:"protocol"`
	Enabled         bool   `json:"enabled" url:"enabled"`
	SplitAddress    bool   `json:"split_address" url:"split_address"`
	UserMapTemplate string `json:"user_map_template" url:"user_map_template"`
}

// AuthServerList holds authentication servers
type AuthServerList struct {
	Items []AuthServer `json:"items"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

// GetAuthServers returns a AuthServerList object
// This contains a paginated list of authentication servers and links
// to the neighbouring pages.
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#listing-authentication-settings
func (c *Client) GetAuthServers(domainID int, opts *ListOptions) (l *AuthServerList, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	l = &AuthServerList{}

	err = c.get(fmt.Sprintf("authservers/%d", domainID), opts, l)

	return
}

// GetAuthServer returns an authentication server
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-authentication-settings
func (c *Client) GetAuthServer(domainID, serverID int) (server *AuthServer, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &AuthServer{}

	err = c.get(fmt.Sprintf("authservers/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateAuthServer creates an authentication server
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-authentication-settings
func (c *Client) CreateAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	v, _ = query.Values(server)

	err = c.post(fmt.Sprintf("authservers/%d", domainID), v, server)

	return
}

// UpdateAuthServer updates an authentication server
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-authentication-settings
func (c *Client) UpdateAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.put(fmt.Sprintf("authservers/%d/%d", domainID, server.ID), v, nil)

	return
}

// DeleteAuthServer deletes an authentication server
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-authentication-settings
func (c *Client) DeleteAuthServer(domainID int, server *AuthServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.delete(fmt.Sprintf("authservers/%d/%d", domainID, server.ID), v)

	return
}
