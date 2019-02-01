// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api Golang bindings for Baruwa REST API
package api

// Meta holds meta
type Meta struct {
	Total int `json:"total"`
}

// Links holds links
type Links struct {
	Pages Pages `json:"pages"`
}

// Pages holds pages
type Pages struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Previous string `json:"prev"`
	Next     string `json:"next"`
}

// ListOptions holds list options
type ListOptions struct {
	Page string
}
