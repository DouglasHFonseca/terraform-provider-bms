package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	adopenapi "go-client-ad-management"
	crativeopenapi "go-client-creative-management"
	"os"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &bmsProviderModel{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &bmsProviderModel{
			version: version,
		}
	}
}

// bmsProviderModel is the provider implementation.
type bmsProviderModel struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version   string
	AccountID types.String `tfsdk:"account_id"`
}

type providerConfig struct {
	clientCreative *crativeopenapi.APIClient
	clientAd       *adopenapi.APIClient
	accountID      string
}

// Metadata returns the provider type name.
func (p *bmsProviderModel) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bms"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *bmsProviderModel) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

// Configure prepares a BMS API clientCreative for data sources and resources.
func (p *bmsProviderModel) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring BMS clientCreative")

	var config bmsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.AccountID.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("accountId"),
			"Unknown BMS API AccountId",
			"The provider cannot create the BMS API clientCreative as there is an unknown configuration value for the BMS API accountId.",
		)
	}

	accountId := os.Getenv("ACCOUNT_ID")

	if !config.AccountID.IsNull() {
		accountId = config.AccountID.ValueString()
	}

	if accountId == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("accountId"),
			"Unknown BMS API AccountId",
			"The provider cannot create the BMS API clientCreative as there is an unknown configuration value for the BMS API accountId.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	p.AccountID = types.StringValue(accountId)

	// Configure the BMS API clientCreative.
	cfgCreative := crativeopenapi.NewConfiguration()
	clientCreative := crativeopenapi.NewAPIClient(cfgCreative)

	// Configure the BMS API clientAd.
	cfgAd := adopenapi.NewConfiguration()
	clientAd := adopenapi.NewAPIClient(cfgAd)

	configClient := &providerConfig{
		clientCreative: clientCreative,
		clientAd:       clientAd,
		accountID:      accountId,
	}

	resp.DataSourceData = configClient
	resp.ResourceData = configClient

	tflog.Info(ctx, "Configured BMS client Creative/Ad", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *bmsProviderModel) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCreativesDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *bmsProviderModel) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCreativeResource,
		NewCreativeGroupResource,
		NewAdResource,
	}
}
