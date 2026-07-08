# RunMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataType** | Pointer to **string** |  | [optional] [readonly] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Target** | Pointer to [**Target**](Target.md) |  | [optional] 

## Methods

### NewRunMetadata

`func NewRunMetadata() *RunMetadata`

NewRunMetadata instantiates a new RunMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunMetadataWithDefaults

`func NewRunMetadataWithDefaults() *RunMetadata`

NewRunMetadataWithDefaults instantiates a new RunMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataType

`func (o *RunMetadata) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *RunMetadata) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *RunMetadata) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *RunMetadata) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetProduct

`func (o *RunMetadata) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RunMetadata) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RunMetadata) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *RunMetadata) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTarget

`func (o *RunMetadata) GetTarget() Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RunMetadata) GetTargetOk() (*Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RunMetadata) SetTarget(v Target)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RunMetadata) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


