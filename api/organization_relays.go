// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// RelaySetting holds relay settings
type RelaySetting struct {
	ID              int     `json:"id,omitempty"`
	Address         string  `json:"address"`
	Username        string  `json:"username"`
	Enabled         bool    `json:"enabled"`
	RequireTLS      bool    `json:"require_tls"`
	Password1       string  `json:"password1,omitempty"`
	Password2       string  `json:"password2,omitempty"`
	Description     string  `json:"description"`
	LowScore        float64 `json:"low_score"`
	HighScore       float64 `json:"high_score"`
	SpamActions     int     `json:"spam_actions"`
	HighSpamActions int     `json:"highspam_actions"`
	BlockMacros     bool    `json:"block_macros"`
	RateLimit       int     `json:"ratelimit"`
}

// GetRelaySetting returns radius settings
// https://www.baruwa.com/docs/api/#retrieve-relay-settings
func (c *Client) GetRelaySetting(id int) (server *RelaySetting, err error) {
	return
}

// CreateRelaySetting creates radius settings
// https://www.baruwa.com/docs/api/#create-relay-settings
func (c *Client) CreateRelaySetting(server *RelaySetting) (err error) {
	return
}

// UpdateRelaySetting updates radius settings
// https://www.baruwa.com/docs/api/#update-relay-settings
func (c *Client) UpdateRelaySetting(server *RelaySetting) (err error) {
	return
}

// DeleteRelaySetting deletes radius settings
// https://www.baruwa.com/docs/api/#delete-relay-settings
func (c *Client) DeleteRelaySetting(id int) (err error) {
	return
}
