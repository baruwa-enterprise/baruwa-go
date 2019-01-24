// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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
func (c *Client) GetDomainDeliveryServer(id int) (domain *DomainDeliveryServer, err error) {
	return
}

// CreateDomainDeliveryServer creates a domain delivery server
// https://www.baruwa.com/docs/api/#create-a-delivery-server
func (c *Client) CreateDomainDeliveryServer(domain *DomainDeliveryServer) (err error) {
	return
}

// UpdateDomainDeliveryServer updates a domain delivery server
// https://www.baruwa.com/docs/api/#update-a-delivery-server
func (c *Client) UpdateDomainDeliveryServer(domain *DomainDeliveryServer) (err error) {
	return
}

// DeleteDomainDeliveryServer deletes a domain delivery server
// https://www.baruwa.com/docs/api/#delete-a-delivery-server
func (c *Client) DeleteDomainDeliveryServer(id int) (err error) {
	return
}
