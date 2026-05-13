package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Client holds configuration for all three Ubiquiti API surfaces.
type Client struct {
	CloudAPIKey   string
	ControllerURL string // optional local controller, e.g. https://192.168.1.1
	LocalAPIKey   string // optional local API key
	SiteID        string // legacy site ID (cloud site manager) OR integration site UUID

	httpClient *http.Client

	// Auto-discovered on first network/protect call when no ControllerURL is set.
	discoverOnce      sync.Once
	consoleID         string // resolved from Site Manager API during discover()
	integrationSiteID string // UUID from /proxy/network/integration/v1/sites
	discoverErr       error
}

// NewClient creates a new Unifi API client.
func NewClient(cloudAPIKey, controllerURL, localAPIKey, siteID string, sslInsecure bool) *Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecure}, //nolint:gosec
	}
	return &Client{
		CloudAPIKey:   cloudAPIKey,
		ControllerURL: controllerURL,
		LocalAPIKey:   localAPIKey,
		SiteID:        siteID,
		httpClient: &http.Client{
			Timeout:   30 * time.Second,
			Transport: transport,
		},
	}
}

// ErrNotFound is returned when the API responds with HTTP 404.
var ErrNotFound = fmt.Errorf("resource not found (404)")

// ── Cloud proxy discovery ─────────────────────────────────────────────────────

// discover resolves consoleID and integrationSiteID exactly once.
// It is called lazily before the first network/protect request when no local
// controller URL is configured.
func (c *Client) discover() error {
	c.discoverOnce.Do(func() {
		// 1. Get the console (host) ID — use the first host returned.
		hosts, err := c.ListHosts()
		if err != nil {
			c.discoverErr = fmt.Errorf("discovering console: %w", err)
			return
		}
		if len(hosts) == 0 {
			c.discoverErr = fmt.Errorf("no consoles found in cloud account")
			return
		}
		c.consoleID = hosts[0].ID

		// 2. Resolve the integration site ID via the cloud proxy.
		type integrationSite struct {
			ID                string `json:"id"`
			InternalReference string `json:"internalReference"`
			Name              string `json:"name"`
		}
		type integrationSitesResp struct {
			Data []integrationSite `json:"data"`
		}
		var sitesResp integrationSitesResp
		path := fmt.Sprintf("/v1/connector/consoles/%s/proxy/network/integration/v1/sites", c.consoleID)
		if err := c.doCloud(path, &sitesResp); err != nil {
			c.discoverErr = fmt.Errorf("discovering integration sites: %w", err)
			return
		}
		if len(sitesResp.Data) == 0 {
			c.discoverErr = fmt.Errorf("no integration sites found for console %s", c.consoleID)
			return
		}
		// If SiteID matches internalReference (e.g. "default") or the UUID, use it.
		// Otherwise fall back to the first site.
		c.integrationSiteID = sitesResp.Data[0].ID
		for _, s := range sitesResp.Data {
			if s.ID == c.SiteID || s.InternalReference == c.SiteID {
				c.integrationSiteID = s.ID
				break
			}
		}
	})
	return c.discoverErr
}

// ── Request helpers ───────────────────────────────────────────────────────────

// doCloud performs a GET against the Site Manager cloud API (api.ui.com).
func (c *Client) doCloud(path string, out interface{}) error {
	url := "https://api.ui.com" + path
	return c.doRequest(url, c.CloudAPIKey, out)
}

// doNetwork performs a GET against the Network integration API.
// Prefers a local controller if configured; otherwise routes through the cloud proxy.
func (c *Client) doNetwork(path string, out interface{}) error {
	if c.ControllerURL != "" {
		url := fmt.Sprintf("%s/proxy/network/integration/v1/sites/%s%s", c.ControllerURL, c.SiteID, path)
		return c.doRequest(url, c.LocalAPIKey, out)
	}
	if c.CloudAPIKey == "" {
		return fmt.Errorf("no controller URL or cloud API key configured")
	}
	if err := c.discover(); err != nil {
		return err
	}
	url := fmt.Sprintf("https://api.ui.com/v1/connector/consoles/%s/proxy/network/integration/v1/sites/%s%s",
		c.consoleID, c.integrationSiteID, path)
	return c.doRequest(url, c.CloudAPIKey, out)
}

