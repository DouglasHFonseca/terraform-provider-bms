# AdPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]Ad**](Ad.md) | The values on this page. | 
**NextPageToken** | Pointer to **NullableString** | The token to be used to retrieve the next page. | [optional] 

## Methods

### NewAdPage

`func NewAdPage(values []Ad, ) *AdPage`

NewAdPage instantiates a new AdPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdPageWithDefaults

`func NewAdPageWithDefaults() *AdPage`

NewAdPageWithDefaults instantiates a new AdPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *AdPage) GetValues() []Ad`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AdPage) GetValuesOk() (*[]Ad, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AdPage) SetValues(v []Ad)`

SetValues sets Values field to given value.


### GetNextPageToken

`func (o *AdPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AdPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AdPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AdPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *AdPage) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *AdPage) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


