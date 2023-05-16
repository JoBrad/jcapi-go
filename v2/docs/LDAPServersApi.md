# \LDAPServersApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphLdapServerAssociationsList**](LDAPServersApi.md#GraphLdapServerAssociationsList) | **Get** /ldapservers/{ldapserver_id}/associations | List the associations of a LDAP Server
[**GraphLdapServerAssociationsPost**](LDAPServersApi.md#GraphLdapServerAssociationsPost) | **Post** /ldapservers/{ldapserver_id}/associations | Manage the associations of a LDAP Server
[**GraphLdapServerTraverseUser**](LDAPServersApi.md#GraphLdapServerTraverseUser) | **Get** /ldapservers/{ldapserver_id}/users | List the Users bound to a LDAP Server
[**GraphLdapServerTraverseUserGroup**](LDAPServersApi.md#GraphLdapServerTraverseUserGroup) | **Get** /ldapservers/{ldapserver_id}/usergroups | List the User Groups bound to a LDAP Server
[**LdapserversGet**](LDAPServersApi.md#LdapserversGet) | **Get** /ldapservers/{id} | Get LDAP Server
[**LdapserversList**](LDAPServersApi.md#LdapserversList) | **Get** /ldapservers | List LDAP Servers
[**LdapserversPatch**](LDAPServersApi.md#LdapserversPatch) | **Patch** /ldapservers/{id} | Update existing LDAP server


# **GraphLdapServerAssociationsList**
> []GraphConnection GraphLdapServerAssociationsList(ctx, ldapserverId, targets, optional)
List the associations of a LDAP Server

This endpoint returns the _direct_ associations of this LDAP Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.  #### Sample Request  ```  curl -X GET 'https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations?targets=user_group \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;ldap_server\&quot; can be associated to. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
 **targets** | [**[]string**](string.md)| Targets which a \&quot;ldap_server\&quot; can be associated to. | 
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

# **GraphLdapServerAssociationsPost**
> GraphLdapServerAssociationsPost(ctx, ldapserverId, optional)
Manage the associations of a LDAP Server

This endpoint allows you to manage the _direct_ associations of a LDAP Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user\",     \"id\": \"{User_ID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
 **body** | [**GraphOperationLdapServer**](GraphOperationLdapServer.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphLdapServerTraverseUser**
> []GraphObjectWithPaths GraphLdapServerTraverseUser(ctx, ldapserverId, optional)
List the Users bound to a LDAP Server

This endpoint will return all Users bound to an LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this LDAP server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this LDAP server instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
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

# **GraphLdapServerTraverseUserGroup**
> []GraphObjectWithPaths GraphLdapServerTraverseUserGroup(ctx, ldapserverId, optional)
List the User Groups bound to a LDAP Server

This endpoint will return all Users Groups bound to a LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this LDAP server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this LDAP server instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/usergroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| ObjectID of the LDAP Server. | 
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

# **LdapserversGet**
> LdapServer LdapserversGet(ctx, id, optional)
Get LDAP Server

This endpoint returns a specific LDAP server.  ##### Sample Request  ```  curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the LDAP server. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**LdapServer**](ldap-server.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversList**
> []LdapServer LdapserversList(ctx, optional)
List LDAP Servers

This endpoint returns the object IDs of your LDAP servers.   ##### Sample Request  ```   curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]LdapServer**](ldap-server.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversPatch**
> InlineResponse20011 LdapserversPatch(ctx, id, optional)
Update existing LDAP server

This endpoint allows updating some attributes of an LDAP server.  Sample Request  ``` curl -X PATCH https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"userLockoutAction\": \"remove\",     \"userPasswordExpirationAction\": \"disable\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the LDAP server. | 
 **body** | [**Body5**](Body5.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

