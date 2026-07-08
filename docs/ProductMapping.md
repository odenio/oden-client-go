# ProductMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewProductMapping

`func NewProductMapping() *ProductMapping`

NewProductMapping instantiates a new ProductMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMappingWithDefaults

`func NewProductMappingWithDefaults() *ProductMapping`

NewProductMappingWithDefaults instantiates a new ProductMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ProductMapping) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductMapping) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductMapping) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ProductMapping) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetLine

`func (o *ProductMapping) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ProductMapping) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ProductMapping) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *ProductMapping) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetMatch

`func (o *ProductMapping) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ProductMapping) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ProductMapping) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *ProductMapping) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


