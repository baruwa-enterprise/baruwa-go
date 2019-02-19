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

func TestGetOrganizationsOK(t *testing.T) {
	data := `{
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
		}
	}
`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetOrganizations(nil)
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
}

func TestGetOrganizationError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetOrganization(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	if ds != nil {
		t.Errorf("Expected %v got %v", nil, ds)
	}
}

func TestGetOrganizationOK(t *testing.T) {
	organizationID := 2
	data := fmt.Sprintf(`
	{
		"domains": [{
			"name": "example.com",
			"id": 2
		}, {
			"name": "example.net",
			"id": 4
		}],
		"name": "My Org",
		"id": %d
	}
	`, organizationID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	o, err := client.GetOrganization(organizationID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if o.ID != organizationID {
		t.Errorf("Expected %d got %d", organizationID, o.ID)
	}
}

func TestCreateOrganizationError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	_, err = client.CreateOrganization(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != formParamError {
		t.Errorf("Expected '%s' got '%s'", formParamError, err)
	}
}

func TestCreateOrganizationOK(t *testing.T) {
	organizationID := 2
	data := fmt.Sprintf(`
	{
		"domains": [{
			"name": "example.com",
			"id": 2
		}, {
			"name": "example.net",
			"id": 4
		}],
		"name": "My Org",
		"id": %d
	}
	`, organizationID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	f := &OrganizationForm{
		Name:    "My Org",
		Domains: []int{2, 4},
	}
	o, err := client.CreateOrganization(f)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if o.ID != organizationID {
		t.Errorf("Expected %d got %d", organizationID, o.ID)
	}
}

func TestUpdateOrganizationError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateOrganization(nil, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != formParamError {
		t.Errorf("Expected '%s' got '%s'", formParamError, err)
	}
	f := &OrganizationForm{
		Name:    "My Org",
		Domains: []int{2, 4},
	}
	err = client.UpdateOrganization(f, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != formSIDError {
		t.Errorf("Expected '%s' got '%s'", formSIDError, err)
	}
	f.ID = 1
	err = client.UpdateOrganization(f, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != orgParamError {
		t.Errorf("Expected '%s' got '%s'", orgParamError, err)
	}
	org := &Organization{
		Name: "My Org",
	}
	err = client.UpdateOrganization(f, org)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != orgSIDError {
		t.Errorf("Expected '%s' got '%s'", orgSIDError, err)
	}
}

func TestUpdateOrganizationOK(t *testing.T) {
	organizationID := 2
	data := fmt.Sprintf(`
	{
		"domains": [{
			"name": "example.com",
			"id": 2
		}, {
			"name": "example.net",
			"id": 4
		}],
		"name": "My Org",
		"id": %d
	}
	`, organizationID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	f := &OrganizationForm{
		ID:      organizationID,
		Name:    "My Org",
		Domains: []int{2, 4},
	}
	o := &Organization{
		ID: organizationID,
	}
	err = client.UpdateOrganization(f, o)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}

func TestDeleteOrganizationError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteOrganization(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
}

func TestDeleteOrganizationOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteOrganization(1)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
