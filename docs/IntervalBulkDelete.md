# IntervalBulkDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**IntervalType**](IntervalType.md) |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**HardDelete** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewIntervalBulkDelete

`func NewIntervalBulkDelete() *IntervalBulkDelete`

NewIntervalBulkDelete instantiates a new IntervalBulkDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalBulkDeleteWithDefaults

`func NewIntervalBulkDeleteWithDefaults() *IntervalBulkDelete`

NewIntervalBulkDeleteWithDefaults instantiates a new IntervalBulkDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntervalBulkDelete) GetType() IntervalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntervalBulkDelete) GetTypeOk() (*IntervalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntervalBulkDelete) SetType(v IntervalType)`

SetType sets Type field to given value.

### HasType

`func (o *IntervalBulkDelete) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLine

`func (o *IntervalBulkDelete) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *IntervalBulkDelete) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *IntervalBulkDelete) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *IntervalBulkDelete) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetStartTime

`func (o *IntervalBulkDelete) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IntervalBulkDelete) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IntervalBulkDelete) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IntervalBulkDelete) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *IntervalBulkDelete) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *IntervalBulkDelete) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *IntervalBulkDelete) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *IntervalBulkDelete) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHardDelete

`func (o *IntervalBulkDelete) GetHardDelete() bool`

GetHardDelete returns the HardDelete field if non-nil, zero value otherwise.

### GetHardDeleteOk

`func (o *IntervalBulkDelete) GetHardDeleteOk() (*bool, bool)`

GetHardDeleteOk returns a tuple with the HardDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDelete

`func (o *IntervalBulkDelete) SetHardDelete(v bool)`

SetHardDelete sets HardDelete field to given value.

### HasHardDelete

`func (o *IntervalBulkDelete) HasHardDelete() bool`

HasHardDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


