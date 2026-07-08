# \DashboardsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteDashboard**](DashboardsAPI.md#ExecuteDashboard) | **Post** /v2/dashboard/execute | Execute a dashboard



## ExecuteDashboard

> []DashboardExecuteResult ExecuteDashboard(ctx).DashboardExecuteRequest(dashboardExecuteRequest).Execute()

Execute a dashboard



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
	dashboardExecuteRequest := *openapiclient.NewDashboardExecuteRequest(*openapiclient.NewDashboardExecuteRequestDashboard("Id_example")) // DashboardExecuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.ExecuteDashboard(context.Background()).DashboardExecuteRequest(dashboardExecuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ExecuteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteDashboard`: []DashboardExecuteResult
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.ExecuteDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardExecuteRequest** | [**DashboardExecuteRequest**](DashboardExecuteRequest.md) |  | 

### Return type

[**[]DashboardExecuteResult**](DashboardExecuteResult.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

