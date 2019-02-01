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

// UserList holds users
type UserList struct {
	Items []User `json:"items"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

// GetUsers returns a UserList object
// This contains a paginated list of user accounts and links
// to the neigbouring pages.
// https://www.baruwa.com/docs/api/#list-all-accounts
func (c *Client) GetUsers(opts *ListOptions) (l *UserList, err error) {
	err = c.get("users", nil, l)

	return
}

// GetUser returns a user account
// https://www.baruwa.com/docs/api/#retrieve-an-existing-account
func (c *Client) GetUser(userID int) (user *User, err error) {
	if userID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	err = c.get(fmt.Sprintf("users/%d", userID), nil, user)

	return
}

// CreateUser creates a user account
// https://www.baruwa.com/docs/api/#create-a-new-account
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
// https://www.baruwa.com/docs/api/#update-an-account
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
// https://www.baruwa.com/docs/api/#delete-an-account
func (c *Client) DeleteUser(userID int) (err error) {
	if userID <= 0 {
		err = fmt.Errorf("The userID param should be > 0")
		return
	}

	err = c.delete(fmt.Sprintf("users/%d", userID), nil)

	return
}
