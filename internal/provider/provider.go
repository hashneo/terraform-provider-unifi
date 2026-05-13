package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
	"github.com/steventaylor/terraform-provider-unifi/internal/provider/datasources"
	"github.com/steventaylor/terraform-provider-unifi/internal/provider/resources"
)

var _ provider.Provider = &UnifiProvider{}

type UnifiProvider struct {
	version string
}

type UnifiProviderModel struct {
	CloudAPIKey   types.String `tfsdk:"cloud_api_key"`
	ControllerURL types.String `tfsdk:"controller_url"`
	LocalAPIKey   types.String `tfsdk:"local_api_key"`
	SiteID        types.String `tfsdk:"site_id"`
	SSLInsecure   types.Bool   `tfsdk:"ssl_insecure"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UnifiProvider{version: version}
	}
}

func (p *UnifiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "unifi"
	resp.Version = p.version
}

func (p *UnifiProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage and inventory Ubiquiti UniFi infrastructure via the Site Manager, Network, and Protect APIs.",
		Attributes: map[string]schema.Attribute{
			"cloud_api_key": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "API key for the Ubiquiti Site Manager cloud API (api.ui.com). Can also be set via UNIFI_CLOUD_API_KEY environment variable.",
			},
			"controller_url": schema.StringAttribute{
				Optional:    true,
				Description: "Base URL of the local UniFi controller (e.g. https://192.168.1.1). Can also be set via UNIFI_CONTROLLER_URL environment variable.",
			},
			"local_api_key": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "API key for the local controller Network/Protect APIs. Can also be set via UNIFI_LOCAL_API_KEY environment variable.",
			},
			"site_id": schema.StringAttribute{
				Optional:    true,
				Description: "UniFi site ID (default: 'default'). Can also be set via UNIFI_SITE_ID environment variable.",
			},
			"ssl_insecure": schema.BoolAttribute{
				Optional:    true,
				Description: "Skip TLS certificate verification for local controller connections (useful for self-signed certs).",
			},
		},
	}
}

func (p *UnifiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config UnifiProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cloudAPIKey := os.Getenv("UNIFI_CLOUD_API_KEY")
	if !config.CloudAPIKey.IsNull() && !config.CloudAPIKey.IsUnknown() {
		cloudAPIKey = config.CloudAPIKey.ValueString()
	}

	controllerURL := os.Getenv("UNIFI_CONTROLLER_URL")
	if !config.ControllerURL.IsNull() && !config.ControllerURL.IsUnknown() {
		controllerURL = config.ControllerURL.ValueString()
	}

	localAPIKey := os.Getenv("UNIFI_LOCAL_API_KEY")
	if !config.LocalAPIKey.IsNull() && !config.LocalAPIKey.IsUnknown() {
		localAPIKey = config.LocalAPIKey.ValueString()
	}

	siteID := os.Getenv("UNIFI_SITE_ID")
	if siteID == "" {
		siteID = "default"
	}
	if !config.SiteID.IsNull() && !config.SiteID.IsUnknown() {
		siteID = config.SiteID.ValueString()
	}

	sslInsecure := false
	if !config.SSLInsecure.IsNull() && !config.SSLInsecure.IsUnknown() {
		sslInsecure = config.SSLInsecure.ValueBool()
	}

	c := client.NewClient(cloudAPIKey, controllerURL, localAPIKey, siteID, sslInsecure)
	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *UnifiProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Site Manager (cloud)
		datasources.NewHostsDataSource,
		datasources.NewSitesDataSource,
		datasources.NewCloudDevicesDataSource,
		datasources.NewSDWANConfigsDataSource,
		// Network (local)
		datasources.NewAdoptedDevicesDataSource,
		datasources.NewPendingDevicesDataSource,
		datasources.NewClientsDataSource,
		datasources.NewNetworksDataSource,
		datasources.NewWiFiBroadcastsDataSource,
		datasources.NewFirewallZonesDataSource,
		datasources.NewFirewallPoliciesDataSource,
		datasources.NewACLRulesDataSource,
		datasources.NewDNSPoliciesDataSource,
		datasources.NewVPNServersDataSource,
		datasources.NewWANInterfacesDataSource,
		datasources.NewDeviceTagsDataSource,
		// Protect (local)
		datasources.NewCamerasDataSource,
		datasources.NewNVRDataSource,
		datasources.NewSensorsDataSource,
		datasources.NewLightsDataSource,
		datasources.NewViewersDataSource,
	}
}

func (p *UnifiProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewNetworkResource,
		resources.NewWiFiBroadcastResource,
		resources.NewFirewallPolicyResource,
		resources.NewACLRuleResource,
	}
}
