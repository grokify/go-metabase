# \DatasetApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryDatabase**](DatasetApi.md#QueryDatabase) | **Post** /api/dataset | Execute a query


# **QueryDatabase**
> DatasetQueryResults QueryDatabase(ctx, datasetQueryJsonQuery)
Execute a query

Execute a query and retrieve the results in the usual format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datasetQueryJsonQuery** | [**DatasetQueryJsonQuery**](DatasetQueryJsonQuery.md)|  | 

### Return type

[**DatasetQueryResults**](DatasetQueryResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

