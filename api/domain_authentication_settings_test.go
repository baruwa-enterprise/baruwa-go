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

func TestGetAuthServersError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	l, err := client.GetAuthServers(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if l != nil {
		t.Errorf("Expected %v got %v", nil, l)
	}
}

func TestGetAuthServersOK(t *testing.T) {
	data := `
	{
		"items": [{
			"protocol": 2,
			"enabled": true,
			"user_map_template": "example_%(user)s",
			"split_address": true,
			"address": "192.168.1.150",
			"id": 2
		}],
		"meta": {
			"total": 1
		}
	}
	`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	s, err := client.GetAuthServers(1, nil)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if len(s.Items) != 1 {
		t.Errorf("Expected %d got %d", 1, len(s.Items))
	}
	if s.Meta.Total != 1 {
		t.Errorf("Expected %d got %d", 1, s.Meta.Total)
	}
}

func TestGetAuthServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetAuthServer(0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if as != nil {
		t.Errorf("Expected %v got %v", nil, as)
	}
	_, err = client.GetAuthServer(1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
}

func TestGetAuthServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"protocol": 2,
		"enabled": true,
		"user_map_template": "example_%%(user)s",
		"split_address": true,
		"address": "192.168.1.151",
		"id": %d
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetAuthServer(domainID, serverID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if as.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, as.ID)
	}
	if !as.Enabled {
		t.Errorf("Expected %t got %t", true, as.Enabled)
	}
}

func TestCreateAuthServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateAuthServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateAuthServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateAuthServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	umpt := "example_%(user)s"
	data := fmt.Sprintf(`
	{
		"protocol": 2,
		"enabled": true,
		"user_map_template": "%s",
		"split_address": true,
		"address": "192.168.1.151",
		"id": %d
	}
	`, umpt, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as := &AuthServer{
		Address:         "192.168.1.151",
		Protocol:        2,
		Enabled:         true,
		SplitAddress:    true,
		UserMapTemplate: umpt,
	}
	err = client.CreateAuthServer(domainID, as)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if as.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, as.ID)
	}
	if !as.Enabled {
		t.Errorf("Expected %t got %t", true, as.Enabled)
	}
	if as.UserMapTemplate != umpt {
		t.Errorf("Expected '%s' got '%s'", umpt, as.UserMapTemplate)
	}
}

func TestUpdateAuthServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateAuthServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateAuthServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	as := &AuthServer{
		Address:      "192.168.1.151",
		Protocol:     2,
		Enabled:      true,
		SplitAddress: true,
	}
	err = client.UpdateAuthServer(1, as)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateAuthServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	umpt := "example_%(user)s"
	data := fmt.Sprintf(`
	{
		"protocol": 2,
		"enabled": false,
		"user_map_template": "%s",
		"split_address": true,
		"address": "192.168.1.151",
		"id": %d
	}
	`, umpt, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as := &AuthServer{
		ID:              serverID,
		Address:         "192.168.1.151",
		Protocol:        2,
		Enabled:         false,
		SplitAddress:    true,
		UserMapTemplate: umpt,
	}
	err = client.UpdateAuthServer(domainID, as)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if as.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, as.ID)
	}
	if as.Enabled {
		t.Errorf("Expected %t got %t", false, as.Enabled)
	}
}

func TestDeleteAuthServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteAuthServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteAuthServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	as := &AuthServer{
		Address:      "192.168.1.151",
		Protocol:     2,
		Enabled:      true,
		SplitAddress: true,
	}
	err = client.DeleteAuthServer(1, as)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteAuthServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as := &AuthServer{
		ID:           serverID,
		Address:      "192.168.1.151",
		Protocol:     2,
		Enabled:      false,
		SplitAddress: true,
	}
	err = client.DeleteAuthServer(domainID, as)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
