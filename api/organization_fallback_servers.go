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

// FallBackServerOrg holds fallback server organization
type FallBackServerOrg struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FallBackServer holds organization fallback servers
type FallBackServer struct {
	ID           int               `json:"id,omitempty"`
	Address      string            `json:"address"`
	Protocol     int               `json:"protocol"`
	Port         int               `json:"port"`
	RequireTLS   bool              `json:"require_tls"`
	Enabled      bool              `json:"enabled"`
	Organization FallBackServerOrg `json:"organization"`
}

// GetFallBackServer returns radius settings
// https://www.baruwa.com/docs/api/#retrieve-a-fallback-server
func (c *Client) GetFallBackServer(serverID int) (server *FallBackServer, err error) {
	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("failbackservers/%d", serverID), server)

	return
}

// CreateFallBackServer creates radius settings
// https://www.baruwa.com/docs/api/#create-a-fallback-server
func (c *Client) CreateFallBackServer(organizationID int, server *FallBackServer) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("failbackservers/%d", organizationID), v, server)

	return
}

// UpdateFallBackServer updates radius settings
// https://www.baruwa.com/docs/api/#update-a-fallback-server
func (c *Client) UpdateFallBackServer(server *FallBackServer) (err error) {
	var v url.Values

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

	err = c.put(fmt.Sprintf("failbackservers/%d", server.ID), v, server)

	return
}

// DeleteFallBackServer deletes radius settings
// https://www.baruwa.com/docs/api/#delete-a-fallback-server
func (c *Client) DeleteFallBackServer(server *FallBackServer) (err error) {
	var v url.Values

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

	err = c.delete(fmt.Sprintf("failbackservers/%d", server.ID), v)

	return
}
