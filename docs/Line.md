# Line

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecondaryName** | Pointer to **string** |  | [optional] 
**Factory** | Pointer to [**Factory**](Factory.md) |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewLine

`func NewLine() *Line`

NewLine instantiates a new Line object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineWithDefaults

`func NewLineWithDefaults() *Line`

NewLineWithDefaults instantiates a new Line object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Line) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Line) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Line) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Line) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Line) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Line) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Line) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Line) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecondaryName

`func (o *Line) GetSecondaryName() string`

GetSecondaryName returns the SecondaryName field if non-nil, zero value otherwise.

### GetSecondaryNameOk

`func (o *Line) GetSecondaryNameOk() (*string, bool)`

GetSecondaryNameOk returns a tuple with the SecondaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryName

`func (o *Line) SetSecondaryName(v string)`

SetSecondaryName sets SecondaryName field to given value.

### HasSecondaryName

`func (o *Line) HasSecondaryName() bool`

HasSecondaryName returns a boolean if a field has been set.

### GetFactory

`func (o *Line) GetFactory() Factory`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *Line) GetFactoryOk() (*Factory, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *Line) SetFactory(v Factory)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *Line) HasFactory() bool`

HasFactory returns a boolean if a field has been set.

### GetMatch

`func (o *Line) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Line) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Line) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Line) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


