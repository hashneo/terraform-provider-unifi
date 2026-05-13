package client

import (
	"bytes"
	"encoding/json"
)

// ── Pagination envelope ───────────────────────────────────────────────────────

type page[T any] struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Count      int `json:"count"`
	TotalCount int `json:"totalCount"`
	Data       []T `json:"data"`
}

// ── Network API types (per openapi.json v10.3.58) ─────────────────────────────

// AdoptedDevice maps to "Adopted device overview".
type AdoptedDevice struct {
	ID                string   `json:"id"`
	MACAddress        string   `json:"macAddress"`
	IPAddress         string   `json:"ipAddress"`
	Name              string   `json:"name"`
	Model             string   `json:"model"`
	State             string   `json:"state"`
	Supported         bool     `json:"supported"`
	FirmwareVersion   string   `json:"firmwareVersion"`
	FirmwareUpdatable bool     `json:"firmwareUpdatable"`
	Features          []string `json:"features"`
}

// PendingDevice maps to "Device pending adoption".
// Returned by GET /v1/pending-devices (no site prefix).
type PendingDevice struct {
	MACAddress          string   `json:"macAddress"`
	IPAddress           string   `json:"ipAddress"`
	Model               string   `json:"model"`
	State               string   `json:"state"`
	Supported           bool     `json:"supported"`
	FirmwareVersion     string   `json:"firmwareVersion"`
	FirmwareUpdatable   bool     `json:"firmwareUpdatable"`
	Features            []string `json:"features"`
	AdoptionTargetSites []string `json:"adoptionTargetSiteIds"`
}

// NetworkClient maps to "Wired/Wireless client overview".
// The API returns a discriminated union; we capture the common fields.
type NetworkClient struct {
	Type          string `json:"type"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ConnectedAt   string `json:"connectedAt"`
	IPAddress     string `json:"ipAddress"`
	MACAddress    string `json:"macAddress"`
	UplinkDeviceID string `json:"uplinkDeviceId"`
}

// Network maps to "Network overview".
type Network struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Management string `json:"management"`
	Enabled    bool   `json:"enabled"`
	VLANId     int    `json:"vlanId"`
	Default    bool   `json:"default"`
}

// WiFiBroadcast maps to "Wifi broadcast overview".
type WiFiBroadcast struct {
	Type    string              `json:"type"`
	ID      string              `json:"id"`
	Name    string              `json:"name"`
	Enabled bool                `json:"enabled"`
	Network *WiFiNetworkRef     `json:"network"`
	Security *WiFiSecurityOverview `json:"securityConfiguration"`
}

type WiFiNetworkRef struct {
	Type string `json:"type"`
}

type WiFiSecurityOverview struct {
	Type string `json:"type"`
}

// FirewallZone maps to "Firewall zone".
type FirewallZone struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	NetworkIDs []string          `json:"networkIds"`
	Metadata   *EntityMetadata   `json:"metadata"`
}

// EntityMetadata captures the origin/configurable metadata present on many objects.
type EntityMetadata struct {
	Origin       string `json:"origin"`
	Configurable bool   `json:"configurable"`
}

// FirewallPolicyZoneRef is a source or destination zone reference in a policy.
type FirewallPolicyZoneRef struct {
	ZoneID string `json:"zoneId"`
}

// FirewallPolicyIPScope captures the IP version scope.
type FirewallPolicyIPScope struct {
	IPVersion string `json:"ipVersion"`
}

// FirewallPolicy maps to "Firewall policy".
type FirewallPolicy struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Enabled         bool                   `json:"enabled"`
	Index           int                    `json:"index"`
	Action          string                 `json:"-"` // flattened from {"type":"ALLOW"}
	AllowReturn     bool                   `json:"-"` // from action.allowReturnTraffic
	Source          *FirewallPolicyZoneRef `json:"source"`
	Destination     *FirewallPolicyZoneRef `json:"destination"`
	IPProtocolScope *FirewallPolicyIPScope `json:"ipProtocolScope"`
	LoggingEnabled  bool                   `json:"loggingEnabled"`
	Description     string                 `json:"description"`
	Metadata        *EntityMetadata        `json:"metadata"`
}

// firewallPolicyRaw is used for unmarshalling the nested action field.
type firewallPolicyRaw struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Enabled         bool                   `json:"enabled"`
	Index           int                    `json:"index"`
	Description     string                 `json:"description"`
	LoggingEnabled  bool                   `json:"loggingEnabled"`
	Source          *FirewallPolicyZoneRef `json:"source"`
	Destination     *FirewallPolicyZoneRef `json:"destination"`
	IPProtocolScope *FirewallPolicyIPScope `json:"ipProtocolScope"`
	Metadata        *EntityMetadata        `json:"metadata"`
	Action          json.RawMessage        `json:"action"`
}

func (p *FirewallPolicy) UnmarshalJSON(b []byte) error {
	var raw firewallPolicyRaw
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	p.ID = raw.ID
	p.Name = raw.Name
	p.Enabled = raw.Enabled
	p.Index = raw.Index
	p.Description = raw.Description
	p.LoggingEnabled = raw.LoggingEnabled
	p.Source = raw.Source
	p.Destination = raw.Destination
	p.IPProtocolScope = raw.IPProtocolScope
	p.Metadata = raw.Metadata
	// action is {"type":"ALLOW","allowReturnTraffic":true} or similar
	var action struct {
		Type             string `json:"type"`
		AllowReturnTraffic bool `json:"allowReturnTraffic"`
	}
	if len(raw.Action) > 0 {
		_ = json.Unmarshal(raw.Action, &action)
	}
	p.Action = action.Type
	p.AllowReturn = action.AllowReturnTraffic
	return nil
}

// ACLRule maps to "ACL rule".
type ACLRule struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	Action      string `json:"action"`
	Index       int    `json:"index"`
}

// DNSPolicy maps to the DNS policy schema.
type DNSPolicy struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Domain  string `json:"domain"`
}

// VPNServer maps to "VPN server overview".
type VPNServer struct {
	Type    string `json:"type"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// WANInterface maps to "WAN overview".
type WANInterface struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DeviceTag maps to "Device tag". Field is deviceIds (not devices).
type DeviceTag struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	DeviceIDs []string `json:"deviceIds"`
}

