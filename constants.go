// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

const (
	// APIVersion of Baruwa API
	APIVersion = "v1"
	// Version of this library
	Version              = "0.0.1"
	timeFmt              = "2006:01:02:15:04:05"
	endpointError        = "The endpoint param is required"
	userIDError          = "The userID param should be > 0"
	aliasIDError         = "The aliasID param should be > 0"
	domainIDError        = "The domainID param should be > 0"
	serverIDError        = "The serverID param should be > 0"
	settingsIDError      = "The settingsID param should be > 0"
	organizationIDError  = "The organizationID param should be > 0"
	relayIDError         = "The relayID param should be > 0"
	aliasSIDError        = "The alias.ID param should be > 0"
	serverSIDError       = "The server.ID param should be > 0"
	settingsSIDError     = "The settings.ID param should be > 0"
	domainSIDError       = "The domain.ID param should be > 0"
	formSIDError         = "The form.ID param should be > 0"
	orgSIDError          = "The org.ID param should be > 0"
	userParamError       = "The user param is required"
	aliasParamError      = "The alias param is required"
	serverParamError     = "The server param is required"
	settingsParamError   = "The settings param is required"
	domainNameParamError = "The domainName param is required"
	domainParamError     = "The domain param is required"
	formParamError       = "The form param is required"
	orgParamError        = "The org param is required"
	clientIDError        = "clientID is required"
	clientSecretError    = "secret is required"
	pwFormError          = "The form param is required"
	// UserListURL - users list paging url fmt string
	UserListURL = "%s/api/%s/users?page=%d"
	// OrgListURL - organization list paging url fmt string
	OrgListURL = "%s/api/%s/organizations?page=%d"
	// OrgSMListURL - organization smarthosts list paging url fmt string
	OrgSMListURL = "%s/api/%s/organizations/smarthosts/%d?page=%d"
	// OrgFSListURL - fallback servers list paging url fmt string
	OrgFSListURL = "%s/api/%s/failbackservers/%d?page=%d"
	// DomainListURL - domains list paging url fmt string
	DomainListURL = "%s/api/%s/domains?page=%d"
	// UDSListURL - user delivery servers list paging url fmt string
	UDSListURL = "%s/api/%s/userdeliveryservers/%d?page=%d"
	// DSMListURL - domain smarthosts list paging url fmt string
	DSMListURL = "%s/api/%s/domains/smarthosts/%d?page=%d"
	// DDSListURL - delivery servers list paging url fmt string
	DDSListURL = "%s/api/%s/deliveryservers/%d?page=%d"
	// DASListURL - domain smarthosts list paging url fmt string
	DASListURL = "%s/api/%s/authservers/%d?page=%d"
	// DAliasListURL - domain aliases list paging url fmt string
	DAliasListURL = "%s/api/%s/domainaliases/%d?page=%d"
)
