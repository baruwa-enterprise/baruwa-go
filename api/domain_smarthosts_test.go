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

func TestGetDomainSmartHostsOK(t *testing.T) {
	domainID := 2
	data := fmt.Sprintf(`{
		"items": [{
			"enabled": true,
			"require_tls": false,
			"id": 2,
			"address": "192.168.1.150",
			"username": "andrew",
			"description": "outbound-archiver",
			"port": 25
		},
		{
			"enabled": true,
			"require_tls": false,
			"id": 2,
			"address": "192.168.2.150",
			"username": "andrew",
			"description": "outbound-archiver2",
			"port": 25
		}],
		"meta": {
			"total": 2
		},
		"links": {
			"pages": {
				"last": "http://baruwa.example.com/api/v1/domains/smarthosts/%d?page=2",
				"next": "http://baruwa.example.com/api/v1/domains/smarthosts/%d?page=2"
			}
		}
	}
`, domainID, domainID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetDomainSmartHosts(domainID, nil)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err.Error())
	}
	if len(u.Items) != 2 {
		t.Errorf("Expected %d got %d", 2, len(u.Items))
	}
	if u.Meta.Total != 2 {
		t.Errorf("Expected %d got %d", 2, u.Meta.Total)
	}
	if u.Links.Pages.First != "" {
		t.Errorf("Expected '' got '%s'", u.Links.Pages.First)
	}
	next := fmt.Sprintf("http://baruwa.example.com/api/v1/domains/smarthosts/%d?page=2", domainID)
	if u.Links.Pages.Next != next {
		t.Errorf("Expected '%s' got '%s'", next, u.Links.Pages.Next)
	}
}

func TestGetDomainSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetDomainSmartHost(0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if ds != nil {
		t.Errorf("Expected %v got %v", nil, ds)
	}
	_, err = client.GetDomainSmartHost(1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
}

func TestGetDomainSmartHostOK(t *testing.T) {
	serverID := 2
	data := fmt.Sprintf(`
	{
		"enabled": true,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.150",
		"username": "andrew",
		"description": "outbound-archiver",
		"port": 25
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetDomainSmartHost(1, serverID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
}

func TestCreateDomainSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateDomainSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateDomainSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateDomainSmartHostOK(t *testing.T) {
	serverID := 2
	data := fmt.Sprintf(`
	{
		"enabled": true,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.150",
		"username": "andrew",
		"description": "outbound-archiver",
		"port": 25
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainSmartHost{
		Username:    "andrew",
		Password:    "p4ssw0rd",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.CreateDomainSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
}

func TestUpdateDomainSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateDomainSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateDomainSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &DomainSmartHost{
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.UpdateDomainSmartHost(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateDomainSmartHostOK(t *testing.T) {
	serverID := 2
	data := fmt.Sprintf(`
	{
		"enabled": false,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.150",
		"username": "andrew",
		"description": "outbound-archiver",
		"port": 25
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainSmartHost{
		ID:          serverID,
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     false,
	}
	err = client.UpdateDomainSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.Enabled {
		t.Errorf("Expected %t got %t", false, ds.Enabled)
	}
}

func TestDeleteDomainSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteDomainSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteDomainSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &DomainSmartHost{
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.DeleteDomainSmartHost(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteDomainSmartHostOK(t *testing.T) {
	serverID := 2
	data := fmt.Sprintf(`
	{
		"enabled": false,
		"require_tls": false,
		"id": %d,
		"address": "192.168.1.150",
		"username": "andrew",
		"description": "outbound-archiver",
		"port": 25
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds := &DomainSmartHost{
		ID:          serverID,
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     false,
	}
	err = client.DeleteDomainSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
