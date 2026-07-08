# MaintenanceWorkOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewMaintenanceWorkOrder

`func NewMaintenanceWorkOrder() *MaintenanceWorkOrder`

NewMaintenanceWorkOrder instantiates a new MaintenanceWorkOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWorkOrderWithDefaults

`func NewMaintenanceWorkOrderWithDefaults() *MaintenanceWorkOrder`

NewMaintenanceWorkOrderWithDefaults instantiates a new MaintenanceWorkOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceWorkOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceWorkOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceWorkOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceWorkOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MaintenanceWorkOrder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintenanceWorkOrder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintenanceWorkOrder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MaintenanceWorkOrder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MaintenanceWorkOrder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceWorkOrder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceWorkOrder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenanceWorkOrder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *MaintenanceWorkOrder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MaintenanceWorkOrder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MaintenanceWorkOrder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MaintenanceWorkOrder) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLine

`func (o *MaintenanceWorkOrder) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *MaintenanceWorkOrder) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *MaintenanceWorkOrder) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *MaintenanceWorkOrder) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetStartedAt

`func (o *MaintenanceWorkOrder) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MaintenanceWorkOrder) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MaintenanceWorkOrder) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MaintenanceWorkOrder) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MaintenanceWorkOrder) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MaintenanceWorkOrder) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MaintenanceWorkOrder) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MaintenanceWorkOrder) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *MaintenanceWorkOrder) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MaintenanceWorkOrder) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MaintenanceWorkOrder) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MaintenanceWorkOrder) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMatch

`func (o *MaintenanceWorkOrder) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *MaintenanceWorkOrder) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *MaintenanceWorkOrder) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *MaintenanceWorkOrder) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


