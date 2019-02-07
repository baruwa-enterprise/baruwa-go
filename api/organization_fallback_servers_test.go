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

func TestGetFallBackServersError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	l, err := client.GetFallBackServers(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	if l != nil {
		t.Errorf("Expected %v got %v", nil, l)
	}
}

func TestGetFallBackServersOK(t *testing.T) {
	data := `
	{
		"items": [{
			"organization": {
				"name": "Baruwa",
				"id": 2
			},
			"protocol": 1,
			"enabled": true,
			"id": 2,
			"address": "192.168.1.150",
			"port": 25
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
	s, err := client.GetFallBackServers(1, nil)
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

func TestGetFallBackServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetFallBackServer(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	if as != nil {
		t.Errorf("Expected %v got %v", nil, as)
	}
}

func TestGetFallBackServerOK(t *testing.T) {
	serverID := 4
	data := fmt.Sprintf(`
	{
		"organization": {
			"name": "Baruwa",
			"id": 2
		},
		"protocol": 1,
		"enabled": true,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.151",
		"port": 25
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetFallBackServer(serverID)
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

func TestCreateFallBackServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateFallBackServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	err = client.CreateFallBackServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateFallBackServerOK(t *testing.T) {
	organizationID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"organization": {
			"name": "Baruwa",
			"id": %d
		},
		"protocol": 1,
		"enabled": true,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.151",
		"port": 25
	}
	`, organizationID, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	fs := &FallBackServer{
		Address:      "192.168.1.151",
		Protocol:     1,
		Enabled:      true,
		RequireTLS:   false,
		Port:         25,
		Organization: &FallBackServerOrg{Name: "Baruwa", ID: organizationID},
	}
	err = client.CreateFallBackServer(organizationID, fs)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if fs.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, fs.ID)
	}
	if !fs.Enabled {
		t.Errorf("Expected %t got %t", true, fs.Enabled)
	}
}

func TestUpdateFallBackServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateFallBackServer(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	fs := &FallBackServer{
		Address:    "192.168.1.151",
		Protocol:   1,
		Enabled:    true,
		RequireTLS: false,
		Port:       25,
	}
	err = client.UpdateFallBackServer(fs)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateFallBackServerOK(t *testing.T) {
	organizationID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"organization": {
			"name": "Baruwa",
			"id": %d
		},
		"protocol": 1,
		"enabled": false,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.151",
		"port": 25
	}
	`, organizationID, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	fs := &FallBackServer{
		ID:         serverID,
		Address:    "192.168.1.151",
		Protocol:   1,
		Enabled:    false,
		RequireTLS: false,
		Port:       25,
	}
	err = client.UpdateFallBackServer(fs)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if fs.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, fs.ID)
	}
	if fs.Enabled {
		t.Errorf("Expected %t got %t", false, fs.Enabled)
	}
}

func TestDeleteFallBackServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteFallBackServer(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	fs := &FallBackServer{
		Address:    "192.168.1.151",
		Protocol:   1,
		Enabled:    true,
		RequireTLS: false,
		Port:       25,
	}
	err = client.DeleteFallBackServer(fs)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteFallBackServerOK(t *testing.T) {
	serverID := 4
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	fs := &FallBackServer{
		ID:         serverID,
		Address:    "192.168.1.151",
		Protocol:   1,
		Enabled:    false,
		RequireTLS: false,
		Port:       25,
	}
	err = client.DeleteFallBackServer(fs)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
