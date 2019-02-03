// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAliasAddressError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	alias, err := client.GetAliasAddress(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if alias != nil {
		t.Errorf("Expected %v got %v", nil, alias)
	}
}

func TestGetAliasAddressOK(t *testing.T) {
	aliasID := 3
	data := fmt.Sprintf(`
	{
		"enabled": false,
		"id": %d,
		"address": "info@example.com"
	}
	`, aliasID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	alias, err := client.GetAliasAddress(aliasID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if alias.ID != aliasID {
		t.Errorf("Expected %d got %d", aliasID, alias.ID)
	}
}

func TestCreateAliasAddressError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateAliasAddress(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
	err = client.CreateAliasAddress(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
}

func TestCreateAliasAddressOK(t *testing.T) {
	aliasID := 3
	data := fmt.Sprintf(`
	{
		"enabled": false,
		"id": %d,
		"address": "info@example.com"
	}
	`, aliasID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &AliasAddress{
		Address: "a@example.com",
		Enabled: false,
	}
	err = client.CreateAliasAddress(1, a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if a.ID != aliasID {
		t.Errorf("Expected %d got %d", aliasID, a.ID)
	}
}

func TestUpdateAliasAddressError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateAliasAddress(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
	a := &AliasAddress{
		Address: "a@example.com",
		Enabled: false,
	}
	err = client.UpdateAliasAddress(a)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasSIDError {
		t.Errorf("Expected '%s' got '%s'", aliasSIDError, err)
	}
}

func TestUpdateAliasAddressOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &AliasAddress{
		ID:      2,
		Address: "a@example.com",
		Enabled: true,
	}
	err = client.UpdateAliasAddress(a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}

func TestDeleteAliasAddressError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteAliasAddress(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
	a := &AliasAddress{
		Address: "a@example.com",
		Enabled: false,
	}
	err = client.DeleteAliasAddress(a)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasSIDError {
		t.Errorf("Expected '%s' got '%s'", aliasSIDError, err)
	}
}

func TestDeleteAliasAddressOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &AliasAddress{
		ID:      2,
		Address: "a@example.com",
		Enabled: true,
	}
	err = client.DeleteAliasAddress(a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
