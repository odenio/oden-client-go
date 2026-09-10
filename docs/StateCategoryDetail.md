# StateCategoryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**StateCategoryType**](StateCategoryType.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Builtin** | Pointer to **bool** | True for Oden&#39;s global categories, which have no organization and cannot be modified. | [optional] 

## Methods

### NewStateCategoryDetail

`func NewStateCategoryDetail() *StateCategoryDetail`

NewStateCategoryDetail instantiates a new StateCategoryDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateCategoryDetailWithDefaults

`func NewStateCategoryDetailWithDefaults() *StateCategoryDetail`

NewStateCategoryDetailWithDefaults instantiates a new StateCategoryDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StateCategoryDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateCategoryDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateCategoryDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StateCategoryDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StateCategoryDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateCategoryDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateCategoryDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StateCategoryDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StateCategoryDetail) GetType() StateCategoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StateCategoryDetail) GetTypeOk() (*StateCategoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StateCategoryDetail) SetType(v StateCategoryType)`

SetType sets Type field to given value.

### HasType

`func (o *StateCategoryDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *StateCategoryDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StateCategoryDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StateCategoryDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StateCategoryDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBuiltin

`func (o *StateCategoryDetail) GetBuiltin() bool`

GetBuiltin returns the Builtin field if non-nil, zero value otherwise.

### GetBuiltinOk

`func (o *StateCategoryDetail) GetBuiltinOk() (*bool, bool)`

GetBuiltinOk returns a tuple with the Builtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltin

`func (o *StateCategoryDetail) SetBuiltin(v bool)`

SetBuiltin sets Builtin field to given value.

### HasBuiltin

`func (o *StateCategoryDetail) HasBuiltin() bool`

HasBuiltin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


