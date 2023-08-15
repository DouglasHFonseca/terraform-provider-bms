package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "go-client-creative-management"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &creativesDataSource{}
	_ datasource.DataSourceWithConfigure = &creativesDataSource{}
)

// NewCreativesDataSource is a helper function to simplify the provider implementation.
func NewCreativesDataSource() datasource.DataSource {
	return &creativesDataSource{}
}

// creativesDataSource is the data source implementation.
type creativesDataSource struct {
	client    *openapi.APIClient
	accountID string
}

// creativesDataSourceModel maps the data source schema data.
type creativesDataSourceModel struct {
	Creatives []creativeModel `tfsdk:"creatives"`
	Filters   filtersModel    `tfsdk:"filters"`
}

// filtersModel maps the data source filters schema data.
type filtersModel struct {
	Tag              types.String   `tfsdk:"tag"`
	Domain           types.String   `tfsdk:"domain"`
	Archived         types.Bool     `tfsdk:"archived"`
	Name             types.String   `tfsdk:"name"`
	Spec             *filterSpec    `tfsdk:"spec"`
	CreativeIds      []types.String `tfsdk:"creative_ids"`
	CreativeGroupIds []types.String `tfsdk:"creative_group_ids"`
	Status           types.String   `tfsdk:"status"`
	Enabled          types.Bool     `tfsdk:"enabled"`
	Search           types.String   `tfsdk:"search"`
}

// filterSpec maps the data source filter spec schema data.
type filterSpec struct {
	Banner filterBanner `tfsdk:"banner"`
}

// filterBanner maps the data source filter banner schema data.
type filterBanner struct {
	Width  types.Int64 `tfsdk:"width"`
	Height types.Int64 `tfsdk:"height"`
}

// creativeModel maps creative schema data.
type creativeModel struct {
	CreativeId types.String   `tfsdk:"creative_id"`
	Status     types.String   `tfsdk:"status"`
	AccountId  types.String   `tfsdk:"account_id"`
	Enabled    types.Bool     `tfsdk:"enabled"`
	Tags       []types.String `tfsdk:"tags"`
	Domain     types.String   `tfsdk:"domain"`
	Name       types.String   `tfsdk:"name"`
	Banner     []bannerModel  `tfsdk:"banner"`
	Archived   types.Bool     `tfsdk:"archived"`
}

// bannerModel maps banner data.
type bannerModel struct {
	Width   types.Int64  `tfsdk:"width"`
	Height  types.Int64  `tfsdk:"height"`
	Snippet types.String `tfsdk:"snippet"`
}

// Configure adds the provider configured client to the data source.
func (d *creativesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*providerConfig)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *swagger.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = config.client
	d.accountID = config.accountID
}

// Metadata returns the data source type name.
func (d *creativesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_creatives"
}

// Schema defines the schema for the data source.
func (d *creativesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"tag": schema.StringAttribute{
						Optional: true,
					},
					"domain": schema.StringAttribute{
						Optional: true,
					},
					"archived": schema.BoolAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Optional: true,
					},
					"spec": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"banner": schema.SingleNestedAttribute{
								Optional: true,
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
					"creative_ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"creative_group_ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"status": schema.StringAttribute{
						Optional: true,
					},
					"enabled": schema.BoolAttribute{
						Optional: true,
					},
					"search": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"creatives": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"creative_id": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"account_id": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"domain": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"banner": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"width": schema.Int64Attribute{
									Computed: true,
								},
								"height": schema.Int64Attribute{
									Computed: true,
								},
								"snippet": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"archived": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *creativesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state creativesDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	jsonStringifyFilters, err := JSONStringifyFilters(state.Filters)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Convert Filters to JSON",
			fmt.Sprintf("Error converting filters to JSON: %s", err.Error()),
		)
		return
	}

	data, _, err := d.client.CreativesAPI.ListCreatives(ctx, d.accountID).Filters(jsonStringifyFilters).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to READ BMS Creatives",
			fmt.Sprintf("Unable to retrieve creatives: %s", err.Error()),
		)
		return
	}

	creatives := data.Values

	for _, creative := range creatives {
		creativeModel := creativeModel{
			CreativeId: types.StringValue(creative.CreativeId),
			Status:     types.StringValue(string(creative.Status)),
			AccountId:  types.StringValue(creative.AccountId),
			Enabled:    types.BoolValue(creative.Enabled),
			Domain:     types.StringValue(creative.Domain),
			Name:       types.StringValue(creative.Name),
			Archived:   types.BoolValue(creative.Archived),
		}

		if creative.Banner != nil {
			creativeModel.Banner = append(creativeModel.Banner, bannerModel{
				Width:   types.Int64Value(int64(creative.Banner.Width)),
				Height:  types.Int64Value(int64(creative.Banner.Height)),
				Snippet: types.StringValue(creative.Banner.Snippet),
			})
		}

		for _, tag := range creative.Tags {
			creativeModel.Tags = append(creativeModel.Tags, types.StringValue(tag))
		}

		state.Creatives = append(state.Creatives, creativeModel)
	}

	diagsSet := resp.State.Set(ctx, &state)

	resp.Diagnostics.Append(diagsSet...)

	if resp.Diagnostics.HasError() {
		return
	}
}

// JSONStringifyFilters converts the filters to a JSON string
func JSONStringifyFilters(filters filtersModel) (string, error) {
	readFilterFields := *openapi.NewCreativeFilters()

	readFilterFields.Search = filters.Search.ValueStringPointer()
	readFilterFields.Name = filters.Name.ValueStringPointer()
	readFilterFields.Domain = filters.Domain.ValueStringPointer()
	readFilterFields.Archived = filters.Archived.ValueBoolPointer()
	readFilterFields.Tag = filters.Tag.ValueStringPointer()
	readFilterFields.Status = filters.Status.ValueStringPointer()
	readFilterFields.Enabled = filters.Enabled.ValueBoolPointer()

	if filters.CreativeIds != nil {
		var creativeIds []string
		for _, creativeId := range filters.CreativeIds {
			creativeIds = append(creativeIds, creativeId.ValueString())
		}
		readFilterFields.CreativeIds = creativeIds
	}

	if filters.CreativeGroupIds != nil {
		var creativeGroupIds []string
		for _, creativeGroupId := range filters.CreativeGroupIds {
			creativeGroupIds = append(creativeGroupIds, creativeGroupId.ValueString())
		}
		readFilterFields.CreativeGroupIds = creativeGroupIds
	}

	if filters.Spec != nil {
		readFilterFields.Spec = &openapi.CreativeSpec{
			Banner: &openapi.BannerSpec{
				Width:  float32(filters.Spec.Banner.Width.ValueInt64()),
				Height: float32(filters.Spec.Banner.Height.ValueInt64()),
			},
		}
	}

	jsonData, err := json.Marshal(readFilterFields)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
