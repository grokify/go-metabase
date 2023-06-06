# \DatabaseApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabases**](DatabaseApi.md#ListDatabases) | **Get** /api/database | List Databases



## ListDatabases

> []Database ListDatabases(ctx).IncludeTables(includeTables).IncludeCards(includeCards).Execute()

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
    includeTables := true // bool | value may be nil, or if non-nil, value must be a valid boolean string ('true' or 'false'). (optional) (default to false)
    includeCards := true // bool | value may be nil, or if non-nil, value must be a valid boolean string ('true' or 'false'). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.ListDatabases(context.Background()).IncludeTables(includeTables).IncludeCards(includeCards).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: []Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeTables** | **bool** | value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | [default to false]
 **includeCards** | **bool** | value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | [default to false]

### Return type

[**[]Database**](Database.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

