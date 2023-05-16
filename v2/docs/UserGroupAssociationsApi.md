# \UserGroupAssociationsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphUserGroupAssociationsList**](UserGroupAssociationsApi.md#GraphUserGroupAssociationsList) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
[**GraphUserGroupAssociationsPost**](UserGroupAssociationsApi.md#GraphUserGroupAssociationsPost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
[**GraphUserGroupTraverseActiveDirectory**](UserGroupAssociationsApi.md#GraphUserGroupTraverseActiveDirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
[**GraphUserGroupTraverseApplication**](UserGroupAssociationsApi.md#GraphUserGroupTraverseApplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
[**GraphUserGroupTraverseDirectory**](UserGroupAssociationsApi.md#GraphUserGroupTraverseDirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
[**GraphUserGroupTraverseGSuite**](UserGroupAssociationsApi.md#GraphUserGroupTraverseGSuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
[**GraphUserGroupTraverseLdapServer**](UserGroupAssociationsApi.md#GraphUserGroupTraverseLdapServer) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
[**GraphUserGroupTraverseOffice365**](UserGroupAssociationsApi.md#GraphUserGroupTraverseOffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
[**GraphUserGroupTraverseRadiusServer**](UserGroupAssociationsApi.md#GraphUserGroupTraverseRadiusServer) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
[**GraphUserGroupTraverseSystem**](UserGroupAssociationsApi.md#GraphUserGroupTraverseSystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
[**GraphUserGroupTraverseSystemGroup**](UserGroupAssociationsApi.md#GraphUserGroupTraverseSystemGroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups


# **GraphUserGroupAssociationsList**
> []GraphConnection GraphUserGroupAssociationsList(ctx, groupId, targets, optional)
List the associations of a User Group.

This endpoint returns the _direct_ associations of this User Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.   #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations?targets=system \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;user_group\&quot; can be associated to. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
 **targets** | [**[]string**](string.md)| Targets which a \&quot;user_group\&quot; can be associated to. | 
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

# **GraphUserGroupAssociationsPost**
> GraphUserGroupAssociationsPost(ctx, groupId, optional)
Manage the associations of a User Group

This endpoint manages the _direct_ associations of this User Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"system\",     \"id\": \"{SystemID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
 **body** | [**GraphOperationUserGroup**](GraphOperationUserGroup.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphUserGroupTraverseActiveDirectory**
> []GraphObjectWithPaths GraphUserGroupTraverseActiveDirectory(ctx, groupId, optional)
List the Active Directories bound to a User Group

This endpoint will return all Active Directory Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding Active Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Active Directory from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/activedirectories \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseApplication**
> []GraphObjectWithPaths GraphUserGroupTraverseApplication(ctx, groupId, optional)
List the Applications bound to a User Group

This endpoint will return all Applications bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/applications \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseDirectory**
> []GraphObjectWithPaths GraphUserGroupTraverseDirectory(ctx, groupId, optional)
List the Directories bound to a User Group

This endpoint will return all Directories bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directories from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/directories \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseGSuite**
> []GraphObjectWithPaths GraphUserGroupTraverseGSuite(ctx, groupId, optional)
List the G Suite instances bound to a User Group

This endpoint will return all G Suite Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID/gsuites \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseLdapServer**
> []GraphObjectWithPaths GraphUserGroupTraverseLdapServer(ctx, groupId, optional)
List the LDAP Servers bound to a User Group

This endpoint will return all LDAP Servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/ldapservers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseOffice365**
> []GraphObjectWithPaths GraphUserGroupTraverseOffice365(ctx, groupId, optional)
List the Office 365 instances bound to a User Group

This endpoint will return all Office 365 instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/office365s \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseRadiusServer**
> []GraphObjectWithPaths GraphUserGroupTraverseRadiusServer(ctx, groupId, optional)
List the RADIUS Servers bound to a User Group

This endpoint will return all RADIUS servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/radiusservers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseSystem**
> []GraphObjectWithPaths GraphUserGroupTraverseSystem(ctx, groupId, optional)
List the Systems bound to a User Group

This endpoint will return all Systems bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

# **GraphUserGroupTraverseSystemGroup**
> []GraphObjectWithPaths GraphUserGroupTraverseSystemGroup(ctx, groupId, optional)
List the System Groups bound to User Groups

This endpoint will return all System Groups bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this User Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this User Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systemgroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the User Group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the User Group. | 
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

