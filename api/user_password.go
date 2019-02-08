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

// PasswordForm sends password update
type PasswordForm struct {
	Password1 string `json:"password1" url:"password1"`
	Password2 string `json:"password2" url:"password2"`
}

// ChangeUserPassword changes a users account password
// https://www.baruwa.com/docs/api/#change-a-password
func (c *Client) ChangeUserPassword(userID int, form *PasswordForm) (err error) {
	var v url.Values

	if form == nil {
		err = fmt.Errorf(pwFormError)
		return
	}

	if userID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	if v, err = query.Values(form); err != nil {
		return
	}

	err = c.post(fmt.Sprintf("users/chpw/%d", userID), v, nil)

	return
}
