# BulkDeleteQualityTestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBulkDeleteQualityTestsRequest

`func NewBulkDeleteQualityTestsRequest() *BulkDeleteQualityTestsRequest`

NewBulkDeleteQualityTestsRequest instantiates a new BulkDeleteQualityTestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteQualityTestsRequestWithDefaults

`func NewBulkDeleteQualityTestsRequestWithDefaults() *BulkDeleteQualityTestsRequest`

NewBulkDeleteQualityTestsRequestWithDefaults instantiates a new BulkDeleteQualityTestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BulkDeleteQualityTestsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkDeleteQualityTestsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkDeleteQualityTestsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *BulkDeleteQualityTestsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLine

`func (o *BulkDeleteQualityTestsRequest) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *BulkDeleteQualityTestsRequest) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *BulkDeleteQualityTestsRequest) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *BulkDeleteQualityTestsRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetStartTime

`func (o *BulkDeleteQualityTestsRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BulkDeleteQualityTestsRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BulkDeleteQualityTestsRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BulkDeleteQualityTestsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *BulkDeleteQualityTestsRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BulkDeleteQualityTestsRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BulkDeleteQualityTestsRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BulkDeleteQualityTestsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


