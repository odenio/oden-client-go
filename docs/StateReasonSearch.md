# StateReasonSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | [**[]Line**](Line.md) |  | 
**Category** | Pointer to [**StateCategory**](StateCategory.md) |  | [optional] 

## Methods

### NewStateReasonSearch

`func NewStateReasonSearch(lines []Line, ) *StateReasonSearch`

NewStateReasonSearch instantiates a new StateReasonSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateReasonSearchWithDefaults

`func NewStateReasonSearchWithDefaults() *StateReasonSearch`

NewStateReasonSearchWithDefaults instantiates a new StateReasonSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *StateReasonSearch) GetLines() []Line`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *StateReasonSearch) GetLinesOk() (*[]Line, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *StateReasonSearch) SetLines(v []Line)`

SetLines sets Lines field to given value.


### GetCategory

`func (o *StateReasonSearch) GetCategory() StateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StateReasonSearch) GetCategoryOk() (*StateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StateReasonSearch) SetCategory(v StateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StateReasonSearch) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


