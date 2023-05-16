# \ApplicationsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsDeleteLogo**](ApplicationsApi.md#ApplicationsDeleteLogo) | **Delete** /applications/{application_id}/logo | Delete application image
[**ApplicationsGet**](ApplicationsApi.md#ApplicationsGet) | **Get** /applications/{application_id} | Get an Application
[**ApplicationsPostLogo**](ApplicationsApi.md#ApplicationsPostLogo) | **Post** /applications/{application_id}/logo | Save application logo
[**GraphApplicationAssociationsList**](ApplicationsApi.md#GraphApplicationAssociationsList) | **Get** /applications/{application_id}/associations | List the associations of an Application
[**GraphApplicationAssociationsPost**](ApplicationsApi.md#GraphApplicationAssociationsPost) | **Post** /applications/{application_id}/associations | Manage the associations of an Application
[**GraphApplicationTraverseUser**](ApplicationsApi.md#GraphApplicationTraverseUser) | **Get** /applications/{application_id}/users | List the Users bound to an Application
[**GraphApplicationTraverseUserGroup**](ApplicationsApi.md#GraphApplicationTraverseUserGroup) | **Get** /applications/{application_id}/usergroups | List the User Groups bound to an Application
[**ImportCreate**](ApplicationsApi.md#ImportCreate) | **Post** /applications/{application_id}/import/jobs | Create an import job
[**ImportUsers**](ApplicationsApi.md#ImportUsers) | **Get** /applications/{application_id}/import/users | Get a list of users to import from an Application IdM service provider


# **ApplicationsDeleteLogo**
> ApplicationsDeleteLogo(ctx, applicationId, optional)
Delete application image

Deletes the specified image from an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsGet**
> interface{} ApplicationsGet(ctx, applicationId, optional)
Get an Application

The endpoint retrieves an Application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationsPostLogo**
> ApplicationsPostLogo(ctx, applicationId, optional)
Save application logo

This endpoint sets the logo for an application.  #### Sample Request ``` curl -X POST 'https://console.jumpcloud.com/api/v2/applications/{Application_ID}/logo \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**|  | 
 **image** | ***os.File**| The file to upload. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphApplicationAssociationsList**
> []GraphConnection GraphApplicationAssociationsList(ctx, applicationId, targets, optional)
List the associations of an Application

This endpoint will return the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Applications and User Groups.   #### Sample Request ``` curl -X GET 'https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations?targets=user_group \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;application\&quot; can be associated to. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **targets** | [**[]string**](string.md)| Targets which a \&quot;application\&quot; can be associated to. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphApplicationAssociationsPost**
> GraphApplicationAssociationsPost(ctx, applicationId, optional)
Manage the associations of an Application

This endpoint allows you to manage the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Application and User Groups.  #### Sample Request ``` curl -X POST 'https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations' \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user_group\",     \"id\": \"{Group_ID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **body** | [**GraphOperationApplication**](GraphOperationApplication.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphApplicationTraverseUser**
> []GraphObjectWithPaths GraphApplicationTraverseUser(ctx, applicationId, optional)
List the Users bound to an Application

This endpoint will return all Users bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Application to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Application.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/users \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphApplicationTraverseUserGroup**
> []GraphObjectWithPaths GraphApplicationTraverseUserGroup(ctx, applicationId, optional)
List the User Groups bound to an Application

This endpoint will return all Users Groups bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates  each path from this Application to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Application.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/usergroups \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportCreate**
> JobId ImportCreate(ctx, applicationId, optional)
Create an import job

This endpoint allows you to create a user import job that will import new users and/or update existing users in JumpCloud from the application.  The endpoint can currently only be used for applications that have an active Identity Management custom API integration.  The request will  fail with a “Not found” error for applications if that type of integration is not configured.  To learn more about configuring this type of integration, read [Import users from an external identity source using a custom API integration](https://support.jumpcloud.com/support/s/article/Import-users-from-a-custom-rest-API-integration). #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/applications/{application_id}/import/jobs \\ -H 'Accept: application/json' \\ -H 'Content-Type: application/json' \\ -H 'x-api-key: {API_KEY}' \\ -H 'x-org-id: {ORG_ID}' \\ -d '{     \"allowUserReactivation\": true,     \"operations\": [         \"users.create\",         \"users.update\"     ]     \"queryString\": \"location=Chicago&department=IT\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **body** | [**ImportUsersRequest**](ImportUsersRequest.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**JobId**](job-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportUsers**
> ImportUsersResponse ImportUsers(ctx, applicationId, optional)
Get a list of users to import from an Application IdM service provider

Get a list of users to import from an Application IdM service provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string**| ObjectID of the Application. | 
 **filter** | **string**| Filter users by a search term | [default to ]
 **query** | **string**| URL query to merge with the service provider request | [default to ]
 **sort** | **string**| Sort users by supported fields | [default to ]
 **sortOrder** | **string**|  | [default to asc]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**ImportUsersResponse**](importUsersResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

