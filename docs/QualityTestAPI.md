# \QualityTestAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2QualitySchemaSearchPost**](QualityTestAPI.md#V2QualitySchemaSearchPost) | **Post** /v2/quality_schema/search | 
[**V2QualityTestDeletePost**](QualityTestAPI.md#V2QualityTestDeletePost) | **Post** /v2/quality_test/delete | 
[**V2QualityTestSearchPost**](QualityTestAPI.md#V2QualityTestSearchPost) | **Post** /v2/quality_test/search | 
[**V2QualityTestSetPost**](QualityTestAPI.md#V2QualityTestSetPost) | **Post** /v2/quality_test/set | 
[**V2QualityTestsDeletePost**](QualityTestAPI.md#V2QualityTestsDeletePost) | **Post** /v2/quality_tests/delete | 



## V2QualitySchemaSearchPost

> V2QualitySchemaSearchPost(ctx).QualitySchema(qualitySchema).Execute()





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
	qualitySchema := *openapiclient.NewQualitySchema() // QualitySchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.V2QualitySchemaSearchPost(context.Background()).QualitySchema(qualitySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.V2QualitySchemaSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2QualitySchemaSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualitySchema** | [**QualitySchema**](QualitySchema.md) |  | 

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


## V2QualityTestDeletePost

> V2QualityTestDeletePost(ctx).QualityTest(qualityTest).Execute()





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
	qualityTest := *openapiclient.NewQualityTest() // QualityTest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.V2QualityTestDeletePost(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.V2QualityTestDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2QualityTestDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityTest** | [**QualityTest**](QualityTest.md) |  | 

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


## V2QualityTestSearchPost

> V2QualityTestSearchPost(ctx).QualityTest(qualityTest).Execute()





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
	qualityTest := *openapiclient.NewQualityTest() // QualityTest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.V2QualityTestSearchPost(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.V2QualityTestSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2QualityTestSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityTest** | [**QualityTest**](QualityTest.md) |  | 

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


## V2QualityTestSetPost

> V2QualityTestSetPost(ctx).QualityTest(qualityTest).Execute()





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
	qualityTest := *openapiclient.NewQualityTest() // QualityTest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.V2QualityTestSetPost(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.V2QualityTestSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2QualityTestSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityTest** | [**QualityTest**](QualityTest.md) |  | 

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


## V2QualityTestsDeletePost

> V2QualityTestsDeletePost(ctx).V2QualityTestsDeletePostRequest(v2QualityTestsDeletePostRequest).Execute()





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
	v2QualityTestsDeletePostRequest := *openapiclient.NewV2QualityTestsDeletePostRequest() // V2QualityTestsDeletePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.V2QualityTestsDeletePost(context.Background()).V2QualityTestsDeletePostRequest(v2QualityTestsDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.V2QualityTestsDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2QualityTestsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2QualityTestsDeletePostRequest** | [**V2QualityTestsDeletePostRequest**](V2QualityTestsDeletePostRequest.md) |  | 

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

