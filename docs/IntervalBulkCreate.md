# IntervalBulkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**IntervalType**](IntervalType.md) |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**Intervals** | Pointer to [**[]Interval**](Interval.md) |  | [optional] 

## Methods

### NewIntervalBulkCreate

`func NewIntervalBulkCreate() *IntervalBulkCreate`

NewIntervalBulkCreate instantiates a new IntervalBulkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalBulkCreateWithDefaults

`func NewIntervalBulkCreateWithDefaults() *IntervalBulkCreate`

NewIntervalBulkCreateWithDefaults instantiates a new IntervalBulkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntervalBulkCreate) GetType() IntervalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntervalBulkCreate) GetTypeOk() (*IntervalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntervalBulkCreate) SetType(v IntervalType)`

SetType sets Type field to given value.

### HasType

`func (o *IntervalBulkCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLine

`func (o *IntervalBulkCreate) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *IntervalBulkCreate) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *IntervalBulkCreate) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *IntervalBulkCreate) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetIntervals

`func (o *IntervalBulkCreate) GetIntervals() []Interval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *IntervalBulkCreate) GetIntervalsOk() (*[]Interval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *IntervalBulkCreate) SetIntervals(v []Interval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *IntervalBulkCreate) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


