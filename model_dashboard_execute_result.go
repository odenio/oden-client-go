/*
Oden API

The Oden Private Partner API exposes RESTful API endpoints for clients to get, create and update data on the Oden Platform.  The API is based on the OpenAPI 3.0 specification. ### Current Version The URL, and host, for the current version is [https://api.oden.app/v2](https://api.oden.app/v2).  ### Oden's Data Model - **Organization**: This represents the Organization registered as an Oden customer. An organization has an associated collection of users, factories, lines, etc. This is the entity a specific authentication token is associated with. - **Asset** or **Machinegroup**: Assets, or machinegroups, are collections of machines, which may either be a **Factory** or a **Line**:   - **Factory**: Factories are collections of lines, representing a particular manufacturing location.     - **Line**: Lines are collections of machines, often representing a particular production line. Lines may also have **Products** mapped to them, indicating what is currently being manufactured by the specific line.       - **Machine**: Machines are the physical machines that make up a line - **Product**: Products capture what entities a manufacturer produces - **Interval**: An interval is a period of time that takes place on a manufacturing line and expresses some business concern. It's Oden's way of making metrics aggregatable, traceable, and relatable to a manufacturer.   - **Run**: A run is a production interval that labels a period of production as being work on some single product   - **Batch**: A batch is a production interval that represents a portion of a particular run   - **State**: A state is an interval that tracks the availability or utilization of a line     - **State Category**: A state category describes what state a line is in - such as ex. uptime, downtime, scrapping, etc.     - **State Reason**: A state reason describes why a line is in a particular state - such as \"maintenance\" being a reason for the category \"downtime\".   - **Custom**: A custom interval can track any other type of interval-based data a manufacturer might want to analyze. These are created on a per-factory basis. - **Target**: Targets specify values and upper/lower thresholds for metrics when specific products are running on specific lines - **Scrap/Yield**: Scrap/yield output specifies amount of produced product on a line during either a run or batch interval. Oden will categorize all output as either scrap or yield - as specified by the Scrap Yield Schema for a given factory. If you have other categories, like rework/blocked/off-grade, you must choose between categorizing those amounts as either good or bad production by specifying as scrap or yield. Clients may also add scrap codes (i.e., reasons) to a given Scrap Yield Data entry.   - **Scrap Code**: A scrap code is a code that explains the reason for a scrap/yield raw data input - such as \"Rework\" - **Quality Test**: Quality Tests are results of quality assurance tests done on site, and uploaded to Oden. They may be attached to a single Batch or Run. - **Metric**: Known in factories as \"tags\", metrics are the raw data that is collected by Oden from the machines and devices on the factory floor. - **Metric Group**: Metric groups are metrics that represent the same thing across different lines. They provide common display names for tags and allow labeling groups of tags as measuring key types of data like performance or production. - **Maintenance Work Order**: A maintenance work order can be used to track work orders maintained in MaintainX and associate them with an Oden line.   ### Best Practices Under the current implementation, the Oden API does not rate limit requests from clients.  However, rate limiting will be introduced in the near future and it is recommended that users design their API clients to not exceed a request rate of **one per second**.  ### Schema All v2 API access is over HTTPS and accessed from https://api.oden.app/v2 All data is sent and received as JSON.  API requests with duplicate keys will process only the data for the first key detected and ignore the rest, so it's not recommended. Batching multiple messages in this way is currently not possible.   - Example of duplicate key in JSON: {\"raw_data\":{\"scrap\":\"10\",\"scrap\":\"100\"}}  All timestamps are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format:    `YYYY-MM-DDTHH:MM:SSZ`  All durations are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format with the largest unit of time being the hour:     `PT[n]H[n]M[n]S`  All timestamps sent to the API in POST requests should be in ISO 8601 format. ### HTTP Verbs The ONLY HTTP call type (sometimes called *verb* or *method*) used within Oden's API is **POST**. There are three actions supported via a **POST**; call, search, set, and delete, together supporting CRUD operations;   - **search** requests are used to search for and *read* objects in the Oden Platform       - All Oden Objects may be uniquely identified by some combination of, or a single, parameter.         - Ex a `line` my be identified by either:           - `id`           - `name` AND `factory`   - **set** requests are used to *create* or *update* objects   - **delete** requests are used to *delete* objects. If a delete endpoint is not yet implemented for a given object, users may choose to update the values of a specific entity to null or 0 values.  ### URI Components All endpoints may be accessed with the URI pattern: `https://api.oden.app/v2/{object}/{action}`  Where:   - `object` is the name of the object being requested:        - `factory`, `quality_test`, `interval`, `line`, etc...   - `action` is the name of the action being requested     - `search` , `set` , `delete`  e.g. `https://api.oden.app/v2/factory/search`  # Authentication Clients can authenticate through the v2 API using a Token provided by Oden. Tokens are opaque strings similar to [Bearer](https://swagger.io/docs/specification/authentication/bearer-authentication/) tokens that the client must pass in the [HTTP Authorization request header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) in every request. The syntax is as follows:  `Authorization: <type> <credentials>`  Where \\<type\\> is \"Token\" and \\<credentials\\> is the Token string. For example:  `Authorization: Token tokenStringProvidedByOden`  Authenticating with an invalid Token will return `401 Unauthorized Error`.  Authenticating with a Token that is not authorized to read requested data will return `403 Forbidden Error`.  Some endpoints may require requests to be broken out by machinegroup (i.e., line or factory) and the number of requests would scale accordingly. This multiplicity should be taken into consideration when deciding on the frequency the API client makes requests to the Oden endpoints.  To authenticate in this [UI](https://api.oden.app/v2/ui/), click the Lock icon, and copy/paste the token into the Authorize box. 

API version: 2.0.0
Contact: support@oden.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oden

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DashboardExecuteResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardExecuteResult{}

// DashboardExecuteResult Executed output of a single dashboard module.
type DashboardExecuteResult struct {
	ModuleId string `json:"module_id"`
	ModuleName string `json:"module_name"`
	// The module's stored visualization (e.g. `table`, `line_chart`, `bar_chart`). Type label only — does not change the response shape. 
	ModuleType string `json:"module_type"`
	Range *DashboardExecuteResultRange `json:"range,omitempty"`
	// Echo of the filter dimensions that were applied, resolved to human-readable values where possible (e.g. line names instead of IDs). 
	FiltersApplied map[string]interface{} `json:"filters_applied,omitempty"`
	// Column metadata. `type` is derived from the first non-null cell in the column. Null when the module errored. 
	Columns []DashboardColumnSpec `json:"columns,omitempty"`
	// Row data as objects keyed by column name (not positional arrays). Values are typed JSON natively. Null when the module errored. 
	Rows []map[string]interface{} `json:"rows,omitempty"`
	// Set to a short message when the module failed to parse or execute. Null on success. 
	Error NullableString `json:"error,omitempty"`
}

type _DashboardExecuteResult DashboardExecuteResult

// NewDashboardExecuteResult instantiates a new DashboardExecuteResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardExecuteResult(moduleId string, moduleName string, moduleType string) *DashboardExecuteResult {
	this := DashboardExecuteResult{}
	this.ModuleId = moduleId
	this.ModuleName = moduleName
	this.ModuleType = moduleType
	return &this
}

// NewDashboardExecuteResultWithDefaults instantiates a new DashboardExecuteResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardExecuteResultWithDefaults() *DashboardExecuteResult {
	this := DashboardExecuteResult{}
	return &this
}

// GetModuleId returns the ModuleId field value
func (o *DashboardExecuteResult) GetModuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value
// and a boolean to check if the value has been set.
func (o *DashboardExecuteResult) GetModuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleId, true
}

// SetModuleId sets field value
func (o *DashboardExecuteResult) SetModuleId(v string) {
	o.ModuleId = v
}

// GetModuleName returns the ModuleName field value
func (o *DashboardExecuteResult) GetModuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value
// and a boolean to check if the value has been set.
func (o *DashboardExecuteResult) GetModuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleName, true
}

// SetModuleName sets field value
func (o *DashboardExecuteResult) SetModuleName(v string) {
	o.ModuleName = v
}

// GetModuleType returns the ModuleType field value
func (o *DashboardExecuteResult) GetModuleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *DashboardExecuteResult) GetModuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *DashboardExecuteResult) SetModuleType(v string) {
	o.ModuleType = v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *DashboardExecuteResult) GetRange() DashboardExecuteResultRange {
	if o == nil || IsNil(o.Range) {
		var ret DashboardExecuteResultRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteResult) GetRangeOk() (*DashboardExecuteResultRange, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *DashboardExecuteResult) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given DashboardExecuteResultRange and assigns it to the Range field.
func (o *DashboardExecuteResult) SetRange(v DashboardExecuteResultRange) {
	o.Range = &v
}

// GetFiltersApplied returns the FiltersApplied field value if set, zero value otherwise.
func (o *DashboardExecuteResult) GetFiltersApplied() map[string]interface{} {
	if o == nil || IsNil(o.FiltersApplied) {
		var ret map[string]interface{}
		return ret
	}
	return o.FiltersApplied
}

// GetFiltersAppliedOk returns a tuple with the FiltersApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteResult) GetFiltersAppliedOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FiltersApplied) {
		return map[string]interface{}{}, false
	}
	return o.FiltersApplied, true
}

// HasFiltersApplied returns a boolean if a field has been set.
func (o *DashboardExecuteResult) HasFiltersApplied() bool {
	if o != nil && !IsNil(o.FiltersApplied) {
		return true
	}

	return false
}

// SetFiltersApplied gets a reference to the given map[string]interface{} and assigns it to the FiltersApplied field.
func (o *DashboardExecuteResult) SetFiltersApplied(v map[string]interface{}) {
	o.FiltersApplied = v
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardExecuteResult) GetColumns() []DashboardColumnSpec {
	if o == nil {
		var ret []DashboardColumnSpec
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardExecuteResult) GetColumnsOk() ([]DashboardColumnSpec, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DashboardExecuteResult) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []DashboardColumnSpec and assigns it to the Columns field.
func (o *DashboardExecuteResult) SetColumns(v []DashboardColumnSpec) {
	o.Columns = v
}

// GetRows returns the Rows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardExecuteResult) GetRows() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardExecuteResult) GetRowsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *DashboardExecuteResult) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []map[string]interface{} and assigns it to the Rows field.
func (o *DashboardExecuteResult) SetRows(v []map[string]interface{}) {
	o.Rows = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardExecuteResult) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardExecuteResult) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *DashboardExecuteResult) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *DashboardExecuteResult) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *DashboardExecuteResult) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *DashboardExecuteResult) UnsetError() {
	o.Error.Unset()
}

func (o DashboardExecuteResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardExecuteResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["module_id"] = o.ModuleId
	toSerialize["module_name"] = o.ModuleName
	toSerialize["module_type"] = o.ModuleType
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.FiltersApplied) {
		toSerialize["filters_applied"] = o.FiltersApplied
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

func (o *DashboardExecuteResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module_id",
		"module_name",
		"module_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDashboardExecuteResult := _DashboardExecuteResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboardExecuteResult)

	if err != nil {
		return err
	}

	*o = DashboardExecuteResult(varDashboardExecuteResult)

	return err
}

type NullableDashboardExecuteResult struct {
	value *DashboardExecuteResult
	isSet bool
}

func (v NullableDashboardExecuteResult) Get() *DashboardExecuteResult {
	return v.value
}

func (v *NullableDashboardExecuteResult) Set(val *DashboardExecuteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardExecuteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardExecuteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardExecuteResult(val *DashboardExecuteResult) *NullableDashboardExecuteResult {
	return &NullableDashboardExecuteResult{value: val, isSet: true}
}

func (v NullableDashboardExecuteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardExecuteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


