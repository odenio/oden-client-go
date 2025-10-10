# StateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataType** | **string** |  | 
**Reason** | Pointer to [**StateReason**](StateReason.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**StateCategory**](StateCategory.md) |  | [optional] 

## Methods

### NewStateMetadata

`func NewStateMetadata(metadataType string, ) *StateMetadata`

NewStateMetadata instantiates a new StateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateMetadataWithDefaults

`func NewStateMetadataWithDefaults() *StateMetadata`

NewStateMetadataWithDefaults instantiates a new StateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataType

`func (o *StateMetadata) GetMetadataType() string`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *StateMetadata) GetMetadataTypeOk() (*string, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *StateMetadata) SetMetadataType(v string)`

SetMetadataType sets MetadataType field to given value.


### GetReason

`func (o *StateMetadata) GetReason() StateReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StateMetadata) GetReasonOk() (*StateReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StateMetadata) SetReason(v StateReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StateMetadata) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *StateMetadata) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StateMetadata) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StateMetadata) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StateMetadata) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCategory

`func (o *StateMetadata) GetCategory() StateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StateMetadata) GetCategoryOk() (*StateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StateMetadata) SetCategory(v StateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StateMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


