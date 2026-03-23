# \MaintenanceWorkOrdersAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MaintenanceWorkOrderDeletePost**](MaintenanceWorkOrdersAPI.md#V2MaintenanceWorkOrderDeletePost) | **Post** /v2/maintenance_work_order/delete | 
[**V2MaintenanceWorkOrderSearchPost**](MaintenanceWorkOrdersAPI.md#V2MaintenanceWorkOrderSearchPost) | **Post** /v2/maintenance_work_order/search | 
[**V2MaintenanceWorkOrderSetPost**](MaintenanceWorkOrdersAPI.md#V2MaintenanceWorkOrderSetPost) | **Post** /v2/maintenance_work_order/set | 



## V2MaintenanceWorkOrderDeletePost

> []MaintenanceWorkOrder V2MaintenanceWorkOrderDeletePost(ctx).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()





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
	maintenanceWorkOrder := *openapiclient.NewMaintenanceWorkOrder() // MaintenanceWorkOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderDeletePost(context.Background()).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MaintenanceWorkOrderDeletePost`: []MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MaintenanceWorkOrderDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWorkOrder** | [**MaintenanceWorkOrder**](MaintenanceWorkOrder.md) |  | 

### Return type

[**[]MaintenanceWorkOrder**](MaintenanceWorkOrder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MaintenanceWorkOrderSearchPost

> []MaintenanceWorkOrder V2MaintenanceWorkOrderSearchPost(ctx).V2MaintenanceWorkOrderSearchPostRequest(v2MaintenanceWorkOrderSearchPostRequest).Execute()





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
	v2MaintenanceWorkOrderSearchPostRequest := *openapiclient.NewV2MaintenanceWorkOrderSearchPostRequest() // V2MaintenanceWorkOrderSearchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSearchPost(context.Background()).V2MaintenanceWorkOrderSearchPostRequest(v2MaintenanceWorkOrderSearchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MaintenanceWorkOrderSearchPost`: []MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MaintenanceWorkOrderSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2MaintenanceWorkOrderSearchPostRequest** | [**V2MaintenanceWorkOrderSearchPostRequest**](V2MaintenanceWorkOrderSearchPostRequest.md) |  | 

### Return type

[**[]MaintenanceWorkOrder**](MaintenanceWorkOrder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MaintenanceWorkOrderSetPost

> MaintenanceWorkOrder V2MaintenanceWorkOrderSetPost(ctx).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()





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
	maintenanceWorkOrder := *openapiclient.NewMaintenanceWorkOrder() // MaintenanceWorkOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSetPost(context.Background()).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MaintenanceWorkOrderSetPost`: MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.V2MaintenanceWorkOrderSetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MaintenanceWorkOrderSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWorkOrder** | [**MaintenanceWorkOrder**](MaintenanceWorkOrder.md) |  | 

### Return type

[**MaintenanceWorkOrder**](MaintenanceWorkOrder.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

