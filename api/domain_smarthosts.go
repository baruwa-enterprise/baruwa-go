// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// DomainSmartHost holds domain smarthosts
type DomainSmartHost struct {
	ID          int    `json:"id,omitempty"`
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Port        int    `json:"port"`
	RequireTLS  bool   `json:"require_tls"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

// GetDomainSmartHost returns a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-a-domain-smarthost
func (c *Client) GetDomainSmartHost(id int) (server *DomainSmartHost, err error) {
	return
}

// CreateDomainSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-a-domain-smarthost
func (c *Client) CreateDomainSmartHost(server *DomainSmartHost) (err error) {
	return
}

// UpdateDomainSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-domain-smarthost
func (c *Client) UpdateDomainSmartHost(server *DomainSmartHost) (err error) {
	return
}

// DeleteDomainSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#delete-a-domain-smarthost
func (c *Client) DeleteDomainSmartHost(id int) (err error) {
	return
}
