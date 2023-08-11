# Creative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this creative is going to redirect to. | 
**Name** | **string** | The name of the creative. | 
**Tags** | **[]string** | Tags that the user can add to facilitate searching. | 
**Banner** | Pointer to [**CreativeCreateFieldsAllOfBanner**](CreativeCreateFieldsAllOfBanner.md) |  | [optional] 
**CreativeId** | **string** | The ID of the creative. | 
**Status** | **string** | Indicates the status of the creative. | 
**AccountId** | **string** | The ID of the account that owns this creative. | 
**Archived** | **bool** | Indicates if it is archived. | 
**Enabled** | **bool** | Indicates if the creative is enabled for usage. | 

## Methods

### NewCreative

`func NewCreative(domain string, name string, tags []string, creativeId string, status string, accountId string, archived bool, enabled bool, ) *Creative`

NewCreative instantiates a new Creative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeWithDefaults

`func NewCreativeWithDefaults() *Creative`

NewCreativeWithDefaults instantiates a new Creative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Creative) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Creative) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Creative) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *Creative) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Creative) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Creative) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *Creative) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Creative) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Creative) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetBanner

`func (o *Creative) GetBanner() CreativeCreateFieldsAllOfBanner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *Creative) GetBannerOk() (*CreativeCreateFieldsAllOfBanner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *Creative) SetBanner(v CreativeCreateFieldsAllOfBanner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *Creative) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetCreativeId

`func (o *Creative) GetCreativeId() string`

GetCreativeId returns the CreativeId field if non-nil, zero value otherwise.

### GetCreativeIdOk

`func (o *Creative) GetCreativeIdOk() (*string, bool)`

GetCreativeIdOk returns a tuple with the CreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeId

`func (o *Creative) SetCreativeId(v string)`

SetCreativeId sets CreativeId field to given value.


### GetStatus

`func (o *Creative) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Creative) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Creative) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *Creative) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Creative) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Creative) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArchived

`func (o *Creative) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Creative) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Creative) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetEnabled

`func (o *Creative) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Creative) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Creative) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


