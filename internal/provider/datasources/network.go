package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

// ── unifi_adopted_devices ─────────────────────────────────────────────────────

type AdoptedDevicesDataSource struct{ client *client.Client }

func NewAdoptedDevicesDataSource() datasource.DataSource { return &AdoptedDevicesDataSource{} }
func (d *AdoptedDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_adopted_devices"
}
func (d *AdoptedDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all adopted UniFi devices on the site (switches, APs, gateways).",
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":                  schema.StringAttribute{Computed: true},
				"mac_address":         schema.StringAttribute{Computed: true},
				"ip_address":          schema.StringAttribute{Computed: true},
				"name":                schema.StringAttribute{Computed: true},
				"model":               schema.StringAttribute{Computed: true},
				"state":               schema.StringAttribute{Computed: true},
				"supported":           schema.BoolAttribute{Computed: true},
				"firmware_version":    schema.StringAttribute{Computed: true},
				"firmware_updatable":  schema.BoolAttribute{Computed: true},
				"features":            schema.ListAttribute{Computed: true, ElementType: types.StringType},
			}}},
		},
	}
}
func (d *AdoptedDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type adoptedDevicesModel struct {
	Devices []adoptedDeviceModel `tfsdk:"devices"`
}
type adoptedDeviceModel struct {
	ID                types.String `tfsdk:"id"`
	MACAddress        types.String `tfsdk:"mac_address"`
	IPAddress         types.String `tfsdk:"ip_address"`
	Name              types.String `tfsdk:"name"`
	Model             types.String `tfsdk:"model"`
	State             types.String `tfsdk:"state"`
	Supported         types.Bool   `tfsdk:"supported"`
	FirmwareVersion   types.String `tfsdk:"firmware_version"`
	FirmwareUpdatable types.Bool   `tfsdk:"firmware_updatable"`
	Features          types.List   `tfsdk:"features"`
}

func (d *AdoptedDevicesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	devs, err := d.client.ListAdoptedDevices()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list adopted devices", err.Error())
		return
	}
	state := adoptedDevicesModel{}
	for _, dev := range devs {
		feats, diag := types.ListValueFrom(ctx, types.StringType, dev.Features)
		resp.Diagnostics.Append(diag...)
		state.Devices = append(state.Devices, adoptedDeviceModel{
			ID: types.StringValue(dev.ID), MACAddress: types.StringValue(dev.MACAddress),
			IPAddress: types.StringValue(dev.IPAddress), Name: types.StringValue(dev.Name),
			Model: types.StringValue(dev.Model), State: types.StringValue(dev.State),
			Supported: types.BoolValue(dev.Supported), FirmwareVersion: types.StringValue(dev.FirmwareVersion),
			FirmwareUpdatable: types.BoolValue(dev.FirmwareUpdatable), Features: feats,
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_pending_devices ─────────────────────────────────────────────────────

type PendingDevicesDataSource struct{ client *client.Client }

func NewPendingDevicesDataSource() datasource.DataSource { return &PendingDevicesDataSource{} }
func (d *PendingDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pending_devices"
}
func (d *PendingDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists UniFi devices awaiting adoption.",
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"mac_address": schema.StringAttribute{Computed: true},
				"ip_address":  schema.StringAttribute{Computed: true},
				"model":       schema.StringAttribute{Computed: true},
				"state":       schema.StringAttribute{Computed: true},
				"supported":   schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *PendingDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type pendingDevicesModel struct {
	Devices []pendingDeviceModel `tfsdk:"devices"`
}
type pendingDeviceModel struct {
	MACAddress types.String `tfsdk:"mac_address"`
	IPAddress  types.String `tfsdk:"ip_address"`
	Model      types.String `tfsdk:"model"`
	State      types.String `tfsdk:"state"`
	Supported  types.Bool   `tfsdk:"supported"`
}

func (d *PendingDevicesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	devs, err := d.client.ListPendingDevices()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list pending devices", err.Error())
		return
	}
	state := pendingDevicesModel{}
	for _, dev := range devs {
		state.Devices = append(state.Devices, pendingDeviceModel{
			MACAddress: types.StringValue(dev.MACAddress), IPAddress: types.StringValue(dev.IPAddress),
			Model: types.StringValue(dev.Model), State: types.StringValue(dev.State),
			Supported: types.BoolValue(dev.Supported),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_clients ─────────────────────────────────────────────────────────────

type ClientsDataSource struct{ client *client.Client }

func NewClientsDataSource() datasource.DataSource { return &ClientsDataSource{} }
func (d *ClientsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clients"
}
func (d *ClientsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all connected network clients.",
		Attributes: map[string]schema.Attribute{
			"clients": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":               schema.StringAttribute{Computed: true},
				"type":             schema.StringAttribute{Computed: true},
				"name":             schema.StringAttribute{Computed: true},
				"mac_address":      schema.StringAttribute{Computed: true},
				"ip_address":       schema.StringAttribute{Computed: true},
				"connected_at":     schema.StringAttribute{Computed: true},
				"uplink_device_id": schema.StringAttribute{Computed: true},
			}}},
		},
	}
}
func (d *ClientsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type clientsModel struct {
	Clients []networkClientModel `tfsdk:"clients"`
}
type networkClientModel struct {
	ID             types.String `tfsdk:"id"`
	Type           types.String `tfsdk:"type"`
	Name           types.String `tfsdk:"name"`
	MACAddress     types.String `tfsdk:"mac_address"`
	IPAddress      types.String `tfsdk:"ip_address"`
	ConnectedAt    types.String `tfsdk:"connected_at"`
	UplinkDeviceID types.String `tfsdk:"uplink_device_id"`
}

func (d *ClientsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	clients, err := d.client.ListClients()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list clients", err.Error())
		return
	}
	state := clientsModel{}
	for _, c := range clients {
		state.Clients = append(state.Clients, networkClientModel{
			ID: types.StringValue(c.ID), Type: types.StringValue(c.Type),
			Name: types.StringValue(c.Name), MACAddress: types.StringValue(c.MACAddress),
			IPAddress: types.StringValue(c.IPAddress), ConnectedAt: types.StringValue(c.ConnectedAt),
			UplinkDeviceID: types.StringValue(c.UplinkDeviceID),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_networks ────────────────────────────────────────────────────────────

type NetworksDataSource struct{ client *client.Client }

func NewNetworksDataSource() datasource.DataSource { return &NetworksDataSource{} }
func (d *NetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks"
}
func (d *NetworksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all networks (VLANs) on the site.",
		Attributes: map[string]schema.Attribute{
			"networks": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":         schema.StringAttribute{Computed: true},
				"name":       schema.StringAttribute{Computed: true},
				"management": schema.StringAttribute{Computed: true},
				"enabled":    schema.BoolAttribute{Computed: true},
				"vlan_id":    schema.Int64Attribute{Computed: true},
				"default":    schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *NetworksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type networksModel struct {
	Networks []networkModel `tfsdk:"networks"`
}
type networkModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Management types.String `tfsdk:"management"`
	Enabled    types.Bool   `tfsdk:"enabled"`
	VLANId     types.Int64  `tfsdk:"vlan_id"`
	Default    types.Bool   `tfsdk:"default"`
}

func (d *NetworksDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	nets, err := d.client.ListNetworks()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list networks", err.Error())
		return
	}
	state := networksModel{}
	for _, n := range nets {
		state.Networks = append(state.Networks, networkModel{
			ID: types.StringValue(n.ID), Name: types.StringValue(n.Name),
			Management: types.StringValue(n.Management), Enabled: types.BoolValue(n.Enabled),
			VLANId: types.Int64Value(int64(n.VLANId)), Default: types.BoolValue(n.Default),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_wifi_broadcasts ─────────────────────────────────────────────────────

type WiFiBroadcastsDataSource struct{ client *client.Client }

func NewWiFiBroadcastsDataSource() datasource.DataSource { return &WiFiBroadcastsDataSource{} }
func (d *WiFiBroadcastsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wifi_broadcasts"
}
func (d *WiFiBroadcastsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all WiFi broadcasts (SSIDs) on the site.",
		Attributes: map[string]schema.Attribute{
			"wifi_broadcasts": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":           schema.StringAttribute{Computed: true},
				"type":         schema.StringAttribute{Computed: true},
				"name":         schema.StringAttribute{Computed: true},
				"enabled":      schema.BoolAttribute{Computed: true},
				"network_type": schema.StringAttribute{Computed: true},
				"security":     schema.StringAttribute{Computed: true},
			}}},
		},
	}
}
func (d *WiFiBroadcastsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type wifiBroadcastsModel struct {
	WiFiBroadcasts []wifiBroadcastModel `tfsdk:"wifi_broadcasts"`
}
type wifiBroadcastModel struct {
	ID          types.String `tfsdk:"id"`
	Type        types.String `tfsdk:"type"`
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	NetworkType types.String `tfsdk:"network_type"`
	Security    types.String `tfsdk:"security"`
}

func (d *WiFiBroadcastsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	wifis, err := d.client.ListWiFiBroadcasts()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list WiFi broadcasts", err.Error())
		return
	}
	state := wifiBroadcastsModel{}
	for _, w := range wifis {
		networkType := ""
		if w.Network != nil {
			networkType = w.Network.Type
		}
		security := ""
		if w.Security != nil {
			security = w.Security.Type
		}
		state.WiFiBroadcasts = append(state.WiFiBroadcasts, wifiBroadcastModel{
			ID: types.StringValue(w.ID), Type: types.StringValue(w.Type),
			Name: types.StringValue(w.Name), Enabled: types.BoolValue(w.Enabled),
			NetworkType: types.StringValue(networkType), Security: types.StringValue(security),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_firewall_zones ──────────────────────────────────────────────────────

type FirewallZonesDataSource struct{ client *client.Client }

func NewFirewallZonesDataSource() datasource.DataSource { return &FirewallZonesDataSource{} }
func (d *FirewallZonesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_zones"
}
func (d *FirewallZonesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all firewall zones on the site.",
		Attributes: map[string]schema.Attribute{
			"zones": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":          schema.StringAttribute{Computed: true},
				"name":        schema.StringAttribute{Computed: true},
				"network_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType},
			}}},
		},
	}
}
func (d *FirewallZonesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type firewallZonesModel struct {
	Zones []firewallZoneModel `tfsdk:"zones"`
}
type firewallZoneModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	NetworkIDs types.List   `tfsdk:"network_ids"`
}

func (d *FirewallZonesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	zones, err := d.client.ListFirewallZones()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list firewall zones", err.Error())
		return
	}
	state := firewallZonesModel{}
	for _, z := range zones {
		nids, diag := types.ListValueFrom(ctx, types.StringType, z.NetworkIDs)
		resp.Diagnostics.Append(diag...)
		state.Zones = append(state.Zones, firewallZoneModel{
			ID: types.StringValue(z.ID), Name: types.StringValue(z.Name), NetworkIDs: nids,
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_firewall_policies ───────────────────────────────────────────────────

type FirewallPoliciesDataSource struct{ client *client.Client }

func NewFirewallPoliciesDataSource() datasource.DataSource { return &FirewallPoliciesDataSource{} }
func (d *FirewallPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policies"
}
func (d *FirewallPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all firewall policies on the site.",
		Attributes: map[string]schema.Attribute{
			"policies": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":          schema.StringAttribute{Computed: true},
				"name":        schema.StringAttribute{Computed: true},
				"enabled":     schema.BoolAttribute{Computed: true},
				"action":      schema.StringAttribute{Computed: true},
				"description": schema.StringAttribute{Computed: true},
				"index":       schema.Int64Attribute{Computed: true},
			}}},
		},
	}
}
func (d *FirewallPoliciesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type firewallPoliciesModel struct {
	Policies []firewallPolicyModel `tfsdk:"policies"`
}
type firewallPolicyModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Action      types.String `tfsdk:"action"`
	Description types.String `tfsdk:"description"`
	Index       types.Int64  `tfsdk:"index"`
}

func (d *FirewallPoliciesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	policies, err := d.client.ListFirewallPolicies()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list firewall policies", err.Error())
		return
	}
	state := firewallPoliciesModel{}
	for _, p := range policies {
		state.Policies = append(state.Policies, firewallPolicyModel{
			ID: types.StringValue(p.ID), Name: types.StringValue(p.Name), Enabled: types.BoolValue(p.Enabled),
			Action: types.StringValue(p.Action), Description: types.StringValue(p.Description),
			Index: types.Int64Value(int64(p.Index)),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_acl_rules ───────────────────────────────────────────────────────────

type ACLRulesDataSource struct{ client *client.Client }

func NewACLRulesDataSource() datasource.DataSource { return &ACLRulesDataSource{} }
func (d *ACLRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_rules"
}
func (d *ACLRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all ACL rules on the site.",
		Attributes: map[string]schema.Attribute{
			"rules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":          schema.StringAttribute{Computed: true},
				"type":        schema.StringAttribute{Computed: true},
				"name":        schema.StringAttribute{Computed: true},
				"enabled":     schema.BoolAttribute{Computed: true},
				"action":      schema.StringAttribute{Computed: true},
				"description": schema.StringAttribute{Computed: true},
				"index":       schema.Int64Attribute{Computed: true},
			}}},
		},
	}
}
func (d *ACLRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type aclRulesModel struct {
	Rules []aclRuleModel `tfsdk:"rules"`
}
type aclRuleModel struct {
	ID          types.String `tfsdk:"id"`
	Type        types.String `tfsdk:"type"`
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Action      types.String `tfsdk:"action"`
	Description types.String `tfsdk:"description"`
	Index       types.Int64  `tfsdk:"index"`
}

func (d *ACLRulesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	rules, err := d.client.ListACLRules()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list ACL rules", err.Error())
		return
	}
	state := aclRulesModel{}
	for _, r := range rules {
		state.Rules = append(state.Rules, aclRuleModel{
			ID: types.StringValue(r.ID), Type: types.StringValue(r.Type),
			Name: types.StringValue(r.Name), Enabled: types.BoolValue(r.Enabled),
			Action: types.StringValue(r.Action), Description: types.StringValue(r.Description),
			Index: types.Int64Value(int64(r.Index)),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_dns_policies ────────────────────────────────────────────────────────

type DNSPoliciesDataSource struct{ client *client.Client }

func NewDNSPoliciesDataSource() datasource.DataSource { return &DNSPoliciesDataSource{} }
func (d *DNSPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_policies"
}
func (d *DNSPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all DNS policies on the site.",
		Attributes: map[string]schema.Attribute{
			"policies": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":      schema.StringAttribute{Computed: true},
				"domain":  schema.StringAttribute{Computed: true},
				"enabled": schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *DNSPoliciesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type dnsPoliciesModel struct {
	Policies []dnsPolicyModel `tfsdk:"policies"`
}
type dnsPolicyModel struct {
	ID      types.String `tfsdk:"id"`
	Domain  types.String `tfsdk:"domain"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (d *DNSPoliciesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	policies, err := d.client.ListDNSPolicies()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list DNS policies", err.Error())
		return
	}
	state := dnsPoliciesModel{}
	for _, p := range policies {
		state.Policies = append(state.Policies, dnsPolicyModel{
			ID: types.StringValue(p.ID), Domain: types.StringValue(p.Domain), Enabled: types.BoolValue(p.Enabled),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_vpn_servers ─────────────────────────────────────────────────────────

type VPNServersDataSource struct{ client *client.Client }

func NewVPNServersDataSource() datasource.DataSource { return &VPNServersDataSource{} }
func (d *VPNServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_servers"
}
func (d *VPNServersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all VPN servers on the site.",
		Attributes: map[string]schema.Attribute{
			"vpn_servers": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":      schema.StringAttribute{Computed: true},
				"type":    schema.StringAttribute{Computed: true},
				"name":    schema.StringAttribute{Computed: true},
				"enabled": schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *VPNServersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type vpnServersModel struct {
	VPNServers []vpnServerModel `tfsdk:"vpn_servers"`
}
type vpnServerModel struct {
	ID      types.String `tfsdk:"id"`
	Type    types.String `tfsdk:"type"`
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (d *VPNServersDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	servers, err := d.client.ListVPNServers()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list VPN servers", err.Error())
		return
	}
	state := vpnServersModel{}
	for _, s := range servers {
		state.VPNServers = append(state.VPNServers, vpnServerModel{
			ID: types.StringValue(s.ID), Type: types.StringValue(s.Type),
			Name: types.StringValue(s.Name), Enabled: types.BoolValue(s.Enabled),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_wan_interfaces ──────────────────────────────────────────────────────

type WANInterfacesDataSource struct{ client *client.Client }

func NewWANInterfacesDataSource() datasource.DataSource { return &WANInterfacesDataSource{} }
func (d *WANInterfacesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wan_interfaces"
}
func (d *WANInterfacesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all WAN interfaces on the site gateway.",
		Attributes: map[string]schema.Attribute{
			"interfaces": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":   schema.StringAttribute{Computed: true},
				"name": schema.StringAttribute{Computed: true},
			}}},
		},
	}
}
func (d *WANInterfacesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type wanInterfacesModel struct {
	Interfaces []wanInterfaceModel `tfsdk:"interfaces"`
}
type wanInterfaceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *WANInterfacesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	ifaces, err := d.client.ListWANInterfaces()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list WAN interfaces", err.Error())
		return
	}
	state := wanInterfacesModel{}
	for _, i := range ifaces {
		state.Interfaces = append(state.Interfaces, wanInterfaceModel{
			ID: types.StringValue(i.ID), Name: types.StringValue(i.Name),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_device_tags ─────────────────────────────────────────────────────────

type DeviceTagsDataSource struct{ client *client.Client }

func NewDeviceTagsDataSource() datasource.DataSource { return &DeviceTagsDataSource{} }
func (d *DeviceTagsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_tags"
}
func (d *DeviceTagsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all device tags on the site.",
		Attributes: map[string]schema.Attribute{
			"tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":         schema.StringAttribute{Computed: true},
				"name":       schema.StringAttribute{Computed: true},
				"device_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType},
			}}},
		},
	}
}
func (d *DeviceTagsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	d.client = c
}

type deviceTagsModel struct {
	Tags []deviceTagModel `tfsdk:"tags"`
}
type deviceTagModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	DeviceIDs types.List   `tfsdk:"device_ids"`
}

func (d *DeviceTagsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	tags, err := d.client.ListDeviceTags()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list device tags", err.Error())
		return
	}
	state := deviceTagsModel{}
	for _, t := range tags {
		devs, diag := types.ListValueFrom(ctx, types.StringType, t.DeviceIDs)
		resp.Diagnostics.Append(diag...)
		state.Tags = append(state.Tags, deviceTagModel{
			ID: types.StringValue(t.ID), Name: types.StringValue(t.Name), DeviceIDs: devs,
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
