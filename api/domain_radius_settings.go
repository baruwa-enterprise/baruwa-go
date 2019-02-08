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

// RadiusSettings holds domain radius settings
type RadiusSettings struct {
	ID         int         `json:"id,omitempty" url:"id,omitempty"`
	Secret     string      `json:"secret" url:"secret"`
	Timeout    int         `json:"timeout" url:"timeout"`
	AuthServer *SettingsAS `json:"authserver" url:"authserver"`
}

// GetRadiusSettings returns radius settings
// https://www.baruwa.com/docs/api/#retrieve-radius-settings
func (c *Client) GetRadiusSettings(domainID, serverID, settingsID int) (settings *RadiusSettings, err error) {
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

	settings = &RadiusSettings{}

	err = c.get(fmt.Sprintf("radiussettings/%d/%d/%d", domainID, serverID, settingsID), nil, settings)

	return
}

// CreateRadiusSettings creates radius settings
// https://www.baruwa.com/docs/api/#create-radius-settings
func (c *Client) CreateRadiusSettings(domainID, serverID int, settings *RadiusSettings) (err error) {
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

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("radiussettings/%d/%d", domainID, serverID), v, settings)

	return
}

// UpdateRadiusSettings updates radius settings
// https://www.baruwa.com/docs/api/#update-radius-settings
func (c *Client) UpdateRadiusSettings(domainID, serverID int, settings *RadiusSettings) (err error) {
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

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("radiussettings/%d/%d/%d", domainID, serverID, settings.ID), v, nil)

	return
}

// DeleteRadiusSettings deletes radius settings
// https://www.baruwa.com/docs/api/#delete-radius-settings
func (c *Client) DeleteRadiusSettings(domainID, serverID int, settings *RadiusSettings) (err error) {
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

	if v, err = query.Values(settings); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("radiussettings/%d/%d/%d", domainID, serverID, settings.ID), v)

	return
}
