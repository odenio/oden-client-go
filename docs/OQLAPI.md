# \OQLAPI

All URIs are relative to *https://api.oden.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2OqlQueryPost**](OQLAPI.md#V2OqlQueryPost) | **Post** /v2/oql/query | 



## V2OqlQueryPost

> V2OqlQueryPost(ctx).OQLQuery(oQLQuery).Format(format).Execute()





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
	oQLQuery := *openapiclient.NewOQLQuery("Query_example") // OQLQuery | 
	format := "format_example" // string | Format of the response. Can be `json`, `jsonextended` or `csv`. If unspecified, defaults to `jsonextended`.  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OQLAPI.V2OqlQueryPost(context.Background()).OQLQuery(oQLQuery).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OQLAPI.V2OqlQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2OqlQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oQLQuery** | [**OQLQuery**](OQLQuery.md) |  | 
 **format** | **string** | Format of the response. Can be &#x60;json&#x60;, &#x60;jsonextended&#x60; or &#x60;csv&#x60;. If unspecified, defaults to &#x60;jsonextended&#x60;.  | [default to &quot;json&quot;]

### Return type

 (empty response body)

### Authorization

[APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

