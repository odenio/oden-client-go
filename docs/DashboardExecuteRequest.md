# DashboardExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboard** | [**DashboardExecuteRequestDashboard**](DashboardExecuteRequestDashboard.md) |  | 
**Range** | Pointer to [**DashboardExecuteRange**](DashboardExecuteRange.md) |  | [optional] 
**Filters** | Pointer to [**DashboardExecuteFilters**](DashboardExecuteFilters.md) |  | [optional] 

## Methods

### NewDashboardExecuteRequest

`func NewDashboardExecuteRequest(dashboard DashboardExecuteRequestDashboard, ) *DashboardExecuteRequest`

NewDashboardExecuteRequest instantiates a new DashboardExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardExecuteRequestWithDefaults

`func NewDashboardExecuteRequestWithDefaults() *DashboardExecuteRequest`

NewDashboardExecuteRequestWithDefaults instantiates a new DashboardExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboard

`func (o *DashboardExecuteRequest) GetDashboard() DashboardExecuteRequestDashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *DashboardExecuteRequest) GetDashboardOk() (*DashboardExecuteRequestDashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *DashboardExecuteRequest) SetDashboard(v DashboardExecuteRequestDashboard)`

SetDashboard sets Dashboard field to given value.


### GetRange

`func (o *DashboardExecuteRequest) GetRange() DashboardExecuteRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *DashboardExecuteRequest) GetRangeOk() (*DashboardExecuteRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *DashboardExecuteRequest) SetRange(v DashboardExecuteRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *DashboardExecuteRequest) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetFilters

`func (o *DashboardExecuteRequest) GetFilters() DashboardExecuteFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DashboardExecuteRequest) GetFiltersOk() (*DashboardExecuteFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DashboardExecuteRequest) SetFilters(v DashboardExecuteFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DashboardExecuteRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


