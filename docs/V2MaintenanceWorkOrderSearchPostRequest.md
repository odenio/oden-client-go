# V2MaintenanceWorkOrderSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Line** | Pointer to [**Line**](Line.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewV2MaintenanceWorkOrderSearchPostRequest

`func NewV2MaintenanceWorkOrderSearchPostRequest() *V2MaintenanceWorkOrderSearchPostRequest`

NewV2MaintenanceWorkOrderSearchPostRequest instantiates a new V2MaintenanceWorkOrderSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MaintenanceWorkOrderSearchPostRequestWithDefaults

`func NewV2MaintenanceWorkOrderSearchPostRequestWithDefaults() *V2MaintenanceWorkOrderSearchPostRequest`

NewV2MaintenanceWorkOrderSearchPostRequestWithDefaults instantiates a new V2MaintenanceWorkOrderSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLine

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetStartTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMatch

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *V2MaintenanceWorkOrderSearchPostRequest) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *V2MaintenanceWorkOrderSearchPostRequest) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *V2MaintenanceWorkOrderSearchPostRequest) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