// ── Network API methods ───────────────────────────────────────────────────────

func (c *Client) ListAdoptedDevices() ([]AdoptedDevice, error) {
	return listAllNetwork[AdoptedDevice](c, "/devices")
}

// ListPendingDevices uses /v1/pending-devices — no site prefix.
func (c *Client) ListPendingDevices() ([]PendingDevice, error) {
	var p page[PendingDevice]
	if err := c.doNetworkRoot("/pending-devices", &p); err != nil {
		return []PendingDevice{}, nil
	}
	return p.Data, nil
}

func (c *Client) ListClients() ([]NetworkClient, error) {
	return listAllNetwork[NetworkClient](c, "/clients")
}

func (c *Client) ListNetworks() ([]Network, error) {
	return listAllNetwork[Network](c, "/networks")
}

func (c *Client) GetNetwork(id string) (*Network, error) {
	var n Network
	if err := c.doNetwork("/networks/"+id, &n); err != nil {
		return nil, err
	}
	return &n, nil
}

func (c *Client) CreateNetwork(n Network) (*Network, error) {
	body, _ := json.Marshal(n)
	var out Network
	if err := c.doNetworkWrite("POST", "/networks", bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateNetwork(id string, n Network) (*Network, error) {
	body, _ := json.Marshal(n)
	var out Network
	if err := c.doNetworkWrite("PUT", "/networks/"+id, bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteNetwork(id string) error {
	return c.doNetworkWrite("DELETE", "/networks/"+id, nil, nil)
}

func (c *Client) ListWiFiBroadcasts() ([]WiFiBroadcast, error) {
	items, err := listAllNetwork[WiFiBroadcast](c, "/wifi/broadcasts")
	if err != nil {
		return []WiFiBroadcast{}, nil
	}
	return items, nil
}

func (c *Client) GetWiFiBroadcast(id string) (*WiFiBroadcast, error) {
	var w WiFiBroadcast
	if err := c.doNetwork("/wifi/broadcasts/"+id, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *Client) CreateWiFiBroadcast(w WiFiBroadcast) (*WiFiBroadcast, error) {
	body, _ := json.Marshal(w)
	var out WiFiBroadcast
	if err := c.doNetworkWrite("POST", "/wifi/broadcasts", bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateWiFiBroadcast(id string, w WiFiBroadcast) (*WiFiBroadcast, error) {
	body, _ := json.Marshal(w)
	var out WiFiBroadcast
	if err := c.doNetworkWrite("PUT", "/wifi/broadcasts/"+id, bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteWiFiBroadcast(id string) error {
	return c.doNetworkWrite("DELETE", "/wifi/broadcasts/"+id, nil, nil)
}

func (c *Client) ListFirewallZones() ([]FirewallZone, error) {
	items, err := listAllNetwork[FirewallZone](c, "/firewall/zones")
	if err != nil {
		return []FirewallZone{}, nil
	}
	return items, nil
}

func (c *Client) ListFirewallPolicies() ([]FirewallPolicy, error) {
	items, err := listAllNetwork[FirewallPolicy](c, "/firewall/policies")
	if err != nil {
		return []FirewallPolicy{}, nil
	}
	return items, nil
}

func (c *Client) GetFirewallPolicy(id string) (*FirewallPolicy, error) {
	var p FirewallPolicy
	if err := c.doNetwork("/firewall/policies/"+id, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *Client) CreateFirewallPolicy(p FirewallPolicy) (*FirewallPolicy, error) {
	body, _ := json.Marshal(p)
	var out FirewallPolicy
	if err := c.doNetworkWrite("POST", "/firewall/policies", bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateFirewallPolicy(id string, p FirewallPolicy) (*FirewallPolicy, error) {
	body, _ := json.Marshal(p)
	var out FirewallPolicy
	if err := c.doNetworkWrite("PUT", "/firewall/policies/"+id, bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteFirewallPolicy(id string) error {
	return c.doNetworkWrite("DELETE", "/firewall/policies/"+id, nil, nil)
}

func (c *Client) ListACLRules() ([]ACLRule, error) {
	items, err := listAllNetwork[ACLRule](c, "/acl-rules")
	if err != nil {
		return []ACLRule{}, nil
	}
	return items, nil
}

func (c *Client) GetACLRule(id string) (*ACLRule, error) {
	var r ACLRule
	if err := c.doNetwork("/acl-rules/"+id, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) CreateACLRule(r ACLRule) (*ACLRule, error) {
	body, _ := json.Marshal(r)
	var out ACLRule
	if err := c.doNetworkWrite("POST", "/acl-rules", bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateACLRule(id string, r ACLRule) (*ACLRule, error) {
	body, _ := json.Marshal(r)
	var out ACLRule
	if err := c.doNetworkWrite("PUT", "/acl-rules/"+id, bytes.NewReader(body), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteACLRule(id string) error {
	return c.doNetworkWrite("DELETE", "/acl-rules/"+id, nil, nil)
}

func (c *Client) ListDNSPolicies() ([]DNSPolicy, error) {
	return listAllNetwork[DNSPolicy](c, "/dns/policies")
}

func (c *Client) ListVPNServers() ([]VPNServer, error) {
	return listAllNetwork[VPNServer](c, "/vpn/servers")
}

func (c *Client) ListWANInterfaces() ([]WANInterface, error) {
	items, err := listAllNetwork[WANInterface](c, "/wans")
	if err != nil {
		return []WANInterface{}, nil
	}
	return items, nil
}

func (c *Client) ListDeviceTags() ([]DeviceTag, error) {
	return listAllNetwork[DeviceTag](c, "/device-tags")
}
