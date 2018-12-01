# \DatabasesApi

All URIs are relative to *http://example.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabases**](DatabasesApi.md#ListDatabases) | **Get** /api/database | List Databases


# **ListDatabases**
> []Database ListDatabases(ctx, optional)
List Databases

Fetch all Databases. include_tables means we should hydrate the Tables belonging to each DB. include_cards here means we should also include virtual Table entries for saved Questions, e.g. so we can easily use them as source Tables in queries. Default for both is false.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDatabasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDatabasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeTables** | **optional.Bool**| value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | 
 **includeCards** | **optional.Bool**| value may be nil, or if non-nil, value must be a valid boolean string (&#39;true&#39; or &#39;false&#39;). | 

### Return type

[**[]Database**](Database.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

