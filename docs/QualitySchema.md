# QualitySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Factory** | Pointer to [**Factory**](Factory.md) |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewQualitySchema

`func NewQualitySchema() *QualitySchema`

NewQualitySchema instantiates a new QualitySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualitySchemaWithDefaults

`func NewQualitySchemaWithDefaults() *QualitySchema`

NewQualitySchemaWithDefaults instantiates a new QualitySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualitySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualitySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualitySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QualitySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFactory

`func (o *QualitySchema) GetFactory() Factory`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *QualitySchema) GetFactoryOk() (*Factory, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *QualitySchema) SetFactory(v Factory)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *QualitySchema) HasFactory() bool`

HasFactory returns a boolean if a field has been set.

### GetSchema

`func (o *QualitySchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *QualitySchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *QualitySchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *QualitySchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDescription

`func (o *QualitySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QualitySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QualitySchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QualitySchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMatch

`func (o *QualitySchema) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *QualitySchema) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *QualitySchema) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *QualitySchema) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


