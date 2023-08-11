# \CreativeGroupsAPI

All URIs are relative to *http://api.creative.ads.sd.internal.us-east-2.bluemsdev.team*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreativeGroup**](CreativeGroupsAPI.md#CreateCreativeGroup) | **Post** /v1/accounts/{accountId}/creative-groups | Create Creative Group
[**DeleteCreativeGroup**](CreativeGroupsAPI.md#DeleteCreativeGroup) | **Delete** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Delete Creative Group
[**GetCreativeGroup**](CreativeGroupsAPI.md#GetCreativeGroup) | **Get** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Get Creative Group
[**GetCreativeGroupAvailableCreatives**](CreativeGroupsAPI.md#GetCreativeGroupAvailableCreatives) | **Get** /v1/accounts/{accountId}/creative-groups/{creativeGroupId}/creatives:available | Get Creative Group Available Creatives
[**ListCreativeGroups**](CreativeGroupsAPI.md#ListCreativeGroups) | **Get** /v1/accounts/{accountId}/creative-groups | List Creative Groups
[**PatchCreativeGroup**](CreativeGroupsAPI.md#PatchCreativeGroup) | **Patch** /v1/accounts/{accountId}/creative-groups/{creativeGroupId} | Patch Creative Group



## CreateCreativeGroup

> CreativeGroup CreateCreativeGroup(ctx, accountId).CreativeGroupCreateFields(creativeGroupCreateFields).Execute()

Create Creative Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    creativeGroupCreateFields := *openapiclient.NewCreativeGroupCreateFields("example.com", "Example.com weekends discount banners 300x600", *openapiclient.NewCreativeSpec(), []openapiclient.WeightedCreative{*openapiclient.NewWeightedCreative("0ujssxh0cECutqzMgbtXSGnjorm", float32(100))}) // CreativeGroupCreateFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativeGroupsAPI.CreateCreativeGroup(context.Background(), accountId).CreativeGroupCreateFields(creativeGroupCreateFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.CreateCreativeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreativeGroup`: CreativeGroup
    fmt.Fprintf(os.Stdout, "Response from `CreativeGroupsAPI.CreateCreativeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreativeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creativeGroupCreateFields** | [**CreativeGroupCreateFields**](CreativeGroupCreateFields.md) |  | 

### Return type

[**CreativeGroup**](CreativeGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCreativeGroup

> DeleteCreativeGroup(ctx, accountId, creativeGroupId).Execute()

Delete Creative Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    creativeGroupId := "0ujsszwN8NRY24YaXiTIE2VWDTS" // string | The ID of the creative group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CreativeGroupsAPI.DeleteCreativeGroup(context.Background(), accountId, creativeGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.DeleteCreativeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeGroupId** | **string** | The ID of the creative group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreativeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreativeGroup

> CreativeGroup GetCreativeGroup(ctx, accountId, creativeGroupId).Execute()

Get Creative Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    creativeGroupId := "0ujsszwN8NRY24YaXiTIE2VWDTS" // string | The ID of the creative group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativeGroupsAPI.GetCreativeGroup(context.Background(), accountId, creativeGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.GetCreativeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreativeGroup`: CreativeGroup
    fmt.Fprintf(os.Stdout, "Response from `CreativeGroupsAPI.GetCreativeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeGroupId** | **string** | The ID of the creative group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreativeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreativeGroup**](CreativeGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreativeGroupAvailableCreatives

> []AvailableCreative GetCreativeGroupAvailableCreatives(ctx, accountId, creativeGroupId).Execute()

Get Creative Group Available Creatives



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    creativeGroupId := "0ujsszwN8NRY24YaXiTIE2VWDTS" // string | The ID of the creative group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativeGroupsAPI.GetCreativeGroupAvailableCreatives(context.Background(), accountId, creativeGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.GetCreativeGroupAvailableCreatives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreativeGroupAvailableCreatives`: []AvailableCreative
    fmt.Fprintf(os.Stdout, "Response from `CreativeGroupsAPI.GetCreativeGroupAvailableCreatives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeGroupId** | **string** | The ID of the creative group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreativeGroupAvailableCreativesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AvailableCreative**](AvailableCreative.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreativeGroups

> CreativeGroupPage ListCreativeGroups(ctx, accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()

List Creative Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    filters := "{}" // string | The applicable filters, a json object according with `CreativeGroupFilters`. (optional)
    pageSize := float32(10) // float32 | The page size. (optional) (default to 10)
    pageToken := "AIUHSDJANBDHOH3" // string | The page token from which to continue listing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativeGroupsAPI.ListCreativeGroups(context.Background(), accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.ListCreativeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreativeGroups`: CreativeGroupPage
    fmt.Fprintf(os.Stdout, "Response from `CreativeGroupsAPI.ListCreativeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCreativeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | The applicable filters, a json object according with &#x60;CreativeGroupFilters&#x60;. | 
 **pageSize** | **float32** | The page size. | [default to 10]
 **pageToken** | **string** | The page token from which to continue listing. | 

### Return type

[**CreativeGroupPage**](CreativeGroupPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCreativeGroup

> CreativeGroup PatchCreativeGroup(ctx, accountId, creativeGroupId).CreativeGroupPatchFields(creativeGroupPatchFields).Execute()

Patch Creative Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "359312844241" // string | The ID of the account.
    creativeGroupId := "0ujsszwN8NRY24YaXiTIE2VWDTS" // string | The ID of the creative group.
    creativeGroupPatchFields := *openapiclient.NewCreativeGroupPatchFields() // CreativeGroupPatchFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativeGroupsAPI.PatchCreativeGroup(context.Background(), accountId, creativeGroupId).CreativeGroupPatchFields(creativeGroupPatchFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativeGroupsAPI.PatchCreativeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCreativeGroup`: CreativeGroup
    fmt.Fprintf(os.Stdout, "Response from `CreativeGroupsAPI.PatchCreativeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeGroupId** | **string** | The ID of the creative group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCreativeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **creativeGroupPatchFields** | [**CreativeGroupPatchFields**](CreativeGroupPatchFields.md) |  | 

### Return type

[**CreativeGroup**](CreativeGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

