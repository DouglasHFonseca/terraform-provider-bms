# \ExchangeReviewsAPI

All URIs are relative to *http://api.ad.ads.sd.internal.us-east-2.bluemsdev.team*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutExchangeReview**](ExchangeReviewsAPI.md#PutExchangeReview) | **Put** /v1/accounts/{accountId}/ads/{adId}/exchange-reviews/{exchangeId} | Put Exchange Review
[**RefreshExchangeReviewStatus**](ExchangeReviewsAPI.md#RefreshExchangeReviewStatus) | **Post** /v1/accounts/{accountId}/ads/{adId}/exchange-reviews:refresh | Refresh Exchange Review Status
[**SendAdForExchangeReview**](ExchangeReviewsAPI.md#SendAdForExchangeReview) | **Post** /v1/accounts/{accountId}/ads/{adId}/exchange-reviews | Send Ad for Exchange Review



## PutExchangeReview

> PutExchangeReview(ctx, accountId, adId, exchangeId).ExchangeReviewPutFields(exchangeReviewPutFields).Execute()

Put Exchange Review



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
    exchangeId := "adx" // string | The ID of the exchange.
    exchangeReviewPutFields := *openapiclient.NewExchangeReviewPutFields("prd-ujsDTSszw04YaXiTIE2VN8NRY2W", "approved") // ExchangeReviewPutFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExchangeReviewsAPI.PutExchangeReview(context.Background(), accountId, adId, exchangeId).ExchangeReviewPutFields(exchangeReviewPutFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeReviewsAPI.PutExchangeReview``: %v\n", err)
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
**exchangeId** | **string** | The ID of the exchange. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExchangeReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exchangeReviewPutFields** | [**ExchangeReviewPutFields**](ExchangeReviewPutFields.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshExchangeReviewStatus

> RefreshExchangeReviewStatus(ctx, accountId, adId).ExchangeReviewCreateFields(exchangeReviewCreateFields).Execute()

Refresh Exchange Review Status



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
    exchangeReviewCreateFields := *openapiclient.NewExchangeReviewCreateFields([]string{"adx"}) // ExchangeReviewCreateFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExchangeReviewsAPI.RefreshExchangeReviewStatus(context.Background(), accountId, adId).ExchangeReviewCreateFields(exchangeReviewCreateFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeReviewsAPI.RefreshExchangeReviewStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRefreshExchangeReviewStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exchangeReviewCreateFields** | [**ExchangeReviewCreateFields**](ExchangeReviewCreateFields.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAdForExchangeReview

> SendAdForExchangeReview(ctx, accountId, adId).ExchangeReviewCreateFields(exchangeReviewCreateFields).Execute()

Send Ad for Exchange Review



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
    exchangeReviewCreateFields := *openapiclient.NewExchangeReviewCreateFields([]string{"adx"}) // ExchangeReviewCreateFields | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExchangeReviewsAPI.SendAdForExchangeReview(context.Background(), accountId, adId).ExchangeReviewCreateFields(exchangeReviewCreateFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangeReviewsAPI.SendAdForExchangeReview``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSendAdForExchangeReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exchangeReviewCreateFields** | [**ExchangeReviewCreateFields**](ExchangeReviewCreateFields.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

