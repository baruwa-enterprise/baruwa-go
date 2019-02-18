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

// OrgSmartHost holds domain smarthosts
type OrgSmartHost struct {
	ID          int    `json:"id,omitempty" url:"id,omitempty"`
	Address     string `json:"address" url:"address"`
	Username    string `json:"username" url:"username"`
	Password    string `json:"password,omitempty" url:"password,omitempty"`
	Port        int    `json:"port" url:"port"`
	RequireTLS  bool   `json:"require_tls" url:"require_tls"`
	Enabled     bool   `json:"enabled" url:"enabled"`
	Description string `json:"description" url:"description"`
}

// OrgSmartHostList holds domain smarthosts
type OrgSmartHostList struct {
	Items []OrgSmartHost `json:"items"`
	Links Links          `json:"links"`
	Meta  Meta           `json:"meta"`
}

// GetOrgSmartHosts returns a OrgSmartHostList object
// This contains a paginated list of Organization smarthosts and links
// to the neighbouring pages.
// https://www.baruwa.com/docs/api/#listing-organization-smarthosts
func (c *Client) GetOrgSmartHosts(organizationID int, opts *ListOptions) (l *OrgSmartHostList, err error) {
	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
		return
	}

	l = &OrgSmartHostList{}

	err = c.get(fmt.Sprintf("organizations/smarthosts/%d", organizationID), opts, l)

	return
}

// GetOrgSmartHost returns a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-a-organization-smarthost
func (c *Client) GetOrgSmartHost(organizationID, serverID int) (server *OrgSmartHost, err error) {
	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &OrgSmartHost{}

	err = c.get(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, serverID), nil, server)

	return
}

// CreateOrgSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-an-organization-smarthost
func (c *Client) CreateOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
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

	err = c.post(fmt.Sprintf("organizations/smarthosts/%d", organizationID), v, server)

	return
}

// UpdateOrgSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-organization-smarthost
func (c *Client) UpdateOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
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

	err = c.put(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, server.ID), v, server)

	return
}

// DeleteOrgSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-system-status
func (c *Client) DeleteOrgSmartHost(organizationID int, server *OrgSmartHost) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf(organizationIDError)
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

	err = c.delete(fmt.Sprintf("organizations/smarthosts/%d/%d", organizationID, server.ID), v)

	return
}
