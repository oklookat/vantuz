// Just HTTP Client.
package vantuz

import (
	"net/http"
	"time"
)

const (
	_userAgent = "vantuz"
)

// Client.
func C() *Client {
	p := &Client{}
	p.SetClient(&http.Client{
		Timeout: 20 * time.Second,
	})
	p.SetLogger(&dummyLogger{})
	p.SetGlobalHeader("Content-Type", "application/json")
	p.SetUserAgent(_userAgent)
	return p
}

func (c *Client) SetUserAgent(val string) *Client {
	c.SetGlobalHeader("User-Agent", val)
	return c
}

func (c *Client) SetAuthorization(val string) *Client {
	c.SetGlobalHeader("Authorization", val)
	return c
}
