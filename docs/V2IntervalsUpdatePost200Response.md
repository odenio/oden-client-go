# V2IntervalsUpdatePost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumUpdated** | Pointer to **int32** |  | [optional] 
**UpdatedIntervals** | Pointer to [**[]Interval**](Interval.md) |  | [optional] 
**NumFailed** | Pointer to **int32** |  | [optional] 
**FailedIntervals** | Pointer to [**[]V2IntervalsUpdatePost200ResponseFailedIntervalsInner**](V2IntervalsUpdatePost200ResponseFailedIntervalsInner.md) |  | [optional] 

## Methods

### NewV2IntervalsUpdatePost200Response

`func NewV2IntervalsUpdatePost200Response() *V2IntervalsUpdatePost200Response`

NewV2IntervalsUpdatePost200Response instantiates a new V2IntervalsUpdatePost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2IntervalsUpdatePost200ResponseWithDefaults

`func NewV2IntervalsUpdatePost200ResponseWithDefaults() *V2IntervalsUpdatePost200Response`

NewV2IntervalsUpdatePost200ResponseWithDefaults instantiates a new V2IntervalsUpdatePost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumUpdated

`func (o *V2IntervalsUpdatePost200Response) GetNumUpdated() int32`

GetNumUpdated returns the NumUpdated field if non-nil, zero value otherwise.

### GetNumUpdatedOk

`func (o *V2IntervalsUpdatePost200Response) GetNumUpdatedOk() (*int32, bool)`

GetNumUpdatedOk returns a tuple with the NumUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpdated

`func (o *V2IntervalsUpdatePost200Response) SetNumUpdated(v int32)`

SetNumUpdated sets NumUpdated field to given value.

### HasNumUpdated

`func (o *V2IntervalsUpdatePost200Response) HasNumUpdated() bool`

HasNumUpdated returns a boolean if a field has been set.

### GetUpdatedIntervals

`func (o *V2IntervalsUpdatePost200Response) GetUpdatedIntervals() []Interval`

GetUpdatedIntervals returns the UpdatedIntervals field if non-nil, zero value otherwise.

### GetUpdatedIntervalsOk

`func (o *V2IntervalsUpdatePost200Response) GetUpdatedIntervalsOk() (*[]Interval, bool)`

GetUpdatedIntervalsOk returns a tuple with the UpdatedIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedIntervals

`func (o *V2IntervalsUpdatePost200Response) SetUpdatedIntervals(v []Interval)`

SetUpdatedIntervals sets UpdatedIntervals field to given value.

### HasUpdatedIntervals

`func (o *V2IntervalsUpdatePost200Response) HasUpdatedIntervals() bool`

HasUpdatedIntervals returns a boolean if a field has been set.

### GetNumFailed

`func (o *V2IntervalsUpdatePost200Response) GetNumFailed() int32`

GetNumFailed returns the NumFailed field if non-nil, zero value otherwise.

### GetNumFailedOk

`func (o *V2IntervalsUpdatePost200Response) GetNumFailedOk() (*int32, bool)`

GetNumFailedOk returns a tuple with the NumFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailed

`func (o *V2IntervalsUpdatePost200Response) SetNumFailed(v int32)`

SetNumFailed sets NumFailed field to given value.

### HasNumFailed

`func (o *V2IntervalsUpdatePost200Response) HasNumFailed() bool`

HasNumFailed returns a boolean if a field has been set.

### GetFailedIntervals

`func (o *V2IntervalsUpdatePost200Response) GetFailedIntervals() []V2IntervalsUpdatePost200ResponseFailedIntervalsInner`

GetFailedIntervals returns the FailedIntervals field if non-nil, zero value otherwise.

### GetFailedIntervalsOk

`func (o *V2IntervalsUpdatePost200Response) GetFailedIntervalsOk() (*[]V2IntervalsUpdatePost200ResponseFailedIntervalsInner, bool)`

GetFailedIntervalsOk returns a tuple with the FailedIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIntervals

`func (o *V2IntervalsUpdatePost200Response) SetFailedIntervals(v []V2IntervalsUpdatePost200ResponseFailedIntervalsInner)`

SetFailedIntervals sets FailedIntervals field to given value.

### HasFailedIntervals

`func (o *V2IntervalsUpdatePost200Response) HasFailedIntervals() bool`

HasFailedIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


