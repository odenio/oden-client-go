# \StatesAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchStateCategories**](StatesAPI.md#SearchStateCategories) | **Post** /v2/state_category/search | Search state categories
[**SearchStateReasons**](StatesAPI.md#SearchStateReasons) | **Post** /v2/state_reason/search | Search state reasons on one or more lines
[**SetStateCategory**](StatesAPI.md#SetStateCategory) | **Post** /v2/state_category/set | Create or update a state category
[**SetStateReasons**](StatesAPI.md#SetStateReasons) | **Post** /v2/state_reason/set | Create a state reason on one or more lines



## SearchStateCategories

> []StateCategoryDetail SearchStateCategories(ctx).StateCategorySearch(stateCategorySearch).Execute()

Search state categories



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
	stateCategorySearch := *openapiclient.NewStateCategorySearch() // StateCategorySearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatesAPI.SearchStateCategories(context.Background()).StateCategorySearch(stateCategorySearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatesAPI.SearchStateCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStateCategories`: []StateCategoryDetail
	fmt.Fprintf(os.Stdout, "Response from `StatesAPI.SearchStateCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStateCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateCategorySearch** | [**StateCategorySearch**](StateCategorySearch.md) |  | 

### Return type

[**[]StateCategoryDetail**](StateCategoryDetail.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStateReasons

> []StateReasonDetail SearchStateReasons(ctx).StateReasonSearch(stateReasonSearch).Execute()

Search state reasons on one or more lines



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
	stateReasonSearch := *openapiclient.NewStateReasonSearch([]openapiclient.Line{*openapiclient.NewLine()}) // StateReasonSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatesAPI.SearchStateReasons(context.Background()).StateReasonSearch(stateReasonSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatesAPI.SearchStateReasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStateReasons`: []StateReasonDetail
	fmt.Fprintf(os.Stdout, "Response from `StatesAPI.SearchStateReasons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStateReasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateReasonSearch** | [**StateReasonSearch**](StateReasonSearch.md) |  | 

### Return type

[**[]StateReasonDetail**](StateReasonDetail.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStateCategory

> StateCategoryDetail SetStateCategory(ctx).StateCategorySet(stateCategorySet).Execute()

Create or update a state category



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
	stateCategorySet := *openapiclient.NewStateCategorySet() // StateCategorySet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatesAPI.SetStateCategory(context.Background()).StateCategorySet(stateCategorySet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatesAPI.SetStateCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStateCategory`: StateCategoryDetail
	fmt.Fprintf(os.Stdout, "Response from `StatesAPI.SetStateCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetStateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateCategorySet** | [**StateCategorySet**](StateCategorySet.md) |  | 

### Return type

[**StateCategoryDetail**](StateCategoryDetail.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStateReasons

> []StateReasonDetail SetStateReasons(ctx).StateReasonSet(stateReasonSet).Execute()

Create a state reason on one or more lines



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
	stateReasonSet := *openapiclient.NewStateReasonSet([]openapiclient.Line{*openapiclient.NewLine()}, *openapiclient.NewStateCategory(), "Name_example") // StateReasonSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatesAPI.SetStateReasons(context.Background()).StateReasonSet(stateReasonSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatesAPI.SetStateReasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStateReasons`: []StateReasonDetail
	fmt.Fprintf(os.Stdout, "Response from `StatesAPI.SetStateReasons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetStateReasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateReasonSet** | [**StateReasonSet**](StateReasonSet.md) |  | 

### Return type

[**[]StateReasonDetail**](StateReasonDetail.md)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

