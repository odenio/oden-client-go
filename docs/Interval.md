# Interval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | [**IntervalType**](IntervalType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Line** | [**Line**](Line.md) |  | 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to [**IntervalMetadata**](IntervalMetadata.md) |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewInterval

`func NewInterval(type_ IntervalType, line Line, ) *Interval`

NewInterval instantiates a new Interval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalWithDefaults

`func NewIntervalWithDefaults() *Interval`

NewIntervalWithDefaults instantiates a new Interval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Interval) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interval) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interval) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Interval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Interval) GetType() IntervalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interval) GetTypeOk() (*IntervalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interval) SetType(v IntervalType)`

SetType sets Type field to given value.


### GetName

`func (o *Interval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Interval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLine

`func (o *Interval) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Interval) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Interval) SetLine(v Line)`

SetLine sets Line field to given value.


### GetStartTime

`func (o *Interval) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Interval) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Interval) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Interval) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Interval) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Interval) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Interval) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Interval) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMetadata

`func (o *Interval) GetMetadata() IntervalMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Interval) GetMetadataOk() (*IntervalMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Interval) SetMetadata(v IntervalMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Interval) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMatch

`func (o *Interval) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Interval) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Interval) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Interval) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


