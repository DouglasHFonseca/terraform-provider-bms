# TimeOfWeek

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **float32** | Indicates when the window starts in number of minutes since Sunday, midnight. | 
**End** | **float32** | Indicates when the window ends in number of minutes since Sunday midnight, must be after &#x60;start&#x60;. | 

## Methods

### NewTimeOfWeek

`func NewTimeOfWeek(start float32, end float32, ) *TimeOfWeek`

NewTimeOfWeek instantiates a new TimeOfWeek object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfWeekWithDefaults

`func NewTimeOfWeekWithDefaults() *TimeOfWeek`

NewTimeOfWeekWithDefaults instantiates a new TimeOfWeek object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TimeOfWeek) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TimeOfWeek) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TimeOfWeek) SetStart(v float32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TimeOfWeek) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimeOfWeek) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimeOfWeek) SetEnd(v float32)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


