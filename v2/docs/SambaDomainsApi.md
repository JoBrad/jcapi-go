# \SambaDomainsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LdapserversSambaDomainsDelete**](SambaDomainsApi.md#LdapserversSambaDomainsDelete) | **Delete** /ldapservers/{ldapserver_id}/sambadomains/{id} | Delete Samba Domain
[**LdapserversSambaDomainsGet**](SambaDomainsApi.md#LdapserversSambaDomainsGet) | **Get** /ldapservers/{ldapserver_id}/sambadomains/{id} | Get Samba Domain
[**LdapserversSambaDomainsList**](SambaDomainsApi.md#LdapserversSambaDomainsList) | **Get** /ldapservers/{ldapserver_id}/sambadomains | List Samba Domains
[**LdapserversSambaDomainsPost**](SambaDomainsApi.md#LdapserversSambaDomainsPost) | **Post** /ldapservers/{ldapserver_id}/sambadomains | Create Samba Domain
[**LdapserversSambaDomainsPut**](SambaDomainsApi.md#LdapserversSambaDomainsPut) | **Put** /ldapservers/{ldapserver_id}/sambadomains/{id} | Update Samba Domain


# **LdapserversSambaDomainsDelete**
> string LdapserversSambaDomainsDelete(ctx, ldapserverId, id, optional)
Delete Samba Domain

This endpoint allows you to delete a samba domain from an LDAP server.  ##### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsGet**
> SambaDomain LdapserversSambaDomainsGet(ctx, ldapserverId, id, optional)
Get Samba Domain

This endpoint returns a specific samba domain for an LDAP server.  ##### Sample Request ``` curl -X GET \\   https://console.jumpcloud.com/api/v2/ldapservers/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**SambaDomain**](samba-domain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsList**
> []SambaDomain LdapserversSambaDomainsList(ctx, ldapserverId, optional)
List Samba Domains

This endpoint returns all samba domains for an LDAP server.  ##### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]SambaDomain**](samba-domain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPost**
> SambaDomain LdapserversSambaDomainsPost(ctx, ldapserverId, optional)
Create Samba Domain

This endpoint allows you to create a samba domain for an LDAP server.  ##### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"sid\":\"{SID_ID}\",     \"name\":\"{WORKGROUP_NAME}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **body** | [**SambaDomain**](SambaDomain.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**SambaDomain**](samba-domain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPut**
> SambaDomain LdapserversSambaDomainsPut(ctx, ldapserverId, id, optional)
Update Samba Domain

This endpoint allows you to update the samba domain information for an LDAP server.  ##### Sample Request ``` curl -X PUT https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"sid\":\"{SID_ID}\",     \"name\":\"{WORKGROUP_NAME}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **body** | [**SambaDomain**](SambaDomain.md)|  | 

### Return type

[**SambaDomain**](samba-domain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

