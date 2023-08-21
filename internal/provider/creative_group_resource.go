package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "go-client-creative-management"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &creativeGroupResource{}
	_ resource.ResourceWithConfigure = &creativeGroupResource{}
)

// creativeGroupModel maps creative group schema data.
type creativeGroupModel struct {
	CreativeGroupId types.String     `tfsdk:"creative_group_id"`
	AccountId       types.String     `tfsdk:"account_id"`
	Tags            []types.String   `tfsdk:"tags"`
	Archived        types.Bool       `tfsdk:"archived"`
	Domain          types.String     `tfsdk:"domain"`
	Name            types.String     `tfsdk:"name"`
	Spec            []specModel      `tfsdk:"spec"`
	Creatives       []creativesModel `tfsdk:"creative"`
}

// creativesModel maps creatives data.
type creativesModel struct {
	CreativeId types.String `tfsdk:"creative_id"`
	Weight     types.Int64  `tfsdk:"weight"`
}

// specModel maps spec data.
type specModel struct {
	Banner []bannerGroupModel `tfsdk:"banner"`
}

// bannerGroupModel maps banner data.
type bannerGroupModel struct {
	Width  types.Int64 `tfsdk:"width"`
	Height types.Int64 `tfsdk:"height"`
}

// NewCreativeGroupResource is a helper function to simplify the provider implementation.
func NewCreativeGroupResource() resource.Resource {
	return &creativeGroupResource{}
}

// creativeGroupResource is the resource implementation.
type creativeGroupResource struct {
	client    *openapi.APIClient
	accountID string
}

// Configure adds the provider configured clientCreative to the resource.
func (r *creativeGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = config.clientCreative
	r.accountID = config.accountID
}

// Metadata returns the resource type name.
func (r *creativeGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ads_creative_group"
}

// Schema defines the schema for the resource.
func (r *creativeGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"creative_group_id": schema.StringAttribute{
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
			"tags": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"archived": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
		Blocks: map[string]schema.Block{
			"spec": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"banner": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"width": schema.Int64Attribute{
										Required: true,
									},
									"height": schema.Int64Attribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"creative": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"creative_id": schema.StringAttribute{
							Required: true,
						},
						"weight": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.Between(0, 999),
							},
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *creativeGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan creativeGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var specFields openapi.CreativeSpec

	if plan.Spec != nil {
		if plan.Spec[0].Banner != nil {
			var bannerFields []*openapi.BannerSpec
			for _, banner := range plan.Spec[0].Banner {
				bannerFields = append(bannerFields, openapi.NewBannerSpec(float32(banner.Width.ValueInt64()), float32(banner.Height.ValueInt64())))
			}
			specFields.Banner = bannerFields[0]
		}
	}

	var creativesFields []openapi.WeightedCreative
	for _, creative := range plan.Creatives {
		creativesFields = append(creativesFields, *openapi.NewWeightedCreative(creative.CreativeId.ValueString(), float32(creative.Weight.ValueInt64())))
	}

	creativeGroupCreateFields := *openapi.NewCreativeGroupCreateFields(plan.Domain.ValueString(), plan.Name.ValueString(), specFields, creativesFields)

	if plan.Tags != nil {
		var tags []string
		for _, tag := range plan.Tags {
			tags = append(tags, tag.ValueString())
		}
		creativeGroupCreateFields.Tags = tags
	}

	creative, _, err := r.client.CreativeGroupsAPI.CreateCreativeGroup(ctx, r.accountID).CreativeGroupCreateFields(creativeGroupCreateFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Creative Group",
			fmt.Sprintf("Error creating Creative Group: %s", err),
		)
		return
	}

	// Overwrite the terraform state with the newly created resource.
	plan = r.OverwriteCreativeGroupModel(*creative)

	diags = resp.State.Set(ctx, plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Read refreshes the Terraform state with the latest data.
func (r *creativeGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state creativeGroupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	creative, response, err := r.client.CreativeGroupsAPI.GetCreativeGroup(ctx, r.accountID, state.CreativeGroupId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Response [READ] Creative Group",
			fmt.Sprintf("Error reading Creative Group: %v", response),
		)

		resp.Diagnostics.AddError(
			"Error [READ] Creative Group",
			fmt.Sprintf("Error reading Creative Group: %s", err),
		)
		return
	}

	// Overwrite creative with refreshed state
	state = r.OverwriteCreativeGroupModel(*creative)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *creativeGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan creativeGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	creativeGroupPatchFields := *openapi.NewCreativeGroupPatchFields()

	creativeGroupPatchFields.Name = plan.Name.ValueStringPointer()

	for _, creative := range plan.Creatives {
		creativeGroupPatchFields.Creatives = append(creativeGroupPatchFields.Creatives, *openapi.NewWeightedCreative(creative.CreativeId.ValueString(), float32(creative.Weight.ValueInt64())))
	}

	for _, tag := range plan.Tags {
		creativeGroupPatchFields.Tags = append(creativeGroupPatchFields.Tags, tag.ValueString())
	}
	creativeGroupPatchFields.Archived = plan.Archived.ValueBoolPointer()

	_, response, err := r.client.CreativeGroupsAPI.PatchCreativeGroup(ctx, r.accountID, plan.CreativeGroupId.ValueString()).CreativeGroupPatchFields(creativeGroupPatchFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE] Creative Group",
			fmt.Sprintf("Error updating Creative Group: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [UPDATE] Creative Group",
			fmt.Sprintf("Error updating Creative Group: %v", response),
		)
		return
	}

	creative, _, err := r.client.CreativeGroupsAPI.GetCreativeGroup(ctx, r.accountID, plan.CreativeGroupId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE -> READ] Creative",
			fmt.Sprintf("Error reading Creative: %v\n", err),
		)
		return
	}

	plan = r.OverwriteCreativeGroupModel(*creative)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *creativeGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state creativeGroupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.CreativeGroupsAPI.DeleteCreativeGroup(ctx, r.accountID, state.CreativeGroupId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [DELETE] Creative Group",
			fmt.Sprintf("Error deleting Creative Group: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [DELETE] Creative Group",
			fmt.Sprintf("Error deleting Creative Group: %v", response),
		)
		return
	}
}

// OverwriteCreativeGroupModel overwrites the creative model with the latest data.
func (r *creativeGroupResource) OverwriteCreativeGroupModel(creativeGroup openapi.CreativeGroup) creativeGroupModel {
	plan := creativeGroupModel{
		CreativeGroupId: types.StringValue(creativeGroup.CreativeGroupId),
		AccountId:       types.StringValue(creativeGroup.AccountId),
		Archived:        types.BoolValue(creativeGroup.Archived),
		Domain:          types.StringValue(creativeGroup.Domain),
		Name:            types.StringValue(creativeGroup.Name),
	}

	plan.Tags = []types.String{}
	for _, tag := range creativeGroup.Tags {
		plan.Tags = append(plan.Tags,
			types.StringValue(tag),
		)
	}

	plan.Spec = []specModel{}
	plan.Spec = append(plan.Spec, specModel{
		Banner: append([]bannerGroupModel{}, bannerGroupModel{
			Height: types.Int64Value(int64(creativeGroup.Spec.Banner.Height)),
			Width:  types.Int64Value(int64(creativeGroup.Spec.Banner.Width)),
		}),
	})

	plan.Creatives = []creativesModel{}
	for _, creative := range creativeGroup.Creatives {
		plan.Creatives = append(plan.Creatives, creativesModel{
			CreativeId: types.StringValue(creative.CreativeId),
			Weight:     types.Int64Value(int64(creative.Weight)),
		})
	}

	return plan
}
