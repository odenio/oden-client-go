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

// checks if the ProductMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMapping{}

// ProductMapping An entity representing a product/line mapping. Mappings are simple 1:1 relationships, and as such cannot be modified, only created. 
type ProductMapping struct {
	Product *Product `json:"product,omitempty"`
	Line *Line `json:"line,omitempty"`
	Match *Match `json:"match,omitempty"`
}

// NewProductMapping instantiates a new ProductMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMapping() *ProductMapping {
	this := ProductMapping{}
	var match Match = UNIQUE
	this.Match = &match
	return &this
}

// NewProductMappingWithDefaults instantiates a new ProductMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMappingWithDefaults() *ProductMapping {
	this := ProductMapping{}
	var match Match = UNIQUE
	this.Match = &match
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ProductMapping) GetProduct() Product {
	if o == nil || IsNil(o.Product) {
		var ret Product
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMapping) GetProductOk() (*Product, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ProductMapping) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given Product and assigns it to the Product field.
func (o *ProductMapping) SetProduct(v Product) {
	o.Product = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ProductMapping) GetLine() Line {
	if o == nil || IsNil(o.Line) {
		var ret Line
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMapping) GetLineOk() (*Line, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ProductMapping) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given Line and assigns it to the Line field.
func (o *ProductMapping) SetLine(v Line) {
	o.Line = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *ProductMapping) GetMatch() Match {
	if o == nil || IsNil(o.Match) {
		var ret Match
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMapping) GetMatchOk() (*Match, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *ProductMapping) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given Match and assigns it to the Match field.
func (o *ProductMapping) SetMatch(v Match) {
	o.Match = &v
}

func (o ProductMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return toSerialize, nil
}

type NullableProductMapping struct {
	value *ProductMapping
	isSet bool
}

func (v NullableProductMapping) Get() *ProductMapping {
	return v.value
}

func (v *NullableProductMapping) Set(val *ProductMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMapping(val *ProductMapping) *NullableProductMapping {
	return &NullableProductMapping{value: val, isSet: true}
}

func (v NullableProductMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


