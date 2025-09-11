# IntervalType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Factory** | Pointer to [**Factory**](Factory.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewIntervalType

`func NewIntervalType() *IntervalType`

NewIntervalType instantiates a new IntervalType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalTypeWithDefaults

`func NewIntervalTypeWithDefaults() *IntervalType`

NewIntervalTypeWithDefaults instantiates a new IntervalType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntervalType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntervalType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntervalType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntervalType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFactory

`func (o *IntervalType) GetFactory() Factory`

GetFactory returns the Factory field if non-nil, zero value otherwise.

### GetFactoryOk

`func (o *IntervalType) GetFactoryOk() (*Factory, bool)`

GetFactoryOk returns a tuple with the Factory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactory

`func (o *IntervalType) SetFactory(v Factory)`

SetFactory sets Factory field to given value.

### HasFactory

`func (o *IntervalType) HasFactory() bool`

HasFactory returns a boolean if a field has been set.

### GetName

`func (o *IntervalType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntervalType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntervalType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntervalType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatch

`func (o *IntervalType) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *IntervalType) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *IntervalType) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *IntervalType) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


