# IntervalMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataType** | Pointer to **string** |  | [optional] [readonly] 
**Run** | Pointer to [**Interval**](Interval.md) |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Target** | Pointer to [**Target**](Target.md) |  | [optional] 
**Reason** | Pointer to [**StateReason**](StateReason.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**StateCategory**](StateCategory.md) |  | [optional] 

## Methods

### NewIntervalMetadata

`func NewIntervalMetadata() *IntervalMetadata`

NewIntervalMetadata instantiates a new IntervalMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalMetadataWithDefaults

`func NewIntervalMetadataWithDefaults() *IntervalMetadata`

NewIntervalMetadataWithDefaults instantiates a new IntervalMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataType

`func (o *IntervalMetadata) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *IntervalMetadata) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *IntervalMetadata) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *IntervalMetadata) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetRun

`func (o *IntervalMetadata) GetRun() Interval`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *IntervalMetadata) GetRunOk() (*Interval, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *IntervalMetadata) SetRun(v Interval)`

SetRun sets Run field to given value.

### HasRun

`func (o *IntervalMetadata) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetProduct

`func (o *IntervalMetadata) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *IntervalMetadata) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *IntervalMetadata) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *IntervalMetadata) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTarget

`func (o *IntervalMetadata) GetTarget() Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IntervalMetadata) GetTargetOk() (*Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IntervalMetadata) SetTarget(v Target)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *IntervalMetadata) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetReason

`func (o *IntervalMetadata) GetReason() StateReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IntervalMetadata) GetReasonOk() (*StateReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IntervalMetadata) SetReason(v StateReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IntervalMetadata) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *IntervalMetadata) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IntervalMetadata) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IntervalMetadata) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IntervalMetadata) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCategory

`func (o *IntervalMetadata) GetCategory() StateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntervalMetadata) GetCategoryOk() (*StateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntervalMetadata) SetCategory(v StateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IntervalMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


