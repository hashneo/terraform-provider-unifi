package client

import "errors"

// ── Site Manager types ────────────────────────────────────────────────────────

type HostsResponse struct {
	Data []Host `json:"data"`
}

type Host struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	HardwareID   string            `json:"hardwareId"`
	IsBlocked    bool              `json:"isBlocked"`
	LatestBackup *LatestBackup     `json:"latestBackup"`
	Owner        bool              `json:"owner"`
	ReportedState *HostReportedState `json:"reportedState"`
	UserPermissions []string       `json:"userPermissions"`
}

type LatestBackup struct {
	CompletionTime string `json:"completionTime"`
}

type HostReportedState struct {
	Hostname       string `json:"hostname"`
	IP             string `json:"ip"`
	IsConfigured   bool   `json:"isConfigured"`
	IsSetup        bool   `json:"isSetup"`
	Name           string `json:"name"`
	HardwareID     string `json:"hardwareId"`
	FirmwareVersion string `json:"firmwareVersion"`
}

type SitesResponse struct {
	Data []Site `json:"data"`
}

type Site struct {
	HostID      string `json:"hostId"`
	IsOwner     bool   `json:"isOwner"`
	MetaData    *SiteMetaData `json:"meta"`
	SiteID      string `json:"siteId"`
}

type SiteMetaData struct {
	GatewayMAC  string `json:"gatewayMac"`
	Name        string `json:"name"`
}

type CloudDevicesResponse struct {
	Data []CloudDevice `json:"data"`
}

type CloudDevice struct {
	ID          string `json:"id"`
	MAC         string `json:"mac"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	IP          string `json:"ip"`
	IsAdopted   bool   `json:"isAdopted"`
	IsConnected bool   `json:"isConnected"`
	FirmwareVersion string `json:"firmwareVersion"`
	ProductLine string `json:"productLine"`
	SiteID      string `json:"siteId"`
	HostID      string `json:"hostId"`
}

type SDWANConfigsResponse struct {
	Data []SDWANConfig `json:"data"`
}

type SDWANConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

// ── Site Manager API methods ──────────────────────────────────────────────────

func (c *Client) ListHosts() ([]Host, error) {
	var resp HostsResponse
	if err := c.doCloud("/v1/hosts", &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) ListSites() ([]Site, error) {
	var resp SitesResponse
	if err := c.doCloud("/v1/sites", &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) ListCloudDevices() ([]CloudDevice, error) {
	var resp CloudDevicesResponse
	if err := c.doCloud("/v1/devices", &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) ListSDWANConfigs() ([]SDWANConfig, error) {
	var resp SDWANConfigsResponse
	if err := c.doCloud("/v1/sd-wan/configs", &resp); err != nil {
		if errors.Is(err, ErrNotFound) {
			return []SDWANConfig{}, nil
		}
		return nil, err
	}
	return resp.Data, nil
}
