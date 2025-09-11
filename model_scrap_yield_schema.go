/*
Oden API

The Oden Private Partner API exposes RESTful API endpoints for clients to get, create and update data on the Oden Platform.  The API is based on the OpenAPI 3.0 specification. ### Current Version The URL, and host, for the current version is [https://api.oden.app/v2](https://api.oden.app/v2).  ### Oden's Data Model - **Organization**: This represents the Organization registered as an Oden customer. An organization has an associated collection of users, factories, lines, etc. This is the entity a specific authentication token is associated with. - **Asset** or **Machinegroup**: Assets, or machinegroups, are collections of machines, which may either be a **Factory** or a **Line**:   - **Factory**: Factories are collections of lines, representing a particular manufacturing location.     - **Line**: Lines are collections of machines, often representing a particular production line. Lines may also have **Products** mapped to them, indicating what is currently being manufactured by the specific line.       - **Machine**: Machines are the physical machines that make up a line - **Product**: Products capture what entities a manufacturer produces - **Interval**: An interval is a period of time that takes place on a manufacturing line and expresses some business concern. It's Oden's way of making metrics aggregatable, traceable, and relatable to a manufacturer.   - **Run**: A run is a production interval that labels a period of production as being work on some single product   - **Batch**: A batch is a production interval that represents a portion of a particular run   - **State**: A state is an interval that tracks the availability or utilization of a line     - **State Category**: A state category describes what state a line is in - such as ex. uptime, downtime, scrapping, etc.     - **State Reason**: A state reason describes why a line is in a particular state - such as \"maintenance\" being a reason for the category \"downtime\".   - **Custom**: A custom interval can track any other type of interval-based data a manufacturer might want to analyze. These are created on a per-factory basis. - **Target**: Targets specify values and upper/lower thresholds for metrics when specific products are running on specific lines - **Scrap/Yield**: Scrap/yield output specifies amount of produced product on a line during either a run or batch interval. Oden will categorize all output as either scrap or yield - as specified by the Scrap Yield Schema for a given factory. If you have other categories, like rework/blocked/off-grade, you must choose between categorizing those amounts as either good or bad production by specifying as scrap or yield. Clients may also add scrap codes (i.e., reasons) to a given Scrap Yield Data entry.   - **Scrap Code**: A scrap code is a code that explains the reason for a scrap/yield raw data input - such as \"Rework\" - **Quality Test**: Quality Tests are results of quality assurance tests done on site, and uploaded to Oden. They may be attached to a single Batch or Run. - **Metric**: Known in factories as \"tags\", metrics are the raw data that is collected by Oden from the machines and devices on the factory floor. - **Metric Group**: Metric groups are metrics that represent the same thing across different lines. They provide common display names for tags and allow labeling groups of tags as measuring key types of data like performance or production.  ### Best Practices Under the current implementation, the Oden API does not rate limit requests from clients.  However, rate limiting will be introduced in the near future and it is recommended that users design their API clients to not exceed a request rate of **one per second**.  ### Schema All v2 API access is over HTTPS and accessed from https://api.oden.app/v2 All data is sent and received as JSON.  API requests with duplicate keys will process only the data for the first key detected and ignore the rest, so it's not recommended. Batching multiple messages in this way is currently not possible.   - Example of duplicate key in JSON: {\"raw_data\":{\"scrap\":\"10\",\"scrap\":\"100\"}}  All timestamps are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format:    `YYYY-MM-DDTHH:MM:SSZ`  All durations are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format with the largest unit of time being the hour:     `PT[n]H[n]M[n]S`  All timestamps sent to the API in POST requests should be in ISO 8601 format. ### HTTP Verbs The ONLY HTTP call type (sometimes called *verb* or *method*) used within Oden's API is **POST**. There are three actions supported via a **POST**; call, search, set, and delete, together supporting CRUD operations;   - **search** requests are used to search for and *read* objects in the Oden Platform       - All Oden Objects may be uniquely identified by some combination of, or a single, parameter.         - Ex a `line` my be identified by either:           - `id`           - `name` AND `factory`   - **set** requests are used to *create* or *update* objects   - **delete** requests are used to *delete* objects. If a delete endpoint is not yet implemented for a given object, users may choose to update the values of a specific entity to null or 0 values.  ### URI Components All endpoints may be accessed with the URI pattern: `https://api.oden.app/v2/{object}/{action}`  Where:   - `object` is the name of the object being requested:        - `factory`, `quality_test`, `interval`, `line`, etc...   - `action` is the name of the action being requested     - `search` , `set` , `delete`  e.g. `https://api.oden.app/v2/factory/search`  # Authentication Clients can authenticate through the v2 API using a Token provided by Oden. Tokens are opaque strings similar to [Bearer](https://swagger.io/docs/specification/authentication/bearer-authentication/) tokens that the client must pass in the [HTTP Authorization request header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) in every request. The syntax is as follows:  `Authorization: <type> <credentials>`  Where \\<type\\> is \"Token\" and \\<credentials\\> is the Token string. For example:  `Authorization: Token tokenStringProvidedByOden`  Authenticating with an invalid Token will return `401 Unauthorized Error`.  Authenticating with a Token that is not authorized to read requested data will return `403 Forbidden Error`.  Some endpoints may require requests to be broken out by machinegroup (i.e., line or factory) and the number of requests would scale accordingly. This multiplicity should be taken into consideration when deciding on the frequency the API client makes requests to the Oden endpoints.  To authenticate in this [UI](https://api.oden.app/v2/ui/), click the Lock icon, and copy/paste the token into the Authorize box. 

API version: 2.0.0
Contact: support@oden.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oden

import (
	"encoding/json"
)

// checks if the ScrapYieldSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScrapYieldSchema{}

// ScrapYieldSchema Scrap yield schema represents a factory's scrap/yield data ingestion configuration. 
type ScrapYieldSchema struct {
	Id *string `json:"id,omitempty"`
	Factory *Factory `json:"factory,omitempty"`
	ScrapConversionFactor *float64 `json:"scrap_conversion_factor,omitempty"`
	ScrapUnit *Unit `json:"scrap_unit,omitempty"`
	YieldConversionFactor *float64 `json:"yield_conversion_factor,omitempty"`
	YieldUnit *Unit `json:"yield_unit,omitempty"`
}

// NewScrapYieldSchema instantiates a new ScrapYieldSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScrapYieldSchema() *ScrapYieldSchema {
	this := ScrapYieldSchema{}
	return &this
}

// NewScrapYieldSchemaWithDefaults instantiates a new ScrapYieldSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScrapYieldSchemaWithDefaults() *ScrapYieldSchema {
	this := ScrapYieldSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScrapYieldSchema) SetId(v string) {
	o.Id = &v
}

// GetFactory returns the Factory field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetFactory() Factory {
	if o == nil || IsNil(o.Factory) {
		var ret Factory
		return ret
	}
	return *o.Factory
}

// GetFactoryOk returns a tuple with the Factory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetFactoryOk() (*Factory, bool) {
	if o == nil || IsNil(o.Factory) {
		return nil, false
	}
	return o.Factory, true
}

// HasFactory returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasFactory() bool {
	if o != nil && !IsNil(o.Factory) {
		return true
	}

	return false
}

// SetFactory gets a reference to the given Factory and assigns it to the Factory field.
func (o *ScrapYieldSchema) SetFactory(v Factory) {
	o.Factory = &v
}

// GetScrapConversionFactor returns the ScrapConversionFactor field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetScrapConversionFactor() float64 {
	if o == nil || IsNil(o.ScrapConversionFactor) {
		var ret float64
		return ret
	}
	return *o.ScrapConversionFactor
}

// GetScrapConversionFactorOk returns a tuple with the ScrapConversionFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetScrapConversionFactorOk() (*float64, bool) {
	if o == nil || IsNil(o.ScrapConversionFactor) {
		return nil, false
	}
	return o.ScrapConversionFactor, true
}

// HasScrapConversionFactor returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasScrapConversionFactor() bool {
	if o != nil && !IsNil(o.ScrapConversionFactor) {
		return true
	}

	return false
}

// SetScrapConversionFactor gets a reference to the given float64 and assigns it to the ScrapConversionFactor field.
func (o *ScrapYieldSchema) SetScrapConversionFactor(v float64) {
	o.ScrapConversionFactor = &v
}

// GetScrapUnit returns the ScrapUnit field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetScrapUnit() Unit {
	if o == nil || IsNil(o.ScrapUnit) {
		var ret Unit
		return ret
	}
	return *o.ScrapUnit
}

// GetScrapUnitOk returns a tuple with the ScrapUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetScrapUnitOk() (*Unit, bool) {
	if o == nil || IsNil(o.ScrapUnit) {
		return nil, false
	}
	return o.ScrapUnit, true
}

// HasScrapUnit returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasScrapUnit() bool {
	if o != nil && !IsNil(o.ScrapUnit) {
		return true
	}

	return false
}

// SetScrapUnit gets a reference to the given Unit and assigns it to the ScrapUnit field.
func (o *ScrapYieldSchema) SetScrapUnit(v Unit) {
	o.ScrapUnit = &v
}

// GetYieldConversionFactor returns the YieldConversionFactor field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetYieldConversionFactor() float64 {
	if o == nil || IsNil(o.YieldConversionFactor) {
		var ret float64
		return ret
	}
	return *o.YieldConversionFactor
}

// GetYieldConversionFactorOk returns a tuple with the YieldConversionFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetYieldConversionFactorOk() (*float64, bool) {
	if o == nil || IsNil(o.YieldConversionFactor) {
		return nil, false
	}
	return o.YieldConversionFactor, true
}

// HasYieldConversionFactor returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasYieldConversionFactor() bool {
	if o != nil && !IsNil(o.YieldConversionFactor) {
		return true
	}

	return false
}

// SetYieldConversionFactor gets a reference to the given float64 and assigns it to the YieldConversionFactor field.
func (o *ScrapYieldSchema) SetYieldConversionFactor(v float64) {
	o.YieldConversionFactor = &v
}

// GetYieldUnit returns the YieldUnit field value if set, zero value otherwise.
func (o *ScrapYieldSchema) GetYieldUnit() Unit {
	if o == nil || IsNil(o.YieldUnit) {
		var ret Unit
		return ret
	}
	return *o.YieldUnit
}

// GetYieldUnitOk returns a tuple with the YieldUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScrapYieldSchema) GetYieldUnitOk() (*Unit, bool) {
	if o == nil || IsNil(o.YieldUnit) {
		return nil, false
	}
	return o.YieldUnit, true
}

// HasYieldUnit returns a boolean if a field has been set.
func (o *ScrapYieldSchema) HasYieldUnit() bool {
	if o != nil && !IsNil(o.YieldUnit) {
		return true
	}

	return false
}

// SetYieldUnit gets a reference to the given Unit and assigns it to the YieldUnit field.
func (o *ScrapYieldSchema) SetYieldUnit(v Unit) {
	o.YieldUnit = &v
}

func (o ScrapYieldSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScrapYieldSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Factory) {
		toSerialize["factory"] = o.Factory
	}
	if !IsNil(o.ScrapConversionFactor) {
		toSerialize["scrap_conversion_factor"] = o.ScrapConversionFactor
	}
	if !IsNil(o.ScrapUnit) {
		toSerialize["scrap_unit"] = o.ScrapUnit
	}
	if !IsNil(o.YieldConversionFactor) {
		toSerialize["yield_conversion_factor"] = o.YieldConversionFactor
	}
	if !IsNil(o.YieldUnit) {
		toSerialize["yield_unit"] = o.YieldUnit
	}
	return toSerialize, nil
}

type NullableScrapYieldSchema struct {
	value *ScrapYieldSchema
	isSet bool
}

func (v NullableScrapYieldSchema) Get() *ScrapYieldSchema {
	return v.value
}

func (v *NullableScrapYieldSchema) Set(val *ScrapYieldSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableScrapYieldSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableScrapYieldSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScrapYieldSchema(val *ScrapYieldSchema) *NullableScrapYieldSchema {
	return &NullableScrapYieldSchema{value: val, isSet: true}
}

func (v NullableScrapYieldSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScrapYieldSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


