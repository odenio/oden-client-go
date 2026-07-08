# MetricGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**UnitKindOfQuantity** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewMetricGroup

`func NewMetricGroup() *MetricGroup`

NewMetricGroup instantiates a new MetricGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGroupWithDefaults

`func NewMetricGroupWithDefaults() *MetricGroup`

NewMetricGroupWithDefaults instantiates a new MetricGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MetricGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUnitKindOfQuantity

`func (o *MetricGroup) GetUnitKindOfQuantity() string`

GetUnitKindOfQuantity returns the UnitKindOfQuantity field if non-nil, zero value otherwise.

### GetUnitKindOfQuantityOk

`func (o *MetricGroup) GetUnitKindOfQuantityOk() (*string, bool)`

GetUnitKindOfQuantityOk returns a tuple with the UnitKindOfQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitKindOfQuantity

`func (o *MetricGroup) SetUnitKindOfQuantity(v string)`

SetUnitKindOfQuantity sets UnitKindOfQuantity field to given value.

### HasUnitKindOfQuantity

`func (o *MetricGroup) HasUnitKindOfQuantity() bool`

HasUnitKindOfQuantity returns a boolean if a field has been set.

### GetMatch

`func (o *MetricGroup) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *MetricGroup) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *MetricGroup) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *MetricGroup) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


