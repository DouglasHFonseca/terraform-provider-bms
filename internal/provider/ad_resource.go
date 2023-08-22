package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "go-client-ad-management"
	"regexp"
	"time"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &adResource{}
	_ resource.ResourceWithConfigure = &adResource{}
)

// adModel maps ad schema data.
type adModel struct {
	AdId            types.String          `tfsdk:"ad_id"`
	AccountId       types.String          `tfsdk:"account_id"`
	Name            types.String          `tfsdk:"name"`
	Archived        types.Bool            `tfsdk:"archived"`
	Domain          types.String          `tfsdk:"domain"`
	ExchangeReviews []exchangeReviewModel `tfsdk:"exchange_reviews"`
	Rules           []rulesModel          `tfsdk:"rule"`
	Tags            []types.String        `tfsdk:"tags"`
	Spec            []specModel           `tfsdk:"spec"`
}

// rulesModel maps rule schema data.
type rulesModel struct {
	AdRuleId        types.String     `tfsdk:"ad_rule_id"`
	Name            types.String     `tfsdk:"name"`
	CreativeGroupId types.String     `tfsdk:"creative_group_id"`
	Enabled         types.Bool       `tfsdk:"enabled"`
	Conditions      []conditionModel `tfsdk:"schedule_condition"`
}

// conditionModel maps condition schema data.
type conditionModel struct {
	Type        types.String       `tfsdk:"type"`
	Start       types.String       `tfsdk:"start"`
	End         types.String       `tfsdk:"end"`
	TimesOfWeek []timesOfWeekModel `tfsdk:"time_of_week"`
}

// timesOfWeekModel maps times of week schema data.
type timesOfWeekModel struct {
	Start types.Int64 `tfsdk:"start"`
	End   types.Int64 `tfsdk:"end"`
}

// exchangeReviewModel maps exchange review schema data.
type exchangeReviewModel struct {
	UpdatedAt        types.String `tfsdk:"updated_at"`
	ExchangeId       types.String `tfsdk:"exchange_id"`
	ExchangeReviewId types.String `tfsdk:"exchange_review_id"`
	Status           types.String `tfsdk:"status"`
	Comments         types.String `tfsdk:"comments"`
}

// NewAdResource is a helper function to simplify the provider implementation.
func NewAdResource() resource.Resource {
	return &adResource{}
}

// adResource is the resource implementation.
type adResource struct {
	client    *openapi.APIClient
	accountID string
}

