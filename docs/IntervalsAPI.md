# \IntervalsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2IntervalDeletePost**](IntervalsAPI.md#V2IntervalDeletePost) | **Post** /v2/interval/delete | 
[**V2IntervalSearchPost**](IntervalsAPI.md#V2IntervalSearchPost) | **Post** /v2/interval/search | 
[**V2IntervalSetPost**](IntervalsAPI.md#V2IntervalSetPost) | **Post** /v2/interval/set | 
[**V2IntervalTypeSearchPost**](IntervalsAPI.md#V2IntervalTypeSearchPost) | **Post** /v2/interval_type/search | 
[**V2IntervalsDeletePost**](IntervalsAPI.md#V2IntervalsDeletePost) | **Post** /v2/intervals/delete | 
[**V2IntervalsSetPost**](IntervalsAPI.md#V2IntervalsSetPost) | **Post** /v2/intervals/set | 



## V2IntervalDeletePost

> []Interval V2IntervalDeletePost(ctx).Interval(interval).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalDeletePost(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalDeletePost`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2IntervalSearchPost

> []Interval V2IntervalSearchPost(ctx).Interval(interval).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalSearchPost(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalSearchPost`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2IntervalSetPost

> Interval V2IntervalSetPost(ctx).Interval(interval).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalSetPost(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalSetPost`: Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2IntervalTypeSearchPost

> []IntervalType V2IntervalTypeSearchPost(ctx).IntervalType(intervalType).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	intervalType := *openapiclient.NewIntervalType() // IntervalType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalTypeSearchPost(context.Background()).IntervalType(intervalType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalTypeSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalTypeSearchPost`: []IntervalType
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalTypeSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalTypeSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalType** | [**IntervalType**](IntervalType.md) |  | 

### Return type

[**[]IntervalType**](IntervalType.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2IntervalsDeletePost

> V2IntervalsDeletePost200Response V2IntervalsDeletePost(ctx).IntervalBulkDelete(intervalBulkDelete).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	intervalBulkDelete := *openapiclient.NewIntervalBulkDelete() // IntervalBulkDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalsDeletePost(context.Background()).IntervalBulkDelete(intervalBulkDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalsDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalsDeletePost`: V2IntervalsDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalsDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalBulkDelete** | [**IntervalBulkDelete**](IntervalBulkDelete.md) |  | 

### Return type

[**V2IntervalsDeletePost200Response**](V2IntervalsDeletePost200Response.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2IntervalsSetPost

> V2IntervalsSetPost200Response V2IntervalsSetPost(ctx).IntervalBulkCreate(intervalBulkCreate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/odenio/oden-client-go"
)

func main() {
	intervalBulkCreate := *openapiclient.NewIntervalBulkCreate() // IntervalBulkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.V2IntervalsSetPost(context.Background()).IntervalBulkCreate(intervalBulkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.V2IntervalsSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2IntervalsSetPost`: V2IntervalsSetPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.V2IntervalsSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2IntervalsSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalBulkCreate** | [**IntervalBulkCreate**](IntervalBulkCreate.md) |  | 

### Return type

[**V2IntervalsSetPost200Response**](V2IntervalsSetPost200Response.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

