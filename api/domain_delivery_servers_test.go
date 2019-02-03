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

func TestGetDomainDeliveryServersError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	l, err := client.GetDomainDeliveryServers(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if l != nil {
		t.Errorf("Expected %v got %v", nil, l)
	}
}

func TestGetDomainDeliveryServersOK(t *testing.T) {
	data := `
	{
		"items": [{
			"domain": {
				"name": "example.com",
				"id": 2
			},
			"protocol": 1,
			"enabled": true,
			"verification_only": false,
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
	s, err := client.GetDomainDeliveryServers(1, nil)
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

func TestGetDomainDeliveryServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetDomainDeliveryServer(0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if as != nil {
		t.Errorf("Expected %v got %v", nil, as)
	}
	as, err = client.GetDomainDeliveryServer(1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
}

func TestGetDomainDeliveryServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"domain": {
			"name": "example.com",
			"id": 2
		},
		"protocol": 1,
		"enabled": true,
		"require_tls": false,
		"verification_only": false,
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
	as, err := client.GetDomainDeliveryServer(domainID, serverID)
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

func TestCreateDomainDeliveryServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateDomainDeliveryServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateDomainDeliveryServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateDomainDeliveryServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"domain": {
			"name": "example.com",
			"id": %d
		},
		"protocol": 1,
		"enabled": true,
		"require_tls": false,
		"verification_only": false,
		"id": %d,
		"address": "192.168.1.151",
		"port": 25
	}
	`, domainID, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainDeliveryServer{
		Address:          "192.168.1.151",
		Protocol:         1,
		Enabled:          true,
		RequireTLS:       false,
		VerificationOnly: false,
		Port:             25,
		Domain:           &AliasDomain{Name: "example.com", ID: domainID},
	}
	err = client.CreateDomainDeliveryServer(domainID, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
	if !ds.Enabled {
		t.Errorf("Expected %t got %t", true, ds.Enabled)
	}
}

func TestUpdateDomainDeliveryServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateDomainDeliveryServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateDomainDeliveryServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &DomainDeliveryServer{
		Address:          "192.168.1.151",
		Protocol:         1,
		Enabled:          true,
		RequireTLS:       false,
		VerificationOnly: false,
		Port:             25,
	}
	err = client.UpdateDomainDeliveryServer(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateDomainDeliveryServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := fmt.Sprintf(`
	{
		"domain": {
			"name": "example.com",
			"id": %d
		},
		"protocol": 1,
		"enabled": false,
		"require_tls": false,
		"verification_only": false,
		"id": %d,
		"address": "192.168.1.151",
		"port": 25
	}
	`, domainID, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainDeliveryServer{
		ID:               serverID,
		Address:          "192.168.1.151",
		Protocol:         1,
		Enabled:          false,
		RequireTLS:       false,
		VerificationOnly: false,
		Port:             25,
	}
	err = client.UpdateDomainDeliveryServer(domainID, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
	if ds.Enabled {
		t.Errorf("Expected %t got %t", false, ds.Enabled)
	}
}

func TestDeleteDomainDeliveryServerError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteDomainDeliveryServer(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteDomainDeliveryServer(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &DomainDeliveryServer{
		Address:          "192.168.1.151",
		Protocol:         1,
		Enabled:          true,
		RequireTLS:       false,
		VerificationOnly: false,
		Port:             25,
	}
	err = client.DeleteDomainDeliveryServer(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteDomainDeliveryServerOK(t *testing.T) {
	domainID := 2
	serverID := 4
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainDeliveryServer{
		ID:               serverID,
		Address:          "192.168.1.151",
		Protocol:         1,
		Enabled:          false,
		RequireTLS:       false,
		VerificationOnly: false,
		Port:             25,
	}
	err = client.DeleteDomainDeliveryServer(domainID, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
