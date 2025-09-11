# \ProductsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ProductDeletePost**](ProductsAPI.md#V2ProductDeletePost) | **Post** /v2/product/delete | 
[**V2ProductSearchPost**](ProductsAPI.md#V2ProductSearchPost) | **Post** /v2/product/search | 
[**V2ProductSetPost**](ProductsAPI.md#V2ProductSetPost) | **Post** /v2/product/set | 



## V2ProductDeletePost

> V2ProductDeletePost(ctx).Product(product).Execute()





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
	product := *openapiclient.NewProduct() // Product | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductsAPI.V2ProductDeletePost(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.V2ProductDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

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


## V2ProductSearchPost

> []Product V2ProductSearchPost(ctx).Product(product).Execute()





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
	product := *openapiclient.NewProduct() // Product | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.V2ProductSearchPost(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.V2ProductSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ProductSearchPost`: []Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.V2ProductSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

### Return type

[**[]Product**](Product.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ProductSetPost

> V2ProductSetPost(ctx).Product(product).Execute()





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
	product := *openapiclient.NewProduct() // Product | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductsAPI.V2ProductSetPost(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.V2ProductSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ProductSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

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

