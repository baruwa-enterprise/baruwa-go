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

func TestGetDomainsError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusInternalServerError, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	_, err = client.GetDomains(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
}

func TestGetDomainsOK(t *testing.T) {
	data := `{
		"items": [{
			"signatures": [{
				"type": 1,
				"id": 1
			}],
			"highspam_actions": 2,
			"delivery_mode": 1,
			"virus_checks": true,
			"ldap_callout": false,
			"dkimkeys": [],
			"timezone": "Africa/Johannesburg",
			"spam_actions": 2,
			"id": 2,
			"deliveryservers": [{
				"address": "192.168.1.150",
				"id": 2,
				"port": 25
			}],
			"site_url": "https://mail.example.com",
			"authservers": [{
				"protocol": 2,
				"id": 2,
				"address": "mail.example.com"
			}],
			"report_every": 3,
			"aliases": [{
				"name": "mojo.example.com",
				"id": 2
			}],
			"status": true,
			"accept_inbound": true,
			"discard_mail": false,
			"virus_checks_at_smtp": true,
			"low_score": 10.0,
			"name": "example.com",
			"language": "en",
			"spam_checks": false,
			"smtp_callout": false,
			"message_size": "0",
			"high_score": 20.0,
			"virus_actions": 2
		}, {
			"signatures": [],
			"highspam_actions": 2,
			"delivery_mode": 1,
			"virus_checks": true,
			"ldap_callout": false,
			"dkimkeys": [],
			"timezone": "Africa/Johannesburg",
			"spam_actions": 2,
			"id": 4,
			"deliveryservers": [{
				"address": "192.168.1.150",
				"id": 4,
				"port": 25
			}],
			"site_url": "https://mail.example.net",
			"authservers": [],
			"report_every": 3,
			"aliases": [],
			"status": true,
			"discard_mail": false,
			"virus_checks_at_smtp": false,
			"low_score": 0.0,
			"name": "example.net",
			"language": "en",
			"spam_checks": true,
			"smtp_callout": true,
			"message_size": "0",
			"high_score": 0.0,
			"virus_actions": 2
		}],
		"meta": {
			"total": 2
		},
		"links": {
			"pages": {
				"last": "http://baruwa.example.com/api/v1/domains?page=2",
				"next": "http://baruwa.example.com/api/v1/domains?page=2"
			}
		}
	}`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetDomains(nil)
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
	next := "http://baruwa.example.com/api/v1/domains?page=2"
	if u.Links.Pages.Next != next {
		t.Errorf("Expected '%s' got '%s'", next, u.Links.Pages.Next)
	}
}

func TestGetDomainError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetDomain(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
	if ds != nil {
		t.Errorf("Expected %v got %v", nil, ds)
	}
}

func TestGetDomainOK(t *testing.T) {
	domainID := 4
	data := fmt.Sprintf(`
	{
		"signatures": [{
			"type": 1,
			"id": 1
		}],
		"highspam_actions": 2,
		"delivery_mode": 1,
		"virus_checks": true,
		"ldap_callout": false,
		"dkimkeys": [],
		"timezone": "Africa/Johannesburg",
		"spam_actions": 2,
		"id": %d,
		"deliveryservers": [{
			"address": "192.168.1.150",
			"id": 2,
			"port": 25
		}],
		"site_url": "https://mail.example.com",
		"authservers": [{
			"protocol": 2,
			"id": 2,
			"address": "mail.example.com"
		}],
		"report_every": 3,
		"aliases": [{
			"name": "mojo.example.com",
			"id": 2
		}],
		"status": true,
		"discard_mail": false,
		"virus_checks_at_smtp": true,
		"low_score": 10.0,
		"name": "example.com",
		"language": "en",
		"spam_checks": false,
		"smtp_callout": false,
		"message_size": "0",
		"high_score": 20.0,
		"virus_actions": 2
	}
	`, domainID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	d, err := client.GetDomain(domainID)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	if d.ID != domainID {
		t.Errorf("Expected %d got %d", domainID, d.ID)
	}
}

func TestGetDomainByNameError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	ds, err := client.GetDomainByName("")
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainNameParamError {
		t.Errorf("Expected '%s' got '%s'", domainNameParamError, err)
	}
	if ds != nil {
		t.Errorf("Expected %v got %v", nil, ds)
	}
}

func TestGetDomainByNameOK(t *testing.T) {
	domainName := "example.net"
	data := fmt.Sprintf(`
	{
		"signatures": [{
			"type": 1,
			"id": 1
		}],
		"highspam_actions": 2,
		"delivery_mode": 1,
		"virus_checks": true,
		"ldap_callout": false,
		"dkimkeys": [],
		"timezone": "Africa/Johannesburg",
		"spam_actions": 2,
		"id": 1,
		"deliveryservers": [{
			"address": "192.168.1.150",
			"id": 2,
			"port": 25
		}],
		"site_url": "https://mail.example.com",
		"authservers": [{
			"protocol": 2,
			"id": 2,
			"address": "mail.example.com"
		}],
		"report_every": 3,
		"aliases": [{
			"name": "mojo.example.com",
			"id": 2
		}],
		"status": true,
		"discard_mail": false,
		"virus_checks_at_smtp": true,
		"low_score": 10.0,
		"name": "%s",
		"language": "en",
		"spam_checks": false,
		"smtp_callout": false,
		"message_size": "0",
		"high_score": 20.0,
		"virus_actions": 2
	}
	`, domainName)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	d, err := client.GetDomainByName(domainName)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	if d.Name != domainName {
		t.Errorf("Expected '%s' got '%s'", domainName, d.Name)
	}
}

