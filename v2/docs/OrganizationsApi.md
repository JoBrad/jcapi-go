# \OrganizationsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdministratorOrganizationsCreateByAdministrator**](OrganizationsApi.md#AdministratorOrganizationsCreateByAdministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
[**AdministratorOrganizationsListByAdministrator**](OrganizationsApi.md#AdministratorOrganizationsListByAdministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
[**AdministratorOrganizationsListByOrganization**](OrganizationsApi.md#AdministratorOrganizationsListByOrganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
[**AdministratorOrganizationsRemoveByAdministrator**](OrganizationsApi.md#AdministratorOrganizationsRemoveByAdministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
[**OrganizationsListCases**](OrganizationsApi.md#OrganizationsListCases) | **Get** /organizations/cases | Get all cases (Support/Feature requests) for organization


# **AdministratorOrganizationsCreateByAdministrator**
> AdministratorOrganizationLink AdministratorOrganizationsCreateByAdministrator(ctx, id, optional)
Allow Adminstrator access to an Organization.

This endpoint allows you to grant Administrator access to an Organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **body** | [**AdministratorOrganizationLinkReq**](AdministratorOrganizationLinkReq.md)|  | 

### Return type

[**AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsListByAdministrator**
> []AdministratorOrganizationLink AdministratorOrganizationsListByAdministrator(ctx, id, optional)
List the association links between an Administrator and Organizations.

This endpoint returns the association links between an Administrator and Organizations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsListByOrganization**
> []AdministratorOrganizationLink AdministratorOrganizationsListByOrganization(ctx, id, optional)
List the association links between an Organization and Administrators.

This endpoint returns the association links between an Organization and Administrators.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsRemoveByAdministrator**
> AdministratorOrganizationsRemoveByAdministrator(ctx, administratorId, id)
Remove association between an Administrator and an Organization.

This endpoint removes the association link between an Administrator and an Organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **administratorId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsListCases**
> OrganizationCasesResponse OrganizationsListCases(ctx, optional)
Get all cases (Support/Feature requests) for organization

This endpoint returns the cases (Support/Feature requests) for the organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**OrganizationCasesResponse**](OrganizationCasesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

