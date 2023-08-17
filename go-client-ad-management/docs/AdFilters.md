# AdFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | Generic search that can match against id, name, tag and domain. | [optional] 
**AdIds** | Pointer to **[]string** | The IDs of the ads to search for. | [optional] 
**Domain** | Pointer to **string** | The domain to search for. | [optional] 
**CreativeGroupIds** | Pointer to **[]string** | The list of creative group IDs to search usage for. | [optional] 
**Name** | Pointer to **string** | The name to search for. | [optional] 
**Archived** | Pointer to **bool** | &#x60;true&#x60; to include only archived creative groups, &#x60;false&#x60; to include only active, leave empty to include all. | [optional] 
**Tag** | Pointer to **string** | The tags to look for. | [optional] 
**Spec** | Pointer to [**AdsCreativeCreativeSpec**](AdsCreativeCreativeSpec.md) |  | [optional] 
**ExchangeReviewStatus** | Pointer to **string** | Include only ads that have at least one exchange review with the indicated status. | [optional] 

## Methods

### NewAdFilters

`func NewAdFilters() *AdFilters`

NewAdFilters instantiates a new AdFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdFiltersWithDefaults

`func NewAdFiltersWithDefaults() *AdFilters`

NewAdFiltersWithDefaults instantiates a new AdFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *AdFilters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AdFilters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AdFilters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AdFilters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetAdIds

`func (o *AdFilters) GetAdIds() []string`

GetAdIds returns the AdIds field if non-nil, zero value otherwise.

### GetAdIdsOk

`func (o *AdFilters) GetAdIdsOk() (*[]string, bool)`

GetAdIdsOk returns a tuple with the AdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdIds

`func (o *AdFilters) SetAdIds(v []string)`

SetAdIds sets AdIds field to given value.

### HasAdIds

`func (o *AdFilters) HasAdIds() bool`

HasAdIds returns a boolean if a field has been set.

### GetDomain

`func (o *AdFilters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AdFilters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AdFilters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AdFilters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCreativeGroupIds

`func (o *AdFilters) GetCreativeGroupIds() []string`

GetCreativeGroupIds returns the CreativeGroupIds field if non-nil, zero value otherwise.

### GetCreativeGroupIdsOk

`func (o *AdFilters) GetCreativeGroupIdsOk() (*[]string, bool)`

GetCreativeGroupIdsOk returns a tuple with the CreativeGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeGroupIds

`func (o *AdFilters) SetCreativeGroupIds(v []string)`

SetCreativeGroupIds sets CreativeGroupIds field to given value.

### HasCreativeGroupIds

`func (o *AdFilters) HasCreativeGroupIds() bool`

HasCreativeGroupIds returns a boolean if a field has been set.

### GetName

`func (o *AdFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *AdFilters) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AdFilters) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AdFilters) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AdFilters) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTag

`func (o *AdFilters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AdFilters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AdFilters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AdFilters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSpec

`func (o *AdFilters) GetSpec() AdsCreativeCreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AdFilters) GetSpecOk() (*AdsCreativeCreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AdFilters) SetSpec(v AdsCreativeCreativeSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AdFilters) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetExchangeReviewStatus

`func (o *AdFilters) GetExchangeReviewStatus() string`

GetExchangeReviewStatus returns the ExchangeReviewStatus field if non-nil, zero value otherwise.

### GetExchangeReviewStatusOk

`func (o *AdFilters) GetExchangeReviewStatusOk() (*string, bool)`

GetExchangeReviewStatusOk returns a tuple with the ExchangeReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeReviewStatus

`func (o *AdFilters) SetExchangeReviewStatus(v string)`

SetExchangeReviewStatus sets ExchangeReviewStatus field to given value.

### HasExchangeReviewStatus

`func (o *AdFilters) HasExchangeReviewStatus() bool`

HasExchangeReviewStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


