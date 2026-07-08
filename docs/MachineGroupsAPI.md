# \MachineGroupsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2FactorySearchPost**](MachineGroupsAPI.md#V2FactorySearchPost) | **Post** /v2/factory/search | 
[**V2LineSearchPost**](MachineGroupsAPI.md#V2LineSearchPost) | **Post** /v2/line/search | 



## V2FactorySearchPost

> []Factory V2FactorySearchPost(ctx).Factory(factory).Execute()





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
	factory := *openapiclient.NewFactory() // Factory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineGroupsAPI.V2FactorySearchPost(context.Background()).Factory(factory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineGroupsAPI.V2FactorySearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FactorySearchPost`: []Factory
	fmt.Fprintf(os.Stdout, "Response from `MachineGroupsAPI.V2FactorySearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2FactorySearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **factory** | [**Factory**](Factory.md) |  | 

### Return type

[**[]Factory**](Factory.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LineSearchPost

> []Line V2LineSearchPost(ctx).Line(line).Execute()





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
	line := *openapiclient.NewLine() // Line | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineGroupsAPI.V2LineSearchPost(context.Background()).Line(line).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineGroupsAPI.V2LineSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LineSearchPost`: []Line
	fmt.Fprintf(os.Stdout, "Response from `MachineGroupsAPI.V2LineSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2LineSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **line** | [**Line**](Line.md) |  | 

### Return type

[**[]Line**](Line.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

