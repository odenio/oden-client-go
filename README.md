# Go API client for oden

The Oden Private Partner API exposes RESTful API endpoints for clients to get, create and update data on the Oden Platform.

The API is based on the OpenAPI 3.0 specification.
### Current Version
The URL, and host, for the current version is [https://api.oden.app/v2](https://api.oden.app/v2).

### Oden's Data Model
- **Organization**: This represents the Organization registered as an Oden customer. An organization has an associated collection of users, factories, lines, etc. This is the entity a specific authentication token is associated with.
- **Asset** or **Machinegroup**: Assets, or machinegroups, are collections of machines, which may either be a **Factory** or a **Line**:
  - **Factory**: Factories are collections of lines, representing a particular manufacturing location.
    - **Line**: Lines are collections of machines, often representing a particular production line. Lines may also have **Products** mapped to them, indicating what is currently being manufactured by the specific line.
      - **Machine**: Machines are the physical machines that make up a line
- **Product**: Products capture what entities a manufacturer produces
- **Interval**: An interval is a period of time that takes place on a manufacturing line and expresses some business concern. It's Oden's way of making metrics aggregatable, traceable, and relatable to a manufacturer.
  - **Run**: A run is a production interval that labels a period of production as being work on some single product
  - **Batch**: A batch is a production interval that represents a portion of a particular run
  - **State**: A state is an interval that tracks the availability or utilization of a line
    - **State Category**: A state category describes what state a line is in - such as ex. uptime, downtime, scrapping, etc.
    - **State Reason**: A state reason describes why a line is in a particular state - such as \"maintenance\" being a reason for the category \"downtime\".
  - **Custom**: A custom interval can track any other type of interval-based data a manufacturer might want to analyze. These are created on a per-factory basis.
- **Target**: Targets specify values and upper/lower thresholds for metrics when specific products are running on specific lines
- **Scrap/Yield**: Scrap/yield output specifies amount of produced product on a line during either a run or batch interval. Oden will categorize all output as either scrap or yield - as specified by the Scrap Yield Schema for a given factory. If you have other categories, like rework/blocked/off-grade, you must choose between categorizing those amounts as either good or bad production by specifying as scrap or yield. Clients may also add scrap codes (i.e., reasons) to a given Scrap Yield Data entry.
  - **Scrap Code**: A scrap code is a code that explains the reason for a scrap/yield raw data input - such as \"Rework\"
- **Quality Test**: Quality Tests are results of quality assurance tests done on site, and uploaded to Oden. They may be attached to a single Batch or Run.
- **Metric**: Known in factories as \"tags\", metrics are the raw data that is collected by Oden from the machines and devices on the factory floor.
- **Metric Group**: Metric groups are metrics that represent the same thing across different lines. They provide common display names for tags and allow labeling groups of tags as measuring key types of data like performance or production.

### Best Practices
Under the current implementation, the Oden API does not rate limit requests from clients.

However, rate limiting will be introduced in the near future and it is recommended that users design their API
clients to not exceed a request rate of **one per second**.

### Schema
All v2 API access is over HTTPS and accessed from https://api.oden.app/v2
All data is sent and received as JSON.

API requests with duplicate keys will process only the data for the first key detected and ignore the rest, so it's not recommended. Batching multiple messages in this way is currently not possible.
  - Example of duplicate key in JSON: {\"raw_data\":{\"scrap\":\"10\",\"scrap\":\"100\"}}

All timestamps are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format:

  `YYYY-MM-DDTHH:MM:SSZ`

All durations are returned in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Times) format with the largest unit of time being the hour:

   `PT[n]H[n]M[n]S`

All timestamps sent to the API in POST requests should be in ISO 8601 format.
### HTTP Verbs
The ONLY HTTP call type (sometimes called *verb* or *method*) used within Oden's API is **POST**.
There are three actions supported via a **POST**; call, search, set, and delete, together supporting CRUD operations;
  - **search** requests are used to search for and *read* objects in the Oden Platform
      - All Oden Objects may be uniquely identified by some combination of, or a single, parameter.
        - Ex a `line` my be identified by either:
          - `id`
          - `name` AND `factory`
  - **set** requests are used to *create* or *update* objects
  - **delete** requests are used to *delete* objects. If a delete endpoint is not yet implemented for a given object, users may choose to update the values of a specific entity to null or 0 values.

### URI Components
All endpoints may be accessed with the URI pattern:
`https://api.oden.app/v2/{object}/{action}`

Where:
  - `object` is the name of the object being requested:
       - `factory`, `quality_test`, `interval`, `line`, etc...
  - `action` is the name of the action being requested
    - `search` , `set` , `delete`

e.g. `https://api.oden.app/v2/factory/search`

# Authentication
Clients can authenticate through the v2 API using a Token provided by Oden. Tokens are opaque strings similar to
[Bearer](https://swagger.io/docs/specification/authentication/bearer-authentication/) tokens that the client must
pass in the [HTTP Authorization request header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) in every request.
The syntax is as follows:

`Authorization: <type> <credentials>`

Where \\<type\\> is \"Token\" and \\<credentials\\> is the Token string. For example:

`Authorization: Token tokenStringProvidedByOden`

Authenticating with an invalid Token will return `401 Unauthorized Error`.

Authenticating with a Token that is not authorized to read requested data will return `403 Forbidden Error`.

Some endpoints may require requests to be broken out by machinegroup (i.e., line or factory) and the number of
requests would scale accordingly. This multiplicity should be taken into consideration when deciding on the
frequency the API client makes requests to the Oden endpoints.

To authenticate in this [UI](https://api.oden.app/v2/ui/), click the Lock icon, and copy/paste the token into the Authorize box.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Generator version: 7.15.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://oden.io/contact/](https://oden.io/contact/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import oden "github.com/odenio/oden-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `oden.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), oden.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `oden.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), oden.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `oden.ContextOperationServerIndices` and `oden.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), oden.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), oden.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.oden.app*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IntervalsAPI* | [**V2IntervalDeletePost**](docs/IntervalsAPI.md#v2intervaldeletepost) | **Post** /v2/interval/delete | 
*IntervalsAPI* | [**V2IntervalSearchPost**](docs/IntervalsAPI.md#v2intervalsearchpost) | **Post** /v2/interval/search | 
*IntervalsAPI* | [**V2IntervalSetPost**](docs/IntervalsAPI.md#v2intervalsetpost) | **Post** /v2/interval/set | 
*IntervalsAPI* | [**V2IntervalTypeSearchPost**](docs/IntervalsAPI.md#v2intervaltypesearchpost) | **Post** /v2/interval_type/search | 
*IntervalsAPI* | [**V2IntervalUpdatePost**](docs/IntervalsAPI.md#v2intervalupdatepost) | **Post** /v2/interval/update | 
*IntervalsAPI* | [**V2IntervalsDeletePost**](docs/IntervalsAPI.md#v2intervalsdeletepost) | **Post** /v2/intervals/delete | 
*IntervalsAPI* | [**V2IntervalsSetPost**](docs/IntervalsAPI.md#v2intervalssetpost) | **Post** /v2/intervals/set | 
*IntervalsAPI* | [**V2IntervalsUpdatePost**](docs/IntervalsAPI.md#v2intervalsupdatepost) | **Post** /v2/intervals/update | 
*MachineGroupsAPI* | [**V2FactorySearchPost**](docs/MachineGroupsAPI.md#v2factorysearchpost) | **Post** /v2/factory/search | 
*MachineGroupsAPI* | [**V2LineSearchPost**](docs/MachineGroupsAPI.md#v2linesearchpost) | **Post** /v2/line/search | 
*MetricGroupsAPI* | [**V2MetricGroupSearchPost**](docs/MetricGroupsAPI.md#v2metricgroupsearchpost) | **Post** /v2/metric_group/search | 
*OQLAPI* | [**V2OqlQueryPost**](docs/OQLAPI.md#v2oqlquerypost) | **Post** /v2/oql/query | 
*ProductAttributesAPI* | [**V2ProductAttributeSearchPost**](docs/ProductAttributesAPI.md#v2productattributesearchpost) | **Post** /v2/product_attribute/search | 
*ProductAttributesAPI* | [**V2ProductAttributeSetPost**](docs/ProductAttributesAPI.md#v2productattributesetpost) | **Post** /v2/product_attribute/set | 
*ProductMappingsAPI* | [**V2ProductMappingSearchPost**](docs/ProductMappingsAPI.md#v2productmappingsearchpost) | **Post** /v2/product_mapping/search | 
*ProductMappingsAPI* | [**V2ProductMappingSetPost**](docs/ProductMappingsAPI.md#v2productmappingsetpost) | **Post** /v2/product_mapping/set | 
*ProductsAPI* | [**V2ProductDeletePost**](docs/ProductsAPI.md#v2productdeletepost) | **Post** /v2/product/delete | 
*ProductsAPI* | [**V2ProductSearchPost**](docs/ProductsAPI.md#v2productsearchpost) | **Post** /v2/product/search | 
*ProductsAPI* | [**V2ProductSetPost**](docs/ProductsAPI.md#v2productsetpost) | **Post** /v2/product/set | 
*QualityTestAPI* | [**V2QualitySchemaSearchPost**](docs/QualityTestAPI.md#v2qualityschemasearchpost) | **Post** /v2/quality_schema/search | 
*QualityTestAPI* | [**V2QualityTestDeletePost**](docs/QualityTestAPI.md#v2qualitytestdeletepost) | **Post** /v2/quality_test/delete | 
*QualityTestAPI* | [**V2QualityTestSearchPost**](docs/QualityTestAPI.md#v2qualitytestsearchpost) | **Post** /v2/quality_test/search | 
*QualityTestAPI* | [**V2QualityTestSetPost**](docs/QualityTestAPI.md#v2qualitytestsetpost) | **Post** /v2/quality_test/set | 
*QualityTestAPI* | [**V2QualityTestsDeletePost**](docs/QualityTestAPI.md#v2qualitytestsdeletepost) | **Post** /v2/quality_tests/delete | 
*ScrapYieldDataAPI* | [**V2ScrapYieldDeletePost**](docs/ScrapYieldDataAPI.md#v2scrapyielddeletepost) | **Post** /v2/scrap_yield/delete | 
*ScrapYieldDataAPI* | [**V2ScrapYieldSearchPost**](docs/ScrapYieldDataAPI.md#v2scrapyieldsearchpost) | **Post** /v2/scrap_yield/search | 
*ScrapYieldDataAPI* | [**V2ScrapYieldSetPost**](docs/ScrapYieldDataAPI.md#v2scrapyieldsetpost) | **Post** /v2/scrap_yield/set | 
*TargetsAPI* | [**V2TargetSearchPost**](docs/TargetsAPI.md#v2targetsearchpost) | **Post** /v2/target/search | 
*TargetsAPI* | [**V2TargetSetPost**](docs/TargetsAPI.md#v2targetsetpost) | **Post** /v2/target/set | 


## Documentation For Models

 - [BatchMetadata](docs/BatchMetadata.md)
 - [CustomMetadata](docs/CustomMetadata.md)
 - [Factory](docs/Factory.md)
 - [GenericError](docs/GenericError.md)
 - [Interval](docs/Interval.md)
 - [IntervalBulkCreate](docs/IntervalBulkCreate.md)
 - [IntervalBulkDelete](docs/IntervalBulkDelete.md)
 - [IntervalBulkUpdate](docs/IntervalBulkUpdate.md)
 - [IntervalMetadata](docs/IntervalMetadata.md)
 - [IntervalType](docs/IntervalType.md)
 - [Line](docs/Line.md)
 - [Match](docs/Match.md)
 - [MetricGroup](docs/MetricGroup.md)
 - [OQLQuery](docs/OQLQuery.md)
 - [Product](docs/Product.md)
 - [ProductAttribute](docs/ProductAttribute.md)
 - [ProductMapping](docs/ProductMapping.md)
 - [QualitySchema](docs/QualitySchema.md)
 - [QualityTest](docs/QualityTest.md)
 - [RunMetadata](docs/RunMetadata.md)
 - [ScrapYieldData](docs/ScrapYieldData.md)
 - [ScrapYieldSchema](docs/ScrapYieldSchema.md)
 - [StateCategory](docs/StateCategory.md)
 - [StateMetadata](docs/StateMetadata.md)
 - [StateReason](docs/StateReason.md)
 - [Target](docs/Target.md)
 - [Unit](docs/Unit.md)
 - [V2IntervalsDeletePost200Response](docs/V2IntervalsDeletePost200Response.md)
 - [V2IntervalsUpdatePost200Response](docs/V2IntervalsUpdatePost200Response.md)
 - [V2IntervalsUpdatePost200ResponseFailedIntervalsInner](docs/V2IntervalsUpdatePost200ResponseFailedIntervalsInner.md)
 - [V2LineSearchPost400Response](docs/V2LineSearchPost400Response.md)
 - [V2LineSearchPost409Response](docs/V2LineSearchPost409Response.md)
 - [V2LineSearchPost500Response](docs/V2LineSearchPost500Response.md)
 - [V2QualityTestsDeletePostRequest](docs/V2QualityTestsDeletePostRequest.md)
 - [V2ScrapYieldSearchPostRequest](docs/V2ScrapYieldSearchPostRequest.md)
 - [V2ScrapYieldSetPostRequest](docs/V2ScrapYieldSetPostRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### APIKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APIKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		oden.ContextAPIKeys,
		map[string]oden.APIKey{
			"APIKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@oden.io

