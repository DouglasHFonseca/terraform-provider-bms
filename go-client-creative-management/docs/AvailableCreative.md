# AvailableCreative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreativeId** | **string** | The ID of the creative. | 
**Weight** | **float32** | The weight of the creative. | 
**Banner** | Pointer to [**BannerFields**](BannerFields.md) |  | [optional] 

## Methods

### NewAvailableCreative

`func NewAvailableCreative(creativeId string, weight float32, ) *AvailableCreative`

NewAvailableCreative instantiates a new AvailableCreative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableCreativeWithDefaults

`func NewAvailableCreativeWithDefaults() *AvailableCreative`

NewAvailableCreativeWithDefaults instantiates a new AvailableCreative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreativeId

`func (o *AvailableCreative) GetCreativeId() string`

GetCreativeId returns the CreativeId field if non-nil, zero value otherwise.

### GetCreativeIdOk

`func (o *AvailableCreative) GetCreativeIdOk() (*string, bool)`

GetCreativeIdOk returns a tuple with the CreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeId

`func (o *AvailableCreative) SetCreativeId(v string)`

SetCreativeId sets CreativeId field to given value.


### GetWeight

`func (o *AvailableCreative) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AvailableCreative) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AvailableCreative) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetBanner

`func (o *AvailableCreative) GetBanner() BannerFields`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *AvailableCreative) GetBannerOk() (*BannerFields, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *AvailableCreative) SetBanner(v BannerFields)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *AvailableCreative) HasBanner() bool`

HasBanner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


