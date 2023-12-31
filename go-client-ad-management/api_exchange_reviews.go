/*
AdServer Ad Management

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ExchangeReviewsAPIService ExchangeReviewsAPI service
type ExchangeReviewsAPIService service

type ApiPutExchangeReviewRequest struct {
	ctx context.Context
	ApiService *ExchangeReviewsAPIService
	accountId string
	adId string
	exchangeId string
	exchangeReviewPutFields *ExchangeReviewPutFields
}

func (r ApiPutExchangeReviewRequest) ExchangeReviewPutFields(exchangeReviewPutFields ExchangeReviewPutFields) ApiPutExchangeReviewRequest {
	r.exchangeReviewPutFields = &exchangeReviewPutFields
	return r
}

func (r ApiPutExchangeReviewRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutExchangeReviewExecute(r)
}

/*
PutExchangeReview Put Exchange Review

Updates the exchange review in a given ad for a given exchange.

**Notes:**
1. This endpoint will be called by the exchange integrations when they become aware that a review has been updated.
1. Publishes event `AdPatchedEvent` with the patched ad as message.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account.
 @param adId The ID of the ad.
 @param exchangeId The ID of the exchange.
 @return ApiPutExchangeReviewRequest
*/
func (a *ExchangeReviewsAPIService) PutExchangeReview(ctx context.Context, accountId string, adId string, exchangeId string) ApiPutExchangeReviewRequest {
	return ApiPutExchangeReviewRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		adId: adId,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
func (a *ExchangeReviewsAPIService) PutExchangeReviewExecute(r ApiPutExchangeReviewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeReviewsAPIService.PutExchangeReview")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{accountId}/ads/{adId}/exchange-reviews/{exchangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exchangeId"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exchangeReviewPutFields == nil {
		return nil, reportError("exchangeReviewPutFields is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.exchangeReviewPutFields
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRefreshExchangeReviewStatusRequest struct {
	ctx context.Context
	ApiService *ExchangeReviewsAPIService
	accountId string
	adId string
	exchangeReviewCreateFields *ExchangeReviewCreateFields
}

func (r ApiRefreshExchangeReviewStatusRequest) ExchangeReviewCreateFields(exchangeReviewCreateFields ExchangeReviewCreateFields) ApiRefreshExchangeReviewStatusRequest {
	r.exchangeReviewCreateFields = &exchangeReviewCreateFields
	return r
}

func (r ApiRefreshExchangeReviewStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.RefreshExchangeReviewStatusExecute(r)
}

/*
RefreshExchangeReviewStatus Refresh Exchange Review Status

Refreshes the review status of the ad on the requested ad exchanges.

**Notes:**
1. Publishes command `RefreshExchangeReviewStatusCommand` with the message: 
    ```
    { 
       accountId: string, 
       adId: string, 
       exchangeIds: string[] 
    }
    ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account.
 @param adId The ID of the ad.
 @return ApiRefreshExchangeReviewStatusRequest
*/
func (a *ExchangeReviewsAPIService) RefreshExchangeReviewStatus(ctx context.Context, accountId string, adId string) ApiRefreshExchangeReviewStatusRequest {
	return ApiRefreshExchangeReviewStatusRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		adId: adId,
	}
}

// Execute executes the request
func (a *ExchangeReviewsAPIService) RefreshExchangeReviewStatusExecute(r ApiRefreshExchangeReviewStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeReviewsAPIService.RefreshExchangeReviewStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{accountId}/ads/{adId}/exchange-reviews:refresh"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exchangeReviewCreateFields == nil {
		return nil, reportError("exchangeReviewCreateFields is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.exchangeReviewCreateFields
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSendAdForExchangeReviewRequest struct {
	ctx context.Context
	ApiService *ExchangeReviewsAPIService
	accountId string
	adId string
	exchangeReviewCreateFields *ExchangeReviewCreateFields
}

func (r ApiSendAdForExchangeReviewRequest) ExchangeReviewCreateFields(exchangeReviewCreateFields ExchangeReviewCreateFields) ApiSendAdForExchangeReviewRequest {
	r.exchangeReviewCreateFields = &exchangeReviewCreateFields
	return r
}

func (r ApiSendAdForExchangeReviewRequest) Execute() (*http.Response, error) {
	return r.ApiService.SendAdForExchangeReviewExecute(r)
}

/*
SendAdForExchangeReview Send Ad for Exchange Review

Initiates the review process for the given ad.

**Notes:**
1. You can only perform this action if the ad has active rules (as in `GET rules:active`). If the ad has no active rules, an error should be returned.
1. If the ad already has a review for the exchange with status `approved`, an error should be returned.
1. Adds an entry to the ad's exchangeReviews array for each exchangeId with the status `queued`. If an exchange review already exists, its status should be changed to `queued`.
1. Publishes event `AdPatchedEvent` with the patched ad as message.
1. Publishes command `SendAdForExchangeReviewCommand` with the message: 
    ```
    { 
       accountId: string, 
       adId: string, 
       exchangeIds: string[] 
    }
    ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account.
 @param adId The ID of the ad.
 @return ApiSendAdForExchangeReviewRequest
*/
func (a *ExchangeReviewsAPIService) SendAdForExchangeReview(ctx context.Context, accountId string, adId string) ApiSendAdForExchangeReviewRequest {
	return ApiSendAdForExchangeReviewRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		adId: adId,
	}
}

// Execute executes the request
func (a *ExchangeReviewsAPIService) SendAdForExchangeReviewExecute(r ApiSendAdForExchangeReviewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeReviewsAPIService.SendAdForExchangeReview")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{accountId}/ads/{adId}/exchange-reviews"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"adId"+"}", url.PathEscape(parameterValueToString(r.adId, "adId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exchangeReviewCreateFields == nil {
		return nil, reportError("exchangeReviewCreateFields is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.exchangeReviewCreateFields
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
