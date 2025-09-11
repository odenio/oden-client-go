# \ScrapYieldDataAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ScrapYieldDeletePost**](ScrapYieldDataAPI.md#V2ScrapYieldDeletePost) | **Post** /v2/scrap_yield/delete | 
[**V2ScrapYieldSearchPost**](ScrapYieldDataAPI.md#V2ScrapYieldSearchPost) | **Post** /v2/scrap_yield/search | 
[**V2ScrapYieldSetPost**](ScrapYieldDataAPI.md#V2ScrapYieldSetPost) | **Post** /v2/scrap_yield/set | 



## V2ScrapYieldDeletePost

> V2ScrapYieldDeletePost(ctx).V2ScrapYieldSearchPostRequest(v2ScrapYieldSearchPostRequest).Execute()





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
	v2ScrapYieldSearchPostRequest := *openapiclient.NewV2ScrapYieldSearchPostRequest(*openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // V2ScrapYieldSearchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.V2ScrapYieldDeletePost(context.Background()).V2ScrapYieldSearchPostRequest(v2ScrapYieldSearchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.V2ScrapYieldDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ScrapYieldDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ScrapYieldSearchPostRequest** | [**V2ScrapYieldSearchPostRequest**](V2ScrapYieldSearchPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ScrapYieldSearchPost

> V2ScrapYieldSearchPost(ctx).V2ScrapYieldSearchPostRequest(v2ScrapYieldSearchPostRequest).Execute()





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
	v2ScrapYieldSearchPostRequest := *openapiclient.NewV2ScrapYieldSearchPostRequest(*openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // V2ScrapYieldSearchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.V2ScrapYieldSearchPost(context.Background()).V2ScrapYieldSearchPostRequest(v2ScrapYieldSearchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.V2ScrapYieldSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ScrapYieldSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ScrapYieldSearchPostRequest** | [**V2ScrapYieldSearchPostRequest**](V2ScrapYieldSearchPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ScrapYieldSetPost

> V2ScrapYieldSetPost(ctx).V2ScrapYieldSetPostRequest(v2ScrapYieldSetPostRequest).Pattern(pattern).Execute()





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
	v2ScrapYieldSetPostRequest := *openapiclient.NewV2ScrapYieldSetPostRequest(*openapiclient.NewScrapYieldData(), *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // V2ScrapYieldSetPostRequest | 
	pattern := "pattern_example" // string | Optional pattern type to use for matching: - `exact` for exact match - `contains` for the string to be contained in the query - `regex` to match based on a regular expression  (optional) (default to "exact")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.V2ScrapYieldSetPost(context.Background()).V2ScrapYieldSetPostRequest(v2ScrapYieldSetPostRequest).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.V2ScrapYieldSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ScrapYieldSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ScrapYieldSetPostRequest** | [**V2ScrapYieldSetPostRequest**](V2ScrapYieldSetPostRequest.md) |  | 
 **pattern** | **string** | Optional pattern type to use for matching: - &#x60;exact&#x60; for exact match - &#x60;contains&#x60; for the string to be contained in the query - &#x60;regex&#x60; to match based on a regular expression  | [default to &quot;exact&quot;]

### Return type

 (empty response body)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

