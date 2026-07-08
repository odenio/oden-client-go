# SearchLines400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Name** | **string** |  | 
**Retryable** | **bool** |  | 

## Methods

### NewSearchLines400Response

`func NewSearchLines400Response(error_ string, name string, retryable bool, ) *SearchLines400Response`

NewSearchLines400Response instantiates a new SearchLines400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLines400ResponseWithDefaults

`func NewSearchLines400ResponseWithDefaults() *SearchLines400Response`

NewSearchLines400ResponseWithDefaults instantiates a new SearchLines400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SearchLines400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchLines400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchLines400Response) SetError(v string)`

SetError sets Error field to given value.


### GetName

`func (o *SearchLines400Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchLines400Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchLines400Response) SetName(v string)`

SetName sets Name field to given value.


### GetRetryable

`func (o *SearchLines400Response) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *SearchLines400Response) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *SearchLines400Response) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


