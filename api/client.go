// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type expirationTime int64

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	*e = expirationTime(i)
	return nil
}

// Client represents the Baruwa API client
type Client struct {
	BaseURL   *url.URL
	UserAgent string
	client    *http.Client
	token     string
}

// Options represents optional settings and flags that can be passed to New
type Options struct {
	// HTTP client for communication with the Baruwa API
	HTTPClient *http.Client
	// User agent for HTTP client
	UserAgent string
}

// TokenResponse is for API response for the /oauth2/token endpoint
type TokenResponse struct {
	RefreshToken string         `json:"refresh_token"`
	Token        string         `json:"access_token"`
	Type         string         `json:"token_type"`
	Scope        string         `json:"score"`
	ExpiresIn    expirationTime `json:"expires_in"`
}

// ErrorResponse https://www.baruwa.com/docs/api/#errors
type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Message  string         `json:"error,omitempty"`
	Code     int            `json:"code,omitempty"`
}

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Code, r.Message)
}

func (c *Client) newRequest(method, path string, opts *ListOptions, body io.Reader) (req *http.Request, err error) {
	var k, p string
	var q url.Values
	var u, nu, rel *url.URL

	if rel, err = url.Parse(path); err != nil {
		return
	}

	u = c.BaseURL.ResolveReference(rel)

	if method == http.MethodGet && opts != nil && opts.Page != "" {
		if nu, err = url.Parse(opts.Page); err == nil {
			if strings.HasPrefix(nu.String(), u.String()) {
				q = nu.Query()
				if len(q) >= 1 {
					for k = range q {
						if k != "page" {
							q.Del(k)
						}
					}
					if p = q.Get("page"); p != "" {
						u.RawQuery = q.Encode()
					}
				}
			}
		}
	}

	if req, err = http.NewRequest(method, u.String(), body); err != nil {
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return
}

func (c *Client) get(path string, opts *ListOptions, data interface{}) (err error) {
	var req *http.Request

	if req, err = c.newRequest(http.MethodGet, apiPath(path), opts, nil); err != nil {
		return
	}

	err = c.doWithOAuth(req, data)

	return
}

func (c *Client) post(p string, v url.Values, data interface{}) (err error) {
	var req *http.Request

	// fmt.Println(v.Encode())

	if req, err = c.newRequest(http.MethodPost, apiPath(p), nil, strings.NewReader(v.Encode())); err != nil {
		return
	}

	err = c.doWithOAuth(req, data)

	return
}

func (c *Client) put(p string, v url.Values, data interface{}) (err error) {
	var req *http.Request

	fmt.Println(v.Encode())

	if req, err = c.newRequest(http.MethodPut, apiPath(p), nil, strings.NewReader(v.Encode())); err != nil {
		return
	}

	err = c.doWithOAuth(req, data)

	return
}

func (c *Client) delete(p string, v url.Values) (err error) {
	var req *http.Request

	if v == nil {
		if req, err = c.newRequest(http.MethodDelete, apiPath(p), nil, nil); err != nil {
			return
		}
	} else {
		if req, err = c.newRequest(http.MethodDelete, apiPath(p), nil, strings.NewReader(v.Encode())); err != nil {
			return
		}
	}

	err = c.doWithOAuth(req, nil)

	return
}

// GetAccessToken returns a token
func (c *Client) GetAccessToken(clientID, secret string) (token *TokenResponse, err error) {
	var req *http.Request
	var buf *bytes.Buffer

	if clientID == "" {
		err = fmt.Errorf(clientIDError)
		return
	}

	if secret == "" {
		err = fmt.Errorf(clientSecretError)
		return
	}

	buf = bytes.NewBuffer([]byte("grant_type=password"))
	if req, err = c.newRequest(http.MethodPost, "oauth2/token", nil, buf); err != nil {
		return
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secret)

	token = &TokenResponse{}

	err = c.do(req, token)

	return
}

func (c *Client) doWithOAuth(req *http.Request, data interface{}) (err error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	err = c.do(req, data)

	return
}

func (c *Client) do(req *http.Request, v interface{}) (err error) {
	var data []byte
	var errResp *ErrorResponse
	var resp *http.Response

	if resp, err = c.client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp = &ErrorResponse{Response: resp}
		errResp.Code = resp.StatusCode
		errResp.Message = resp.Status
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			if err = json.Unmarshal(data, errResp); err == nil {
				err = errResp
			}
		} else {
			err = errResp
		}

		return
	}

	if v == nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(v)

	return
}

func apiPath(p string) string {
	return fmt.Sprintf("/api/%s/%s", APIVersion, p)
}

// New creates a new Baruwa API client.Options are optional and can be nil.
func New(endpoint, token string, options *Options) (c *Client, err error) {
	var ua string
	var baseurl *url.URL
	var client *http.Client
	var transport *http.Transport

	if endpoint == "" {
		err = fmt.Errorf(endpointError)
		return
	}

	ua = fmt.Sprintf("baruwa-go/%s", Version)
	transport = &http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	}
	client = http.DefaultClient
	client.Transport = transport
	if baseurl, err = url.Parse(endpoint); err != nil {
		return
	}

	if options != nil {
		if options.HTTPClient != nil {
			client = options.HTTPClient
		}
		if options.UserAgent != "" {
			ua = options.UserAgent
		}
	}

	c = &Client{
		BaseURL:   baseurl,
		UserAgent: ua,
		client:    client,
		token:     token,
	}

	return
}
