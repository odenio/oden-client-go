# ProductAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**RuleBased** | Pointer to **bool** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewProductAttribute

`func NewProductAttribute() *ProductAttribute`

NewProductAttribute instantiates a new ProductAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAttributeWithDefaults

`func NewProductAttributeWithDefaults() *ProductAttribute`

NewProductAttributeWithDefaults instantiates a new ProductAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ProductAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProductAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProductAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProductAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *ProductAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetProduct

`func (o *ProductAttribute) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductAttribute) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductAttribute) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ProductAttribute) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRuleBased

`func (o *ProductAttribute) GetRuleBased() bool`

GetRuleBased returns the RuleBased field if non-nil, zero value otherwise.

### GetRuleBasedOk

`func (o *ProductAttribute) GetRuleBasedOk() (*bool, bool)`

GetRuleBasedOk returns a tuple with the RuleBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleBased

`func (o *ProductAttribute) SetRuleBased(v bool)`

SetRuleBased sets RuleBased field to given value.

### HasRuleBased

`func (o *ProductAttribute) HasRuleBased() bool`

HasRuleBased returns a boolean if a field has been set.

### GetMatch

`func (o *ProductAttribute) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ProductAttribute) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ProductAttribute) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *ProductAttribute) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


