// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

// SystemStatus holds system status
type SystemStatus struct {
	Inbound  int   `json:"inbound"`
	Status   bool  `json:"status"`
	Total    []int `json:"total"`
	Outbound int   `json:"outbound"`
}

// GetSystemStatus returns radius settings
// https://www.baruwa.com/docs/api/#retrieve-system-status
func (c *Client) GetSystemStatus() (status *SystemStatus, err error) {
	err = c.get("status", nil, status)

	return
}