// doProtect performs a GET against the Protect integration API.
// Prefers a local controller if configured; otherwise routes through the cloud proxy.
func (c *Client) doProtect(path string, out interface{}) error {
	if c.ControllerURL != "" {
		url := fmt.Sprintf("%s/proxy/protect/integration/v1%s", c.ControllerURL, path)
		return c.doRequest(url, c.LocalAPIKey, out)
	}
	if c.CloudAPIKey == "" {
		return fmt.Errorf("no controller URL or cloud API key configured")
	}
	if err := c.discover(); err != nil {
		return err
	}
	url := fmt.Sprintf("https://api.ui.com/v1/connector/consoles/%s/proxy/protect/integration/v1%s",
		c.consoleID, path)
	return c.doRequest(url, c.CloudAPIKey, out)
}

func (c *Client) doRequest(url, apiKey string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode == 404 {
		return ErrNotFound
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d from %s: %s", resp.StatusCode, url, string(body))
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("decoding response from %s: %w", url, err)
	}
	return nil
}

// listAllNetwork fetches every page from a paginated Network API endpoint.
func listAllNetwork[T any](c *Client, path string) ([]T, error) {
	const pageSize = 100
	var all []T
	offset := 0
	for {
		var p page[T]
		sep := "?"
		if strings.Contains(path, "?") {
			sep = "&"
		}
		pagedPath := fmt.Sprintf("%s%soffset=%d&limit=%d", path, sep, offset, pageSize)
		if err := c.doNetwork(pagedPath, &p); err != nil {
			return nil, err
		}
		all = append(all, p.Data...)
		if len(all) >= p.TotalCount || len(p.Data) == 0 {
			break
		}
		offset += pageSize
	}
	return all, nil
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 ||
		func() bool {
			for i := 0; i <= len(s)-len(sub); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		}())
}

// doNetworkRoot performs a GET against the Network integration API at the root level
// (no site prefix) — used for endpoints like /v1/pending-devices.
func (c *Client) doNetworkRoot(path string, out interface{}) error {
	if c.ControllerURL != "" {
		url := fmt.Sprintf("%s/proxy/network/integration/v1%s", c.ControllerURL, path)
		return c.doRequest(url, c.LocalAPIKey, out)
	}
	if c.CloudAPIKey == "" {
		return fmt.Errorf("no controller URL or cloud API key configured")
	}
	if err := c.discover(); err != nil {
		return err
	}
	url := fmt.Sprintf("https://api.ui.com/v1/connector/consoles/%s/proxy/network/integration/v1%s",
		c.consoleID, path)
	return c.doRequest(url, c.CloudAPIKey, out)
}

// doNetworkWrite performs a write (POST/PUT/DELETE) against the Network integration API.
func (c *Client) doNetworkWrite(method, path string, body io.Reader, out interface{}) error {
	if err := c.discover(); err != nil {
		return err
	}
	var url string
	var apiKey string
	if c.ControllerURL != "" {
		url = fmt.Sprintf("%s/proxy/network/integration/v1/sites/%s%s", c.ControllerURL, c.SiteID, path)
		apiKey = c.LocalAPIKey
	} else {
		url = fmt.Sprintf("https://api.ui.com/v1/connector/consoles/%s/proxy/network/integration/v1/sites/%s%s",
			c.consoleID, c.integrationSiteID, path)
		apiKey = c.CloudAPIKey
	}
	return c.doWriteRequest(method, url, apiKey, body, out)
}

func (c *Client) doWriteRequest(method, url, apiKey string, body io.Reader, out interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d from %s: %s", resp.StatusCode, url, string(respBody))
	}

	if out != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, out); err != nil {
			return fmt.Errorf("decoding response from %s: %w", url, err)
		}
	}
	return nil
}
