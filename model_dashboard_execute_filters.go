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
)

// checks if the DashboardExecuteFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardExecuteFilters{}

// DashboardExecuteFilters Optional filter overrides applied to every module. Every field is optional; omitting one means \"no override on that dimension\". 
type DashboardExecuteFilters struct {
	// Lines to restrict to. Each entry must supply `id`, `name`, or both; entries that supply neither are rejected. Other Line fields (factory, secondary_name, match) are not used here and are intentionally omitted so generated clients don't suggest them as inputs. 
	Lines []DashboardExecuteFiltersLinesInner `json:"lines,omitempty"`
	Shifts []int32 `json:"shifts,omitempty"`
	ProductIds []string `json:"product_ids,omitempty"`
	ProductAttributeValueIds []string `json:"product_attribute_value_ids,omitempty"`
	ScrapCategories []string `json:"scrap_categories,omitempty"`
	States *DashboardExecuteFiltersStates `json:"states,omitempty"`
	CustomIntervals []DashboardExecuteFiltersCustomIntervalsInner `json:"custom_intervals,omitempty"`
}

// NewDashboardExecuteFilters instantiates a new DashboardExecuteFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardExecuteFilters() *DashboardExecuteFilters {
	this := DashboardExecuteFilters{}
	return &this
}

// NewDashboardExecuteFiltersWithDefaults instantiates a new DashboardExecuteFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardExecuteFiltersWithDefaults() *DashboardExecuteFilters {
	this := DashboardExecuteFilters{}
	return &this
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetLines() []DashboardExecuteFiltersLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []DashboardExecuteFiltersLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetLinesOk() ([]DashboardExecuteFiltersLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []DashboardExecuteFiltersLinesInner and assigns it to the Lines field.
func (o *DashboardExecuteFilters) SetLines(v []DashboardExecuteFiltersLinesInner) {
	o.Lines = v
}

// GetShifts returns the Shifts field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetShifts() []int32 {
	if o == nil || IsNil(o.Shifts) {
		var ret []int32
		return ret
	}
	return o.Shifts
}

// GetShiftsOk returns a tuple with the Shifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetShiftsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Shifts) {
		return nil, false
	}
	return o.Shifts, true
}

// HasShifts returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasShifts() bool {
	if o != nil && !IsNil(o.Shifts) {
		return true
	}

	return false
}

// SetShifts gets a reference to the given []int32 and assigns it to the Shifts field.
func (o *DashboardExecuteFilters) SetShifts(v []int32) {
	o.Shifts = v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *DashboardExecuteFilters) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetProductAttributeValueIds returns the ProductAttributeValueIds field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetProductAttributeValueIds() []string {
	if o == nil || IsNil(o.ProductAttributeValueIds) {
		var ret []string
		return ret
	}
	return o.ProductAttributeValueIds
}

// GetProductAttributeValueIdsOk returns a tuple with the ProductAttributeValueIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetProductAttributeValueIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductAttributeValueIds) {
		return nil, false
	}
	return o.ProductAttributeValueIds, true
}

// HasProductAttributeValueIds returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasProductAttributeValueIds() bool {
	if o != nil && !IsNil(o.ProductAttributeValueIds) {
		return true
	}

	return false
}

// SetProductAttributeValueIds gets a reference to the given []string and assigns it to the ProductAttributeValueIds field.
func (o *DashboardExecuteFilters) SetProductAttributeValueIds(v []string) {
	o.ProductAttributeValueIds = v
}

// GetScrapCategories returns the ScrapCategories field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetScrapCategories() []string {
	if o == nil || IsNil(o.ScrapCategories) {
		var ret []string
		return ret
	}
	return o.ScrapCategories
}

// GetScrapCategoriesOk returns a tuple with the ScrapCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetScrapCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.ScrapCategories) {
		return nil, false
	}
	return o.ScrapCategories, true
}

// HasScrapCategories returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasScrapCategories() bool {
	if o != nil && !IsNil(o.ScrapCategories) {
		return true
	}

	return false
}

// SetScrapCategories gets a reference to the given []string and assigns it to the ScrapCategories field.
func (o *DashboardExecuteFilters) SetScrapCategories(v []string) {
	o.ScrapCategories = v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetStates() DashboardExecuteFiltersStates {
	if o == nil || IsNil(o.States) {
		var ret DashboardExecuteFiltersStates
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetStatesOk() (*DashboardExecuteFiltersStates, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasStates() bool {
	if o != nil && !IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given DashboardExecuteFiltersStates and assigns it to the States field.
func (o *DashboardExecuteFilters) SetStates(v DashboardExecuteFiltersStates) {
	o.States = &v
}

// GetCustomIntervals returns the CustomIntervals field value if set, zero value otherwise.
func (o *DashboardExecuteFilters) GetCustomIntervals() []DashboardExecuteFiltersCustomIntervalsInner {
	if o == nil || IsNil(o.CustomIntervals) {
		var ret []DashboardExecuteFiltersCustomIntervalsInner
		return ret
	}
	return o.CustomIntervals
}

// GetCustomIntervalsOk returns a tuple with the CustomIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardExecuteFilters) GetCustomIntervalsOk() ([]DashboardExecuteFiltersCustomIntervalsInner, bool) {
	if o == nil || IsNil(o.CustomIntervals) {
		return nil, false
	}
	return o.CustomIntervals, true
}

// HasCustomIntervals returns a boolean if a field has been set.
func (o *DashboardExecuteFilters) HasCustomIntervals() bool {
	if o != nil && !IsNil(o.CustomIntervals) {
		return true
	}

	return false
}

// SetCustomIntervals gets a reference to the given []DashboardExecuteFiltersCustomIntervalsInner and assigns it to the CustomIntervals field.
func (o *DashboardExecuteFilters) SetCustomIntervals(v []DashboardExecuteFiltersCustomIntervalsInner) {
	o.CustomIntervals = v
}

func (o DashboardExecuteFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardExecuteFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Shifts) {
		toSerialize["shifts"] = o.Shifts
	}
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	if !IsNil(o.ProductAttributeValueIds) {
		toSerialize["product_attribute_value_ids"] = o.ProductAttributeValueIds
	}
	if !IsNil(o.ScrapCategories) {
		toSerialize["scrap_categories"] = o.ScrapCategories
	}
	if !IsNil(o.States) {
		toSerialize["states"] = o.States
	}
	if !IsNil(o.CustomIntervals) {
		toSerialize["custom_intervals"] = o.CustomIntervals
	}
	return toSerialize, nil
}

type NullableDashboardExecuteFilters struct {
	value *DashboardExecuteFilters
	isSet bool
}

func (v NullableDashboardExecuteFilters) Get() *DashboardExecuteFilters {
	return v.value
}

func (v *NullableDashboardExecuteFilters) Set(val *DashboardExecuteFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardExecuteFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardExecuteFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardExecuteFilters(val *DashboardExecuteFilters) *NullableDashboardExecuteFilters {
	return &NullableDashboardExecuteFilters{value: val, isSet: true}
}

func (v NullableDashboardExecuteFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardExecuteFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


