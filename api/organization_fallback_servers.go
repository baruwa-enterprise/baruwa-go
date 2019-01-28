// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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
func (c *Client) GetFallBackServer(id int) (server *FallBackServer, err error) {
	return
}

// CreateFallBackServer creates radius settings
// https://www.baruwa.com/docs/api/#create-a-fallback-server
func (c *Client) CreateFallBackServer(server *FallBackServer) (err error) {
	return
}

// UpdateFallBackServer updates radius settings
// https://www.baruwa.com/docs/api/#update-a-fallback-server
func (c *Client) UpdateFallBackServer(server *FallBackServer) (err error) {
	return
}

// DeleteFallBackServer deletes radius settings
// https://www.baruwa.com/docs/api/#delete-a-fallback-server
func (c *Client) DeleteFallBackServer(id int) (err error) {
	return
}
