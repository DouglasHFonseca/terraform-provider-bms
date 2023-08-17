# AdsCreativeAvailableCreative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreativeId** | **string** | The ID of the creative. | 
**Weight** | **float32** | The weight of the creative. | 
**Banner** | Pointer to [**AdsCreativeBannerFields**](AdsCreativeBannerFields.md) |  | [optional] 

## Methods

### NewAdsCreativeAvailableCreative

`func NewAdsCreativeAvailableCreative(creativeId string, weight float32, ) *AdsCreativeAvailableCreative`

NewAdsCreativeAvailableCreative instantiates a new AdsCreativeAvailableCreative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdsCreativeAvailableCreativeWithDefaults

`func NewAdsCreativeAvailableCreativeWithDefaults() *AdsCreativeAvailableCreative`

NewAdsCreativeAvailableCreativeWithDefaults instantiates a new AdsCreativeAvailableCreative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreativeId

`func (o *AdsCreativeAvailableCreative) GetCreativeId() string`

GetCreativeId returns the CreativeId field if non-nil, zero value otherwise.

### GetCreativeIdOk

`func (o *AdsCreativeAvailableCreative) GetCreativeIdOk() (*string, bool)`

GetCreativeIdOk returns a tuple with the CreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeId

`func (o *AdsCreativeAvailableCreative) SetCreativeId(v string)`

SetCreativeId sets CreativeId field to given value.


### GetWeight

`func (o *AdsCreativeAvailableCreative) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AdsCreativeAvailableCreative) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AdsCreativeAvailableCreative) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetBanner

`func (o *AdsCreativeAvailableCreative) GetBanner() AdsCreativeBannerFields`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *AdsCreativeAvailableCreative) GetBannerOk() (*AdsCreativeBannerFields, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *AdsCreativeAvailableCreative) SetBanner(v AdsCreativeBannerFields)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *AdsCreativeAvailableCreative) HasBanner() bool`

HasBanner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


