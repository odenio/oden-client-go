# \ProductMappingsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ProductMappingSearchPost**](ProductMappingsAPI.md#V2ProductMappingSearchPost) | **Post** /v2/product_mapping/search | 
[**V2ProductMappingSetPost**](ProductMappingsAPI.md#V2ProductMappingSetPost) | **Post** /v2/product_mapping/set | 



## V2ProductMappingSearchPost

> []ProductMapping V2ProductMappingSearchPost(ctx).ProductMapping(productMapping).Execute()





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
	productMapping := *openapiclient.NewProductMapping() // ProductMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductMappingsAPI.V2ProductMappingSearchPost(context.Background()).ProductMapping(productMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductMappingsAPI.V2ProductMappingSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ProductMappingSearchPost`: []ProductMapping
	fmt.Fprintf(os.Stdout, "Response from `ProductMappingsAPI.V2ProductMappingSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductMappingSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productMapping** | [**ProductMapping**](ProductMapping.md) |  | 

### Return type

[**[]ProductMapping**](ProductMapping.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ProductMappingSetPost

> V2ProductMappingSetPost(ctx).ProductMapping(productMapping).Execute()





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
	productMapping := *openapiclient.NewProductMapping() // ProductMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductMappingsAPI.V2ProductMappingSetPost(context.Background()).ProductMapping(productMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductMappingsAPI.V2ProductMappingSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductMappingSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productMapping** | [**ProductMapping**](ProductMapping.md) |  | 

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

