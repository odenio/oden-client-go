# SearchLines409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityName** | **string** |  | 
**NumMatches** | **int32** |  | 
**Error** | **string** |  | 
**Retryable** | **bool** |  | 

## Methods

### NewSearchLines409Response

`func NewSearchLines409Response(entityName string, numMatches int32, error_ string, retryable bool, ) *SearchLines409Response`

NewSearchLines409Response instantiates a new SearchLines409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLines409ResponseWithDefaults

`func NewSearchLines409ResponseWithDefaults() *SearchLines409Response`

NewSearchLines409ResponseWithDefaults instantiates a new SearchLines409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityName

`func (o *SearchLines409Response) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *SearchLines409Response) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *SearchLines409Response) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### GetNumMatches

`func (o *SearchLines409Response) GetNumMatches() int32`

GetNumMatches returns the NumMatches field if non-nil, zero value otherwise.

### GetNumMatchesOk

`func (o *SearchLines409Response) GetNumMatchesOk() (*int32, bool)`

GetNumMatchesOk returns a tuple with the NumMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMatches

`func (o *SearchLines409Response) SetNumMatches(v int32)`

SetNumMatches sets NumMatches field to given value.


### GetError

`func (o *SearchLines409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchLines409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchLines409Response) SetError(v string)`

SetError sets Error field to given value.


### GetRetryable

`func (o *SearchLines409Response) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *SearchLines409Response) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *SearchLines409Response) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


