# ScheduleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicates the type of condition. | 
**Start** | Pointer to **time.Time** | The point in time when this condition becomes valid. | [optional] 
**End** | Pointer to **time.Time** | The point in time when this condition stops being valid, must be after &#x60;start&#x60; if it has been defined. | [optional] 
**TimesOfWeek** | Pointer to [**[]TimeOfWeek**](TimeOfWeek.md) | The windows within a week where this condition is applicable. The time zone used for reference will be the one in use by the user viewing the ad. | [optional] 

## Methods

### NewScheduleCondition

`func NewScheduleCondition(type_ string, ) *ScheduleCondition`

NewScheduleCondition instantiates a new ScheduleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleConditionWithDefaults

`func NewScheduleConditionWithDefaults() *ScheduleCondition`

NewScheduleConditionWithDefaults instantiates a new ScheduleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleCondition) SetType(v string)`

SetType sets Type field to given value.


### GetStart

`func (o *ScheduleCondition) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ScheduleCondition) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ScheduleCondition) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ScheduleCondition) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ScheduleCondition) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ScheduleCondition) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ScheduleCondition) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ScheduleCondition) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTimesOfWeek

`func (o *ScheduleCondition) GetTimesOfWeek() []TimeOfWeek`

GetTimesOfWeek returns the TimesOfWeek field if non-nil, zero value otherwise.

### GetTimesOfWeekOk

`func (o *ScheduleCondition) GetTimesOfWeekOk() (*[]TimeOfWeek, bool)`

GetTimesOfWeekOk returns a tuple with the TimesOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesOfWeek

`func (o *ScheduleCondition) SetTimesOfWeek(v []TimeOfWeek)`

SetTimesOfWeek sets TimesOfWeek field to given value.

### HasTimesOfWeek

`func (o *ScheduleCondition) HasTimesOfWeek() bool`

HasTimesOfWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


