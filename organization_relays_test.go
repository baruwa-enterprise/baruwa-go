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

func TestGetRelaySettingError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs, err := client.GetRelaySetting(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != relayIDError {
		t.Errorf("Expected '%s' got '%s'", relayIDError, err)
	}
	if rs != nil {
		t.Errorf("Expected %v got %v", nil, rs)
	}
}

func TestGetRelaySettingOK(t *testing.T) {
	serverID := 3
	data := fmt.Sprintf(`
	{
		"username": "outboundsmtp",
		"description": "Backup-outbound-smtp",
		"enabled": true,
		"require_tls": false,
		"spam_actions": 2,
		"low_score": 10.0,
		"high_score": 15.0,
		"address": "192.168.1.20",
		"id": %d,
		"highspam_actions": 3
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs, err := client.GetRelaySetting(serverID)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	if rs.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, rs.ID)
	}
}

func TestCreateRelaySettingError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateRelaySetting(0, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != organizationIDError {
		t.Errorf("Expected '%s' got '%s'", organizationIDError, err)
	}
	err = client.CreateRelaySetting(1, nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
}

func TestCreateRelaySettingOK(t *testing.T) {
	serverID := 3
	data := fmt.Sprintf(`
	{
		"username": "outboundsmtp",
		"description": "Backup-outbound-smtp",
		"enabled": true,
		"require_tls": false,
		"spam_actions": 2,
		"low_score": 10.0,
		"high_score": 15.0,
		"address": "192.168.1.20",
		"id": %d,
		"highspam_actions": 3
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs := &RelaySetting{
		Username:        "outboundsmtp",
		Description:     "Backup-outbound-smtp",
		Enabled:         true,
		SpamActions:     2,
		LowScore:        10.0,
		HighScore:       15.0,
		Address:         "192.168.1.20",
		HighSpamActions: 3,
	}
	err = client.CreateRelaySetting(1, rs)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	if rs.ID != serverID {
		t.Errorf("Expected %d got %d", serverID, rs.ID)
	}
}

func TestUpdateRelaySettingError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateRelaySetting(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	rs := &RelaySetting{
		Username:        "outboundsmtp",
		Description:     "Backup-outbound-smtp",
		Enabled:         true,
		SpamActions:     2,
		LowScore:        10.0,
		HighScore:       15.0,
		Address:         "192.168.1.20",
		HighSpamActions: 3,
	}
	err = client.UpdateRelaySetting(rs)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestUpdateRelaySettingOK(t *testing.T) {
	serverID := 3
	data := fmt.Sprintf(`
	{
		"username": "outboundsmtp",
		"description": "Backup-outbound-smtp",
		"enabled": true,
		"require_tls": false,
		"spam_actions": 2,
		"low_score": 10.0,
		"high_score": 15.0,
		"address": "192.168.1.20",
		"id": %d,
		"highspam_actions": 3
	}
	`, serverID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs := &RelaySetting{
		ID:              serverID,
		Username:        "outboundsmtp",
		Description:     "Backup-outbound-smtp",
		Enabled:         true,
		SpamActions:     2,
		LowScore:        10.0,
		HighScore:       15.0,
		Address:         "192.168.1.20",
		HighSpamActions: 3,
	}
	err = client.UpdateRelaySetting(rs)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
}

func TestDeleteRelaySettingError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteRelaySetting(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverParamError {
		t.Errorf("Expected '%s' got '%s'", serverParamError, err)
	}
	rs := &RelaySetting{
		Username:        "outboundsmtp",
		Description:     "Backup-outbound-smtp",
		Enabled:         true,
		SpamActions:     2,
		LowScore:        10.0,
		HighScore:       15.0,
		Address:         "192.168.1.20",
		HighSpamActions: 3,
	}
	err = client.DeleteRelaySetting(rs)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != serverSIDError {
		t.Errorf("Expected '%s' got '%s'", serverSIDError, err)
	}
}

func TestDeleteRelaySettingOK(t *testing.T) {
	serverID := 3
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	rs := &RelaySetting{
		ID:              serverID,
		Username:        "outboundsmtp",
		Description:     "Backup-outbound-smtp",
		Enabled:         true,
		SpamActions:     2,
		LowScore:        10.0,
		HighScore:       15.0,
		Address:         "192.168.1.20",
		HighSpamActions: 3,
	}
	err = client.DeleteRelaySetting(rs)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
}
