# \MachineGroupsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFactories**](MachineGroupsAPI.md#SearchFactories) | **Post** /v2/factory/search | Search factories
[**SearchLines**](MachineGroupsAPI.md#SearchLines) | **Post** /v2/line/search | Search production lines



## SearchFactories

> []Factory SearchFactories(ctx).Factory(factory).Execute()

Search factories



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
	resp, r, err := apiClient.MachineGroupsAPI.SearchFactories(context.Background()).Factory(factory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineGroupsAPI.SearchFactories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFactories`: []Factory
	fmt.Fprintf(os.Stdout, "Response from `MachineGroupsAPI.SearchFactories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFactoriesRequest struct via the builder pattern


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


## SearchLines

> []Line SearchLines(ctx).Line(line).Execute()

Search production lines



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
	resp, r, err := apiClient.MachineGroupsAPI.SearchLines(context.Background()).Line(line).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineGroupsAPI.SearchLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLines`: []Line
	fmt.Fprintf(os.Stdout, "Response from `MachineGroupsAPI.SearchLines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLinesRequest struct via the builder pattern


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

