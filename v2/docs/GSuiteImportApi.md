# \GSuiteImportApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GsuitesListImportJumpcloudUsers**](GSuiteImportApi.md#GsuitesListImportJumpcloudUsers) | **Get** /gsuites/{gsuite_id}/import/jumpcloudusers | Get a list of users in Jumpcloud format to import from a Google Workspace account.
[**GsuitesListImportUsers**](GSuiteImportApi.md#GsuitesListImportUsers) | **Get** /gsuites/{gsuite_id}/import/users | Get a list of users to import from a G Suite instance


# **GsuitesListImportJumpcloudUsers**
> InlineResponse2001 GsuitesListImportJumpcloudUsers(ctx, gsuiteId, optional)
Get a list of users in Jumpcloud format to import from a Google Workspace account.

Lists available G Suite users for import, translated to the Jumpcloud user schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**|  | 
 **maxResults** | **int32**| Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **orderBy** | **string**| Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **pageToken** | **string**| Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **query** | **string**| Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. | 
 **sortOrder** | **string**| Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GsuitesListImportUsers**
> InlineResponse2002 GsuitesListImportUsers(ctx, gsuiteId, optional)
Get a list of users to import from a G Suite instance

Lists G Suite users available for import.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **maxResults** | **int32**| Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **orderBy** | **string**| Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **pageToken** | **string**| Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **query** | **string**| Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. | 
 **sortOrder** | **string**| Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

