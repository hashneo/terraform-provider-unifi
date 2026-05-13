package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

var _ resource.Resource = &WiFiBroadcastResource{}

type WiFiBroadcastResource struct{ client *client.Client }

func NewWiFiBroadcastResource() resource.Resource { return &WiFiBroadcastResource{} }

func (r *WiFiBroadcastResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wifi_broadcast"
}

func (r *WiFiBroadcastResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi WiFi broadcast (SSID).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":    schema.StringAttribute{Required: true, Description: "WiFi broadcast name."},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
		},
	}
}

func (r *WiFiBroadcastResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data type", fmt.Sprintf("expected *client.Client, got %T", req.ProviderData))
		return
	}
	r.client = c
}

type wifiBroadcastResourceModel struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (r *WiFiBroadcastResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan wifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	w, err := r.client.CreateWiFiBroadcast(client.WiFiBroadcast{
		Name: plan.Name.ValueString(), Enabled: plan.Enabled.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to create WiFi broadcast", err.Error())
		return
	}
	plan.ID = types.StringValue(w.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WiFiBroadcastResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state wifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	w, err := r.client.GetWiFiBroadcast(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read WiFi broadcast", err.Error())
		return
	}
	state.Name = types.StringValue(w.Name)
	state.Enabled = types.BoolValue(w.Enabled)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WiFiBroadcastResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan wifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.UpdateWiFiBroadcast(plan.ID.ValueString(), client.WiFiBroadcast{
		Name: plan.Name.ValueString(), Enabled: plan.Enabled.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to update WiFi broadcast", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WiFiBroadcastResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state wifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteWiFiBroadcast(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to delete WiFi broadcast", err.Error())
	}
}
