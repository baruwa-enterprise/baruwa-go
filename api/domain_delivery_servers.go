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

// DomainDeliveryServer holds domain delivery servers
type DomainDeliveryServer struct {
	ID               int          `json:"id,omitempty" url:"id,omitempty"`
	Address          string       `json:"address" url:"address"`
	Protocol         int          `json:"protocol" url:"protocol"`
	Port             int          `json:"port" url:"port"`
	RequireTLS       bool         `json:"require_tls" url:"require_tls"`
	VerificationOnly bool         `json:"verification_only" url:"verification_only"`
	Enabled          bool         `json:"enabled" url:"enabled"`
	Domain           *AliasDomain `json:"domain,omitempty" url:"domain,omitempty"`
}

// DomainDeliveryServerList holds domain delivery servers
type DomainDeliveryServerList struct {
	Items []DomainDeliveryServer `json:"items"`
	Links Links                  `json:"links"`
	Meta  Meta                   `json:"meta"`
}

// GetDomainDeliveryServers returns a DomainDeliveryServerList object
// This contains a paginated list of domain delivery servers and links
// to the neighbouring pages.
// https://www.baruwa.com/docs/api/#listing-delivery-servers
func (c *Client) GetDomainDeliveryServers(domainID int, opts *ListOptions) (l *DomainDeliveryServerList, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	l = &DomainDeliveryServerList{}

	err = c.get(fmt.Sprintf("deliveryservers/%d", domainID), opts, l)

	return
}

// GetDomainDeliveryServer returns a domain delivery server
// https://www.baruwa.com/docs/api/#retrieve-a-delivery-server
func (c *Client) GetDomainDeliveryServer(domainID, serverID int) (server *DomainDeliveryServer, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &DomainDeliveryServer{}

	err = c.get(fmt.Sprintf("deliveryservers/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateDomainDeliveryServer creates a domain delivery server
// https://www.baruwa.com/docs/api/#create-a-delivery-server
func (c *Client) CreateDomainDeliveryServer(domainID int, server *DomainDeliveryServer) (err error) {
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

	err = c.post(fmt.Sprintf("deliveryservers/%d", domainID), v, server)

	return
}

// UpdateDomainDeliveryServer updates a domain delivery server
// https://www.baruwa.com/docs/api/#update-a-delivery-server
func (c *Client) UpdateDomainDeliveryServer(domainID int, server *DomainDeliveryServer) (err error) {
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

	err = c.put(fmt.Sprintf("deliveryservers/%d/%d", domainID, server.ID), v, nil)

	return
}

// DeleteDomainDeliveryServer deletes a domain delivery server
// https://www.baruwa.com/docs/api/#delete-a-delivery-server
func (c *Client) DeleteDomainDeliveryServer(domainID int, server *DomainDeliveryServer) (err error) {
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

	err = c.delete(fmt.Sprintf("deliveryservers/%d/%d", domainID, server.ID), v)

	return
}
