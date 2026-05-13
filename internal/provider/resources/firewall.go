package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

var _ resource.Resource                = &FirewallPolicyResource{}
var _ resource.ResourceWithImportState = &FirewallPolicyResource{}

type FirewallPolicyResource struct{ client *client.Client }

func NewFirewallPolicyResource() resource.Resource { return &FirewallPolicyResource{} }

func (r *FirewallPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policy"
}

func (r *FirewallPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi firewall policy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                schema.StringAttribute{Required: true},
			"enabled":             schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
		"action":      schema.StringAttribute{Required: true, Description: "Action: 'allow' or 'block'."},
		"description": schema.StringAttribute{Optional: true, Computed: true},
			"index":               schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(0)},
		},
	}
}

func (r *FirewallPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type firewallPolicyResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Action      types.String `tfsdk:"action"`
	Description types.String `tfsdk:"description"`
	Index       types.Int64  `tfsdk:"index"`
}

func policyFromPlan(plan firewallPolicyResourceModel) client.FirewallPolicy {
	return client.FirewallPolicy{
		Name: plan.Name.ValueString(), Enabled: plan.Enabled.ValueBool(),
		Action:      plan.Action.ValueString(),
		Description: plan.Description.ValueString(), Index: int(plan.Index.ValueInt64()),
	}
}

func (r *FirewallPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan firewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	p, err := r.client.CreateFirewallPolicy(policyFromPlan(plan))
	if err != nil {
		resp.Diagnostics.AddError("Failed to create firewall policy", err.Error())
		return
	}
	plan.ID = types.StringValue(p.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FirewallPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state firewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	p, err := r.client.GetFirewallPolicy(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read firewall policy", err.Error())
		return
	}
	state.Name = types.StringValue(p.Name)
	state.Enabled = types.BoolValue(p.Enabled)
	state.Action = types.StringValue(p.Action)
	state.Description = types.StringValue(p.Description)
	state.Index = types.Int64Value(int64(p.Index))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FirewallPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan firewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.UpdateFirewallPolicy(plan.ID.ValueString(), policyFromPlan(plan))
	if err != nil {
		resp.Diagnostics.AddError("Failed to update firewall policy", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FirewallPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state firewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteFirewallPolicy(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to delete firewall policy", err.Error())
	}
}

func (r *FirewallPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
