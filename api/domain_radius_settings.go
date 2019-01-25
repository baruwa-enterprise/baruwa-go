// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// RadiusSettings holds domain radius settings
type RadiusSettings struct {
	ID      int    `json:"id,omitempty"`
	Secret  string `json:"secret"`
	Timeout int    `json:"timeout"`
}

// GetRadiusSettings returns radius settings
// https://www.baruwa.com/docs/api/#retrieve-radius-settings
func (c *Client) GetRadiusSettings(id int) (server *RadiusSettings, err error) {
	return
}

// CreateRadiusSettings creates radius settings
// https://www.baruwa.com/docs/api/#create-radius-settings
func (c *Client) CreateRadiusSettings(server *RadiusSettings) (err error) {
	return
}

// UpdateRadiusSettings updates radius settings
// https://www.baruwa.com/docs/api/#update-radius-settings
func (c *Client) UpdateRadiusSettings(server *RadiusSettings) (err error) {
	return
}

// DeleteRadiusSettings deletes radius settings
// https://www.baruwa.com/docs/api/#delete-radius-settings
func (c *Client) DeleteRadiusSettings(id int) (err error) {
	return
}
