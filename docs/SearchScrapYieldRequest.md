# SearchScrapYieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ScrapYieldData**](ScrapYieldData.md) |  | [optional] 
**Interval** | [**Interval**](Interval.md) |  | 

## Methods

### NewSearchScrapYieldRequest

`func NewSearchScrapYieldRequest(interval Interval, ) *SearchScrapYieldRequest`

NewSearchScrapYieldRequest instantiates a new SearchScrapYieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchScrapYieldRequestWithDefaults

`func NewSearchScrapYieldRequestWithDefaults() *SearchScrapYieldRequest`

NewSearchScrapYieldRequestWithDefaults instantiates a new SearchScrapYieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchScrapYieldRequest) GetData() ScrapYieldData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchScrapYieldRequest) GetDataOk() (*ScrapYieldData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchScrapYieldRequest) SetData(v ScrapYieldData)`

SetData sets Data field to given value.

### HasData

`func (o *SearchScrapYieldRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInterval

`func (o *SearchScrapYieldRequest) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SearchScrapYieldRequest) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SearchScrapYieldRequest) SetInterval(v Interval)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


