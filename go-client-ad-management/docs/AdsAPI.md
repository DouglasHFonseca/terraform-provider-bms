# \AdsAPI

All URIs are relative to *http://api.ad.ads.sd.internal.us-east-2.bluemsdev.team*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAd**](AdsAPI.md#CreateAd) | **Post** /v1/accounts/{accountId}/ads | Create Ad
[**DeleteAd**](AdsAPI.md#DeleteAd) | **Delete** /v1/accounts/{accountId}/ads/{adId} | Delete Ad
[**GetAd**](AdsAPI.md#GetAd) | **Get** /v1/accounts/{accountId}/ads/{adId} | Get Ad
[**GetAdActiveRules**](AdsAPI.md#GetAdActiveRules) | **Get** /v1/accounts/{accountId}/ads/{adId}/rules:active | Get Ad Active Rules
[**ListAds**](AdsAPI.md#ListAds) | **Get** /v1/accounts/{accountId}/ads | List Ads
[**PatchAd**](AdsAPI.md#PatchAd) | **Patch** /v1/accounts/{accountId}/ads/{adId} | Patch Ad



## CreateAd

> Ad CreateAd(ctx, accountId).AdCreateFields(adCreateFields).Execute()

Create Ad



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
    adCreateFields := *openapiclient.NewAdCreateFields("example.com", "Example.com summer banner 728x90", *openapiclient.NewAdsCreativeCreativeSpec(), []openapiclient.AdRule{*openapiclient.NewAdRule("NXiWszwDTS8NR0ujsTIE2VY24Ya", "Weekdays, 14-16h", []openapiclient.Condition{*openapiclient.NewCondition("schedule")}, "0ujsszwN8NRY24YaXiTIE2VWDTS", false)}) // AdCreateFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsAPI.CreateAd(context.Background(), accountId).AdCreateFields(adCreateFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.CreateAd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAd`: Ad
    fmt.Fprintf(os.Stdout, "Response from `AdsAPI.CreateAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adCreateFields** | [**AdCreateFields**](AdCreateFields.md) |  | 

### Return type

[**Ad**](Ad.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAd

> DeleteAd(ctx, accountId, adId).Execute()

Delete Ad



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
    adId := "ujsDTSszw04YaXiTIE2VN8NRY2W" // string | The ID of the ad.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdsAPI.DeleteAd(context.Background(), accountId, adId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.DeleteAd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**adId** | **string** | The ID of the ad. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdRequest struct via the builder pattern


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


## GetAd

> Ad GetAd(ctx, accountId, adId).Execute()

Get Ad



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
    adId := "ujsDTSszw04YaXiTIE2VN8NRY2W" // string | The ID of the ad.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsAPI.GetAd(context.Background(), accountId, adId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAd`: Ad
    fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**adId** | **string** | The ID of the ad. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Ad**](Ad.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdActiveRules

> []ActiveAdRule GetAdActiveRules(ctx, accountId, adId).Execute()

Get Ad Active Rules



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
    adId := "ujsDTSszw04YaXiTIE2VN8NRY2W" // string | The ID of the ad.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsAPI.GetAdActiveRules(context.Background(), accountId, adId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.GetAdActiveRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdActiveRules`: []ActiveAdRule
    fmt.Fprintf(os.Stdout, "Response from `AdsAPI.GetAdActiveRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**adId** | **string** | The ID of the ad. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdActiveRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ActiveAdRule**](ActiveAdRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAds

> AdPage ListAds(ctx, accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()

List Ads



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
    filters := "{}" // string | The applicable filters, a json object according with `AdFilters`. (optional)
    pageSize := float32(10) // float32 | The page size. (optional) (default to 10)
    pageToken := "AIUHSDJANBDHOH3" // string | The page token from which to continue listing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsAPI.ListAds(context.Background(), accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.ListAds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAds`: AdPage
    fmt.Fprintf(os.Stdout, "Response from `AdsAPI.ListAds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | The applicable filters, a json object according with &#x60;AdFilters&#x60;. | 
 **pageSize** | **float32** | The page size. | [default to 10]
 **pageToken** | **string** | The page token from which to continue listing. | 

### Return type

[**AdPage**](AdPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAd

> Ad PatchAd(ctx, accountId, adId).AdPatchFields(adPatchFields).Execute()

Patch Ad



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
    adId := "ujsDTSszw04YaXiTIE2VN8NRY2W" // string | The ID of the ad.
    adPatchFields := *openapiclient.NewAdPatchFields() // AdPatchFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsAPI.PatchAd(context.Background(), accountId, adId).AdPatchFields(adPatchFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsAPI.PatchAd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAd`: Ad
    fmt.Fprintf(os.Stdout, "Response from `AdsAPI.PatchAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 
**adId** | **string** | The ID of the ad. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adPatchFields** | [**AdPatchFields**](AdPatchFields.md) |  | 

### Return type

[**Ad**](Ad.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

