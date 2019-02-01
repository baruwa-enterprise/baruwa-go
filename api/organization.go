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
func (c *Client) GetOrganization(organizationID int) (org *Organization, err error) {
	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("organizations/%d", organizationID), nil, org)

	return
}

// CreateOrganization creates an organization
// https://www.baruwa.com/docs/api/#create-an-organization
func (c *Client) CreateOrganization(form *OrganizationForm, org Organization) (err error) {
	var v url.Values

	if form == nil {
		err = fmt.Errorf("The form param cannot be nil")
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	err = c.post("organizations", v, org)

	return
}

// UpdateOrganization updates an organization
// https://www.baruwa.com/docs/api/#update-an-organization
func (c *Client) UpdateOrganization(form *OrganizationForm, org Organization) (err error) {
	var v url.Values

	if form == nil {
		err = fmt.Errorf("The form param cannot be nil")
		return
	}

	if form.ID <= 0 {
		err = fmt.Errorf("The form.ID param should be > 0")
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("organizations/%d", form.ID), v, org)

	return
}

// DeleteOrganization deletes an organization
// https://www.baruwa.com/docs/api/#delete-an-organization
func (c *Client) DeleteOrganization(organizationID int) (err error) {
	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("organizations/%d", organizationID), nil)

	return
}
