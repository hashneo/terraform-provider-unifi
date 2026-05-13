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

var _ resource.Resource = &NetworkResource{}

type NetworkResource struct{ client *client.Client }

func NewNetworkResource() resource.Resource { return &NetworkResource{} }

func (r *NetworkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (r *NetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi network (VLAN/subnet).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Network ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":    schema.StringAttribute{Required: true, Description: "Network name."},
			"vlan_id": schema.Int64Attribute{Optional: true, Computed: true, Description: "VLAN ID (1–4094)."},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
		},
	}
}

func (r *NetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type networkResourceModel struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	VLANId  types.Int64  `tfsdk:"vlan_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan networkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	n, err := r.client.CreateNetwork(client.Network{
		Name: plan.Name.ValueString(), VLANId: int(plan.VLANId.ValueInt64()),
		Enabled: plan.Enabled.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to create network", err.Error())
		return
	}
	plan.ID = types.StringValue(n.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state networkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	n, err := r.client.GetNetwork(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read network", err.Error())
		return
	}
	state.Name = types.StringValue(n.Name)
	state.VLANId = types.Int64Value(int64(n.VLANId))
	state.Enabled = types.BoolValue(n.Enabled)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan networkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	n, err := r.client.UpdateNetwork(plan.ID.ValueString(), client.Network{
		Name: plan.Name.ValueString(), VLANId: int(plan.VLANId.ValueInt64()),
		Enabled: plan.Enabled.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to update network", err.Error())
		return
	}
	plan.Name = types.StringValue(n.Name)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state networkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteNetwork(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to delete network", err.Error())
	}
}