func TestCreateDomainError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.CreateDomain(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainParamError {
		t.Errorf("Expected '%s' got '%s'", domainParamError, err)
	}
}

func TestCreateDomainOK(t *testing.T) {
	domainID := 2
	data := fmt.Sprintf(`
	{
		"signatures": [],
		"highspam_actions": 3,
		"delivery_mode": 1,
		"virus_checks": true,
		"ldap_callout": true,
		"dkimkeys": [],
		"timezone": "Africa/Johannesburg",
		"spam_actions": 3,
		"id": %d,
		"deliveryservers": [],
		"site_url": "http://baruwa.example.net",
		"authservers": [],
		"report_every": 3,
		"aliases": [],
		"status": true,
		"discard_mail": false,
		"virus_checks_at_smtp": true,
		"low_score": 0.0,
		"name": "example.net",
		"language": "en",
		"spam_checks": true,
		"smtp_callout": true,
		"message_size": "0",
		"high_score": 0.0,
		"virus_actions": 3
	}
	`, domainID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	d := &Domain{
		Name:              "example.net",
		Language:          "en",
		SpamChecks:        true,
		SMTPCallout:       true,
		MessageSize:       "0",
		HighScore:         0.0,
		VirusActions:      3,
		HighspamActions:   3,
		DeliveryMode:      1,
		VirusChecks:       true,
		LdapCallout:       true,
		Timezone:          "Africa/Johannesburg",
		SpamActions:       3,
		SiteURL:           "http://baruwa.example.net",
		ReportEvery:       3,
		Status:            true,
		DiscardMail:       false,
		VirusChecksAtSMTP: true,
		LowScore:          0.0,
	}
	err = client.CreateDomain(d)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if d.ID != domainID {
		t.Errorf("Expected %d got %d", domainID, d.ID)
	}
}

func TestUpdateDomainError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateDomain(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainParamError {
		t.Errorf("Expected '%s' got '%s'", domainParamError, err)
	}
	d := &Domain{
		Name:              "example.net",
		Language:          "en",
		SpamChecks:        true,
		SMTPCallout:       true,
		MessageSize:       "0",
		HighScore:         0.0,
		VirusActions:      3,
		HighspamActions:   3,
		DeliveryMode:      1,
		VirusChecks:       true,
		LdapCallout:       true,
		Timezone:          "Africa/Johannesburg",
		SpamActions:       3,
		SiteURL:           "http://baruwa.example.net",
		ReportEvery:       3,
		Status:            true,
		DiscardMail:       false,
		VirusChecksAtSMTP: true,
		LowScore:          0.0,
	}
	err = client.UpdateDomain(d)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainSIDError {
		t.Errorf("Expected '%s' got '%s'", domainSIDError, err)
	}
}

func TestUpdateDomainOK(t *testing.T) {
	domainID := 2
	data := fmt.Sprintf(`
	{
		"signatures": [],
		"highspam_actions": 3,
		"delivery_mode": 1,
		"virus_checks": true,
		"ldap_callout": true,
		"dkimkeys": [],
		"timezone": "Africa/Johannesburg",
		"spam_actions": 3,
		"id": %d,
		"deliveryservers": [],
		"site_url": "http://baruwa.example.net",
		"authservers": [],
		"report_every": 3,
		"aliases": [],
		"status": true,
		"discard_mail": false,
		"virus_checks_at_smtp": true,
		"low_score": 0.0,
		"name": "example.net",
		"language": "en",
		"spam_checks": true,
		"smtp_callout": true,
		"message_size": "0",
		"high_score": 0.0,
		"virus_actions": 3
	}
	`, domainID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	d := &Domain{
		ID:                domainID,
		Name:              "example.net",
		Language:          "en",
		SpamChecks:        true,
		SMTPCallout:       true,
		MessageSize:       "0",
		HighScore:         0.0,
		VirusActions:      3,
		HighspamActions:   3,
		DeliveryMode:      1,
		VirusChecks:       true,
		LdapCallout:       true,
		Timezone:          "Africa/Johannesburg",
		SpamActions:       3,
		SiteURL:           "http://baruwa.example.net",
		ReportEvery:       3,
		Status:            true,
		DiscardMail:       false,
		VirusChecksAtSMTP: true,
		LowScore:          0.0,
	}
	err = client.UpdateDomain(d)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
}

func TestDeleteDomainError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteDomain(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != domainIDError {
		t.Errorf("Expected '%s' got '%s'", domainIDError, err)
	}
}

func TestDeleteDomainOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteDomain(1)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
}
