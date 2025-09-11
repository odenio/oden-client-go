# V2ScrapYieldSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ScrapYieldData**](ScrapYieldData.md) |  | [optional] 
**Interval** | [**Interval**](Interval.md) |  | 

## Methods

### NewV2ScrapYieldSearchPostRequest

`func NewV2ScrapYieldSearchPostRequest(interval Interval, ) *V2ScrapYieldSearchPostRequest`

NewV2ScrapYieldSearchPostRequest instantiates a new V2ScrapYieldSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ScrapYieldSearchPostRequestWithDefaults

`func NewV2ScrapYieldSearchPostRequestWithDefaults() *V2ScrapYieldSearchPostRequest`

NewV2ScrapYieldSearchPostRequestWithDefaults instantiates a new V2ScrapYieldSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2ScrapYieldSearchPostRequest) GetData() ScrapYieldData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2ScrapYieldSearchPostRequest) GetDataOk() (*ScrapYieldData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2ScrapYieldSearchPostRequest) SetData(v ScrapYieldData)`

SetData sets Data field to given value.

### HasData

`func (o *V2ScrapYieldSearchPostRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInterval

`func (o *V2ScrapYieldSearchPostRequest) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V2ScrapYieldSearchPostRequest) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V2ScrapYieldSearchPostRequest) SetInterval(v Interval)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


