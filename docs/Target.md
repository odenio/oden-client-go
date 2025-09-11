# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricGroup** | Pointer to [**MetricGroup**](MetricGroup.md) |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**TargetValue** | Pointer to **float64** |  | [optional] 
**Lsl** | Pointer to **float64** |  | [optional] 
**Usl** | Pointer to **float64** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricGroup

`func (o *Target) GetMetricGroup() MetricGroup`

GetMetricGroup returns the MetricGroup field if non-nil, zero value otherwise.

### GetMetricGroupOk

`func (o *Target) GetMetricGroupOk() (*MetricGroup, bool)`

GetMetricGroupOk returns a tuple with the MetricGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroup

`func (o *Target) SetMetricGroup(v MetricGroup)`

SetMetricGroup sets MetricGroup field to given value.

### HasMetricGroup

`func (o *Target) HasMetricGroup() bool`

HasMetricGroup returns a boolean if a field has been set.

### GetProduct

`func (o *Target) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Target) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Target) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Target) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetLine

`func (o *Target) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Target) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Target) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *Target) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetTargetValue

`func (o *Target) GetTargetValue() float64`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *Target) GetTargetValueOk() (*float64, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *Target) SetTargetValue(v float64)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *Target) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetLsl

`func (o *Target) GetLsl() float64`

GetLsl returns the Lsl field if non-nil, zero value otherwise.

### GetLslOk

`func (o *Target) GetLslOk() (*float64, bool)`

GetLslOk returns a tuple with the Lsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsl

`func (o *Target) SetLsl(v float64)`

SetLsl sets Lsl field to given value.

### HasLsl

`func (o *Target) HasLsl() bool`

HasLsl returns a boolean if a field has been set.

### GetUsl

`func (o *Target) GetUsl() float64`

GetUsl returns the Usl field if non-nil, zero value otherwise.

### GetUslOk

`func (o *Target) GetUslOk() (*float64, bool)`

GetUslOk returns a tuple with the Usl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsl

`func (o *Target) SetUsl(v float64)`

SetUsl sets Usl field to given value.

### HasUsl

`func (o *Target) HasUsl() bool`

HasUsl returns a boolean if a field has been set.

### GetMatch

`func (o *Target) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Target) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Target) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Target) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


