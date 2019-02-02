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
	Version           = "0.0.1"
	timeFmt           = "2006:01:02:15:04:05"
	endpointError     = "The endpoint param is required"
	userIDError       = "The userID param should be > 0"
	userParamError    = "The user param cannot be nil"
	clientIDError     = "clientID is required"
	clientSecretError = "secret is required"
	pwFormError       = "The form param is required"
)
