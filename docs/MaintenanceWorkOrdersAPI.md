# \MaintenanceWorkOrdersAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceWorkOrder**](MaintenanceWorkOrdersAPI.md#DeleteMaintenanceWorkOrder) | **Post** /v2/maintenance_work_order/delete | Delete a maintenance work order
[**SearchMaintenanceWorkOrders**](MaintenanceWorkOrdersAPI.md#SearchMaintenanceWorkOrders) | **Post** /v2/maintenance_work_order/search | Search maintenance work orders
[**SetMaintenanceWorkOrder**](MaintenanceWorkOrdersAPI.md#SetMaintenanceWorkOrder) | **Post** /v2/maintenance_work_order/set | Create or update a maintenance work order



## DeleteMaintenanceWorkOrder

> []MaintenanceWorkOrder DeleteMaintenanceWorkOrder(ctx).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()

Delete a maintenance work order



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
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.DeleteMaintenanceWorkOrder(context.Background()).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.DeleteMaintenanceWorkOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMaintenanceWorkOrder`: []MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.DeleteMaintenanceWorkOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaintenanceWorkOrderRequest struct via the builder pattern


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


## SearchMaintenanceWorkOrders

> []MaintenanceWorkOrder SearchMaintenanceWorkOrders(ctx).SearchMaintenanceWorkOrdersRequest(searchMaintenanceWorkOrdersRequest).Execute()

Search maintenance work orders



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
	searchMaintenanceWorkOrdersRequest := *openapiclient.NewSearchMaintenanceWorkOrdersRequest() // SearchMaintenanceWorkOrdersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.SearchMaintenanceWorkOrders(context.Background()).SearchMaintenanceWorkOrdersRequest(searchMaintenanceWorkOrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.SearchMaintenanceWorkOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMaintenanceWorkOrders`: []MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.SearchMaintenanceWorkOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMaintenanceWorkOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchMaintenanceWorkOrdersRequest** | [**SearchMaintenanceWorkOrdersRequest**](SearchMaintenanceWorkOrdersRequest.md) |  | 

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


## SetMaintenanceWorkOrder

> MaintenanceWorkOrder SetMaintenanceWorkOrder(ctx).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()

Create or update a maintenance work order



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
	resp, r, err := apiClient.MaintenanceWorkOrdersAPI.SetMaintenanceWorkOrder(context.Background()).MaintenanceWorkOrder(maintenanceWorkOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWorkOrdersAPI.SetMaintenanceWorkOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMaintenanceWorkOrder`: MaintenanceWorkOrder
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWorkOrdersAPI.SetMaintenanceWorkOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMaintenanceWorkOrderRequest struct via the builder pattern


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

