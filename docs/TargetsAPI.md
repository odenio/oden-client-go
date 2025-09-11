# \TargetsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2TargetSearchPost**](TargetsAPI.md#V2TargetSearchPost) | **Post** /v2/target/search | 
[**V2TargetSetPost**](TargetsAPI.md#V2TargetSetPost) | **Post** /v2/target/set | 



## V2TargetSearchPost

> []Target V2TargetSearchPost(ctx).Target(target).Execute()





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
	resp, r, err := apiClient.TargetsAPI.V2TargetSearchPost(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.V2TargetSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetSearchPost`: []Target
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.V2TargetSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetSearchPostRequest struct via the builder pattern


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


## V2TargetSetPost

> Target V2TargetSetPost(ctx).Target(target).Execute()





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
	resp, r, err := apiClient.TargetsAPI.V2TargetSetPost(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.V2TargetSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetSetPost`: Target
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.V2TargetSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetSetPostRequest struct via the builder pattern


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

