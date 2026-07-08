# \IntervalsAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteIntervals**](IntervalsAPI.md#BulkDeleteIntervals) | **Post** /v2/intervals/delete | Delete intervals in a time range
[**BulkSetIntervals**](IntervalsAPI.md#BulkSetIntervals) | **Post** /v2/intervals/set | Create a set of intervals
[**BulkUpdateIntervals**](IntervalsAPI.md#BulkUpdateIntervals) | **Post** /v2/intervals/update | Update a set of intervals
[**DeleteInterval**](IntervalsAPI.md#DeleteInterval) | **Post** /v2/interval/delete | Delete an interval
[**DeleteIntervalType**](IntervalsAPI.md#DeleteIntervalType) | **Post** /v2/interval_type/delete | Delete a custom interval type
[**SearchIntervalTypes**](IntervalsAPI.md#SearchIntervalTypes) | **Post** /v2/interval_type/search | Search interval types
[**SearchIntervals**](IntervalsAPI.md#SearchIntervals) | **Post** /v2/interval/search | Search intervals on a line
[**SetInterval**](IntervalsAPI.md#SetInterval) | **Post** /v2/interval/set | Create or update an interval
[**SetIntervalType**](IntervalsAPI.md#SetIntervalType) | **Post** /v2/interval_type/set | Create or update a custom interval type
[**UpdateInterval**](IntervalsAPI.md#UpdateInterval) | **Post** /v2/interval/update | Update an interval



## BulkDeleteIntervals

> BulkDeleteIntervals200Response BulkDeleteIntervals(ctx).IntervalBulkDelete(intervalBulkDelete).Execute()

Delete intervals in a time range



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
	intervalBulkDelete := *openapiclient.NewIntervalBulkDelete() // IntervalBulkDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.BulkDeleteIntervals(context.Background()).IntervalBulkDelete(intervalBulkDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.BulkDeleteIntervals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteIntervals`: BulkDeleteIntervals200Response
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.BulkDeleteIntervals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteIntervalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalBulkDelete** | [**IntervalBulkDelete**](IntervalBulkDelete.md) |  | 

### Return type

[**BulkDeleteIntervals200Response**](BulkDeleteIntervals200Response.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetIntervals

> []string BulkSetIntervals(ctx).IntervalBulkCreate(intervalBulkCreate).Execute()

Create a set of intervals



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
	intervalBulkCreate := *openapiclient.NewIntervalBulkCreate() // IntervalBulkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.BulkSetIntervals(context.Background()).IntervalBulkCreate(intervalBulkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.BulkSetIntervals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkSetIntervals`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.BulkSetIntervals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetIntervalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalBulkCreate** | [**IntervalBulkCreate**](IntervalBulkCreate.md) |  | 

### Return type

**[]string**

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateIntervals

> BulkUpdateIntervals200Response BulkUpdateIntervals(ctx).IntervalBulkUpdate(intervalBulkUpdate).Execute()

Update a set of intervals



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
	intervalBulkUpdate := *openapiclient.NewIntervalBulkUpdate() // IntervalBulkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.BulkUpdateIntervals(context.Background()).IntervalBulkUpdate(intervalBulkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.BulkUpdateIntervals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateIntervals`: BulkUpdateIntervals200Response
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.BulkUpdateIntervals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateIntervalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalBulkUpdate** | [**IntervalBulkUpdate**](IntervalBulkUpdate.md) |  | 

### Return type

[**BulkUpdateIntervals200Response**](BulkUpdateIntervals200Response.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterval

> []Interval DeleteInterval(ctx).Interval(interval).Execute()

Delete an interval



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
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.DeleteInterval(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.DeleteInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInterval`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.DeleteInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntervalType

> IntervalType DeleteIntervalType(ctx).IntervalType(intervalType).Execute()

Delete a custom interval type



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
	intervalType := *openapiclient.NewIntervalType() // IntervalType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.DeleteIntervalType(context.Background()).IntervalType(intervalType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.DeleteIntervalType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntervalType`: IntervalType
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.DeleteIntervalType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntervalTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalType** | [**IntervalType**](IntervalType.md) |  | 

### Return type

[**IntervalType**](IntervalType.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIntervalTypes

> []IntervalType SearchIntervalTypes(ctx).IntervalType(intervalType).Execute()

Search interval types



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
	intervalType := *openapiclient.NewIntervalType() // IntervalType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.SearchIntervalTypes(context.Background()).IntervalType(intervalType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.SearchIntervalTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIntervalTypes`: []IntervalType
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.SearchIntervalTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIntervalTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalType** | [**IntervalType**](IntervalType.md) |  | 

### Return type

[**[]IntervalType**](IntervalType.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIntervals

> []Interval SearchIntervals(ctx).Interval(interval).Execute()

Search intervals on a line



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
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.SearchIntervals(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.SearchIntervals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIntervals`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.SearchIntervals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIntervalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInterval

> []Interval SetInterval(ctx).Interval(interval).Execute()

Create or update an interval



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
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.SetInterval(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.SetInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetInterval`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.SetInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIntervalType

> IntervalType SetIntervalType(ctx).IntervalTypeSet(intervalTypeSet).Execute()

Create or update a custom interval type



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
	intervalTypeSet := *openapiclient.NewIntervalTypeSet() // IntervalTypeSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.SetIntervalType(context.Background()).IntervalTypeSet(intervalTypeSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.SetIntervalType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetIntervalType`: IntervalType
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.SetIntervalType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetIntervalTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intervalTypeSet** | [**IntervalTypeSet**](IntervalTypeSet.md) |  | 

### Return type

[**IntervalType**](IntervalType.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterval

> []Interval UpdateInterval(ctx).Interval(interval).Execute()

Update an interval



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
	interval := *openapiclient.NewInterval(*openapiclient.NewIntervalType(), *openapiclient.NewLine()) // Interval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntervalsAPI.UpdateInterval(context.Background()).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntervalsAPI.UpdateInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInterval`: []Interval
	fmt.Fprintf(os.Stdout, "Response from `IntervalsAPI.UpdateInterval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**Interval**](Interval.md) |  | 

### Return type

[**[]Interval**](Interval.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

