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

// LDAPSettings holds Domain LDAP settings
type LDAPSettings struct {
	ID                int        `json:"id,omitempty" url:"id,omitempty"`
	Basedn            string     `json:"basedn" url:"basedn"`
	NameAttribute     string     `json:"nameattribute" url:"nameattribute"`
	EmailAttribute    string     `json:"emailattribute" url:"emailattribute"`
	BindDN            string     `json:"binddn" url:"binddn"`
	BindPw            string     `json:"bindpw,omitempty" url:"bindpw,omitempty"`
	UseTLS            bool       `json:"usetls" url:"usetls"`
	UseSearch         bool       `json:"usesearch" url:"usesearch"`
	SearchFilter      string     `json:"searchfilter" url:"searchfilter"`
	SearchScope       string     `json:"search_scope" url:"search_scope"`
	EmailSearchFilter string     `json:"emailsearchfilter" url:"emailsearchfilter"`
	EmailSearchScope  string     `json:"emailsearch_scope" url:"emailsearch_scope"`
	AuthServer        SettingsAS `json:"authserver,omitempty" url:"authserver,omitempty"`
}

// GetLDAPSettings returns a domain LDAP settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-ad-ldap-settings
func (c *Client) GetLDAPSettings(domainID, serverID, settingsID int) (settings *LDAPSettings, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	if settingsID <= 0 {
		err = fmt.Errorf(settingsIDError)
		return
	}

	settings = &LDAPSettings{}

	err = c.get(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settingsID), nil, settings)

	return
}

// CreateLDAPSettings creates a domain LDAP settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-ad-ldap-settings
func (c *Client) CreateLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	if settings == nil {
		err = fmt.Errorf(settingsParamError)
		return
	}

	v, _ = query.Values(settings)

	err = c.post(fmt.Sprintf("ldapsettings/%d/%d", domainID, serverID), v, settings)

	return
}

// UpdateLDAPSettings updates a domain LDAP settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-ad-ldap-settings
func (c *Client) UpdateLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	if settings == nil {
		err = fmt.Errorf(settingsParamError)
		return
	}

	if settings.ID <= 0 {
		err = fmt.Errorf(settingsSIDError)
		return
	}

	v, _ = query.Values(settings)

	err = c.put(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settings.ID), v, nil)

	return
}

// DeleteLDAPSettings deletes a domain LDAP settings
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-ad-ldap-settings
func (c *Client) DeleteLDAPSettings(domainID, serverID int, settings *LDAPSettings) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	if settings == nil {
		err = fmt.Errorf(settingsParamError)
		return
	}

	if settings.ID <= 0 {
		err = fmt.Errorf(settingsSIDError)
		return
	}

	v, _ = query.Values(settings)

	err = c.delete(fmt.Sprintf("ldapsettings/%d/%d/%d", domainID, serverID, settings.ID), v)

	return
}
