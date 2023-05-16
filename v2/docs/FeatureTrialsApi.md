# \FeatureTrialsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeatureTrialsGetFeatureTrials**](FeatureTrialsApi.md#FeatureTrialsGetFeatureTrials) | **Get** /featureTrials/{feature_code} | Check current feature trial usage for a specific feature


# **FeatureTrialsGetFeatureTrials**
> FeatureTrialData FeatureTrialsGetFeatureTrials(ctx, featureCode)
Check current feature trial usage for a specific feature

This endpoint get's the current state of a feature trial for an org.  #### Sample Request  ```   curl -X GET \\   https://console.jumpcloud.local/api/v2/featureTrials/zeroTrust \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **featureCode** | **string**|  | 

### Return type

[**FeatureTrialData**](FeatureTrialData.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

