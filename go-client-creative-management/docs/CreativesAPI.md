# \CreativesAPI

All URIs are relative to *http://api.creative.ads.sd.internal.us-east-2.bluemsdev.team*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreative**](CreativesAPI.md#CreateCreative) | **Post** /v1/accounts/{accountId}/creatives | Create Creative
[**DeleteCreative**](CreativesAPI.md#DeleteCreative) | **Delete** /v1/accounts/{accountId}/creatives/{creativeId} | Delete Creative
[**GetCreative**](CreativesAPI.md#GetCreative) | **Get** /v1/accounts/{accountId}/creatives/{creativeId} | Get Creative
[**ListCreatives**](CreativesAPI.md#ListCreatives) | **Get** /v1/accounts/{accountId}/creatives | List Creatives
[**PatchCreative**](CreativesAPI.md#PatchCreative) | **Patch** /v1/accounts/{accountId}/creatives/{creativeId} | Patch Creative
[**SendCreativeForReview**](CreativesAPI.md#SendCreativeForReview) | **Post** /v1/accounts/{accountId}/creatives/{creativeId}:send-for-review | Send Creative for Review



## CreateCreative

> Creative CreateCreative(ctx, accountId).CreativeCreateFields(creativeCreateFields).Execute()

Create Creative



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
    creativeCreateFields := *openapiclient.NewCreativeCreateFields("example.com", "Example.com weekends discount banner 300x600") // CreativeCreateFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesAPI.CreateCreative(context.Background(), accountId).CreativeCreateFields(creativeCreateFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.CreateCreative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreative`: Creative
    fmt.Fprintf(os.Stdout, "Response from `CreativesAPI.CreateCreative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creativeCreateFields** | [**CreativeCreateFields**](CreativeCreateFields.md) |  | 

### Return type

[**Creative**](Creative.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCreative

> DeleteCreative(ctx, accountId, creativeId).Execute()

Delete Creative



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
    creativeId := "0ujssxh0cECutqzMgbtXSGnjorm" // string | The ID of the creative.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CreativesAPI.DeleteCreative(context.Background(), accountId, creativeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.DeleteCreative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeId** | **string** | The ID of the creative. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreativeRequest struct via the builder pattern


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


## GetCreative

> Creative GetCreative(ctx, accountId, creativeId).Execute()

Get Creative



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
    creativeId := "0ujssxh0cECutqzMgbtXSGnjorm" // string | The ID of the creative.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesAPI.GetCreative(context.Background(), accountId, creativeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.GetCreative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreative`: Creative
    fmt.Fprintf(os.Stdout, "Response from `CreativesAPI.GetCreative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeId** | **string** | The ID of the creative. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Creative**](Creative.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreatives

> CreativePage ListCreatives(ctx, accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()

List Creatives



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
    filters := "{}" // string | The applicable filters, a json object according with `CreativeFilters`. (optional)
    pageSize := float32(10) // float32 | The page size. (optional) (default to 10)
    pageToken := "AIUHSDJANBDHOH3" // string | The page token from which to continue listing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesAPI.ListCreatives(context.Background(), accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.ListCreatives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatives`: CreativePage
    fmt.Fprintf(os.Stdout, "Response from `CreativesAPI.ListCreatives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCreativesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | The applicable filters, a json object according with &#x60;CreativeFilters&#x60;. | 
 **pageSize** | **float32** | The page size. | [default to 10]
 **pageToken** | **string** | The page token from which to continue listing. | 

### Return type

[**CreativePage**](CreativePage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCreative

> Creative PatchCreative(ctx, accountId, creativeId).CreativePatchFields(creativePatchFields).Execute()

Patch Creative



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
    creativeId := "0ujssxh0cECutqzMgbtXSGnjorm" // string | The ID of the creative.
    creativePatchFields := *openapiclient.NewCreativePatchFields() // CreativePatchFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreativesAPI.PatchCreative(context.Background(), accountId, creativeId).CreativePatchFields(creativePatchFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.PatchCreative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCreative`: Creative
    fmt.Fprintf(os.Stdout, "Response from `CreativesAPI.PatchCreative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeId** | **string** | The ID of the creative. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCreativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **creativePatchFields** | [**CreativePatchFields**](CreativePatchFields.md) |  | 

### Return type

[**Creative**](Creative.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCreativeForReview

> SendCreativeForReview(ctx, accountId, creativeId).Execute()

Send Creative for Review



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
    creativeId := "0ujssxh0cECutqzMgbtXSGnjorm" // string | The ID of the creative.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CreativesAPI.SendCreativeForReview(context.Background(), accountId, creativeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreativesAPI.SendCreativeForReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**creativeId** | **string** | The ID of the creative. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCreativeForReviewRequest struct via the builder pattern


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

