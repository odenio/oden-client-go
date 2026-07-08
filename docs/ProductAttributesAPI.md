# \ProductAttributesAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchProductAttributes**](ProductAttributesAPI.md#SearchProductAttributes) | **Post** /v2/product_attribute/search | Search product attributes
[**SetProductAttribute**](ProductAttributesAPI.md#SetProductAttribute) | **Post** /v2/product_attribute/set | Create or update a product attribute



## SearchProductAttributes

> []ProductAttribute SearchProductAttributes(ctx).ProductAttribute(productAttribute).Execute()

Search product attributes



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
	productAttribute := *openapiclient.NewProductAttribute() // ProductAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.SearchProductAttributes(context.Background()).ProductAttribute(productAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.SearchProductAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProductAttributes`: []ProductAttribute
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.SearchProductAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAttribute** | [**ProductAttribute**](ProductAttribute.md) |  | 

### Return type

[**[]ProductAttribute**](ProductAttribute.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProductAttribute

> []ProductAttribute SetProductAttribute(ctx).ProductAttribute(productAttribute).Execute()

Create or update a product attribute



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
	productAttribute := *openapiclient.NewProductAttribute() // ProductAttribute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAttributesAPI.SetProductAttribute(context.Background()).ProductAttribute(productAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.SetProductAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProductAttribute`: []ProductAttribute
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.SetProductAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetProductAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAttribute** | [**ProductAttribute**](ProductAttribute.md) |  | 

### Return type

[**[]ProductAttribute**](ProductAttribute.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

