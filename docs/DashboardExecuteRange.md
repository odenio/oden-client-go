# DashboardExecuteRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**Anchor** | Pointer to **string** | Anchor expression for a relative range, e.g. &#x60;now&#x60;. | [optional] 
**Offset** | Pointer to **string** | Offset expression for a relative range, e.g. &#x60;-7D&#x60;. | [optional] 
**Timezone** | Pointer to **string** | IANA timezone identifier (defaults to UTC). | [optional] 

## Methods

### NewDashboardExecuteRange

`func NewDashboardExecuteRange() *DashboardExecuteRange`

NewDashboardExecuteRange instantiates a new DashboardExecuteRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardExecuteRangeWithDefaults

`func NewDashboardExecuteRangeWithDefaults() *DashboardExecuteRange`

NewDashboardExecuteRangeWithDefaults instantiates a new DashboardExecuteRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *DashboardExecuteRange) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DashboardExecuteRange) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DashboardExecuteRange) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *DashboardExecuteRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *DashboardExecuteRange) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DashboardExecuteRange) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DashboardExecuteRange) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DashboardExecuteRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetAnchor

`func (o *DashboardExecuteRange) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *DashboardExecuteRange) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *DashboardExecuteRange) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *DashboardExecuteRange) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetOffset

`func (o *DashboardExecuteRange) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DashboardExecuteRange) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DashboardExecuteRange) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DashboardExecuteRange) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimezone

`func (o *DashboardExecuteRange) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DashboardExecuteRange) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DashboardExecuteRange) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DashboardExecuteRange) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


