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
	ID   int    `json:"id" url:"id"`
	Name string `json:"name" url:"name"`
}

// UserOrganization holds user organizations
type UserOrganization struct {
	ID   int    `json:"id" url:"id"`
	Name string `json:"name" url:"name"`
}

// UserAddress addresses
type UserAddress struct {
}

// User holds users
type User struct {
	ID            int                `json:"id,omitempty" url:"id,omitempty"`
	Username      string             `json:"username" url:"username"`
	Firstname     string             `json:"firstname" url:"firstname"`
	Lastname      string             `json:"lastname" url:"lastname"`
	Email         string             `json:"email" url:"email"`
	Timezone      string             `json:"timezone" url:"timezone"`
	AccountType   int                `json:"account_type" url:"account_type"`
	Enabled       bool               `json:"active" url:"active"`
	SendReport    bool               `json:"send_report" url:"send_report"`
	SpamChecks    bool               `json:"spam_checks" url:"spam_checks"`
	LowScore      LocalFloat64       `json:"low_score" url:"low_score"`
	HighScore     LocalFloat64       `json:"high_score" url:"high_score"`
	BlockMacros   bool               `json:"block_macros" url:"block_macros"`
	CreatedOn     MyTime             `json:"created_on" url:"created_on"`
	LastLogin     MyTime             `json:"last_login" url:"last_login"`
	Domains       []UserDomain       `json:"domains,omitempty" url:"domains,omitempty"`
	Organizations []UserOrganization `json:"organizations,omitempty" url:"organizations,omitempty"`
}

// UserForm holds users
type UserForm struct {
	ID            *int          `json:"id,omitempty" url:"id,omitempty"`
	Username      *string       `json:"username" url:"username,omitempty"`
	Firstname     *string       `json:"firstname" url:"firstname,omitempty"`
	Lastname      *string       `json:"lastname" url:"lastname,omitempty"`
	Password1     *string       `json:"password1" url:"password1,omitempty"`
	Password2     *string       `json:"password2" url:"password2,omitempty"`
	Email         *string       `json:"email" url:"email,omitempty"`
	Timezone      *string       `json:"timezone" url:"timezone,omitempty"`
	AccountType   *int          `json:"account_type" url:"account_type,omitempty"`
	Enabled       *bool         `json:"active" url:"active,omitempty"`
	SendReport    *bool         `json:"send_report" url:"send_report,omitempty"`
	SpamChecks    *bool         `json:"spam_checks" url:"spam_checks,omitempty"`
	LowScore      *LocalFloat64 `json:"low_score" url:"low_score,omitempty,omitempty"`
	HighScore     *LocalFloat64 `json:"high_score" url:"high_score,omitempty"`
	BlockMacros   *bool         `json:"block_macros" url:"block_macros,omitempty"`
	Domains       []int         `json:"domains,omitempty" url:"domains,omitempty"`
	Organizations []int         `json:"organizations,omitempty" url:"organizations,omitempty"`
}

// UserList holds users
type UserList struct {
	Items []User `json:"items"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

// GetUsers returns a UserList object
// This contains a paginated list of user accounts and links
// to the neighbouring pages.
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#list-all-accounts
func (c *Client) GetUsers(opts *ListOptions) (l *UserList, err error) {
	l = &UserList{}

	err = c.get("users", opts, l)

	return
}

// GetUser returns a user account
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-an-existing-account
func (c *Client) GetUser(userID int) (user *User, err error) {
	if userID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	user = &User{}

	err = c.get(fmt.Sprintf("users/%d", userID), nil, user)

	return
}

// CreateUser creates a user account
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-a-new-account
func (c *Client) CreateUser(user *UserForm) (u *User, err error) {
	var v url.Values

	if user == nil {
		err = fmt.Errorf(userParamError)
		return
	}

	v, _ = query.Values(user)

	u = &User{}

	err = c.post("users", v, u)

	return
}

// UpdateUser updates a user account
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-an-account
func (c *Client) UpdateUser(user *UserForm) (err error) {
	var v url.Values

	if user == nil {
		err = fmt.Errorf(userParamError)
		return
	}

	if user.ID == nil || *user.ID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	v, _ = query.Values(user)

	err = c.put(fmt.Sprintf("users/%d", *user.ID), v, nil)

	return
}

// DeleteUser deletes a user account
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-an-account
func (c *Client) DeleteUser(userID int) (err error) {
	if userID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	err = c.delete(fmt.Sprintf("users/%d", userID), nil)

	return
}
