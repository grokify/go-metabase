# \DatasetApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryDatabase**](DatasetApi.md#QueryDatabase) | **Post** /api/dataset | Execute a query



## QueryDatabase

> DatasetQueryResults QueryDatabase(ctx).DatasetQueryJsonQuery(datasetQueryJsonQuery).Execute()

Execute a query



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
    datasetQueryJsonQuery := *openapiclient.NewDatasetQueryJsonQuery() // DatasetQueryJsonQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetApi.QueryDatabase(context.Background()).DatasetQueryJsonQuery(datasetQueryJsonQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetApi.QueryDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryDatabase`: DatasetQueryResults
    fmt.Fprintf(os.Stdout, "Response from `DatasetApi.QueryDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetQueryJsonQuery** | [**DatasetQueryJsonQuery**](DatasetQueryJsonQuery.md) |  | 

### Return type

[**DatasetQueryResults**](DatasetQueryResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

