# Ad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this ad is going to redirect to. | 
**Name** | **string** | The name of this ad. | 
**Spec** | [**AdsCreativeCreativeSpec**](AdsCreativeCreativeSpec.md) |  | 
**Rules** | [**[]AdRule**](AdRule.md) | The rules that specify what are the conditions to display a creative group. | 
**Tags** | **[]string** | Tags that the user can add to facilitate searching. | 
**AdId** | **string** | The ID of the ad. | 
**AccountId** | **string** | The ID of the account that owns this ad. | 
**ExchangeReviews** | [**[]ExchangeReview**](ExchangeReview.md) | Reviews for all applicable exchanges. | 
**Archived** | **bool** | Indicates if it is archived. | 

## Methods

### NewAd

`func NewAd(domain string, name string, spec AdsCreativeCreativeSpec, rules []AdRule, tags []string, adId string, accountId string, exchangeReviews []ExchangeReview, archived bool, ) *Ad`

NewAd instantiates a new Ad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdWithDefaults

`func NewAdWithDefaults() *Ad`

NewAdWithDefaults instantiates a new Ad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Ad) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Ad) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Ad) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *Ad) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ad) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ad) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *Ad) GetSpec() AdsCreativeCreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Ad) GetSpecOk() (*AdsCreativeCreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Ad) SetSpec(v AdsCreativeCreativeSpec)`

SetSpec sets Spec field to given value.


### GetRules

`func (o *Ad) GetRules() []AdRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Ad) GetRulesOk() (*[]AdRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Ad) SetRules(v []AdRule)`

SetRules sets Rules field to given value.


### GetTags

`func (o *Ad) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Ad) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Ad) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetAdId

`func (o *Ad) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *Ad) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *Ad) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetAccountId

`func (o *Ad) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Ad) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Ad) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetExchangeReviews

`func (o *Ad) GetExchangeReviews() []ExchangeReview`

GetExchangeReviews returns the ExchangeReviews field if non-nil, zero value otherwise.

### GetExchangeReviewsOk

`func (o *Ad) GetExchangeReviewsOk() (*[]ExchangeReview, bool)`

GetExchangeReviewsOk returns a tuple with the ExchangeReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReviews

`func (o *Ad) SetExchangeReviews(v []ExchangeReview)`

SetExchangeReviews sets ExchangeReviews field to given value.


### GetArchived

`func (o *Ad) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Ad) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Ad) SetArchived(v bool)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


