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

func Test_User_NotFoundError(t *testing.T) {
	nf := "Not Found"
	server, client, err := getTestServerAndClient(http.StatusNotFound, fmt.Sprintf(`{"code":%d, "error":"%s"}`, http.StatusNotFound, nf))
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusNotFound {
			t.Errorf("Expected %d got %d", http.StatusNotFound, v.Code)
		}
		if v.Message != nf {
			t.Errorf("Expected '%s' got '%s'", nf, v.Message)
		}
	default:
		t.Errorf("Expected *ErrorResponse got %v", v)
	}
	if u.ID != 0 {
		t.Errorf("Expected %d got %d", 0, u.ID)
	}
}

func Test_User_ServerError(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusInternalServerError, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusInternalServerError {
			t.Errorf("Expected %d got %d", http.StatusInternalServerError, v.Code)
		}
	default:
		t.Errorf("Expected error got %v", v)
	}
	if u.ID != 0 {
		t.Errorf("Expected %d got %d", 0, u.ID)
	}
}

func Test_User_UnAuthError(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusUnauthorized, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(5)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	switch v := err.(type) {
	case *ErrorResponse:
		if v.Code != http.StatusUnauthorized {
			t.Errorf("Expected %d got %d", http.StatusUnauthorized, v.Code)
		}
	default:
		t.Errorf("Expected error got %v", v)
	}
	if u.ID != 0 {
		t.Errorf("Expected %d got %d", 0, u.ID)
	}
}

func Test_User_InvalidID(t *testing.T) {
	server, client, err := getTestServerAndClient(http.StatusUnauthorized, ``)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(0)
	if err == nil {
		t.Fatalf("An error should be returned: %v", u)
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
	if u != nil {
		t.Errorf("Expected %v got %v", nil, u)
	}
}

func Test_GetUserOK(t *testing.T) {
	data := `{
	"username": "rowdyrough",
	"send_report": false,
	"account_type": 3,
	"addresses": [],
	"firstname": "Rowdy",
	"organizations": [],
	"lastname": "Rough",
	"spam_checks": false,
	"email": "rowdyrough@example.com",
	"low_score": 0.0,
	"high_score": 0.0,
	"created_on": "2014:10:07:06:35:48",
	"last_login": "2014:10:11:22:38:11",
	"active": true,
	"timezone": "Africa/Johannesburg",
	"local": true,
	"id": 2,
	"domains": [{
		"name": "example.com",
		"id": 4
	}]
	}
`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUser(2)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err.Error())
	}
	if u.ID != 2 {
		t.Errorf("Expected %d got %d", 2, u.ID)
	}
}

func Test_GetUsersOK(t *testing.T) {
	data := `{
		"items": [{
			"username": "fuzzy@example.com",
			"send_report": false,
			"account_type": 3,
			"addresses": [],
			"firstname": "Fuzzy",
			"organizations": [],
			"lastname": "Lumpkins",
			"spam_checks": false,
			"email": "fuzzy@example.com",
			"low_score": 2.0,
			"high_score": 12.0,
			"created_on": "2014:09:20:15:14:30",
			"last_login": "2014:10:03:08:54:28",
			"active": true,
			"timezone": "Africa/Abidjan",
			"local": true,
			"id": 4,
			"domains": [{
				"name": "example.com",
				"id": 4
			}]
		}, {
			"username": "rowdyrough",
			"send_report": false,
			"account_type": 3,
			"addresses": [],
			"firstname": "Rowdy",
			"organizations": [],
			"lastname": "Rough",
			"spam_checks": false,
			"email": "rowdyrough@example.com",
			"low_score": 0.0,
			"high_score": 0.0,
			"created_on": "2014:10:07:06:35:48",
			"last_login": "2014:10:11:22:38:11",
			"active": true,
			"timezone": "Africa/Johannesburg",
			"local": true,
			"id": 5,
			"domains": [{
				"name": "example.com",
				"id": 4
			}]
		}],
		"meta": {
			"total": 2
		},
		"links": {
			"pages": {
				"last": "http://baruwa.example.com/api/v1/users?page=2",
				"next": "http://baruwa.example.com/api/v1/users?page=2"
			}
		}
	}
`
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	u, err := client.GetUsers(nil)
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
	next := "http://baruwa.example.com/api/v1/users?page=2"
	if u.Links.Pages.Next != next {
		t.Errorf("Expected '%s' got '%s'", next, u.Links.Pages.Next)
	}
	t.Log(u)
}

func Test_CreateUserError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	_, err = client.CreateUser(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userParamError {
		t.Errorf("Expected '%s' got '%s'", userParamError, err)
	}
}

func Test_CreateUserOK(t *testing.T) {
	userID := 2
	data := fmt.Sprintf(`{
		"username": "rowdyrough",
		"send_report": false,
		"account_type": 3,
		"addresses": [],
		"firstname": "Rowdy",
		"organizations": [],
		"lastname": "Rough",
		"spam_checks": false,
		"email": "rowdyrough@example.com",
		"low_score": 0.0,
		"high_score": 0.0,
		"created_on": "2014:10:07:06:35:48",
		"last_login": "2014:10:11:22:38:11",
		"active": true,
		"timezone": "Africa/Johannesburg",
		"local": true,
		"id": %d,
		"domains": [{
			"name": "example.com",
			"id": 4
		}]
		}
	`, userID)
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	n := "andrew"
	e := "andrew@example.com"
	tz := "Africa/Johannesburg"
	u := &UserForm{
		Username: &n,
		Email:    &e,
		Timezone: &tz,
	}
	user, err := client.CreateUser(u)
	if err != nil {
		t.Fatalf("An error should not be returned: %s", err)
	}
	if user.ID != userID {
		t.Errorf("Expected %d got %d", userID, user.ID)
	}
}

func Test_UpdateUserError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.UpdateUser(nil)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userParamError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
	n := "andrew"
	e := "andrew@example.com"
	tz := "Africa/Johannesburg"
	u := &UserForm{
		Username: &n,
		Email:    &e,
		Timezone: &tz,
	}
	err = client.UpdateUser(u)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
}

func Test_UpdateUserOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	id := 2
	n := "andrew"
	e := "andrew@example.com"
	tz := "Africa/Johannesburg"
	u := &UserForm{
		ID:       &id,
		Username: &n,
		Email:    &e,
		Timezone: &tz,
	}
	err = client.UpdateUser(u)
	if err != nil {
		t.Fatalf("An error not should be returned: %s", err)
	}
}

func Test_DeleteUserError(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteUser(0)
	if err == nil {
		t.Fatalf("An error should be returned")
	}
	if err.Error() != userIDError {
		t.Errorf("Expected '%s' got '%s'", userIDError, err)
	}
}

func Test_DeleteUserOK(t *testing.T) {
	data := ``
	server, client, err := getTestServerAndClient(http.StatusOK, data)
	if err != nil {
		t.Fatalf("An error should not be returned")
	}
	defer server.Close()
	err = client.DeleteUser(2)
	if err != nil {
		t.Fatalf("An error not should be returned: %s", err)
	}
}
