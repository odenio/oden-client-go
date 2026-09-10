# StateCategorySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**StateCategoryType**](StateCategoryType.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewStateCategorySet

`func NewStateCategorySet() *StateCategorySet`

NewStateCategorySet instantiates a new StateCategorySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateCategorySetWithDefaults

`func NewStateCategorySetWithDefaults() *StateCategorySet`

NewStateCategorySetWithDefaults instantiates a new StateCategorySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StateCategorySet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateCategorySet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateCategorySet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StateCategorySet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StateCategorySet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateCategorySet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateCategorySet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StateCategorySet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StateCategorySet) GetType() StateCategoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StateCategorySet) GetTypeOk() (*StateCategoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StateCategorySet) SetType(v StateCategoryType)`

SetType sets Type field to given value.

### HasType

`func (o *StateCategorySet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *StateCategorySet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StateCategorySet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StateCategorySet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StateCategorySet) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


