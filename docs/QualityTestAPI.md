# \QualityTestAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteQualityTests**](QualityTestAPI.md#BulkDeleteQualityTests) | **Post** /v2/quality_tests/delete | Delete multiple quality tests
[**DeleteQualityTest**](QualityTestAPI.md#DeleteQualityTest) | **Post** /v2/quality_test/delete | Delete a quality test
[**SearchQualitySchemas**](QualityTestAPI.md#SearchQualitySchemas) | **Post** /v2/quality_schema/search | Search quality schemas for a factory
[**SearchQualityTests**](QualityTestAPI.md#SearchQualityTests) | **Post** /v2/quality_test/search | Search quality tests
[**SetQualityTest**](QualityTestAPI.md#SetQualityTest) | **Post** /v2/quality_test/set | Create or update a quality test result



## BulkDeleteQualityTests

> BulkDeleteQualityTests(ctx).BulkDeleteQualityTestsRequest(bulkDeleteQualityTestsRequest).Execute()

Delete multiple quality tests



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
	bulkDeleteQualityTestsRequest := *openapiclient.NewBulkDeleteQualityTestsRequest() // BulkDeleteQualityTestsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityTestAPI.BulkDeleteQualityTests(context.Background()).BulkDeleteQualityTestsRequest(bulkDeleteQualityTestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.BulkDeleteQualityTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteQualityTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkDeleteQualityTestsRequest** | [**BulkDeleteQualityTestsRequest**](BulkDeleteQualityTestsRequest.md) |  | 

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


## DeleteQualityTest

> DeleteQualityTest(ctx).QualityTest(qualityTest).Execute()

Delete a quality test



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
	r, err := apiClient.QualityTestAPI.DeleteQualityTest(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.DeleteQualityTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQualityTestRequest struct via the builder pattern


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


## SearchQualitySchemas

> SearchQualitySchemas(ctx).QualitySchema(qualitySchema).Execute()

Search quality schemas for a factory



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
	r, err := apiClient.QualityTestAPI.SearchQualitySchemas(context.Background()).QualitySchema(qualitySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.SearchQualitySchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchQualitySchemasRequest struct via the builder pattern


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


## SearchQualityTests

> SearchQualityTests(ctx).QualityTest(qualityTest).Execute()

Search quality tests



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
	r, err := apiClient.QualityTestAPI.SearchQualityTests(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.SearchQualityTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchQualityTestsRequest struct via the builder pattern


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


## SetQualityTest

> SetQualityTest(ctx).QualityTest(qualityTest).Execute()

Create or update a quality test result



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
	r, err := apiClient.QualityTestAPI.SetQualityTest(context.Background()).QualityTest(qualityTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityTestAPI.SetQualityTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQualityTestRequest struct via the builder pattern


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

