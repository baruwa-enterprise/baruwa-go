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

func TestGetOrgSmartHostsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	_, err = client.GetOrgSmartHosts(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
}

func TestGetOrgSmartHostsOK(t *testing.T) {
	organizationID := 2
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
				"last": "http://baruwa.example.com/api/v1/organizations/smarthosts/%d?page=2",
				"next": "http://baruwa.example.com/api/v1/organizations/smarthosts/%d?page=2"
			}
		}
	}
`, organizationID, organizationID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetOrgSmartHosts(organizationID, nil)
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
	next := fmt.Sprintf("http://baruwa.example.com/api/v1/organizations/smarthosts/%d?page=2", organizationID)
	if u.Links.Pages.Next != next {
		t.Errorf("Expected '%s' got '%s'", next, u.Links.Pages.Next)
	}
}

func TestGetOrgSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetOrgSmartHost(0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	if ds != nil {
		t.Errorf("Expected %v got %v", nil, ds)
	}
	_, err = client.GetOrgSmartHost(1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
}

func TestGetOrgSmartHostOK(t *testing.T) {
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
	ds, err := client.GetOrgSmartHost(1, serverID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
}

func TestCreateOrgSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateOrgSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	err = client.CreateOrgSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateOrgSmartHostOK(t *testing.T) {
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
	ds := &OrgSmartHost{
		Username:    "andrew",
		Password:    "p4ssw0rd",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.CreateOrgSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ds.ID)
	}
}

func TestUpdateOrgSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateOrgSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	err = client.UpdateOrgSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &OrgSmartHost{
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.UpdateOrgSmartHost(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateOrgSmartHostOK(t *testing.T) {
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
	ds := &OrgSmartHost{
		ID:          serverID,
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     false,
	}
	err = client.UpdateOrgSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ds.Enabled {
		t.Errorf("Expected %t got %t", false, ds.Enabled)
	}
}

func TestDeleteOrgSmartHostError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteOrgSmartHost(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	err = client.DeleteOrgSmartHost(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	ds := &OrgSmartHost{
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     true,
	}
	err = client.DeleteOrgSmartHost(1, ds)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteOrgSmartHostOK(t *testing.T) {
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
	ds := &OrgSmartHost{
		ID:          serverID,
		Username:    "andrew",
		Description: "outbound-archiver",
		Address:     "192.168.1.150",
		Port:        25,
		Enabled:     false,
	}
	err = client.DeleteOrgSmartHost(1, ds)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
