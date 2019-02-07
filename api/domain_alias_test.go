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

func TestGetDomainAliasError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	alias, err := client.GetDomainAlias(0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if alias != nil {
		t.Errorf("Expected %v got %v", nil, alias)
	}
	_, err = client.GetDomainAlias(1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasIDError {
		t.Errorf("Expected '%s' got '%s'", aliasIDError, err)
	}
}

func TestGetDomainAliasOK(t *testing.T) {
	domainID := 2
	aliasID := 4
	data := fmt.Sprintf(`
	{
		"status": true,
		"domain": {
			"name": "example.com",
			"id": %d
		},
		"accept_inbound": true,
		"id": %d,
		"name": "example.net"
	}
	`, domainID, aliasID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	alias, err := client.GetDomainAlias(domainID, aliasID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if alias.ID != aliasID {
		t.Errorf("Expected %d got %d", aliasID, alias.ID)
	}
	if !alias.Enabled {
		t.Errorf("Expected %t got %t", true, alias.Enabled)
	}
}

func TestCreateDomainAliasError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateDomainAlias(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateDomainAlias(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
}

func TestCreateDomainAliasOK(t *testing.T) {
	domainID := 2
	aliasID := 4
	data := fmt.Sprintf(`
	{
		"status": true,
		"domain": {
			"name": "example.com",
			"id": %d
		},
		"accept_inbound": true,
		"id": %d,
		"name": "example.net"
	}
	`, domainID, aliasID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &DomainAlias{
		Address:       "example.net",
		Enabled:       true,
		AcceptInbound: true,
		Domain:        &AliasDomain{ID: 2, Name: "example.com"},
	}
	err = client.CreateDomainAlias(2, a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if a.ID != aliasID {
		t.Errorf("Expected %d got %d", aliasID, a.ID)
	}
}

func TestUpdateDomainAliasError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateDomainAlias(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateDomainAlias(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
	a := &DomainAlias{
		Address:       "example.net",
		Enabled:       true,
		AcceptInbound: false,
		Domain:        &AliasDomain{ID: 2, Name: "example.com"},
	}
	err = client.UpdateDomainAlias(1, a)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasSIDError {
		t.Errorf("Expected '%s' got '%s'", aliasSIDError, err)
	}
}

func TestUpdateDomainAliasOK(t *testing.T) {
	domainID := 2
	aliasID := 4
	data := fmt.Sprintf(`
	{
		"status": true,
		"domain": {
			"name": "example.com",
			"id": %d
		},
		"accept_inbound": false,
		"id": %d,
		"name": "example.net"
	}
	`, domainID, aliasID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &DomainAlias{
		ID:            aliasID,
		Address:       "example.net",
		Enabled:       true,
		AcceptInbound: false,
		Domain:        &AliasDomain{ID: domainID, Name: "example.com"},
	}
	err = client.UpdateDomainAlias(2, a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}

func TestDeleteDomainAliasError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteDomainAlias(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteDomainAlias(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasParamError {
		t.Errorf("Expected '%s' got '%s'", aliasParamError, err)
	}
	a := &DomainAlias{
		Address:       "example.net",
		Enabled:       true,
		AcceptInbound: false,
		Domain:        &AliasDomain{ID: 2, Name: "example.com"},
	}
	err = client.DeleteDomainAlias(1, a)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != aliasSIDError {
		t.Errorf("Expected '%s' got '%s'", aliasSIDError, err)
	}
}

func TestDeleteDomainAliasOK(t *testing.T) {
	domainID := 2
	aliasID := 4
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	a := &DomainAlias{
		ID:            aliasID,
		Address:       "example.net",
		Enabled:       true,
		AcceptInbound: false,
		Domain:        &AliasDomain{ID: domainID, Name: "example.com"},
	}
	err = client.DeleteDomainAlias(2, a)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
