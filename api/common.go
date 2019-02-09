// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var nilTime = (time.Time{}).UnixNano()

// SettingsAS hold the authentication server id
type SettingsAS struct {
	ID int `json:"id"`
}

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

// MyTime custom date formater
type MyTime struct {
	time.Time
}

// UnmarshalJSON unmarshals the custom date
func (mt *MyTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.RFC3339Nano, s)

	if err != nil {
		t, err = time.Parse(timeFmt, s)
	}

	mt.Time = t

	return
}

// MarshalJSON marshals the custom date
func (mt *MyTime) MarshalJSON() ([]byte, error) {
	if mt.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", mt.Time.Format(timeFmt))), nil
}

// LocalFloat64 allow for overriding in cli
type LocalFloat64 float64

// Set is required by cli
func (f *LocalFloat64) Set(v string) (err error) {
	val, err := strconv.ParseFloat(v, 64)
	*f = LocalFloat64(val)
	return
}

func (f *LocalFloat64) String() string {
	if f == nil {
		return "0.0"
	}
	return fmt.Sprintf("%.1f", *f)
}
