# DashboardExecuteFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**[]DashboardExecuteFiltersLinesInner**](DashboardExecuteFiltersLinesInner.md) | Lines to restrict to. Each entry must supply &#x60;id&#x60;, &#x60;name&#x60;, or both; entries that supply neither are rejected. Other Line fields (factory, secondary_name, match) are not used here and are intentionally omitted so generated clients don&#39;t suggest them as inputs.  | [optional] 
**Shifts** | Pointer to **[]int32** |  | [optional] 
**ProductIds** | Pointer to **[]string** |  | [optional] 
**ProductAttributeValueIds** | Pointer to **[]string** |  | [optional] 
**ScrapCategories** | Pointer to **[]string** |  | [optional] 
**States** | Pointer to [**DashboardExecuteFiltersStates**](DashboardExecuteFiltersStates.md) |  | [optional] 
**CustomIntervals** | Pointer to [**[]DashboardExecuteFiltersCustomIntervalsInner**](DashboardExecuteFiltersCustomIntervalsInner.md) |  | [optional] 

## Methods

### NewDashboardExecuteFilters

`func NewDashboardExecuteFilters() *DashboardExecuteFilters`

NewDashboardExecuteFilters instantiates a new DashboardExecuteFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardExecuteFiltersWithDefaults

`func NewDashboardExecuteFiltersWithDefaults() *DashboardExecuteFilters`

NewDashboardExecuteFiltersWithDefaults instantiates a new DashboardExecuteFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *DashboardExecuteFilters) GetLines() []DashboardExecuteFiltersLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *DashboardExecuteFilters) GetLinesOk() (*[]DashboardExecuteFiltersLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *DashboardExecuteFilters) SetLines(v []DashboardExecuteFiltersLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *DashboardExecuteFilters) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetShifts

`func (o *DashboardExecuteFilters) GetShifts() []int32`

GetShifts returns the Shifts field if non-nil, zero value otherwise.

### GetShiftsOk

`func (o *DashboardExecuteFilters) GetShiftsOk() (*[]int32, bool)`

GetShiftsOk returns a tuple with the Shifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShifts

`func (o *DashboardExecuteFilters) SetShifts(v []int32)`

SetShifts sets Shifts field to given value.

### HasShifts

`func (o *DashboardExecuteFilters) HasShifts() bool`

HasShifts returns a boolean if a field has been set.

### GetProductIds

`func (o *DashboardExecuteFilters) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *DashboardExecuteFilters) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *DashboardExecuteFilters) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *DashboardExecuteFilters) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetProductAttributeValueIds

`func (o *DashboardExecuteFilters) GetProductAttributeValueIds() []string`

GetProductAttributeValueIds returns the ProductAttributeValueIds field if non-nil, zero value otherwise.

### GetProductAttributeValueIdsOk

`func (o *DashboardExecuteFilters) GetProductAttributeValueIdsOk() (*[]string, bool)`

GetProductAttributeValueIdsOk returns a tuple with the ProductAttributeValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAttributeValueIds

`func (o *DashboardExecuteFilters) SetProductAttributeValueIds(v []string)`

SetProductAttributeValueIds sets ProductAttributeValueIds field to given value.

### HasProductAttributeValueIds

`func (o *DashboardExecuteFilters) HasProductAttributeValueIds() bool`

HasProductAttributeValueIds returns a boolean if a field has been set.

### GetScrapCategories

`func (o *DashboardExecuteFilters) GetScrapCategories() []string`

GetScrapCategories returns the ScrapCategories field if non-nil, zero value otherwise.

### GetScrapCategoriesOk

`func (o *DashboardExecuteFilters) GetScrapCategoriesOk() (*[]string, bool)`

GetScrapCategoriesOk returns a tuple with the ScrapCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapCategories

`func (o *DashboardExecuteFilters) SetScrapCategories(v []string)`

SetScrapCategories sets ScrapCategories field to given value.

### HasScrapCategories

`func (o *DashboardExecuteFilters) HasScrapCategories() bool`

HasScrapCategories returns a boolean if a field has been set.

### GetStates

`func (o *DashboardExecuteFilters) GetStates() DashboardExecuteFiltersStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *DashboardExecuteFilters) GetStatesOk() (*DashboardExecuteFiltersStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *DashboardExecuteFilters) SetStates(v DashboardExecuteFiltersStates)`

SetStates sets States field to given value.

### HasStates

`func (o *DashboardExecuteFilters) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetCustomIntervals

`func (o *DashboardExecuteFilters) GetCustomIntervals() []DashboardExecuteFiltersCustomIntervalsInner`

GetCustomIntervals returns the CustomIntervals field if non-nil, zero value otherwise.

### GetCustomIntervalsOk

`func (o *DashboardExecuteFilters) GetCustomIntervalsOk() (*[]DashboardExecuteFiltersCustomIntervalsInner, bool)`

GetCustomIntervalsOk returns a tuple with the CustomIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIntervals

`func (o *DashboardExecuteFilters) SetCustomIntervals(v []DashboardExecuteFiltersCustomIntervalsInner)`

SetCustomIntervals sets CustomIntervals field to given value.

### HasCustomIntervals

`func (o *DashboardExecuteFilters) HasCustomIntervals() bool`

HasCustomIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


