# \CloudInsightsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudInsightsCreateAccount**](CloudInsightsApi.md#CloudInsightsCreateAccount) | **Post** /cloudinsights/accounts | Create AWS account record
[**CloudInsightsCreateSavedView**](CloudInsightsApi.md#CloudInsightsCreateSavedView) | **Post** /cloudinsights/views | 
[**CloudInsightsDeleteSavedView**](CloudInsightsApi.md#CloudInsightsDeleteSavedView) | **Delete** /cloudinsights/views/{view_id} | 
[**CloudInsightsFetchDistinctFieldValues**](CloudInsightsApi.md#CloudInsightsFetchDistinctFieldValues) | **Post** /cloudinsights/events/distinct | 
[**CloudInsightsFetchSavedView**](CloudInsightsApi.md#CloudInsightsFetchSavedView) | **Get** /cloudinsights/views/{view_id} | 
[**CloudInsightsFetchSavedViewsList**](CloudInsightsApi.md#CloudInsightsFetchSavedViewsList) | **Get** /cloudinsights/views | 
[**CloudInsightsGetAccountByID**](CloudInsightsApi.md#CloudInsightsGetAccountByID) | **Get** /cloudinsights/accounts/{cloud_insights_id} | 
[**CloudInsightsListAccounts**](CloudInsightsApi.md#CloudInsightsListAccounts) | **Get** /cloudinsights/accounts | 
[**CloudInsightsListEvents**](CloudInsightsApi.md#CloudInsightsListEvents) | **Post** /cloudinsights/events | 
[**CloudInsightsUpdateAccountByID**](CloudInsightsApi.md#CloudInsightsUpdateAccountByID) | **Put** /cloudinsights/accounts/{cloud_insights_id} | Update Cloud Insights Account record by ID
[**CloudInsightsUpdateSavedView**](CloudInsightsApi.md#CloudInsightsUpdateSavedView) | **Put** /cloudinsights/views/{view_id} | 


# **CloudInsightsCreateAccount**
> CloudInsightsAccountPostResponse CloudInsightsCreateAccount(ctx, optional)
Create AWS account record

Allows you to create a single AWS account record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloudInsightsAccountPost**](CloudInsightsAccountPost.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsAccountPostResponse**](cloud-insights-account-post-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain, text/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsCreateSavedView**
> CloudInsightsSavedViewResponse CloudInsightsCreateSavedView(ctx, optional)


Create new saved view for organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloudInsightsSavedViewPost**](CloudInsightsSavedViewPost.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsSavedViewResponse**](cloud-insights-saved-view-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsDeleteSavedView**
> CloudInsightsSavedViewId CloudInsightsDeleteSavedView(ctx, viewId, optional)


Delete Cloud Insights Saved View record by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsSavedViewId**](cloud-insights-saved-view-id.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsFetchDistinctFieldValues**
> CloudInsightsEventsDistinctResponse CloudInsightsFetchDistinctFieldValues(ctx, optional)


Cloud Insights query to retrieve list of distinct field values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloudInsightsEventsDistinctPost**](CloudInsightsEventsDistinctPost.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsEventsDistinctResponse**](cloud-insights-events-distinct-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsFetchSavedView**
> CloudInsightsSavedViewResponse CloudInsightsFetchSavedView(ctx, viewId, optional)


Get Cloud Insights Saved View By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsSavedViewResponse**](cloud-insights-saved-view-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsFetchSavedViewsList**
> []CloudInsightsSavedViewResponse CloudInsightsFetchSavedViewsList(ctx, optional)


Get Cloud Insights Saved Views List

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

### Return type

[**[]CloudInsightsSavedViewResponse**](cloud-insights-saved-view-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsGetAccountByID**
> CloudInsightsAccountGet CloudInsightsGetAccountByID(ctx, cloudInsightsId, optional)


Get Cloud Insights Account by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudInsightsId** | **string**| objectId of the Cloud Insights instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudInsightsId** | **string**| objectId of the Cloud Insights instance. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsAccountGet**](cloud-insights-account-get.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsListAccounts**
> CloudInsightsAccountListResponse CloudInsightsListAccounts(ctx, optional)


Get list of AWS accounts registered with Cloud Insights

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

### Return type

[**CloudInsightsAccountListResponse**](cloud-insights-account-list-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsListEvents**
> CloudInsightsEventsListResponse CloudInsightsListEvents(ctx, optional)


Get list of events based on query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloudInsightsEventsListPost**](CloudInsightsEventsListPost.md)|  | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsEventsListResponse**](cloud-insights-events-list-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsUpdateAccountByID**
> CloudInsightsAccountGet CloudInsightsUpdateAccountByID(ctx, cloudInsightsId, optional)
Update Cloud Insights Account record by ID

Update record an existing Cloud Insights AWS account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudInsightsId** | **string**| objectId of the Cloud Insights instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudInsightsId** | **string**| objectId of the Cloud Insights instance. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsAccountPut**](CloudInsightsAccountPut.md)|  | 

### Return type

[**CloudInsightsAccountGet**](cloud-insights-account-get.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudInsightsUpdateSavedView**
> CloudInsightsSavedViewResponse CloudInsightsUpdateSavedView(ctx, viewId, optional)


Update Cloud Insights Saved View record by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **string**| objectId of the Cloud Insights Saved View record. | 
 **xOrgId** | **string**| Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsSavedViewPut**](CloudInsightsSavedViewPut.md)|  | 

### Return type

[**CloudInsightsSavedViewResponse**](cloud-insights-saved-view-response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

