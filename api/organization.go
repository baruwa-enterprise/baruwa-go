// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// OrgDomain hold alias domain entries
type OrgDomain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Organization holds organizations
type Organization struct {
	ID      int         `json:"id,omitempty"`
	Name    string      `json:"name"`
	Domains []OrgDomain `json:"domains,omitempty"`
}

// OrganizationForm used for creation and update of organizations
type OrganizationForm struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Domains []int  `json:"domains,omitempty"`
	Admins  []int  `json:"admins,omitempty"`
}

// GetOrganization returns an organization
// https://www.baruwa.com/docs/api/#retrieve-an-existing-organization
func (c *Client) GetOrganization(id int) (org *Organization, err error) {
	return
}

// CreateOrganization creates an organization
// https://www.baruwa.com/docs/api/#create-an-organization
func (c *Client) CreateOrganization(form *OrganizationForm, org Organization) (err error) {
	return
}

// UpdateOrganization updates an organization
// https://www.baruwa.com/docs/api/#update-an-organization
func (c *Client) UpdateOrganization(form *OrganizationForm, org Organization) (err error) {
	return
}

// DeleteOrganization deletes an organization
// https://www.baruwa.com/docs/api/#delete-an-organization
func (c *Client) DeleteOrganization(id int) (err error) {
	return
}
