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

func TestGetRadiusSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	as, err := client.GetRadiusSettings(0, 1, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if as != nil {
		t.Errorf("Expected %v got %v", nil, as)
	}
	_, err = client.GetRadiusSettings(1, 0, 1)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	_, err = client.GetRadiusSettings(1, 1, 0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsIDError {
		t.Errorf("Expected '%s' got '%s'", settingsIDError, err)
	}
}

func TestGetRadiusSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"authserver": {
			"id": %d
		},
		"id": %d,
		"timeout": 30
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs, err := client.GetRadiusSettings(domainID, serverID, settingsID)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if rs.ID != settingsID {
		t.Errorf("Expected %d got %d", settingsID, rs.ID)
	}
	if rs.AuthServer.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, rs.AuthServer.ID)
	}
}

func TestCreateRadiusSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateRadiusSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.CreateRadiusSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.CreateRadiusSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
}

func TestCreateRadiusSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"authserver": {
			"id": %d
		},
		"id": %d,
		"timeout": 30
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &RadiusSettings{
		Secret:     "secret",
		Timeout:    30,
		AuthServer: &SettingsAS{ID: serverID},
	}
	err = client.CreateRadiusSettings(domainID, serverID, ls)
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

func TestUpdateRadiusSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateRadiusSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.UpdateRadiusSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.UpdateRadiusSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
	ls := &RadiusSettings{
		Secret:  "secret",
		Timeout: 30,
	}
	err = client.UpdateRadiusSettings(1, 1, ls)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsSIDError {
		t.Errorf("Expected '%s' got '%s'", settingsSIDError, err)
	}
}

func TestUpdateRadiusSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := fmt.Sprintf(`
	{
		"authserver": {
			"id": %d
		},
		"id": %d,
		"timeout": 30
	}
	`, serverID, settingsID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &RadiusSettings{
		ID:         settingsID,
		Secret:     "secret",
		Timeout:    30,
		AuthServer: &SettingsAS{ID: serverID},
	}
	err = client.UpdateRadiusSettings(domainID, serverID, ls)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}

func TestDeleteRadiusSettingsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteRadiusSettings(0, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	err = client.DeleteRadiusSettings(1, 0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverIDError {
		t.Errorf("Expected '%s' got '%s'", serverIDError, err)
	}
	err = client.DeleteRadiusSettings(1, 1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsParamError {
		t.Errorf("Expected '%s' got '%s'", settingsParamError, err)
	}
	ls := &RadiusSettings{
		Secret:  "secret",
		Timeout: 30,
	}
	err = client.DeleteRadiusSettings(1, 1, ls)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != settingsSIDError {
		t.Errorf("Expected '%s' got '%s'", settingsSIDError, err)
	}
}

func TestDeleteRadiusSettingsOK(t *testing.T) {
	domainID := 2
	serverID := 4
	settingsID := 2
	data := ``
	server, client, err := getTestServerAndClient(http.StatusNoContent, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ls := &RadiusSettings{
		ID:         settingsID,
		Secret:     "secret",
		Timeout:    30,
		AuthServer: &SettingsAS{ID: serverID},
	}
	err = client.DeleteRadiusSettings(domainID, serverID, ls)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}
