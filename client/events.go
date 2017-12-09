// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": events Resource Client
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateEventsPath computes a request path to the create action of events.
func CreateEventsPath() string {

	return fmt.Sprintf("/develop/v1/events")
}

// イベント作成
func (c *Client) CreateEvents(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateEventsRequest create the request corresponding to the create action endpoint of the events resource.
func (c *Client) NewCreateEventsRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteEventsPath computes a request path to the delete action of events.
func DeleteEventsPath(userID string, eventID string) string {
	param0 := userID
	param1 := eventID

	return fmt.Sprintf("/develop/v1/events/%s/%s", param0, param1)
}

// イベント削除
func (c *Client) DeleteEvents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteEventsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteEventsRequest create the request corresponding to the delete action endpoint of the events resource.
func (c *Client) NewDeleteEventsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListEventsPath computes a request path to the list action of events.
func ListEventsPath() string {

	return fmt.Sprintf("/develop/v1/events")
}

// イベント複数取得
func (c *Client) ListEvents(ctx context.Context, path string, limit int, offset int, keyword *string, userID *string) (*http.Response, error) {
	req, err := c.NewListEventsRequest(ctx, path, limit, offset, keyword, userID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEventsRequest create the request corresponding to the list action endpoint of the events resource.
func (c *Client) NewListEventsRequest(ctx context.Context, path string, limit int, offset int, keyword *string, userID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp8 := strconv.Itoa(limit)
	values.Set("limit", tmp8)
	tmp9 := strconv.Itoa(offset)
	values.Set("offset", tmp9)
	if keyword != nil {
		values.Set("keyword", *keyword)
	}
	if userID != nil {
		values.Set("user_id", *userID)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowEventsPath computes a request path to the show action of events.
func ShowEventsPath(userID string, eventID string) string {
	param0 := userID
	param1 := eventID

	return fmt.Sprintf("/develop/v1/events/%s/%s", param0, param1)
}

// イベント情報取得
func (c *Client) ShowEvents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowEventsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowEventsRequest create the request corresponding to the show action endpoint of the events resource.
func (c *Client) NewShowEventsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateEventsPath computes a request path to the update action of events.
func UpdateEventsPath(userID string, eventID string) string {
	param0 := userID
	param1 := eventID

	return fmt.Sprintf("/develop/v1/events/%s/%s", param0, param1)
}

// イベント編集
func (c *Client) UpdateEvents(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateEventsRequest create the request corresponding to the update action endpoint of the events resource.
func (c *Client) NewUpdateEventsRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
