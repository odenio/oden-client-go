# BulkDeleteIntervals200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDeleted** | Pointer to **int32** |  | [optional] 
**FailedIntervals** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkDeleteIntervals200Response

`func NewBulkDeleteIntervals200Response() *BulkDeleteIntervals200Response`

NewBulkDeleteIntervals200Response instantiates a new BulkDeleteIntervals200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteIntervals200ResponseWithDefaults

`func NewBulkDeleteIntervals200ResponseWithDefaults() *BulkDeleteIntervals200Response`

NewBulkDeleteIntervals200ResponseWithDefaults instantiates a new BulkDeleteIntervals200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDeleted

`func (o *BulkDeleteIntervals200Response) GetNumDeleted() int32`

GetNumDeleted returns the NumDeleted field if non-nil, zero value otherwise.

### GetNumDeletedOk

`func (o *BulkDeleteIntervals200Response) GetNumDeletedOk() (*int32, bool)`

GetNumDeletedOk returns a tuple with the NumDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleted

`func (o *BulkDeleteIntervals200Response) SetNumDeleted(v int32)`

SetNumDeleted sets NumDeleted field to given value.

### HasNumDeleted

`func (o *BulkDeleteIntervals200Response) HasNumDeleted() bool`

HasNumDeleted returns a boolean if a field has been set.

### GetFailedIntervals

`func (o *BulkDeleteIntervals200Response) GetFailedIntervals() []string`

GetFailedIntervals returns the FailedIntervals field if non-nil, zero value otherwise.

### GetFailedIntervalsOk

`func (o *BulkDeleteIntervals200Response) GetFailedIntervalsOk() (*[]string, bool)`

GetFailedIntervalsOk returns a tuple with the FailedIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIntervals

`func (o *BulkDeleteIntervals200Response) SetFailedIntervals(v []string)`

SetFailedIntervals sets FailedIntervals field to given value.

### HasFailedIntervals

`func (o *BulkDeleteIntervals200Response) HasFailedIntervals() bool`

HasFailedIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


