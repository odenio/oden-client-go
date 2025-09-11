# Unit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**KindOfQuantity** | Pointer to **string** |  | [optional] 

## Methods

### NewUnit

`func NewUnit() *Unit`

NewUnit instantiates a new Unit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitWithDefaults

`func NewUnitWithDefaults() *Unit`

NewUnitWithDefaults instantiates a new Unit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Unit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Unit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Unit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Unit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSymbol

`func (o *Unit) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Unit) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Unit) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Unit) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *Unit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Unit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Unit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Unit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKindOfQuantity

`func (o *Unit) GetKindOfQuantity() string`

GetKindOfQuantity returns the KindOfQuantity field if non-nil, zero value otherwise.

### GetKindOfQuantityOk

`func (o *Unit) GetKindOfQuantityOk() (*string, bool)`

GetKindOfQuantityOk returns a tuple with the KindOfQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindOfQuantity

`func (o *Unit) SetKindOfQuantity(v string)`

SetKindOfQuantity sets KindOfQuantity field to given value.

### HasKindOfQuantity

`func (o *Unit) HasKindOfQuantity() bool`

HasKindOfQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


