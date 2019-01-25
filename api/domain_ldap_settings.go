// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

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

// GetLDAPSettings returns a domain smarthost
// https://www.baruwa.com/docs/api/#retrieve-ad-ldap-settings
func (c *Client) GetLDAPSettings(id int) (server *LDAPSettings, err error) {
	return
}

// CreateLDAPSettings creates a domain smarthost
// https://www.baruwa.com/docs/api/#create-ad-ldap-settings
func (c *Client) CreateLDAPSettings(server *LDAPSettings) (err error) {
	return
}

// UpdateLDAPSettings updates a domain smarthost
// https://www.baruwa.com/docs/api/#update-ad-ldap-settings
func (c *Client) UpdateLDAPSettings(server *LDAPSettings) (err error) {
	return
}

// DeleteLDAPSettings deletes a domain smarthost
// https://www.baruwa.com/docs/api/#delete-ad-ldap-settings
func (c *Client) DeleteLDAPSettings(id int) (err error) {
	return
}
