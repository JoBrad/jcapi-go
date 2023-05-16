# \CommandResultsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandsListResultsByWorkflow**](CommandResultsApi.md#CommandsListResultsByWorkflow) | **Get** /commandresult/workflows | List all Command Results by Workflow


# **CommandsListResultsByWorkflow**
> CommandResultList CommandsListResultsByWorkflow(ctx, optional)
List all Command Results by Workflow

This endpoint returns all command results, grouped by workflowInstanceId.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/commandresult/workflows \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key:{API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **limit** | **int32**|  | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**CommandResultList**](CommandResultList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

