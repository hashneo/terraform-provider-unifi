package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

// ── unifi_cameras ─────────────────────────────────────────────────────────────

type CamerasDataSource struct{ client *client.Client }

func NewCamerasDataSource() datasource.DataSource { return &CamerasDataSource{} }
func (d *CamerasDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cameras"
}
func (d *CamerasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi Protect cameras.",
		Attributes: map[string]schema.Attribute{
			"cameras": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":             schema.StringAttribute{Computed: true},
				"model_key":      schema.StringAttribute{Computed: true},
				"name":           schema.StringAttribute{Computed: true},
				"mac":            schema.StringAttribute{Computed: true},
				"state":          schema.StringAttribute{Computed: true},
				"is_mic_enabled": schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *CamerasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type camerasModel struct {
	Cameras []cameraModel `tfsdk:"cameras"`
}
type cameraModel struct {
	ID           types.String `tfsdk:"id"`
	ModelKey     types.String `tfsdk:"model_key"`
	Name         types.String `tfsdk:"name"`
	MAC          types.String `tfsdk:"mac"`
	State        types.String `tfsdk:"state"`
	IsMicEnabled types.Bool   `tfsdk:"is_mic_enabled"`
}

func (d *CamerasDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	cameras, err := d.client.ListCameras()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list cameras", err.Error())
		return
	}
	state := camerasModel{}
	for _, c := range cameras {
		state.Cameras = append(state.Cameras, cameraModel{
			ID: types.StringValue(c.ID), ModelKey: types.StringValue(c.ModelKey),
			Name: types.StringValue(c.Name), MAC: types.StringValue(c.MAC),
			State: types.StringValue(c.State), IsMicEnabled: types.BoolValue(c.IsMicEnabled),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_nvr ─────────────────────────────────────────────────────────────────

type NVRDataSource struct{ client *client.Client }

func NewNVRDataSource() datasource.DataSource { return &NVRDataSource{} }
func (d *NVRDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nvr"
}
func (d *NVRDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads the UniFi Protect NVR details.",
		Attributes: map[string]schema.Attribute{
			"id":        schema.StringAttribute{Computed: true},
			"model_key": schema.StringAttribute{Computed: true},
			"name":      schema.StringAttribute{Computed: true},
		},
	}
}
func (d *NVRDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type nvrModel struct {
	ID       types.String `tfsdk:"id"`
	ModelKey types.String `tfsdk:"model_key"`
	Name     types.String `tfsdk:"name"`
}

func (d *NVRDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	nvr, err := d.client.GetNVR()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get NVR", err.Error())
		return
	}
	if nvr == nil {
		resp.Diagnostics.Append(resp.State.Set(ctx, &nvrModel{
			ID: types.StringValue(""), ModelKey: types.StringValue(""), Name: types.StringValue(""),
		})...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &nvrModel{
		ID: types.StringValue(nvr.ID), ModelKey: types.StringValue(nvr.ModelKey), Name: types.StringValue(nvr.Name),
	})...)
}

// ── unifi_sensors ─────────────────────────────────────────────────────────────

type SensorsDataSource struct{ client *client.Client }

func NewSensorsDataSource() datasource.DataSource { return &SensorsDataSource{} }
func (d *SensorsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensors"
}
func (d *SensorsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi Protect sensors.",
		Attributes: map[string]schema.Attribute{
			"sensors": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":                 schema.StringAttribute{Computed: true},
				"model_key":          schema.StringAttribute{Computed: true},
				"name":               schema.StringAttribute{Computed: true},
				"mac":                schema.StringAttribute{Computed: true},
				"state":              schema.StringAttribute{Computed: true},
				"is_opened":          schema.BoolAttribute{Computed: true},
				"is_motion_detected": schema.BoolAttribute{Computed: true},
				"battery_percentage": schema.Int64Attribute{Computed: true},
				"battery_is_low":     schema.BoolAttribute{Computed: true},
				"battery_is_charging": schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *SensorsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type sensorsModel struct {
	Sensors []sensorModel `tfsdk:"sensors"`
}
type sensorModel struct {
	ID               types.String `tfsdk:"id"`
	ModelKey         types.String `tfsdk:"model_key"`
	Name             types.String `tfsdk:"name"`
	MAC              types.String `tfsdk:"mac"`
	State            types.String `tfsdk:"state"`
	IsOpened         types.Bool   `tfsdk:"is_opened"`
	IsMotionDetected types.Bool   `tfsdk:"is_motion_detected"`
	BatteryPct       types.Int64  `tfsdk:"battery_percentage"`
	BatteryIsLow     types.Bool   `tfsdk:"battery_is_low"`
	BatteryIsCharging types.Bool  `tfsdk:"battery_is_charging"`
}

func (d *SensorsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	sensors, err := d.client.ListSensors()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list sensors", err.Error())
		return
	}
	state := sensorsModel{}
	for _, s := range sensors {
		m := sensorModel{
			ID: types.StringValue(s.ID), ModelKey: types.StringValue(s.ModelKey),
			Name: types.StringValue(s.Name), MAC: types.StringValue(s.MAC),
			State: types.StringValue(s.State), IsMotionDetected: types.BoolValue(s.IsMotion),
		}
		if s.IsOpened != nil {
			m.IsOpened = types.BoolValue(*s.IsOpened)
		} else {
			m.IsOpened = types.BoolValue(false)
		}
		if s.Battery != nil {
			m.BatteryPct = types.Int64Value(int64(s.Battery.Percentage))
			m.BatteryIsLow = types.BoolValue(s.Battery.IsLow)
			m.BatteryIsCharging = types.BoolValue(s.Battery.IsCharging)
		} else {
			m.BatteryPct = types.Int64Value(0)
			m.BatteryIsLow = types.BoolValue(false)
			m.BatteryIsCharging = types.BoolValue(false)
		}
		state.Sensors = append(state.Sensors, m)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_lights ──────────────────────────────────────────────────────────────

type LightsDataSource struct{ client *client.Client }

func NewLightsDataSource() datasource.DataSource { return &LightsDataSource{} }
func (d *LightsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lights"
}
func (d *LightsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi Protect lights.",
		Attributes: map[string]schema.Attribute{
			"lights": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":                     schema.StringAttribute{Computed: true},
				"model_key":              schema.StringAttribute{Computed: true},
				"name":                   schema.StringAttribute{Computed: true},
				"mac":                    schema.StringAttribute{Computed: true},
				"state":                  schema.StringAttribute{Computed: true},
				"is_light_on":            schema.BoolAttribute{Computed: true},
				"is_pir_motion_detected": schema.BoolAttribute{Computed: true},
			}}},
		},
	}
}
func (d *LightsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type lightsModel struct {
	Lights []lightModel `tfsdk:"lights"`
}
type lightModel struct {
	ID                  types.String `tfsdk:"id"`
	ModelKey            types.String `tfsdk:"model_key"`
	Name                types.String `tfsdk:"name"`
	MAC                 types.String `tfsdk:"mac"`
	State               types.String `tfsdk:"state"`
	IsLightOn           types.Bool   `tfsdk:"is_light_on"`
	IsPirMotionDetected types.Bool   `tfsdk:"is_pir_motion_detected"`
}

func (d *LightsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	lights, err := d.client.ListLights()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list lights", err.Error())
		return
	}
	state := lightsModel{}
	for _, l := range lights {
		state.Lights = append(state.Lights, lightModel{
			ID: types.StringValue(l.ID), ModelKey: types.StringValue(l.ModelKey),
			Name: types.StringValue(l.Name), MAC: types.StringValue(l.MAC),
			State: types.StringValue(l.State), IsLightOn: types.BoolValue(l.IsLightOn),
			IsPirMotionDetected: types.BoolValue(l.IsPirMotion),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ── unifi_viewers ─────────────────────────────────────────────────────────────

type ViewersDataSource struct{ client *client.Client }

func NewViewersDataSource() datasource.DataSource { return &ViewersDataSource{} }
func (d *ViewersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_viewers"
}
func (d *ViewersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists all UniFi Protect viewers.",
		Attributes: map[string]schema.Attribute{
			"viewers": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"id":        schema.StringAttribute{Computed: true},
				"model_key": schema.StringAttribute{Computed: true},
				"name":      schema.StringAttribute{Computed: true},
				"mac":       schema.StringAttribute{Computed: true},
				"state":     schema.StringAttribute{Computed: true},
			}}},
		},
	}
}
func (d *ViewersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type viewersModel struct {
	Viewers []viewerModel `tfsdk:"viewers"`
}
type viewerModel struct {
	ID       types.String `tfsdk:"id"`
	ModelKey types.String `tfsdk:"model_key"`
	Name     types.String `tfsdk:"name"`
	MAC      types.String `tfsdk:"mac"`
	State    types.String `tfsdk:"state"`
}

func (d *ViewersDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	viewers, err := d.client.ListViewers()
	if err != nil {
		resp.Diagnostics.AddError("Failed to list viewers", err.Error())
		return
	}
	state := viewersModel{}
	for _, v := range viewers {
		state.Viewers = append(state.Viewers, viewerModel{
			ID: types.StringValue(v.ID), ModelKey: types.StringValue(v.ModelKey),
			Name: types.StringValue(v.Name), MAC: types.StringValue(v.MAC),
			State: types.StringValue(v.State),
		})
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
