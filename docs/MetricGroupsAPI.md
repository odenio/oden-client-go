# \MetricGroupsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MetricGroupSearchPost**](MetricGroupsAPI.md#V2MetricGroupSearchPost) | **Post** /v2/metric_group/search | 



## V2MetricGroupSearchPost

> []MetricGroup V2MetricGroupSearchPost(ctx).MetricGroup(metricGroup).Execute()





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
	resp, r, err := apiClient.MetricGroupsAPI.V2MetricGroupSearchPost(context.Background()).MetricGroup(metricGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricGroupsAPI.V2MetricGroupSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MetricGroupSearchPost`: []MetricGroup
	fmt.Fprintf(os.Stdout, "Response from `MetricGroupsAPI.V2MetricGroupSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MetricGroupSearchPostRequest struct via the builder pattern


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

