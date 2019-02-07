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

// UserDeliveryServer holds user delivery servers
type UserDeliveryServer struct {
	ID               int          `json:"id,omitempty"`
	Address          string       `json:"address"`
	Protocol         int          `json:"protocol"`
	Port             int          `json:"port"`
	RequireTLS       bool         `json:"require_tls"`
	VerificationOnly bool         `json:"verification_only"`
	Enabled          bool         `json:"enabled"`
	Domain           *AliasDomain `json:"domain,omitempty"`
}

// UserDeliveryServerList holds user delivery servers
type UserDeliveryServerList struct {
	Items []UserDeliveryServer `json:"items"`
	Links Links                `json:"links"`
	Meta  Meta                 `json:"meta"`
}

// GetUserDeliveryServers returns a UserDeliveryServerList object
// This contains a paginated list of domain delivery servers and links
// to the neighbouring pages.
// https://www.baruwa.com/docs/api/#listing-user-delivery-servers
func (c *Client) GetUserDeliveryServers(domainID int, opts *ListOptions) (l *UserDeliveryServerList, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	l = &UserDeliveryServerList{}

	err = c.get(fmt.Sprintf("userdeliveryservers/%d", domainID), opts, l)

	return
}

// GetUserDeliveryServer returns a user delivery server
// https://www.baruwa.com/docs/api/#retrieve-a-user-delivery-server
func (c *Client) GetUserDeliveryServer(domainID, serverID int) (server *UserDeliveryServer, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &UserDeliveryServer{}

	err = c.get(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateUserDeliveryServer creates a user delivery server
// https://www.baruwa.com/docs/api/#create-a-user-delivery-server
func (c *Client) CreateUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
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

	err = c.post(fmt.Sprintf("userdeliveryservers/%d", domainID), v, server)

	return
}

// UpdateUserDeliveryServer updates a user delivery server
// https://www.baruwa.com/docs/api/#update-a-user-delivery-server
func (c *Client) UpdateUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
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

	err = c.put(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, server.ID), v, server)

	return
}

// DeleteUserDeliveryServer deletes a user delivery server
// https://www.baruwa.com/docs/api/#delete-a-user-delivery-server
func (c *Client) DeleteUserDeliveryServer(domainID int, server *UserDeliveryServer) (err error) {
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

	err = c.delete(fmt.Sprintf("userdeliveryservers/%d/%d", domainID, server.ID), v)

	return
}
