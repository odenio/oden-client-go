# \ProductMappingsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchProductMappings**](ProductMappingsAPI.md#SearchProductMappings) | **Post** /v2/product_mapping/search | Search product-to-line mappings
[**SetProductMapping**](ProductMappingsAPI.md#SetProductMapping) | **Post** /v2/product_mapping/set | Map a product to a line



## SearchProductMappings

> []ProductMapping SearchProductMappings(ctx).ProductMapping(productMapping).Execute()

Search product-to-line mappings



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
	resp, r, err := apiClient.ProductMappingsAPI.SearchProductMappings(context.Background()).ProductMapping(productMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductMappingsAPI.SearchProductMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProductMappings`: []ProductMapping
	fmt.Fprintf(os.Stdout, "Response from `ProductMappingsAPI.SearchProductMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductMappingsRequest struct via the builder pattern


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


## SetProductMapping

> SetProductMapping(ctx).ProductMapping(productMapping).Execute()

Map a product to a line



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
	r, err := apiClient.ProductMappingsAPI.SetProductMapping(context.Background()).ProductMapping(productMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductMappingsAPI.SetProductMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetProductMappingRequest struct via the builder pattern


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

