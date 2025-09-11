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
	"time"
)

// checks if the QualityTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QualityTest{}

// QualityTest An object representing a QA test for a line for a particular run or batch interval. 
type QualityTest struct {
	Id *string `json:"id,omitempty"`
	RawData map[string]interface{} `json:"raw_data,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Interval *Interval `json:"interval,omitempty"`
	QualitySchema *QualitySchema `json:"quality_schema,omitempty"`
	Match *Match `json:"match,omitempty"`
}

// NewQualityTest instantiates a new QualityTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityTest() *QualityTest {
	this := QualityTest{}
	var match Match = UNIQUE
	this.Match = &match
	return &this
}

// NewQualityTestWithDefaults instantiates a new QualityTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityTestWithDefaults() *QualityTest {
	this := QualityTest{}
	var match Match = UNIQUE
	this.Match = &match
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QualityTest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QualityTest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QualityTest) SetId(v string) {
	o.Id = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise.
func (o *QualityTest) GetRawData() map[string]interface{} {
	if o == nil || IsNil(o.RawData) {
		var ret map[string]interface{}
		return ret
	}
	return o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetRawDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RawData) {
		return map[string]interface{}{}, false
	}
	return o.RawData, true
}

// HasRawData returns a boolean if a field has been set.
func (o *QualityTest) HasRawData() bool {
	if o != nil && !IsNil(o.RawData) {
		return true
	}

	return false
}

// SetRawData gets a reference to the given map[string]interface{} and assigns it to the RawData field.
func (o *QualityTest) SetRawData(v map[string]interface{}) {
	o.RawData = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QualityTest) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QualityTest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *QualityTest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *QualityTest) GetInterval() Interval {
	if o == nil || IsNil(o.Interval) {
		var ret Interval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetIntervalOk() (*Interval, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *QualityTest) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given Interval and assigns it to the Interval field.
func (o *QualityTest) SetInterval(v Interval) {
	o.Interval = &v
}

// GetQualitySchema returns the QualitySchema field value if set, zero value otherwise.
func (o *QualityTest) GetQualitySchema() QualitySchema {
	if o == nil || IsNil(o.QualitySchema) {
		var ret QualitySchema
		return ret
	}
	return *o.QualitySchema
}

// GetQualitySchemaOk returns a tuple with the QualitySchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetQualitySchemaOk() (*QualitySchema, bool) {
	if o == nil || IsNil(o.QualitySchema) {
		return nil, false
	}
	return o.QualitySchema, true
}

// HasQualitySchema returns a boolean if a field has been set.
func (o *QualityTest) HasQualitySchema() bool {
	if o != nil && !IsNil(o.QualitySchema) {
		return true
	}

	return false
}

// SetQualitySchema gets a reference to the given QualitySchema and assigns it to the QualitySchema field.
func (o *QualityTest) SetQualitySchema(v QualitySchema) {
	o.QualitySchema = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *QualityTest) GetMatch() Match {
	if o == nil || IsNil(o.Match) {
		var ret Match
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityTest) GetMatchOk() (*Match, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *QualityTest) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given Match and assigns it to the Match field.
func (o *QualityTest) SetMatch(v Match) {
	o.Match = &v
}

func (o QualityTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QualityTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RawData) {
		toSerialize["raw_data"] = o.RawData
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.QualitySchema) {
		toSerialize["quality_schema"] = o.QualitySchema
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return toSerialize, nil
}

type NullableQualityTest struct {
	value *QualityTest
	isSet bool
}

func (v NullableQualityTest) Get() *QualityTest {
	return v.value
}

func (v *NullableQualityTest) Set(val *QualityTest) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityTest) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityTest(val *QualityTest) *NullableQualityTest {
	return &NullableQualityTest{value: val, isSet: true}
}

func (v NullableQualityTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


