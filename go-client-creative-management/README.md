# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.creative.ads.sd.internal.us-east-2.bluemsdev.team*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CreativeGroupsAPI* | [**CreateCreativeGroup**](docs/CreativeGroupsAPI.md#createcreativegroup) | **Post** /v1/accounts/{accountId}/creative-groups | Create Creative Group
*CreativeGroupsAPI* | [**DeleteCreativeGroup**](docs/CreativeGroupsAPI.md#deletecreativegroup) | **Delete** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Delete Creative Group
*CreativeGroupsAPI* | [**GetCreativeGroup**](docs/CreativeGroupsAPI.md#getcreativegroup) | **Get** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Get Creative Group
*CreativeGroupsAPI* | [**GetCreativeGroupAvailableCreatives**](docs/CreativeGroupsAPI.md#getcreativegroupavailablecreatives) | **Get** /v1/accounts/{accountId}/creative-groups/{creativeGroupId}/creatives:available | Get Creative Group Available Creatives
*CreativeGroupsAPI* | [**ListCreativeGroups**](docs/CreativeGroupsAPI.md#listcreativegroups) | **Get** /v1/accounts/{accountId}/creative-groups | List Creative Groups
*CreativeGroupsAPI* | [**PatchCreativeGroup**](docs/CreativeGroupsAPI.md#patchcreativegroup) | **Patch** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Patch Creative Group
*CreativesAPI* | [**CreateCreative**](docs/CreativesAPI.md#createcreative) | **Post** /v1/accounts/{accountId}/creatives | Create Creative
*CreativesAPI* | [**DeleteCreative**](docs/CreativesAPI.md#deletecreative) | **Delete** /v1/accounts/{accountId}/creatives/{creativeId} | Delete Creative
*CreativesAPI* | [**GetCreative**](docs/CreativesAPI.md#getcreative) | **Get** /v1/accounts/{accountId}/creatives/{creativeId} | Get Creative
*CreativesAPI* | [**ListCreatives**](docs/CreativesAPI.md#listcreatives) | **Get** /v1/accounts/{accountId}/creatives | List Creatives
*CreativesAPI* | [**PatchCreative**](docs/CreativesAPI.md#patchcreative) | **Patch** /v1/accounts/{accountId}/creatives/{creativeId} | Patch Creative
*CreativesAPI* | [**SendCreativeForReview**](docs/CreativesAPI.md#sendcreativeforreview) | **Post** /v1/accounts/{accountId}/creatives/{creativeId}:send-for-review | Send Creative for Review
*DomainsAPI* | [**ListDomains**](docs/DomainsAPI.md#listdomains) | **Get** /v1/accounts/{accountId}/domains | List Domains


## Documentation For Models

 - [AvailableCreative](docs/AvailableCreative.md)
 - [BannerFields](docs/BannerFields.md)
 - [BannerSpec](docs/BannerSpec.md)
 - [Creative](docs/Creative.md)
 - [CreativeCreateFields](docs/CreativeCreateFields.md)
 - [CreativeCreateFieldsAllOfBanner](docs/CreativeCreateFieldsAllOfBanner.md)
 - [CreativeFields](docs/CreativeFields.md)
 - [CreativeFilters](docs/CreativeFilters.md)
 - [CreativeGroup](docs/CreativeGroup.md)
 - [CreativeGroupCreateFields](docs/CreativeGroupCreateFields.md)
 - [CreativeGroupFilters](docs/CreativeGroupFilters.md)
 - [CreativeGroupPage](docs/CreativeGroupPage.md)
 - [CreativeGroupPatchFields](docs/CreativeGroupPatchFields.md)
 - [CreativePage](docs/CreativePage.md)
 - [CreativePatchFields](docs/CreativePatchFields.md)
 - [CreativeSpec](docs/CreativeSpec.md)
 - [DomainFilters](docs/DomainFilters.md)
 - [DomainPage](docs/DomainPage.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [WeightedCreative](docs/WeightedCreative.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



