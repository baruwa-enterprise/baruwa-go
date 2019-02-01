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
func (c *Client) GetOrgSmartHost(organizationID, serverID int) (server *OrgSmartHost, err error) {
	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, serverID), nil, server)

	return
}

// CreateOrgSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-an-organization-smarthost
func (c *Client) CreateOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
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

	err = c.post(fmt.Sprintf("organizations/smarthosts/%d", organizationID), v, server)

	return
}

// UpdateOrgSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-organization-smarthost
func (c *Client) UpdateOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
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

	err = c.put(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, server.ID), v, server)

	return
}

// DeleteOrgSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-system-status
func (c *Client) DeleteOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
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

	err = c.delete(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, server.ID), v)

	return
}
