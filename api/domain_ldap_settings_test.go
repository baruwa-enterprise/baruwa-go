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

func TestGetLDAPSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetLDAPSettings(0, 1, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if as != nil {
		t.Errorf("Expected %v got %v", nil, as)
	}
	_, err = client.GetLDAPSettings(1, 0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	_, err = client.GetLDAPSettings(1, 1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsIDError {
		t.Errorf("Expected '%s' got '%s'", settingsIDError, err)
	}
}

func TestGetLDAPSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"binddn": "uid=readonly-admin,ou=Users,dc=example,dc=com",
		"emailsearchfilter": "",
		"emailsearch_scope": "subtree",
		"searchfilter": "",
		"search_scope": "subtree",
		"authserver": {
			"id": %d
		},
		"basedn": "ou=Users,dc=example,dc=com",
		"usetls": true,
		"usesearch": false,
		"emailattribute": "mail",
		"id": %d,
		"nameattribute": "uid"
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls, err := client.GetLDAPSettings(domainID, serverID, settingsID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ls.ID != settingsID {
		t.Errorf("Expected %d got %d", settingsID, ls.ID)
	}
	if ls.AuthServer.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ls.AuthServer.ID)
	}
}

func TestCreateLDAPSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateLDAPSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateLDAPSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.CreateLDAPSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
}

func TestCreateLDAPSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"binddn": "uid=readonly-admin,ou=Users,dc=example,dc=com",
		"emailsearchfilter": "",
		"emailsearch_scope": "subtree",
		"searchfilter": "",
		"search_scope": "subtree",
		"authserver": {
			"id": %d
		},
		"basedn": "ou=Users,dc=example,dc=com",
		"usetls": true,
		"usesearch": false,
		"emailattribute": "mail",
		"id": %d,
		"nameattribute": "uid"
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &LDAPSettings{
		BindDN:           "uid=readonly-admin,ou=Users,dc=example,dc=com",
		EmailSearchScope: "subtree",
		SearchScope:      "subtree",
		Basedn:           "ou=Users,dc=example,dc=com",
		UseTLS:           true,
		UseSearch:        true,
		EmailAttribute:   "mail",
		NameAttribute:    "uid",
	}
	err = client.CreateLDAPSettings(domainID, serverID, ls)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ls.ID != settingsID {
		t.Errorf("Expected %d got %d", settingsID, ls.ID)
	}
	if ls.AuthServer.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, ls.AuthServer.ID)
	}
}

func TestUpdateLDAPSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateLDAPSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateLDAPSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.UpdateLDAPSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
	ls := &LDAPSettings{
		BindDN:           "uid=readonly-admin,ou=Users,dc=example,dc=com",
		EmailSearchScope: "subtree",
		SearchScope:      "subtree",
		Basedn:           "ou=Users,dc=example,dc=com",
		UseTLS:           true,
		UseSearch:        true,
		EmailAttribute:   "mail",
		NameAttribute:    "uid",
	}
	err = client.UpdateLDAPSettings(1, 1, ls)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsSIDError {
		t.Errorf("Expected '%s' got '%s'", settingsSIDError, err)
	}
}

func TestUpdateLDAPSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"binddn": "uid=readonly-admin,ou=Users,dc=example,dc=com",
		"emailsearchfilter": "",
		"emailsearch_scope": "subtree",
		"searchfilter": "",
		"search_scope": "subtree",
		"authserver": {
			"id": %d
		},
		"basedn": "ou=Users,dc=example,dc=com",
		"usetls": false,
		"usesearch": false,
		"emailattribute": "mail",
		"id": %d,
		"nameattribute": "uid"
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &LDAPSettings{
		ID:               settingsID,
		BindDN:           "uid=readonly-admin,ou=Users,dc=example,dc=com",
		EmailSearchScope: "subtree",
		SearchScope:      "subtree",
		Basedn:           "ou=Users,dc=example,dc=com",
		UseTLS:           false,
		UseSearch:        true,
		EmailAttribute:   "mail",
		NameAttribute:    "uid",
	}
	err = client.UpdateLDAPSettings(domainID, serverID, ls)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if ls.UseTLS {
		t.Errorf("Expected %t got %t", false, ls.UseTLS)
	}
}

func TestDeleteLDAPSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteLDAPSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteLDAPSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.DeleteLDAPSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
	ls := &LDAPSettings{
		BindDN:           "uid=readonly-admin,ou=Users,dc=example,dc=com",
		EmailSearchScope: "subtree",
		SearchScope:      "subtree",
		Basedn:           "ou=Users,dc=example,dc=com",
		UseTLS:           true,
		UseSearch:        true,
		EmailAttribute:   "mail",
		NameAttribute:    "uid",
	}
	err = client.DeleteLDAPSettings(1, 1, ls)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsSIDError {
		t.Errorf("Expected '%s' got '%s'", settingsSIDError, err)
	}
}

func TestDeleteLDAPSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"binddn": "uid=readonly-admin,ou=Users,dc=example,dc=com",
		"emailsearchfilter": "",
		"emailsearch_scope": "subtree",
		"searchfilter": "",
		"search_scope": "subtree",
		"authserver": {
			"id": %d
		},
		"basedn": "ou=Users,dc=example,dc=com",
		"usetls": false,
		"usesearch": false,
		"emailattribute": "mail",
		"id": %d,
		"nameattribute": "uid"
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &LDAPSettings{
		ID:               settingsID,
		BindDN:           "uid=readonly-admin,ou=Users,dc=example,dc=com",
		EmailSearchScope: "subtree",
		SearchScope:      "subtree",
		Basedn:           "ou=Users,dc=example,dc=com",
		UseTLS:           false,
		UseSearch:        true,
		EmailAttribute:   "mail",
		NameAttribute:    "uid",
	}
	err = client.DeleteLDAPSettings(domainID, serverID, ls)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
