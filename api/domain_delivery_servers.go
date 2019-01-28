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

// DomainDeliveryServer holds domain delivery servers
type DomainDeliveryServer struct {
	ID               int    `json:"id,omitempty"`
	Address          string `json:"address"`
	Protocol         int    `json:"protocol"`
	Port             int    `json:"port"`
	RequireTLS       bool   `json:"require_tls"`
	VerificationOnly bool   `json:"verification_only"`
	Enabled          bool   `json:"enabled"`
}

// GetDomainDeliveryServer returns a domain delivery server
// https://www.baruwa.com/docs/api/#retrieve-a-delivery-server
func (c *Client) GetDomainDeliveryServer(domainid, serverid int) (server *DomainDeliveryServer, err error) {
	if domainid <= 0 {
		err = fmt.Errorf("The domainid param should be > 0")
		return
	}

	if serverid <= 0 {
		err = fmt.Errorf("The serverid param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("deliveryservers/%d/%d", domainid, serverid), server)

	return
}

// CreateDomainDeliveryServer creates a domain delivery server
// https://www.baruwa.com/docs/api/#create-a-delivery-server
func (c *Client) CreateDomainDeliveryServer(domainid int, server *DomainDeliveryServer) (err error) {
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

	err = c.post(fmt.Sprintf("deliveryservers/%d", domainid), v, server)

	return
}

// UpdateDomainDeliveryServer updates a domain delivery server
// https://www.baruwa.com/docs/api/#update-a-delivery-server
func (c *Client) UpdateDomainDeliveryServer(domainid int, server *DomainDeliveryServer) (err error) {
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

	err = c.post(fmt.Sprintf("deliveryservers/%d/%d", domainid, server.ID), v, server)

	return
}

// DeleteDomainDeliveryServer deletes a domain delivery server
// https://www.baruwa.com/docs/api/#delete-a-delivery-server
func (c *Client) DeleteDomainDeliveryServer(domainid int, server *DomainDeliveryServer) (err error) {
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

	err = c.delete(fmt.Sprintf("deliveryservers/%d/%d", domainid, server.ID), v)

	return
}
