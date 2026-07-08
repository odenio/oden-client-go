# ScrapYieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawData** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**ScrapYieldSchema**](ScrapYieldSchema.md) |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewScrapYieldData

`func NewScrapYieldData() *ScrapYieldData`

NewScrapYieldData instantiates a new ScrapYieldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrapYieldDataWithDefaults

`func NewScrapYieldDataWithDefaults() *ScrapYieldData`

NewScrapYieldDataWithDefaults instantiates a new ScrapYieldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawData

`func (o *ScrapYieldData) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ScrapYieldData) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ScrapYieldData) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ScrapYieldData) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetId

`func (o *ScrapYieldData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScrapYieldData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScrapYieldData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScrapYieldData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *ScrapYieldData) GetSchema() ScrapYieldSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ScrapYieldData) GetSchemaOk() (*ScrapYieldSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ScrapYieldData) SetSchema(v ScrapYieldSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ScrapYieldData) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTimestamp

`func (o *ScrapYieldData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ScrapYieldData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ScrapYieldData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ScrapYieldData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ScrapYieldData) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ScrapYieldData) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetMatch

`func (o *ScrapYieldData) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ScrapYieldData) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ScrapYieldData) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *ScrapYieldData) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


