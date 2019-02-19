// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMyTime(t *testing.T) {
	mt := MyTime{}
	err := json.Unmarshal([]byte(""), &mt)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	b, err := mt.MarshalJSON()
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	if string(b) != "null" {
		t.Errorf("Expected '%s' got '%s'", "null", b)
	}
	now := time.Now()
	mt = MyTime{
		now,
	}
	_, err = mt.MarshalJSON()
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
}

func TestExpirationTime(t *testing.T) {
	var b []byte
	var et expirationTime

	if err := et.UnmarshalJSON(b); err == nil {
		t.Fatalf("An error should be returned")
	}
	if err := et.UnmarshalJSON([]byte("18446744073709551616")); err == nil {
		t.Fatalf("An error should be returned")
	}
	if err := et.UnmarshalJSON([]byte("200")); err != nil {
		t.Fatalf("An error should not be returned")
	}
}

func TestLocalFloat64(t *testing.T) {
	var err error
	var lf LocalFloat64

	if s := lf.String(); s != "0.0" {
		t.Errorf("Expected %s got %s", "0.0", s)
	}
	if err = lf.Set("xxxx"); err == nil {
		t.Fatalf("An error should be returned")
	}
	if s := lf.String(); s != "0.0" {
		t.Errorf("Expected %s got %s", "0.0", s)
	}
	if err = lf.Set("0.124"); err != nil {
		t.Fatalf("An error should not be returned")
	}
	if s := lf.String(); s != "0.1" {
		t.Errorf("Expected %s got %s", "0.1", s)
	}
}
