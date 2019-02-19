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

// FallBackServerOrg holds fallback server organization
type FallBackServerOrg struct {
	ID   int    `json:"id" url:"id"`
	Name string `json:"name" url:"name"`
}

// FallBackServer holds organization fallback servers
type FallBackServer struct {
	ID               int                `json:"id,omitempty" url:"id,omitempty"`
	Address          string             `json:"address" url:"address"`
	Protocol         int                `json:"protocol" url:"protocol"`
	Port             int                `json:"port" url:"port"`
	RequireTLS       bool               `json:"require_tls" url:"require_tls"`
	VerificationOnly bool               `json:"verification_only" url:"verification_only"`
	Enabled          bool               `json:"enabled" url:"enabled"`
	Organization     *FallBackServerOrg `json:"organization" url:"organization"`
}

// FallBackServerList holds users
type FallBackServerList struct {
	Items []FallBackServer `json:"items"`
	Links Links            `json:"links"`
	Meta  Meta             `json:"meta"`
}

// GetFallBackServers returns a FallBackServerList object
// This contains a paginated list of fallback servers and links
// to the neighbouring pages.
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#fallback-servers
func (c *Client) GetFallBackServers(organizationID int, opts *ListOptions) (l *FallBackServerList, err error) {
	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
		return
	}

	l = &FallBackServerList{}

	err = c.get(fmt.Sprintf("fallbackservers/list/%d", organizationID), opts, l)

	return
}

// GetFallBackServer returns radius settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-a-fallback-server
func (c *Client) GetFallBackServer(serverID int) (server *FallBackServer, err error) {
	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &FallBackServer{}

	err = c.get(fmt.Sprintf("fallbackservers/%d", serverID), nil, server)

	return
}

// CreateFallBackServer creates radius settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-a-fallback-server
func (c *Client) CreateFallBackServer(organizationID int, server *FallBackServer) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	v, _ = query.Values(server)

	err = c.post(fmt.Sprintf("fallbackservers/%d", organizationID), v, server)

	return
}

// UpdateFallBackServer updates radius settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-a-fallback-server
func (c *Client) UpdateFallBackServer(server *FallBackServer) (err error) {
	var v url.Values

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.put(fmt.Sprintf("fallbackservers/%d", server.ID), v, server)

	return
}

// DeleteFallBackServer deletes radius settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-a-fallback-server
func (c *Client) DeleteFallBackServer(server *FallBackServer) (err error) {
	var v url.Values

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.delete(fmt.Sprintf("fallbackservers/%d", server.ID), v)

	return
}
