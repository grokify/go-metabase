# \DatabaseApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatabaseMetadata**](DatabaseApi.md#DatabaseMetadata) | **Get** /api/database/{databaseId}/metadata | Get metadata about a Database.
[**GetDatabase**](DatabaseApi.md#GetDatabase) | **Get** /api/database/{databaseId} | Get a Database
[**ListDatabases**](DatabaseApi.md#ListDatabases) | **Get** /api/database | List Databases



## DatabaseMetadata

> Database DatabaseMetadata(ctx, databaseId).IncludeHidden(includeHidden).IncludeEditableDataModel(includeEditableDataModel).Execute()

Get metadata about a Database.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/grokify/go-metabase"
)

func main() {
    databaseId := int32(56) // int32 | The database ID
    includeHidden := true // bool | show hidden tables and fields. (optional) (default to false)
    includeEditableDataModel := true // bool | return tables for which the current user has data model editing permissions. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.DatabaseMetadata(context.Background(), databaseId).IncludeHidden(includeHidden).IncludeEditableDataModel(includeEditableDataModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.DatabaseMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabaseMetadata`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.DatabaseMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | The database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabaseMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHidden** | **bool** | show hidden tables and fields. | [default to false]
 **includeEditableDataModel** | **bool** | return tables for which the current user has data model editing permissions. | [default to false]

### Return type

[**Database**](Database.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> Database GetDatabase(ctx, databaseId).Include(include).IncludeEditableDataModel(includeEditableDataModel).ExcludeUneditableDetails(excludeUneditableDetails).Execute()

Get a Database



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/grokify/go-metabase"
)

func main() {
    databaseId := int32(56) // int32 | The database ID
    include := "include_example" // string | value may be nil, or if non-nil, value must be a valid boolean string ('true' or 'false'). (optional) (default to "tables")
    includeEditableDataModel := true // bool | will only include DBs for which the current user has data model editing permissions. (If include=tables, this also applies to the list of tables in each DB). Should only be used if Enterprise Edition code is available the advanced-permissions feature is enabled. (optional) (default to false)
    excludeUneditableDetails := true // bool | will only include DBs for which the current user can edit the DB details. Has no effect unless Enterprise Edition code is available and the advanced-permissions feature is enabled. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.GetDatabase(context.Background(), databaseId).Include(include).IncludeEditableDataModel(includeEditableDataModel).ExcludeUneditableDetails(excludeUneditableDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.GetDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **int32** | The database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | [default to &quot;tables&quot;]
 **includeEditableDataModel** | **bool** | will only include DBs for which the current user has data model editing permissions. (If include&#x3D;tables, this also applies to the list of tables in each DB). Should only be used if Enterprise Edition code is available the advanced-permissions feature is enabled. | [default to false]
 **excludeUneditableDetails** | **bool** | will only include DBs for which the current user can edit the DB details. Has no effect unless Enterprise Edition code is available and the advanced-permissions feature is enabled. | [default to false]

### Return type

[**Database**](Database.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> DatabaseList ListDatabases(ctx).Include(include).Saved(saved).IncludeEditableDataModel(includeEditableDataModel).ExcludeUneditableDetails(excludeUneditableDetails).Execute()

List Databases



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/grokify/go-metabase"
)

func main() {
    include := "include_example" // string | value may be nil, or if non-nil, value must be a valid boolean string ('true' or 'false'). (optional) (default to "tables")
    saved := true // bool | value may be nil, or if non-nil, value must be a valid boolean string ('true' or 'false'). (optional) (default to false)
    includeEditableDataModel := true // bool | will only include DBs for which the current user has data model editing permissions. (If include=tables, this also applies to the list of tables in each DB). Should only be used if Enterprise Edition code is available the advanced-permissions feature is enabled. (optional) (default to false)
    excludeUneditableDetails := true // bool | will only include DBs for which the current user can edit the DB details. Has no effect unless Enterprise Edition code is available and the advanced-permissions feature is enabled. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.ListDatabases(context.Background()).Include(include).Saved(saved).IncludeEditableDataModel(includeEditableDataModel).ExcludeUneditableDetails(excludeUneditableDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: DatabaseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | [default to &quot;tables&quot;]
 **saved** | **bool** | value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | [default to false]
 **includeEditableDataModel** | **bool** | will only include DBs for which the current user has data model editing permissions. (If include&#x3D;tables, this also applies to the list of tables in each DB). Should only be used if Enterprise Edition code is available the advanced-permissions feature is enabled. | [default to false]
 **excludeUneditableDetails** | **bool** | will only include DBs for which the current user can edit the DB details. Has no effect unless Enterprise Edition code is available and the advanced-permissions feature is enabled. | [default to false]

### Return type

[**DatabaseList**](DatabaseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

