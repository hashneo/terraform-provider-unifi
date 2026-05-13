package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

// ── unifi_hosts ───────────────────────────────────────────────────────────────

type HostsDataSource struct{ client *client.Client }

func NewHostsDataSource() datasource.DataSource { return &HostsDataSource{} }

func (d *HostsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hosts"
}

func (d *HostsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi hosts registered to the cloud account.",
		Attributes: map[string]schema.Attribute{
			"hosts": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":               schema.StringAttribute{Computed: true},
						"type":             schema.StringAttribute{Computed: true},
						"hardware_id":      schema.StringAttribute{Computed: true},
						"is_blocked":       schema.BoolAttribute{Computed: true},
						"owner":            schema.BoolAttribute{Computed: true},
						"hostname":         schema.StringAttribute{Computed: true},
						"ip":               schema.StringAttribute{Computed: true},
						"name":             schema.StringAttribute{Computed: true},
						"firmware_version": schema.StringAttribute{Computed: true},
						"is_configured":    schema.BoolAttribute{Computed: true},
						"is_setup":         schema.BoolAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *HostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type hostsModel struct {
	Hosts []hostModel `tfsdk:"hosts"`
}

type hostModel struct {
	ID              types.String `tfsdk:"id"`
	Type            types.String `tfsdk:"type"`
	HardwareID      types.String `tfsdk:"hardware_id"`
	IsBlocked       types.Bool   `tfsdk:"is_blocked"`
	Owner           types.Bool   `tfsdk:"owner"`
	Hostname        types.String `tfsdk:"hostname"`
	IP              types.String `tfsdk:"ip"`
	Name            types.String `tfsdk:"name"`
	FirmwareVersion types.String `tfsdk:"firmware_version"`
	IsConfigured    types.Bool   `tfsdk:"is_configured"`
	IsSetup         types.Bool   `tfsdk:"is_setup"`
}

func (d *HostsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	hosts, err := d.client.ListHosts()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list hosts", err.Error())
		return
	}
	state := hostsModel{}
	for _, h := range hosts {
		m := hostModel{
			ID:         types.StringValue(h.ID),
			Type:       types.StringValue(h.Type),
			HardwareID: types.StringValue(h.HardwareID),
			IsBlocked:  types.BoolValue(h.IsBlocked),
			Owner:      types.BoolValue(h.Owner),
		}
		if h.ReportedState != nil {
			m.Hostname = types.StringValue(h.ReportedState.Hostname)
			m.IP = types.StringValue(h.ReportedState.IP)
			m.Name = types.StringValue(h.ReportedState.Name)
			m.FirmwareVersion = types.StringValue(h.ReportedState.FirmwareVersion)
			m.IsConfigured = types.BoolValue(h.ReportedState.IsConfigured)
			m.IsSetup = types.BoolValue(h.ReportedState.IsSetup)
		} else {
			m.Hostname = types.StringValue("")
			m.IP = types.StringValue("")
			m.Name = types.StringValue("")
			m.FirmwareVersion = types.StringValue("")
			m.IsConfigured = types.BoolValue(false)
			m.IsSetup = types.BoolValue(false)
		}
		state.Hosts = append(state.Hosts, m)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_sites ───────────────────────────────────────────────────────────────

type SitesDataSource struct{ client *client.Client }

func NewSitesDataSource() datasource.DataSource { return &SitesDataSource{} }

func (d *SitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

func (d *SitesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi sites registered to the cloud account.",
		Attributes: map[string]schema.Attribute{
			"sites": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_id":     schema.StringAttribute{Computed: true},
						"site_id":     schema.StringAttribute{Computed: true},
						"name":        schema.StringAttribute{Computed: true},
						"gateway_mac": schema.StringAttribute{Computed: true},
						"is_owner":    schema.BoolAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *SitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type sitesModel struct {
	Sites []siteModel `tfsdk:"sites"`
}

type siteModel struct {
	HostID     types.String `tfsdk:"host_id"`
	SiteID     types.String `tfsdk:"site_id"`
	Name       types.String `tfsdk:"name"`
	GatewayMAC types.String `tfsdk:"gateway_mac"`
	IsOwner    types.Bool   `tfsdk:"is_owner"`
}

func (d *SitesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	sites, err := d.client.ListSites()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list sites", err.Error())
		return
	}
	state := sitesModel{}
	for _, s := range sites {
		m := siteModel{
			HostID:  types.StringValue(s.HostID),
			SiteID:  types.StringValue(s.SiteID),
			IsOwner: types.BoolValue(s.IsOwner),
		}
		if s.MetaData != nil {
			m.Name = types.StringValue(s.MetaData.Name)
			m.GatewayMAC = types.StringValue(s.MetaData.GatewayMAC)
		} else {
			m.Name = types.StringValue("")
			m.GatewayMAC = types.StringValue("")
		}
		state.Sites = append(state.Sites, m)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_cloud_devices ───────────────────────────────────────────────────────

type CloudDevicesDataSource struct{ client *client.Client }

func NewCloudDevicesDataSource() datasource.DataSource { return &CloudDevicesDataSource{} }

func (d *CloudDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_devices"
}

func (d *CloudDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi devices visible from the cloud account (cross-site).",
		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":               schema.StringAttribute{Computed: true},
						"mac":              schema.StringAttribute{Computed: true},
						"model":            schema.StringAttribute{Computed: true},
						"name":             schema.StringAttribute{Computed: true},
						"ip":               schema.StringAttribute{Computed: true},
						"is_adopted":       schema.BoolAttribute{Computed: true},
						"is_connected":     schema.BoolAttribute{Computed: true},
						"firmware_version": schema.StringAttribute{Computed: true},
						"product_line":     schema.StringAttribute{Computed: true},
						"site_id":          schema.StringAttribute{Computed: true},
						"host_id":          schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *CloudDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type cloudDevicesModel struct {
	Devices []cloudDeviceModel `tfsdk:"devices"`
}

type cloudDeviceModel struct {
	ID              types.String `tfsdk:"id"`
	MAC             types.String `tfsdk:"mac"`
	Model           types.String `tfsdk:"model"`
	Name            types.String `tfsdk:"name"`
	IP              types.String `tfsdk:"ip"`
	IsAdopted       types.Bool   `tfsdk:"is_adopted"`
	IsConnected     types.Bool   `tfsdk:"is_connected"`
	FirmwareVersion types.String `tfsdk:"firmware_version"`
	ProductLine     types.String `tfsdk:"product_line"`
	SiteID          types.String `tfsdk:"site_id"`
	HostID          types.String `tfsdk:"host_id"`
}

func (d *CloudDevicesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	devices, err := d.client.ListCloudDevices()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list cloud devices", err.Error())
		return
	}
	state := cloudDevicesModel{}
	for _, dev := range devices {
		state.Devices = append(state.Devices, cloudDeviceModel{
			ID:              types.StringValue(dev.ID),
			MAC:             types.StringValue(dev.MAC),
			Model:           types.StringValue(dev.Model),
			Name:            types.StringValue(dev.Name),
			IP:              types.StringValue(dev.IP),
			IsAdopted:       types.BoolValue(dev.IsAdopted),
			IsConnected:     types.BoolValue(dev.IsConnected),
			FirmwareVersion: types.StringValue(dev.FirmwareVersion),
			ProductLine:     types.StringValue(dev.ProductLine),
			SiteID:          types.StringValue(dev.SiteID),
			HostID:          types.StringValue(dev.HostID),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_sdwan_configs ───────────────────────────────────────────────────────

type SDWANConfigsDataSource struct{ client *client.Client }

func NewSDWANConfigsDataSource() datasource.DataSource { return &SDWANConfigsDataSource{} }

func (d *SDWANConfigsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sdwan_configs"
}

func (d *SDWANConfigsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all SD-WAN configurations from the cloud account.",
		Attributes: map[string]schema.Attribute{
			"configs": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":          schema.StringAttribute{Computed: true},
						"name":        schema.StringAttribute{Computed: true},
						"description": schema.StringAttribute{Computed: true},
						"enabled":     schema.BoolAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *SDWANConfigsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type sdwanConfigsModel struct {
	Configs []sdwanConfigModel `tfsdk:"configs"`
}

type sdwanConfigModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Enabled     types.Bool   `tfsdk:"enabled"`
}

func (d *SDWANConfigsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	configs, err := d.client.ListSDWANConfigs()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list SD-WAN configs", err.Error())
		return
	}
	state := sdwanConfigsModel{}
	for _, cfg := range configs {
		state.Configs = append(state.Configs, sdwanConfigModel{
			ID:          types.StringValue(cfg.ID),
			Name:        types.StringValue(cfg.Name),
			Description: types.StringValue(cfg.Description),
			Enabled:     types.BoolValue(cfg.Enabled),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

