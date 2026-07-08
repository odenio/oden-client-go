# SearchMaintenanceWorkOrdersRequest

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

### NewSearchMaintenanceWorkOrdersRequest

`func NewSearchMaintenanceWorkOrdersRequest() *SearchMaintenanceWorkOrdersRequest`

NewSearchMaintenanceWorkOrdersRequest instantiates a new SearchMaintenanceWorkOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMaintenanceWorkOrdersRequestWithDefaults

`func NewSearchMaintenanceWorkOrdersRequestWithDefaults() *SearchMaintenanceWorkOrdersRequest`

NewSearchMaintenanceWorkOrdersRequestWithDefaults instantiates a new SearchMaintenanceWorkOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchMaintenanceWorkOrdersRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchMaintenanceWorkOrdersRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchMaintenanceWorkOrdersRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *SearchMaintenanceWorkOrdersRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SearchMaintenanceWorkOrdersRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SearchMaintenanceWorkOrdersRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLine

`func (o *SearchMaintenanceWorkOrdersRequest) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SearchMaintenanceWorkOrdersRequest) SetLine(v Line)`

SetLine sets Line field to given value.

### HasLine

`func (o *SearchMaintenanceWorkOrdersRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetStartTime

`func (o *SearchMaintenanceWorkOrdersRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchMaintenanceWorkOrdersRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SearchMaintenanceWorkOrdersRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *SearchMaintenanceWorkOrdersRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchMaintenanceWorkOrdersRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SearchMaintenanceWorkOrdersRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMatch

`func (o *SearchMaintenanceWorkOrdersRequest) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *SearchMaintenanceWorkOrdersRequest) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *SearchMaintenanceWorkOrdersRequest) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *SearchMaintenanceWorkOrdersRequest) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


