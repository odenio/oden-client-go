# \ProductAttributesAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ProductAttributeSearchPost**](ProductAttributesAPI.md#V2ProductAttributeSearchPost) | **Post** /v2/product_attribute/search | 
[**V2ProductAttributeSetPost**](ProductAttributesAPI.md#V2ProductAttributeSetPost) | **Post** /v2/product_attribute/set | 



## V2ProductAttributeSearchPost

> []ProductAttribute V2ProductAttributeSearchPost(ctx).ProductAttribute(productAttribute).Execute()





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
	resp, r, err := apiClient.ProductAttributesAPI.V2ProductAttributeSearchPost(context.Background()).ProductAttribute(productAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.V2ProductAttributeSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ProductAttributeSearchPost`: []ProductAttribute
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.V2ProductAttributeSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductAttributeSearchPostRequest struct via the builder pattern


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


## V2ProductAttributeSetPost

> []ProductAttribute V2ProductAttributeSetPost(ctx).ProductAttribute(productAttribute).Execute()





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
	resp, r, err := apiClient.ProductAttributesAPI.V2ProductAttributeSetPost(context.Background()).ProductAttribute(productAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAttributesAPI.V2ProductAttributeSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ProductAttributeSetPost`: []ProductAttribute
	fmt.Fprintf(os.Stdout, "Response from `ProductAttributesAPI.V2ProductAttributeSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductAttributeSetPostRequest struct via the builder pattern


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

