# \ProductsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProduct**](ProductsAPI.md#DeleteProduct) | **Post** /v2/product/delete | Delete a product
[**SearchProducts**](ProductsAPI.md#SearchProducts) | **Post** /v2/product/search | Search products
[**SetProduct**](ProductsAPI.md#SetProduct) | **Post** /v2/product/set | Create or update a product



## DeleteProduct

> DeleteProduct(ctx).Product(product).Execute()

Delete a product



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
	r, err := apiClient.ProductsAPI.DeleteProduct(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.DeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## SearchProducts

> []Product SearchProducts(ctx).Product(product).Execute()

Search products



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
	resp, r, err := apiClient.ProductsAPI.SearchProducts(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.SearchProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProducts`: []Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.SearchProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductsRequest struct via the builder pattern


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


## SetProduct

> SetProduct(ctx).Product(product).Execute()

Create or update a product



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
	r, err := apiClient.ProductsAPI.SetProduct(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.SetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetProductRequest struct via the builder pattern


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

