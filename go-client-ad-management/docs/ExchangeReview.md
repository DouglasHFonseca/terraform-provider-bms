# ExchangeReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeReviewId** | **string** | The ID of the review to be used later when submitting bids to display the ad this review refers to. | 
**Status** | **string** | Indicates the status of the review: * &#x60;queued&#x60;: The ad is in the exchange submission queue. * &#x60;reviewing&#x60;: The exchange has received the ad and is reviewing it. * &#x60;approved&#x60;: The exchange has reviewed the ad and approved it. * &#x60;rejected&#x60;: The exchange has reviewed the ad and rejected it. Some details of why or what can be done can be found in the &#x60;comments&#x60; field. * &#x60;revoked&#x60;: The exchange has previously reviewed and approved the ad, but the approval was revoked. Some details of why or what can be done can be found in the &#x60;comments&#x60; field. | 
**Comments** | Pointer to **string** | Any comments that the exchange reviewer deem important. | [optional] 
**ExchangeId** | **string** | The exchange ID for this review. | 
**UpdatedAt** | **time.Time** | When the review was last updated. | 

## Methods

### NewExchangeReview

`func NewExchangeReview(exchangeReviewId string, status string, exchangeId string, updatedAt time.Time, ) *ExchangeReview`

NewExchangeReview instantiates a new ExchangeReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeReviewWithDefaults

`func NewExchangeReviewWithDefaults() *ExchangeReview`

NewExchangeReviewWithDefaults instantiates a new ExchangeReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeReviewId

`func (o *ExchangeReview) GetExchangeReviewId() string`

GetExchangeReviewId returns the ExchangeReviewId field if non-nil, zero value otherwise.

### GetExchangeReviewIdOk

`func (o *ExchangeReview) GetExchangeReviewIdOk() (*string, bool)`

GetExchangeReviewIdOk returns a tuple with the ExchangeReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReviewId

`func (o *ExchangeReview) SetExchangeReviewId(v string)`

SetExchangeReviewId sets ExchangeReviewId field to given value.


### GetStatus

`func (o *ExchangeReview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangeReview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangeReview) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetComments

`func (o *ExchangeReview) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ExchangeReview) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ExchangeReview) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ExchangeReview) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetExchangeId

`func (o *ExchangeReview) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ExchangeReview) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ExchangeReview) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.


### GetUpdatedAt

`func (o *ExchangeReview) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExchangeReview) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExchangeReview) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


