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
func (c *Client) GetRelaySetting(relayID int) (server *RelaySetting, err error) {
	if relayID <= 0 {
		err = fmt.Errorf("The relayID param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("relays/%d", relayID), nil, server)

	return
}

// CreateRelaySetting creates radius settings
// https://www.baruwa.com/docs/api/#create-relay-settings
func (c *Client) CreateRelaySetting(organizationID int, server *RelaySetting) (err error) {
	var v url.Values

	if organizationID <= 0 {
		err = fmt.Errorf("The organizationID param should be > 0")
		return
	}

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("relays/%d", organizationID), v, server)

	return
}

// UpdateRelaySetting updates radius settings
// https://www.baruwa.com/docs/api/#update-relay-settings
func (c *Client) UpdateRelaySetting(server *RelaySetting) (err error) {
	var v url.Values

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf("The server.ID param should be > 0")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("relays/%d", server.ID), v, server)

	return
}

// DeleteRelaySetting deletes radius settings
// https://www.baruwa.com/docs/api/#delete-relay-settings
func (c *Client) DeleteRelaySetting(server *RelaySetting) (err error) {
	var v url.Values

	if server == nil {
		err = fmt.Errorf("The server param cannot be nil")
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf("The server.ID param should be > 0")
		return
	}

	if v, err = query.Values(server); err != nil {
		return
	}

	err = c.delete(fmt.Sprintf("relays/%d", server.ID), v)

	return
}
