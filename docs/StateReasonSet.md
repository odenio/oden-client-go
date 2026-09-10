# StateReasonSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | [**[]Line**](Line.md) |  | 
**Category** | [**StateCategory**](StateCategory.md) |  | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewStateReasonSet

`func NewStateReasonSet(lines []Line, category StateCategory, name string, ) *StateReasonSet`

NewStateReasonSet instantiates a new StateReasonSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateReasonSetWithDefaults

`func NewStateReasonSetWithDefaults() *StateReasonSet`

NewStateReasonSetWithDefaults instantiates a new StateReasonSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *StateReasonSet) GetLines() []Line`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *StateReasonSet) GetLinesOk() (*[]Line, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *StateReasonSet) SetLines(v []Line)`

SetLines sets Lines field to given value.


### GetCategory

`func (o *StateReasonSet) GetCategory() StateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StateReasonSet) GetCategoryOk() (*StateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StateReasonSet) SetCategory(v StateCategory)`

SetCategory sets Category field to given value.


### GetName

`func (o *StateReasonSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateReasonSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateReasonSet) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *StateReasonSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StateReasonSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StateReasonSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StateReasonSet) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


