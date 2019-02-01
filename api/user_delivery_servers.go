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

// UserDeliveryServer holds user delivery servers
type UserDeliveryServer struct {
	DomainDeliveryServer
}

// GetUserDeliveryServer returns a user delivery server
// https://www.baruwa.com/docs/api/#retrieve-a-user-delivery-server
func (c *Client) GetUserDeliveryServer(domainID, serverID int) (server *UserDeliveryServer, err error) {
	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateUserDeliveryServer creates a user delivery server
// https://www.baruwa.com/docs/api/#create-a-user-delivery-server
func (c *Client) CreateUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("userdeliveryservers/%d", domainID), v, server)

	return
}

// UpdateUserDeliveryServer updates a user delivery server
// https://www.baruwa.com/docs/api/#update-a-user-delivery-server
func (c *Client) UpdateUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
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

	err = c.put(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, server.ID), v, server)

	return
}

// DeleteUserDeliveryServer deletes a user delivery server
// https://www.baruwa.com/docs/api/#delete-a-user-delivery-server
func (c *Client) DeleteUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
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

	err = c.delete(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, server.ID), v)

	return
}
