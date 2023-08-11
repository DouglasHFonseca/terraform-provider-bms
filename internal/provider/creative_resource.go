package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "go-client-creative-management"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &creativeResource{}
	_ resource.ResourceWithConfigure = &creativeResource{}
)

// NewCreativeResource is a helper function to simplify the provider implementation.
func NewCreativeResource() resource.Resource {
	return &creativeResource{}
}

// creativeResource is the resource implementation.
type creativeResource struct {
	client    *openapi.APIClient
	accountID string
}

// Configure adds the provider configured client to the resource.
func (r *creativeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*providerConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = config.client
	r.accountID = config.accountID
}

// Metadata returns the resource type name.
func (r *creativeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_creative"
}

// Schema defines the schema for the resource.
func (r *creativeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"creative_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"domain": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"banner": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"width": schema.Int64Attribute{
						Required: true,
					},
					"height": schema.Int64Attribute{
						Required: true,
					},
					"snippet": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"archived": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *creativeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan creativeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createFields := *openapi.NewCreativeCreateFields(plan.Domain.ValueString(), plan.Name.ValueString())

	createFields.Banner = &openapi.CreativeCreateFieldsAllOfBanner{
		Width:   float32(plan.Banner.Width.ValueInt64()),
		Height:  float32(plan.Banner.Height.ValueInt64()),
		Snippet: plan.Banner.Snippet.ValueString(),
	}

	if plan.Tags != nil {
		var tags []string
		for _, tag := range plan.Tags {
			tags = append(tags, tag.ValueString())
		}
		createFields.Tags = tags
	}

	creative, _, err := r.client.CreativesAPI.CreateCreative(ctx, r.accountID).CreativeCreateFields(createFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Creative",
			fmt.Sprintf("Error creating Creative: %s", err),
		)
		return
	}

	// Overwrite the terraform state with the newly created resource.
	plan = r.OverwriteCreativeModel(*creative)

	diags = resp.State.Set(ctx, plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Read refreshes the Terraform state with the latest data.
func (r *creativeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state creativeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	creative, response, err := r.client.CreativesAPI.GetCreative(ctx, r.accountID, state.CreativeId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Response [READ] Creative",
			fmt.Sprintf("Error reading Creative: %v", response),
		)

		resp.Diagnostics.AddError(
			"Error [READ] Creative",
			fmt.Sprintf("Error reading Creative: %s", err),
		)
		return
	}

	// Overwrite creative with refreshed state
	state = r.OverwriteCreativeModel(*creative)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *creativeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan creativeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	creativePatchFields := *openapi.NewCreativePatchFields()

	creativePatchFields.Name = plan.Name.ValueStringPointer()
	creativePatchFields.Banner = &openapi.BannerFields{
		Snippet: plan.Banner.Snippet.ValueString(),
	}
	for _, tag := range plan.Tags {
		creativePatchFields.Tags = append(creativePatchFields.Tags, tag.ValueString())
	}
	creativePatchFields.Archived = plan.Archived.ValueBoolPointer()
	creativePatchFields.Enabled = plan.Enabled.ValueBoolPointer()

	_, response, err := r.client.CreativesAPI.PatchCreative(ctx, r.accountID, plan.CreativeId.ValueString()).CreativePatchFields(creativePatchFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE] Creative",
			fmt.Sprintf("Error updating creative: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [UPDATE] Creative",
			fmt.Sprintf("Error updating Creative: %v", response),
		)
		return
	}

	creative, _, err := r.client.CreativesAPI.GetCreative(ctx, r.accountID, plan.CreativeId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE -> READ] Creative",
			fmt.Sprintf("Error reading Creative: %v\n", err),
		)
		return
	}

	plan = r.OverwriteCreativeModel(*creative)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *creativeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state creativeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.CreativesAPI.DeleteCreative(ctx, r.accountID, state.CreativeId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [DELETE] Creative",
			fmt.Sprintf("Error deleting Creative: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [DELETE] Creative",
			fmt.Sprintf("Error deleting Creative: %v", response),
		)
		return
	}
}

// OverwriteCreativeModel overwrites the creative model with the latest data.
func (r *creativeResource) OverwriteCreativeModel(creative openapi.Creative) creativeModel {
	plan := creativeModel{
		CreativeId: types.StringValue(creative.CreativeId),
		Status:     types.StringValue(creative.Status),
		AccountId:  types.StringValue(creative.AccountId),
		Enabled:    types.BoolValue(creative.Enabled),
		Archived:   types.BoolValue(creative.Archived),
		Name:       types.StringValue(creative.Name),
		Domain:     types.StringValue(creative.Domain),
	}

	plan.Tags = []types.String{}
	for _, tag := range creative.Tags {
		plan.Tags = append(plan.Tags,
			types.StringValue(tag),
		)
	}

	plan.Banner = bannerModel{
		Width:   types.Int64Value(int64(creative.Banner.Width)),
		Height:  types.Int64Value(int64(creative.Banner.Height)),
		Snippet: types.StringValue(creative.Banner.Snippet),
	}

	return plan
}
