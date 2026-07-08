# \MetricGroupsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchMetricGroups**](MetricGroupsAPI.md#SearchMetricGroups) | **Post** /v2/metric_group/search | Search metric groups



## SearchMetricGroups

> []MetricGroup SearchMetricGroups(ctx).MetricGroup(metricGroup).Execute()

Search metric groups



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
	metricGroup := *openapiclient.NewMetricGroup() // MetricGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricGroupsAPI.SearchMetricGroups(context.Background()).MetricGroup(metricGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricGroupsAPI.SearchMetricGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMetricGroups`: []MetricGroup
	fmt.Fprintf(os.Stdout, "Response from `MetricGroupsAPI.SearchMetricGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMetricGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricGroup** | [**MetricGroup**](MetricGroup.md) |  | 

### Return type

[**[]MetricGroup**](MetricGroup.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

