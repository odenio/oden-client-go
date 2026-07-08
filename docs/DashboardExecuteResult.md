# DashboardExecuteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleId** | **string** |  | 
**ModuleName** | **string** |  | 
**ModuleType** | **string** | The module&#39;s stored visualization (e.g. &#x60;table&#x60;, &#x60;line_chart&#x60;, &#x60;bar_chart&#x60;). Type label only — does not change the response shape.  | 
**Range** | Pointer to [**DashboardExecuteResultRange**](DashboardExecuteResultRange.md) |  | [optional] 
**FiltersApplied** | Pointer to **map[string]interface{}** | Echo of the filter dimensions that were applied, resolved to human-readable values where possible (e.g. line names instead of IDs).  | [optional] 
**Columns** | Pointer to [**[]DashboardColumnSpec**](DashboardColumnSpec.md) | Column metadata. &#x60;type&#x60; is derived from the first non-null cell in the column. Null when the module errored.  | [optional] 
**Rows** | Pointer to **[]map[string]interface{}** | Row data as objects keyed by column name (not positional arrays). Values are typed JSON natively. Null when the module errored.  | [optional] 
**Error** | Pointer to **NullableString** | Set to a short message when the module failed to parse or execute. Null on success.  | [optional] 

## Methods

### NewDashboardExecuteResult

`func NewDashboardExecuteResult(moduleId string, moduleName string, moduleType string, ) *DashboardExecuteResult`

NewDashboardExecuteResult instantiates a new DashboardExecuteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardExecuteResultWithDefaults

`func NewDashboardExecuteResultWithDefaults() *DashboardExecuteResult`

NewDashboardExecuteResultWithDefaults instantiates a new DashboardExecuteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleId

`func (o *DashboardExecuteResult) GetModuleId() string`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *DashboardExecuteResult) GetModuleIdOk() (*string, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *DashboardExecuteResult) SetModuleId(v string)`

SetModuleId sets ModuleId field to given value.


### GetModuleName

`func (o *DashboardExecuteResult) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *DashboardExecuteResult) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *DashboardExecuteResult) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.


### GetModuleType

`func (o *DashboardExecuteResult) GetModuleType() string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *DashboardExecuteResult) GetModuleTypeOk() (*string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *DashboardExecuteResult) SetModuleType(v string)`

SetModuleType sets ModuleType field to given value.


### GetRange

`func (o *DashboardExecuteResult) GetRange() DashboardExecuteResultRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *DashboardExecuteResult) GetRangeOk() (*DashboardExecuteResultRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *DashboardExecuteResult) SetRange(v DashboardExecuteResultRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *DashboardExecuteResult) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetFiltersApplied

`func (o *DashboardExecuteResult) GetFiltersApplied() map[string]interface{}`

GetFiltersApplied returns the FiltersApplied field if non-nil, zero value otherwise.

### GetFiltersAppliedOk

`func (o *DashboardExecuteResult) GetFiltersAppliedOk() (*map[string]interface{}, bool)`

GetFiltersAppliedOk returns a tuple with the FiltersApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersApplied

`func (o *DashboardExecuteResult) SetFiltersApplied(v map[string]interface{})`

SetFiltersApplied sets FiltersApplied field to given value.

### HasFiltersApplied

`func (o *DashboardExecuteResult) HasFiltersApplied() bool`

HasFiltersApplied returns a boolean if a field has been set.

### GetColumns

`func (o *DashboardExecuteResult) GetColumns() []DashboardColumnSpec`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DashboardExecuteResult) GetColumnsOk() (*[]DashboardColumnSpec, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DashboardExecuteResult) SetColumns(v []DashboardColumnSpec)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DashboardExecuteResult) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *DashboardExecuteResult) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *DashboardExecuteResult) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetRows

`func (o *DashboardExecuteResult) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DashboardExecuteResult) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DashboardExecuteResult) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *DashboardExecuteResult) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRowsNil

`func (o *DashboardExecuteResult) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *DashboardExecuteResult) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil
### GetError

`func (o *DashboardExecuteResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DashboardExecuteResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DashboardExecuteResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DashboardExecuteResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DashboardExecuteResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DashboardExecuteResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


