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
func (c *Client) GetDomainSmartHost(domainid, serverid int) (server *DomainSmartHost, err error) {
	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if serverid <= 0 {
		err = fmt.Errorf("The serverid param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("domains/smarthosts/%d/%d", domainid, serverid), server)

	return
}

// CreateDomainSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-a-domain-smarthost
func (c *Client) CreateDomainSmartHost(domainid int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("domains/smarthosts/%d", domainid), v, server)

	return
}

// UpdateDomainSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-domain-smarthost
func (c *Client) UpdateDomainSmartHost(domainid int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
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

	err = c.put(fmt.Sprintf("domains/smarthosts/%d/%d", domainid, server.ID), v, server)

	return
}

// DeleteDomainSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#delete-a-domain-smarthost
func (c *Client) DeleteDomainSmartHost(domainid int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
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

	err = c.delete(fmt.Sprintf("domains/smarthosts/%d/%d", domainid, server.ID), v)

	return
}
