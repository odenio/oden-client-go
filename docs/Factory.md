# Factory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecondaryName** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Match** | Pointer to [**Match**](Match.md) |  | [optional] [default to UNIQUE]

## Methods

### NewFactory

`func NewFactory() *Factory`

NewFactory instantiates a new Factory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactoryWithDefaults

`func NewFactoryWithDefaults() *Factory`

NewFactoryWithDefaults instantiates a new Factory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Factory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Factory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Factory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Factory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Factory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Factory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Factory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Factory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecondaryName

`func (o *Factory) GetSecondaryName() string`

GetSecondaryName returns the SecondaryName field if non-nil, zero value otherwise.

### GetSecondaryNameOk

`func (o *Factory) GetSecondaryNameOk() (*string, bool)`

GetSecondaryNameOk returns a tuple with the SecondaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryName

`func (o *Factory) SetSecondaryName(v string)`

SetSecondaryName sets SecondaryName field to given value.

### HasSecondaryName

`func (o *Factory) HasSecondaryName() bool`

HasSecondaryName returns a boolean if a field has been set.

### SetSecondaryNameNil

`func (o *Factory) SetSecondaryNameNil(b bool)`

 SetSecondaryNameNil sets the value for SecondaryName to be an explicit nil

### UnsetSecondaryName
`func (o *Factory) UnsetSecondaryName()`

UnsetSecondaryName ensures that no value is present for SecondaryName, not even an explicit nil
### GetTimezone

`func (o *Factory) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Factory) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Factory) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Factory) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetMatch

`func (o *Factory) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Factory) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Factory) SetMatch(v Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Factory) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


