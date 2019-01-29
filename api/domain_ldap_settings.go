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

// LDAPSettings holds Domain LDAP settings
type LDAPSettings struct {
	ID                int    `json:"id,omitempty"`
	Basedn            string `json:"basedn"`
	NameAttribute     string `json:"nameattribute"`
	EmailAttribute    string `json:"emailattribute"`
	BindDN            string `json:"binddn"`
	BindPw            string `json:"bindpw"`
	UseTLS            bool   `json:"usetls"`
	UseSearch         bool   `json:"usesearch"`
	SearchFilter      string `json:"searchfilter"`
	SearchScope       string `json:"search_scope"`
	EmailSearchFilter string `json:"emailsearchfilter"`
	EmailSearchDcope  string `json:"emailsearch_scope"`
}

// GetLDAPSettings returns a domain LDAP settings
// https://www.baruwa.com/docs/api/#retrieve-ad-ldap-settings
func (c *Client) GetLDAPSettings(domainID, serverID, settingsID int) (settings *LDAPSettings, err error) {
	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	if settingsID <= 0 {
		err = fmt.Errorf("The settingsID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settingsID), settings)

	return
}

// CreateLDAPSettings creates a domain LDAP settings
// https://www.baruwa.com/docs/api/#create-ad-ldap-settings
func (c *Client) CreateLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	if settings == nil {
		err = fmt.Errorf("The settings param cannot be nil")
		return
	}

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("ldapsettings/%d/%d", domainID, serverID), v, settings)

	return
}

// UpdateLDAPSettings updates a domain LDAP settings
// https://www.baruwa.com/docs/api/#update-ad-ldap-settings
func (c *Client) UpdateLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	if settings == nil {
		err = fmt.Errorf("The settings param cannot be nil")
		return
	}

	if settings.ID <= 0 {
		err = fmt.Errorf("The settings.ID param should be > 0")
		return
	}

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settings.ID), v, settings)

	return
}

// DeleteLDAPSettings deletes a domain LDAP settings
// https://www.baruwa.com/docs/api/#delete-ad-ldap-settings
func (c *Client) DeleteLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf("The domainID param should be > 0")
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf("The serverID param should be > 0")
		return
	}

	if settings == nil {
		err = fmt.Errorf("The settings param cannot be nil")
		return
	}

	if settings.ID <= 0 {
		err = fmt.Errorf("The settings.ID param should be > 0")
		return
	}

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settings.ID), v)

	return
}
