// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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
func (c *Client) GetAuthServer(id int) (server AuthServer, err error) {
	return
}

// CreateAuthServer creates an authentication server
// https://www.baruwa.com/docs/api/#create-authentication-settings
func (c *Client) CreateAuthServer(server *AuthServer) (err error) {
	return
}

// UpdateAuthServer updates an authentication server
// https://www.baruwa.com/docs/api/#update-authentication-settings
func (c *Client) UpdateAuthServer(server *AuthServer) (err error) {
	return
}

// DeleteAuthServer deletes an authentication server
// https://www.baruwa.com/docs/api/#delete-authentication-settings
func (c *Client) DeleteAuthServer(id int) (err error) {
	return
}