// Configure adds the provider configured clientCreative to the resource.
func (r *adResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*providerConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Configure Type",
			fmt.Sprintf("Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = config.clientAd
	r.accountID = config.accountID
}

// Metadata returns the resource type name.
func (r *adResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ads_ad"
}

// Schema defines the schema for the resource.
func (r *adResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ad_id": schema.StringAttribute{
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
			"name": schema.StringAttribute{
				Required: true,
			},
			"tags": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"archived": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required: true,
			},
		},
		Blocks: map[string]schema.Block{
			"exchange_reviews": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"exchange_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"exchange_review_id": schema.StringAttribute{
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
						"comments": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"rule": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"ad_rule_id": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"creative_group_id": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
							Default:  booldefault.StaticBool(true),
						},
					},
					Blocks: map[string]schema.Block{
						"schedule_condition": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtLeast(1),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(schedule)$`), "type must be one of: schedule"),
										},
									},
									"start": schema.StringAttribute{
										Optional: true,
									},
									"end": schema.StringAttribute{
										Optional:   true,
										Validators: []validator.String{
											//TODO: validate if end > start
										},
									},
								},
								Blocks: map[string]schema.Block{
									"time_of_week": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"start": schema.Int64Attribute{
													Required: true,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
												"end": schema.Int64Attribute{
													Required: true,
													Validators: []validator.Int64{
														int64validator.AtMost(10080),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"spec": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"banner": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtLeast(1),
								listvalidator.SizeAtMost(1),
							},
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
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *adResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan adModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var rules []openapi.AdRule
	for _, rule := range plan.Rules {
		var conditions []openapi.ScheduleCondition
		for _, condition := range rule.Conditions {
			newCondition := *openapi.NewScheduleCondition(condition.Type.ValueString())

			startValue, err := time.Parse(time.RFC3339, condition.Start.ValueString())

			if err != nil {
				resp.Diagnostics.AddError(
					"Error parsing start time",
					fmt.Sprintf("Error parsing start time: %s", err),
				)
				return
			}
			newCondition.Start = &startValue

			endValue, err2 := time.Parse(time.RFC3339, condition.End.ValueString())

			if err2 != nil {
				resp.Diagnostics.AddError(
					"Error parsing end time",
					fmt.Sprintf("Error parsing end time: %s", err),
				)
				return
			}
			newCondition.End = &endValue

			for _, timeOfWeek := range condition.TimesOfWeek {
				newCondition.TimesOfWeek = append(newCondition.TimesOfWeek, *openapi.NewTimeOfWeek(float32(timeOfWeek.Start.ValueInt64()), float32(timeOfWeek.End.ValueInt64())))
			}
			conditions = append(conditions, newCondition)
		}
		rules = append(rules, *openapi.NewAdRule(rule.AdRuleId.ValueString(), rule.Name.ValueString(), conditions, rule.CreativeGroupId.ValueString(), rule.Enabled.ValueBool()))
	}

	adsCreativeSpec := *openapi.NewAdsCreativeCreativeSpec()
	adsCreativeSpec.Banner = openapi.NewAdsCreativeBannerSpec(float32(plan.Spec[0].Banner[0].Width.ValueInt64()), float32(plan.Spec[0].Banner[0].Height.ValueInt64()))

	adCreateFields := *openapi.NewAdCreateFields(plan.Domain.ValueString(), plan.Name.ValueString(), adsCreativeSpec, rules)

	if plan.Tags != nil {
		var tags []string
		for _, tag := range plan.Tags {
			tags = append(tags, tag.ValueString())
		}
		adCreateFields.Tags = tags
	}

	ad, _, err := r.client.AdsAPI.CreateAd(ctx, r.accountID).AdCreateFields(adCreateFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Ads",
			fmt.Sprintf("Error creating Ad: %s", err),
		)
		return
	}

	// Overwrite the terraform state with the newly created resource.
	plan = r.OverwriteAdModel(*ad)

	diags = resp.State.Set(ctx, plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Read refreshes the Terraform state with the latest data.
func (r *adResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state adModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	creative, response, err := r.client.AdsAPI.GetAd(ctx, r.accountID, state.AdId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Response [READ] Ad",
			fmt.Sprintf("Error reading Ad: %v", response),
		)

		resp.Diagnostics.AddError(
			"Error [READ] Ad",
			fmt.Sprintf("Error reading Ad: %s", err),
		)
		return
	}

	// Overwrite creative with refreshed state
	state = r.OverwriteAdModel(*creative)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *adResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan adModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	adPatchFields := *openapi.NewAdPatchFields()

	adPatchFields.Name = plan.Name.ValueStringPointer()
	adPatchFields.Archived = plan.Archived.ValueBoolPointer()

	for _, rule := range plan.Rules {
		var conditions []openapi.ScheduleCondition
		for _, condition := range rule.Conditions {
			conditions = append(conditions, *openapi.NewScheduleCondition(condition.Type.ValueString()))
		}
		adPatchFields.Rules = append(adPatchFields.Rules, *openapi.NewAdRule(rule.AdRuleId.ValueString(), rule.Name.ValueString(), conditions, rule.CreativeGroupId.ValueString(), rule.Enabled.ValueBool()))
	}

	for _, tag := range plan.Tags {
		adPatchFields.Tags = append(adPatchFields.Tags, tag.ValueString())
	}

	_, response, err := r.client.AdsAPI.PatchAd(ctx, r.accountID, plan.AdId.ValueString()).AdPatchFields(adPatchFields).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE] Ad",
			fmt.Sprintf("Error updating Ad: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [UPDATE] Creative",
			fmt.Sprintf("Error updating Ad: %v", response),
		)
		return
	}

	creative, _, err := r.client.AdsAPI.GetAd(ctx, r.accountID, plan.AdId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [UPDATE -> READ] Ad",
			fmt.Sprintf("Error reading Ad: %v\n", err),
		)
		return
	}

	plan = r.OverwriteAdModel(*creative)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *adResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state adModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.AdsAPI.DeleteAd(ctx, r.accountID, state.AccountId.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error [DELETE] Ad",
			fmt.Sprintf("Error deleting Ad: %s", err),
		)

		resp.Diagnostics.AddError(
			"Response [DELETE] Ad",
			fmt.Sprintf("Error deleting Ad: %v", response),
		)
		return
	}
}

// OverwriteAdModel overwrites the creative model with the latest data.
func (r *adResource) OverwriteAdModel(ad openapi.Ad) adModel {
	plan := adModel{
		AdId:      types.StringValue(ad.AdId),
		Name:      types.StringValue(ad.Name),
		Domain:    types.StringValue(ad.Domain),
		Archived:  types.BoolValue(ad.Archived),
		AccountId: types.StringValue(ad.AccountId),
	}

	plan.Tags = []types.String{}
	for _, tag := range ad.Tags {
		plan.Tags = append(plan.Tags,
			types.StringValue(tag),
		)
	}

	plan.ExchangeReviews = []exchangeReviewModel{}
	for _, exchangeReview := range ad.ExchangeReviews {
		plan.ExchangeReviews = append(plan.ExchangeReviews, exchangeReviewModel{
			UpdatedAt:        types.StringValue(exchangeReview.UpdatedAt.Format(time.RFC3339)),
			ExchangeId:       types.StringValue(exchangeReview.ExchangeId),
			ExchangeReviewId: types.StringValue(exchangeReview.ExchangeReviewId),
			Status:           types.StringValue(exchangeReview.Status),
			Comments:         types.StringValue(*exchangeReview.Comments),
		})
	}

	plan.Spec = []specModel{}
	plan.Spec = append(plan.Spec, specModel{
		Banner: append([]bannerGroupModel{}, bannerGroupModel{
			Height: types.Int64Value(int64(ad.Spec.Banner.Height)),
			Width:  types.Int64Value(int64(ad.Spec.Banner.Width)),
		}),
	})

	plan.Rules = []rulesModel{}
	for _, rule := range ad.Rules {
		plan.Rules = append(plan.Rules, rulesModel{
			AdRuleId:        types.StringValue(rule.AdRuleId),
			Name:            types.StringValue(rule.Name),
			CreativeGroupId: types.StringValue(rule.CreativeGroupId),
			Enabled:         types.BoolValue(rule.Enabled),
		})
		for index, condition := range rule.Conditions {
			plan.Rules[index].Conditions = append(plan.Rules[index].Conditions, conditionModel{
				Type:  types.StringValue(condition.Type),
				Start: types.StringValue(condition.Start.Format(time.RFC3339)),
				End:   types.StringValue(condition.End.Format(time.RFC3339)),
			})

			for _, timeOfWeek := range condition.TimesOfWeek {
				plan.Rules[index].Conditions[index].TimesOfWeek = append(plan.Rules[index].Conditions[index].TimesOfWeek, timesOfWeekModel{
					Start: types.Int64Value(int64(timeOfWeek.Start)),
					End:   types.Int64Value(int64(timeOfWeek.End)),
				})
			}
		}
	}

	return plan
}
