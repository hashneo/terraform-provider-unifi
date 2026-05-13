package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/steventaylor/terraform-provider-unifi/internal/client"
)

var _ resource.Resource = &ACLRuleResource{}

type ACLRuleResource struct{ client *client.Client }

func NewACLRuleResource() resource.Resource { return &ACLRuleResource{} }

func (r *ACLRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_rule"
}

func (r *ACLRuleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi ACL rule.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"enabled":     schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
		"action":      schema.StringAttribute{Required: true, Description: "Action: 'allow' or 'block'."},
		"description": schema.StringAttribute{Optional: true, Computed: true},
			"index":       schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(0)},
		},
	}
}

func (r *ACLRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type aclRuleResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Action      types.String `tfsdk:"action"`
	Description types.String `tfsdk:"description"`
	Index       types.Int64  `tfsdk:"index"`
}

func aclFromPlan(plan aclRuleResourceModel) client.ACLRule {
	return client.ACLRule{
		Name: plan.Name.ValueString(), Enabled: plan.Enabled.ValueBool(),
		Action: plan.Action.ValueString(),
		Description: plan.Description.ValueString(), Index: int(plan.Index.ValueInt64()),
	}
}

func (r *ACLRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan aclRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	rule, err := r.client.CreateACLRule(aclFromPlan(plan))
	if err != nil {
		resp.Diagnostics.AddError("Failed to create ACL rule", err.Error())
		return
	}
	plan.ID = types.StringValue(rule.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ACLRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state aclRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	rule, err := r.client.GetACLRule(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to read ACL rule", err.Error())
		return
	}
	state.Name = types.StringValue(rule.Name)
	state.Enabled = types.BoolValue(rule.Enabled)
	state.Action = types.StringValue(rule.Action)
	state.Description = types.StringValue(rule.Description)
	state.Index = types.Int64Value(int64(rule.Index))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ACLRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan aclRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.UpdateACLRule(plan.ID.ValueString(), aclFromPlan(plan))
	if err != nil {
		resp.Diagnostics.AddError("Failed to update ACL rule", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ACLRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state aclRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteACLRule(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Failed to delete ACL rule", err.Error())
	}
}
