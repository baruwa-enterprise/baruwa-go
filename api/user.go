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

// UserDomain holds user domains
type UserDomain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserOrganization holds user organizations
type UserOrganization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserAddress addresses
type UserAddress struct {
}

// User holds users
type User struct {
	ID            int                `json:"id,omitempty"`
	Username      string             `json:"username"`
	Firstname     string             `json:"firstname"`
	Lastname      string             `json:"lastname"`
	Password1     string             `json:"password1,omitempty"`
	Password2     string             `json:"password2,omitempty"`
	Email         string             `json:"email"`
	Timezone      string             `json:"timezone"`
	AccountType   int                `json:"account_type"`
	Active        bool               `json:"active"`
	SendReport    bool               `json:"send_report"`
	SpamChecks    bool               `json:"spam_checks"`
	LowScore      float64            `json:"low_score"`
	HighScore     float64            `json:"high_score"`
	BlockMacros   bool               `json:"block_macros"`
	Domains       []UserDomain       `json:"domains,omitempty"`
	Organizations []UserOrganization `json:"organizations,omitempty"`
}

// GetUser returns a user account
func (c *Client) GetUser(id int) (user *User, err error) {
	if id <= 0 {
		err = fmt.Errorf("The id param should be > 0")
		return
	}

	err = c.get(fmt.Sprintf("users/%d", id), user)
	return
}

// CreateUser creates a user account
func (c *Client) CreateUser(user *User) (err error) {
	var v url.Values

	if user == nil {
		err = fmt.Errorf("The user param cannot be nil")
		return
	}

	if v, err = query.Values(user); err != nil {
		return
	}

	err = c.post("users", v, user)

	return
}

// UpdateUser updates a user account
func (c *Client) UpdateUser(user *User) (err error) {
	var v url.Values

	if user == nil {
		err = fmt.Errorf("The user param cannot be nil")
		return
	}

	if user.ID <= 0 {
		err = fmt.Errorf("The user.ID param should be > 0")
		return
	}

	if v, err = query.Values(user); err != nil {
		return
	}

	err = c.put(fmt.Sprintf("users/%d", user.ID), v, user)

	return
}

// DeleteUser deletes a user account
func (c *Client) DeleteUser(id int) (err error) {
	if id <= 0 {
		err = fmt.Errorf("The id param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("users/%d", id), nil)

	return
}
