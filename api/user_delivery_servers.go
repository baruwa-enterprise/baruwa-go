// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// UserDeliveryServer holds user delivery servers
type UserDeliveryServer struct {
	DomainDeliveryServer
}

// GetUserDeliveryServer returns a user delivery server
// https://www.baruwa.com/docs/api/#retrieve-a-user-delivery-server
func (c *Client) GetUserDeliveryServer(id int) (server *UserDeliveryServer, err error) {
	return
}

// CreateUserDeliveryServer creates a user delivery server
// https://www.baruwa.com/docs/api/#create-a-user-delivery-server
func (c *Client) CreateUserDeliveryServer(server *UserDeliveryServer) (err error) {
	return
}

// UpdateUserDeliveryServer updates a user delivery server
// https://www.baruwa.com/docs/api/#update-a-user-delivery-server
func (c *Client) UpdateUserDeliveryServer(server *UserDeliveryServer) (err error) {
	return
}

// DeleteUserDeliveryServer deletes a user delivery server
// https://www.baruwa.com/docs/api/#delete-a-user-delivery-server
func (c *Client) DeleteUserDeliveryServer(id int) (err error) {
	return
}
