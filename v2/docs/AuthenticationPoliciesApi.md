# \AuthenticationPoliciesApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthnpoliciesDelete**](AuthenticationPoliciesApi.md#AuthnpoliciesDelete) | **Delete** /authn/policies/{id} | Delete Authentication Policy
[**AuthnpoliciesGet**](AuthenticationPoliciesApi.md#AuthnpoliciesGet) | **Get** /authn/policies/{id} | Get an authentication policy
[**AuthnpoliciesList**](AuthenticationPoliciesApi.md#AuthnpoliciesList) | **Get** /authn/policies | List Authentication Policies
[**AuthnpoliciesPatch**](AuthenticationPoliciesApi.md#AuthnpoliciesPatch) | **Patch** /authn/policies/{id} | Patch Authentication Policy
[**AuthnpoliciesPost**](AuthenticationPoliciesApi.md#AuthnpoliciesPost) | **Post** /authn/policies | Create an Authentication Policy


# **AuthnpoliciesDelete**
> AuthnPolicy AuthnpoliciesDelete(ctx, id, optional)
Delete Authentication Policy

Delete the specified authentication policy.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/authn/policies/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the authentication policy | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the authentication policy | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AuthnPolicy**](AuthnPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthnpoliciesGet**
> AuthnPolicy AuthnpoliciesGet(ctx, id, optional)
Get an authentication policy

Return a specific authentication policy.  #### Sample Request ``` curl https://console.jumpcloud.com/api/v2/authn/policies/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the authentication policy | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the authentication policy | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AuthnPolicy**](AuthnPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthnpoliciesList**
> []AuthnPolicy AuthnpoliciesList(ctx, optional)
List Authentication Policies

Get a list of all authentication policies.  #### Sample Request ``` curl https://console.jumpcloud.com/api/v2/authn/policies \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

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
 **xTotalCount** | **int32**|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**[]AuthnPolicy**](AuthnPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthnpoliciesPatch**
> AuthnPolicy AuthnpoliciesPatch(ctx, id, optional)
Patch Authentication Policy

Patch the specified authentication policy.  #### Sample Request ``` curl -X PATCH https://console.jumpcloud.com/api/v2/authn/policies/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ \"disabled\": false }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the authentication policy | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the authentication policy | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **body** | [**AuthnPolicy**](AuthnPolicy.md)|  | 

### Return type

[**AuthnPolicy**](AuthnPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthnpoliciesPost**
> AuthnPolicy AuthnpoliciesPost(ctx, optional)
Create an Authentication Policy

Create an authentication policy.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/authn/policies \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"Sample Policy\",     \"disabled\": false,     \"effect\": {       \"action\": \"allow\"     },     \"targets\": {       \"users\": {         \"inclusions\": [\"ALL\"]       },       \"userGroups\": {         \"exclusions\": [{USER_GROUP_ID}]       },       \"resources\": [ {\"type\": \"user_portal\" } ]     },     \"conditions\":{       \"ipAddressIn\": [{IP_LIST_ID}]     }   }' ```

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
 **body** | [**AuthnPolicy**](AuthnPolicy.md)|  | 

### Return type

[**AuthnPolicy**](AuthnPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

