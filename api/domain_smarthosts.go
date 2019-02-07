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
func (c *Client) GetDomainSmartHost(domainID, serverID int) (server *DomainSmartHost, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &DomainSmartHost{}

	err = c.get(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateDomainSmartHost creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-a-domain-smarthost
func (c *Client) CreateDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("domains/smarthosts/%d", domainID), v, server)

	return
}

// UpdateDomainSmartHost updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-a-domain-smarthost
func (c *Client) UpdateDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
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

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, server.ID), v, server)

	return
}

// DeleteDomainSmartHost deletes a domain smarthost
// https://www.baruwa.com/docs/api/#delete-a-domain-smarthost
func (c *Client) DeleteDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
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

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, server.ID), v)

	return
}
