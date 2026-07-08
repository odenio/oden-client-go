# QualityTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RawData** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Interval** | Pointer to [**Interval**](Interval.md) |  | [optional] 
**QualitySchema** | Pointer to [**QualitySchema**](QualitySchema.md) |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewQualityTest

`func NewQualityTest() *QualityTest`

NewQualityTest instantiates a new QualityTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityTestWithDefaults

`func NewQualityTestWithDefaults() *QualityTest`

NewQualityTestWithDefaults instantiates a new QualityTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QualityTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRawData

`func (o *QualityTest) GetRawData() map[string]interface{}`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *QualityTest) GetRawDataOk() (*map[string]interface{}, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *QualityTest) SetRawData(v map[string]interface{})`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *QualityTest) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetTimestamp

`func (o *QualityTest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QualityTest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QualityTest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QualityTest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetInterval

`func (o *QualityTest) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *QualityTest) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *QualityTest) SetInterval(v Interval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *QualityTest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetQualitySchema

`func (o *QualityTest) GetQualitySchema() QualitySchema`

GetQualitySchema returns the QualitySchema field if non-nil, zero value otherwise.

### GetQualitySchemaOk

`func (o *QualityTest) GetQualitySchemaOk() (*QualitySchema, bool)`

GetQualitySchemaOk returns a tuple with the QualitySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualitySchema

`func (o *QualityTest) SetQualitySchema(v QualitySchema)`

SetQualitySchema sets QualitySchema field to given value.

### HasQualitySchema

`func (o *QualityTest) HasQualitySchema() bool`

HasQualitySchema returns a boolean if a field has been set.

### GetMatch

`func (o *QualityTest) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *QualityTest) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *QualityTest) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *QualityTest) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


