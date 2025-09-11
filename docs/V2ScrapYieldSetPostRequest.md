# V2ScrapYieldSetPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ScrapYieldData**](ScrapYieldData.md) |  | 
**Interval** | [**Interval**](Interval.md) |  | 

## Methods

### NewV2ScrapYieldSetPostRequest

`func NewV2ScrapYieldSetPostRequest(data ScrapYieldData, interval Interval, ) *V2ScrapYieldSetPostRequest`

NewV2ScrapYieldSetPostRequest instantiates a new V2ScrapYieldSetPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ScrapYieldSetPostRequestWithDefaults

`func NewV2ScrapYieldSetPostRequestWithDefaults() *V2ScrapYieldSetPostRequest`

NewV2ScrapYieldSetPostRequestWithDefaults instantiates a new V2ScrapYieldSetPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2ScrapYieldSetPostRequest) GetData() ScrapYieldData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2ScrapYieldSetPostRequest) GetDataOk() (*ScrapYieldData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2ScrapYieldSetPostRequest) SetData(v ScrapYieldData)`

SetData sets Data field to given value.


### GetInterval

`func (o *V2ScrapYieldSetPostRequest) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V2ScrapYieldSetPostRequest) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V2ScrapYieldSetPostRequest) SetInterval(v Interval)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


