# \BulkJobRequestsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUserExpires**](BulkJobRequestsApi.md#BulkUserExpires) | **Post** /bulk/user/expires | Bulk Expire Users
[**BulkUserStatesCreate**](BulkJobRequestsApi.md#BulkUserStatesCreate) | **Post** /bulk/userstates | Create Scheduled Userstate Job
[**BulkUserStatesDelete**](BulkJobRequestsApi.md#BulkUserStatesDelete) | **Delete** /bulk/userstates/{id} | Delete Scheduled Userstate Job
[**BulkUserStatesGetNextScheduled**](BulkJobRequestsApi.md#BulkUserStatesGetNextScheduled) | **Get** /bulk/userstates/eventlist/next | Get the next scheduled state change for a list of users
[**BulkUserStatesList**](BulkJobRequestsApi.md#BulkUserStatesList) | **Get** /bulk/userstates | List Scheduled Userstate Change Jobs
[**BulkUserUnlocks**](BulkJobRequestsApi.md#BulkUserUnlocks) | **Post** /bulk/user/unlocks | Bulk Unlock Users
[**BulkUsersCreate**](BulkJobRequestsApi.md#BulkUsersCreate) | **Post** /bulk/users | Bulk Users Create
[**BulkUsersCreateResults**](BulkJobRequestsApi.md#BulkUsersCreateResults) | **Get** /bulk/users/{job_id}/results | List Bulk Users Results
[**BulkUsersUpdate**](BulkJobRequestsApi.md#BulkUsersUpdate) | **Patch** /bulk/users | Bulk Users Update


# **BulkUserExpires**
> JobId BulkUserExpires(ctx, optional)
Bulk Expire Users

The endpoint allows you to start a bulk job to asynchronously expire users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]BulkUserExpire**](bulk-user-expire.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**JobId**](job-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUserStatesCreate**
> []ScheduledUserstateResult BulkUserStatesCreate(ctx, optional)
Create Scheduled Userstate Job

This endpoint allows you to create scheduled statechange jobs. #### Sample Request ``` curl -X POST \"https://console.jumpcloud.com/api/v2/bulk/userstates\" \\   -H 'x-api-key: {API_KEY}' \\   -H 'Content-Type: application/json' \\   -H 'Accept: application/json' \\   -d '{     \"user_ids\": [\"{User_ID_1}\", \"{User_ID_2}\", \"{User_ID_3}\"],     \"state\": \"SUSPENDED\",     \"start_date\": \"2000-01-01T00:00:00.000Z\"   }' ```

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
 **body** | [**BulkScheduledStatechangeCreate**](BulkScheduledStatechangeCreate.md)|  | 

### Return type

[**[]ScheduledUserstateResult**](scheduled-userstate-result.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUserStatesDelete**
> BulkUserStatesDelete(ctx, id, optional)
Delete Scheduled Userstate Job

This endpoint deletes a scheduled statechange job. #### Sample Request ``` curl -X DELETE \"https://console.jumpcloud.com/api/v2/bulk/userstates/{ScheduledJob_ID}\" \\   -H 'x-api-key: {API_KEY}' \\   -H 'Content-Type: application/json' \\   -H 'Accept: application/json' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Unique identifier of the scheduled statechange job. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Unique identifier of the scheduled statechange job. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUserStatesGetNextScheduled**
> InlineResponse200 BulkUserStatesGetNextScheduled(ctx, users, optional)
Get the next scheduled state change for a list of users

This endpoint is used to lookup the next upcoming scheduled state change for each user in the given list. The users parameter is limited to 100 items per request. The results are also limited to 100 items. This endpoint returns a max of 1 event per state per user. For example, if a user has 3 ACTIVATED events scheduled it will return the next upcoming activation event. However, if a user also has a SUSPENDED event scheduled along with the ACTIVATED events it will return the next upcoming activation event _and_ the next upcoming suspension event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **users** | [**[]string**](string.md)| A list of system user IDs, limited to 100 items. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | [**[]string**](string.md)| A list of system user IDs, limited to 100 items. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUserStatesList**
> []ScheduledUserstateResult BulkUserStatesList(ctx, optional)
List Scheduled Userstate Change Jobs

The endpoint allows you to list scheduled statechange jobs. #### Sample Request ``` curl -X GET \"https://console.jumpcloud.com/api/v2/bulk/userstates\" \\   -H 'x-api-key: {API_KEY}' \\   -H 'Content-Type: application/json' \\   -H 'Accept: application/json' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **userid** | **string**| The systemuser id to filter by. | 

### Return type

[**[]ScheduledUserstateResult**](scheduled-userstate-result.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUserUnlocks**
> JobId BulkUserUnlocks(ctx, optional)
Bulk Unlock Users

The endpoint allows you to start a bulk job to asynchronously unlock users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]BulkUserUnlock**](bulk-user-unlock.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**JobId**](job-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUsersCreate**
> JobId BulkUsersCreate(ctx, optional)
Bulk Users Create

The endpoint allows you to create a bulk job to asynchronously create users. See [Create a System User](https://docs.jumpcloud.com/api/1.0/index.html#operation/systemusers_post) for the full list of attributes.  #### Default User State The `state` of each user in the request can be explicitly passed in or omitted. If `state` is omitted, then the user will get created using the value returned from the [Get an Organization](https://docs.jumpcloud.com/api/1.0/index.html#operation/organizations_get) endpoint. The default user state for bulk created users depends on the `creation-source` header. For `creation-source:jumpcloud:bulk` the default state is stored in `settings.newSystemUserStateDefaults.csvImport`. For other `creation-source` header values, the default state is stored in `settings.newSystemUserStateDefaults.applicationImport`  These default state values can be changed in the admin portal settings or by using the [Update an Organization](https://docs.jumpcloud.com/api/1.0/index.html#operation/organization_put) endpoint.  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/v2/bulk/users \\ -H 'Accept: application/json' \\ -H 'Content-Type: application/json' \\ -H 'x-api-key: {API_KEY}' \\ -d '[   {     \"email\":\"{email}\",     \"firstname\":\"{firstname}\",     \"lastname\":\"{firstname}\",     \"username\":\"{username}\",     \"attributes\":[       {         \"name\":\"EmployeeID\",         \"value\":\"0000\"       },       {         \"name\":\"Custom\",         \"value\":\"attribute\"       }     ]   } ]' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]BulkUserCreate**](bulk-user-create.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **creationSource** | **string**| Defines the creation-source header for gapps, o365 and workdays requests. If the header isn&#39;t sent, the default value is &#x60;jumpcloud:bulk&#x60;, if you send the header with a malformed value you receive a 400 error.  | [default to jumpcloud:bulk]

### Return type

[**JobId**](job-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUsersCreateResults**
> []JobWorkresult BulkUsersCreateResults(ctx, jobId, optional)
List Bulk Users Results

This endpoint will return the results of particular user import or update job request.  #### Sample Request ``` curl -X GET \\   https://console.jumpcloud.com/api/v2/bulk/users/{ImportJobID}/results \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string**|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]JobWorkresult**](job-workresult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUsersUpdate**
> JobId BulkUsersUpdate(ctx, optional)
Bulk Users Update

The endpoint allows you to create a bulk job to asynchronously update users. See [Update a System User](https://docs.jumpcloud.com/api/1.0/index.html#operation/systemusers_put) for full list of attributes.  #### Sample Request  ``` curl -X PATCH https://console.jumpcloud.com/api/v2/bulk/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '[  {    \"id\":\"5be9fb4ddb01290001e85109\",   \"firstname\":\"{UPDATED_FIRSTNAME}\",   \"department\":\"{UPDATED_DEPARTMENT}\",   \"attributes\":[    {\"name\":\"Custom\",\"value\":\"{ATTRIBUTE_VALUE}\"}   ]  },  {    \"id\":\"5be9fb4ddb01290001e85109\",   \"firstname\":\"{UPDATED_FIRSTNAME}\",   \"costCenter\":\"{UPDATED_COST_CENTER}\",   \"phoneNumbers\":[    {\"type\":\"home\",\"number\":\"{HOME_PHONE_NUMBER}\"},    {\"type\":\"work\",\"number\":\"{WORK_PHONE_NUMBER}\"}   ]  } ] ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]BulkUserUpdate**](bulk-user-update.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**JobId**](job-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

