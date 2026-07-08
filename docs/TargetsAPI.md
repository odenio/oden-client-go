# \TargetsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchTargets**](TargetsAPI.md#SearchTargets) | **Post** /v2/target/search | Search metric targets
[**SetTarget**](TargetsAPI.md#SetTarget) | **Post** /v2/target/set | Create or update a metric target



## SearchTargets

> []Target SearchTargets(ctx).Target(target).Execute()

Search metric targets



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
	target := *openapiclient.NewTarget() // Target | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.SearchTargets(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.SearchTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTargets`: []Target
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.SearchTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | [**Target**](Target.md) |  | 

### Return type

[**[]Target**](Target.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTarget

> Target SetTarget(ctx).Target(target).Execute()

Create or update a metric target



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
	target := *openapiclient.NewTarget() // Target | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.SetTarget(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.SetTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTarget`: Target
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.SetTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | [**Target**](Target.md) |  | 

### Return type

[**Target**](Target.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

