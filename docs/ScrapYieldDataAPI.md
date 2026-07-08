# \ScrapYieldDataAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScrapYield**](ScrapYieldDataAPI.md#DeleteScrapYield) | **Post** /v2/scrap_yield/delete | Delete a scrap/yield record
[**SearchScrapYield**](ScrapYieldDataAPI.md#SearchScrapYield) | **Post** /v2/scrap_yield/search | Search scrap/yield records
[**SetScrapYield**](ScrapYieldDataAPI.md#SetScrapYield) | **Post** /v2/scrap_yield/set | Create or update a scrap/yield record



## DeleteScrapYield

> DeleteScrapYield(ctx).SearchScrapYieldRequest(searchScrapYieldRequest).Execute()

Delete a scrap/yield record



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
	searchScrapYieldRequest := *openapiclient.NewSearchScrapYieldRequest(*openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // SearchScrapYieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.DeleteScrapYield(context.Background()).SearchScrapYieldRequest(searchScrapYieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.DeleteScrapYield``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScrapYieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchScrapYieldRequest** | [**SearchScrapYieldRequest**](SearchScrapYieldRequest.md) |  | 

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


## SearchScrapYield

> SearchScrapYield(ctx).SearchScrapYieldRequest(searchScrapYieldRequest).Execute()

Search scrap/yield records



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
	searchScrapYieldRequest := *openapiclient.NewSearchScrapYieldRequest(*openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // SearchScrapYieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.SearchScrapYield(context.Background()).SearchScrapYieldRequest(searchScrapYieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.SearchScrapYield``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchScrapYieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchScrapYieldRequest** | [**SearchScrapYieldRequest**](SearchScrapYieldRequest.md) |  | 

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


## SetScrapYield

> SetScrapYield(ctx).SetScrapYieldRequest(setScrapYieldRequest).Pattern(pattern).Execute()

Create or update a scrap/yield record



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
	setScrapYieldRequest := *openapiclient.NewSetScrapYieldRequest(*openapiclient.NewScrapYieldData(), *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine())) // SetScrapYieldRequest | 
	pattern := "pattern_example" // string | Optional pattern type to use for matching: - `exact` for exact match - `contains` for the string to be contained in the query - `regex` to match based on a regular expression  (optional) (default to "exact")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScrapYieldDataAPI.SetScrapYield(context.Background()).SetScrapYieldRequest(setScrapYieldRequest).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScrapYieldDataAPI.SetScrapYield``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetScrapYieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setScrapYieldRequest** | [**SetScrapYieldRequest**](SetScrapYieldRequest.md) |  | 
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

