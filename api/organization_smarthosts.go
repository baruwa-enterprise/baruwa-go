// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// OrgSmartHost holds domain smarthosts
type OrgSmartHost struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Port        int    `json:"port"`
	RequireTLS  bool   `json:"require_tls"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

// GetOrgSmartHost returns a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-a-organization-smarthost
func (c *Client) GetOrgSmartHost(id int) (server *OrgSmartHost, err error) {
	return
}

// CreateOrgSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-an-organization-smarthost
func (c *Client) CreateOrgSmartHost(server *OrgSmartHost) (err error) {
	return
}

// UpdateOrgSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-organization-smarthost
func (c *Client) UpdateOrgSmartHost(server *OrgSmartHost) (err error) {
	return
}

// DeleteOrgSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-system-status
func (c *Client) DeleteOrgSmartHost(id int) (err error) {
	return
}
