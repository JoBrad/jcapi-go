# \Office365Api

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDelete**](Office365Api.md#DomainsDelete) | **Delete** /office365s/{resourceObjectId}/domains/{objectId} | Delete domain
[**DomainsInsert**](Office365Api.md#DomainsInsert) | **Post** /office365s/{resourceObjectId}/domains | Insert domain
[**DomainsList**](Office365Api.md#DomainsList) | **Get** /office365s/{resourceObjectId}/domains | List domains
[**GraphOffice365AssociationsList**](Office365Api.md#GraphOffice365AssociationsList) | **Get** /office365s/{office365_id}/associations | List the associations of an Office 365 instance
[**GraphOffice365AssociationsPost**](Office365Api.md#GraphOffice365AssociationsPost) | **Post** /office365s/{office365_id}/associations | Manage the associations of an Office 365 instance
[**GraphOffice365TraverseUser**](Office365Api.md#GraphOffice365TraverseUser) | **Get** /office365s/{office365_id}/users | List the Users bound to an Office 365 instance
[**GraphOffice365TraverseUserGroup**](Office365Api.md#GraphOffice365TraverseUserGroup) | **Get** /office365s/{office365_id}/usergroups | List the User Groups bound to an Office 365 instance
[**Office365sGet**](Office365Api.md#Office365sGet) | **Get** /office365s/{office365_id} | Get Office 365 instance
[**Office365sListImportUsers**](Office365Api.md#Office365sListImportUsers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance
[**Office365sPatch**](Office365Api.md#Office365sPatch) | **Patch** /office365s/{office365_id} | Update existing Office 365 instance.
[**TranslationRulesOffice365Delete**](Office365Api.md#TranslationRulesOffice365Delete) | **Delete** /office365s/{office365_id}/translationrules/{id} | Deletes a Office 365 translation rule
[**TranslationRulesOffice365Get**](Office365Api.md#TranslationRulesOffice365Get) | **Get** /office365s/{office365_id}/translationrules/{id} | Gets a specific Office 365 translation rule
[**TranslationRulesOffice365List**](Office365Api.md#TranslationRulesOffice365List) | **Get** /office365s/{office365_id}/translationrules | List all the Office 365 Translation Rules
[**TranslationRulesOffice365Post**](Office365Api.md#TranslationRulesOffice365Post) | **Post** /office365s/{office365_id}/translationrules | Create a new Office 365 Translation Rule


# **DomainsDelete**
> O365DomainResponse DomainsDelete(ctx, resourceObjectId, objectId)
Delete domain

Delete the domain for the specified o365 integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resourceObjectId** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

[**O365DomainResponse**](o365.DomainResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsInsert**
> O365DomainResponse DomainsInsert(ctx, resourceObjectId, body)
Insert domain

Insert the domain for the specified o365 integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resourceObjectId** | **string**|  | 
  **body** | [**Body6**](Body6.md)|  | 

### Return type

[**O365DomainResponse**](o365.DomainResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsList**
> O365DomainsListResponse DomainsList(ctx, resourceObjectId, optional)
List domains

list the domains configured for an o365 integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resourceObjectId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceObjectId** | **string**|  | 
 **limit** | **string**| The number of records to return at once. Limited to 30. | [default to 30]
 **skip** | **string**| The offset into the records to return. | [default to 0]

### Return type

[**O365DomainsListResponse**](o365.DomainsListResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphOffice365AssociationsList**
> []GraphConnection GraphOffice365AssociationsList(ctx, office365Id, targets, optional)
List the associations of an Office 365 instance

This endpoint returns _direct_ associations of an Office 365 instance.   A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.  #### Sample Request ``` curl -X GET 'https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations?targets=user_group' \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 instance. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;office_365\&quot; can be associated to. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **targets** | [**[]string**](string.md)| Targets which a \&quot;office_365\&quot; can be associated to. | 
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

# **GraphOffice365AssociationsPost**
> GraphOffice365AssociationsPost(ctx, office365Id, optional)
Manage the associations of an Office 365 instance

This endpoint allows you to manage the _direct_ associations of a Office 365 instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user_group\",     \"id\": \"{Group_ID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **body** | [**GraphOperationOffice365**](GraphOperationOffice365.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphOffice365TraverseUser**
> []GraphObjectWithPaths GraphOffice365TraverseUser(ctx, office365Id, optional)
List the Users bound to an Office 365 instance

This endpoint will return all Users bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Office 365 instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Office 365 instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 suite. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 suite. | 
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

# **GraphOffice365TraverseUserGroup**
> []GraphObjectWithPaths GraphOffice365TraverseUserGroup(ctx, office365Id, optional)
List the User Groups bound to an Office 365 instance

This endpoint will return all Users Groups bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Office 365 instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Office 365 instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID/usergroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 suite. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 suite. | 
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

# **Office365sGet**
> Office365 Office365sGet(ctx, office365Id, optional)
Get Office 365 instance

This endpoint returns a specific Office 365 instance.  #####  Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**Office365**](office365.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **Office365sPatch**
> Office365 Office365sPatch(ctx, office365Id, optional)
Update existing Office 365 instance.

This endpoint allows updating some attributes of an Office 365 instance.  #####  Sample Request  ``` curl -X PATCH https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"userLockoutAction\": \"maintain\",     \"userPasswordExpirationAction\": \"suspend\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365Id** | **string**| ObjectID of the Office 365 instance. | 
 **body** | [**Office365**](Office365.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**Office365**](office365.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesOffice365Delete**
> TranslationRulesOffice365Delete(ctx, office365Id, id)
Deletes a Office 365 translation rule

This endpoint allows you to delete a translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.  #### Sample Request  ``` curl -X DELETE https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules/{id} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesOffice365Get**
> Office365TranslationRule TranslationRulesOffice365Get(ctx, office365Id, id)
Gets a specific Office 365 translation rule

This endpoint returns a specific translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.  ###### Sample Request  ```   curl -X GET https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules/{id} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **office365Id** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesOffice365List**
> []Office365TranslationRule TranslationRulesOffice365List(ctx, office365Id, optional)
List all the Office 365 Translation Rules

This endpoint returns all translation rules for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.  ##### Sample Request  ```  curl -X GET  https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

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
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**[]Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesOffice365Post**
> Office365TranslationRule TranslationRulesOffice365Post(ctx, office365Id, optional)
Create a new Office 365 Translation Rule

This endpoint allows you to create a translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.  ##### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     {Translation Rule Parameters}   }' ```

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
 **body** | [**Office365TranslationRuleRequest**](Office365TranslationRuleRequest.md)|  | 

### Return type

[**Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

