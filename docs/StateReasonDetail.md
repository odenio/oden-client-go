# StateReasonDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to [**StateCategory**](StateCategory.md) |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 

## Methods

### NewStateReasonDetail

`func NewStateReasonDetail() *StateReasonDetail`

NewStateReasonDetail instantiates a new StateReasonDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateReasonDetailWithDefaults

`func NewStateReasonDetailWithDefaults() *StateReasonDetail`

NewStateReasonDetailWithDefaults instantiates a new StateReasonDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StateReasonDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateReasonDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateReasonDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StateReasonDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StateReasonDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateReasonDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateReasonDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StateReasonDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *StateReasonDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StateReasonDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StateReasonDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StateReasonDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCategory

`func (o *StateReasonDetail) GetCategory() StateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StateReasonDetail) GetCategoryOk() (*StateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StateReasonDetail) SetCategory(v StateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StateReasonDetail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLine

`func (o *StateReasonDetail) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *StateReasonDetail) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *StateReasonDetail) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *StateReasonDetail) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


