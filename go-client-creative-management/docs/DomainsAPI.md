# \DomainsAPI

All URIs are relative to *http://api.creative.ads.sd.internal.us-east-2.bluemsdev.team*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDomains**](DomainsAPI.md#ListDomains) | **Get** /v1/accounts/{accountId}/domains | List Domains



## ListDomains

> DomainPage ListDomains(ctx, accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()

List Domains



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
    filters := "{"name":"globo","includeSuggestions":true}" // string | The applicable filters, a json object according with `DomainFilters`. (optional)
    pageSize := float32(10) // float32 | The page size. (optional) (default to 10)
    pageToken := "AIUHSDJANBDHOH3" // string | The page token from which to continue listing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.ListDomains(context.Background(), accountId).Filters(filters).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomains`: DomainPage
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | The applicable filters, a json object according with &#x60;DomainFilters&#x60;. | 
 **pageSize** | **float32** | The page size. | [default to 10]
 **pageToken** | **string** | The page token from which to continue listing. | 

### Return type

[**DomainPage**](DomainPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

