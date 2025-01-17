# \Office365ImportApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Office365sListImportUsers**](Office365ImportApi.md#Office365sListImportUsers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance


# **Office365sListImportUsers**
> InlineResponse20012 Office365sListImportUsers(ctx, office365Id, optional)
Get a list of users to import from an Office 365 instance

Lists Office 365 users available for import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**|  | 
 **consistencyLevel** | **string**| Defines the consistency header for O365 requests. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#request-headers | 
 **top** | **int32**| Office 365 API maximum number of results per page. See https://docs.microsoft.com/en-us/graph/paging. | 
 **skipToken** | **string**| Office 365 API token used to access the next page of results. See https://docs.microsoft.com/en-us/graph/paging. | 
 **filter** | **string**| Office 365 API filter parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **search** | **string**| Office 365 API search parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **orderby** | **string**| Office 365 API orderby parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **count** | **bool**| Office 365 API count parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

