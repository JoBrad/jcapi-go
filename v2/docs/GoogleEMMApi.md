# \GoogleEMMApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesEraseDevice**](GoogleEMMApi.md#DevicesEraseDevice) | **Post** /google-emm/devices/{deviceId}/erase-device | Erase the Android Device
[**DevicesGetDevice**](GoogleEMMApi.md#DevicesGetDevice) | **Get** /google-emm/devices/{deviceId} | Get device
[**DevicesGetDeviceAndroidPolicy**](GoogleEMMApi.md#DevicesGetDeviceAndroidPolicy) | **Get** /google-emm/devices/{deviceId}/policy_results | Get the policy JSON of a device
[**DevicesLockDevice**](GoogleEMMApi.md#DevicesLockDevice) | **Post** /google-emm/devices/{deviceId}/lock | Lock device
[**DevicesResetPassword**](GoogleEMMApi.md#DevicesResetPassword) | **Post** /google-emm/devices/{deviceId}/resetpassword | Reset Password of a device
[**EnterprisesCreateEnterprise**](GoogleEMMApi.md#EnterprisesCreateEnterprise) | **Post** /google-emm/enterprises | Create a Google Enterprise
[**EnterprisesDeleteEnterprise**](GoogleEMMApi.md#EnterprisesDeleteEnterprise) | **Delete** /google-emm/enterprises/{enterpriseId} | Delete a Google Enterprise
[**EnterprisesGetConnectionStatus**](GoogleEMMApi.md#EnterprisesGetConnectionStatus) | **Get** /google-emm/enterprises/{enterpriseId}/connection-status | Test connection with Google
[**EnterprisesListEnterprises**](GoogleEMMApi.md#EnterprisesListEnterprises) | **Get** /google-emm/enterprises | List Google Enterprises
[**EnterprisesPatchEnterprise**](GoogleEMMApi.md#EnterprisesPatchEnterprise) | **Patch** /google-emm/enterprises/{enterpriseId} | Update a Google Enterprise
[**SignupURLsCreate**](GoogleEMMApi.md#SignupURLsCreate) | **Post** /google-emm/signup-urls | Get a Signup URL to enroll Google enterprise
[**WebTokensCreateWebToken**](GoogleEMMApi.md#WebTokensCreateWebToken) | **Post** /google-emm/web-tokens | Get a web token to render Google Play iFrame


# **DevicesEraseDevice**
> JumpcloudGoogleEmmCommandResponse DevicesEraseDevice(ctx, deviceId, body)
Erase the Android Device

Removes the work profile and all policies from a personal/company-owned Android 8.0+ device. Company owned devices will be relinquished for personal use. Apps and data associated with the personal profile(s) are preserved.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/erase-device \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deviceId** | **string**|  | 
  **body** | [**interface{}**](interface{}.md)|  | 

### Return type

[**JumpcloudGoogleEmmCommandResponse**](jumpcloud.google_emm.CommandResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesGetDevice**
> JumpcloudGoogleEmmDevice DevicesGetDevice(ctx, deviceId)
Get device

Gets a Google EMM enrolled device details.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deviceId** | **string**|  | 

### Return type

[**JumpcloudGoogleEmmDevice**](jumpcloud.google_emm.Device.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesGetDeviceAndroidPolicy**
> JumpcloudGoogleEmmDeviceAndroidPolicy DevicesGetDeviceAndroidPolicy(ctx, deviceId)
Get the policy JSON of a device

Gets an android JSON policy for a Google EMM enrolled device.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/policy_results \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deviceId** | **string**|  | 

### Return type

[**JumpcloudGoogleEmmDeviceAndroidPolicy**](jumpcloud.google_emm.DeviceAndroidPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesLockDevice**
> JumpcloudGoogleEmmCommandResponse DevicesLockDevice(ctx, deviceId, body)
Lock device

Locks a Google EMM enrolled device, as if the lock screen timeout had expired.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/lock \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deviceId** | **string**|  | 
  **body** | [**interface{}**](interface{}.md)|  | 

### Return type

[**JumpcloudGoogleEmmCommandResponse**](jumpcloud.google_emm.CommandResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesResetPassword**
> JumpcloudGoogleEmmCommandResponse DevicesResetPassword(ctx, deviceId, body)
Reset Password of a device

Reset the user's password of a Google EMM enrolled device.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/resetpassword \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ 'new_password' : 'string' }' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deviceId** | **string**|  | 
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**JumpcloudGoogleEmmCommandResponse**](jumpcloud.google_emm.CommandResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterprisesCreateEnterprise**
> JumpcloudGoogleEmmEnterprise EnterprisesCreateEnterprise(ctx, body)
Create a Google Enterprise

Creates a Google EMM enterprise.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/enterprises \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ 'signupUrlName': 'string', 'enrollmentToken': 'string' }' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JumpcloudGoogleEmmCreateEnterpriseRequest**](JumpcloudGoogleEmmCreateEnterpriseRequest.md)|  | 

### Return type

[**JumpcloudGoogleEmmEnterprise**](jumpcloud.google_emm.Enterprise.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterprisesDeleteEnterprise**
> interface{} EnterprisesDeleteEnterprise(ctx, enterpriseId)
Delete a Google Enterprise

Removes a Google EMM enterprise.   Warning: This is a destructive operation and will remove all data associated with Google EMM enterprise from JumpCloud including devices and applications associated with the given enterprise.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/google-emm/devices/{enterpriseId} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enterpriseId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterprisesGetConnectionStatus**
> JumpcloudGoogleEmmConnectionStatus EnterprisesGetConnectionStatus(ctx, enterpriseId)
Test connection with Google

Gives a connection status between JumpCloud and Google.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{enterpriseId}/connection-status \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enterpriseId** | **string**|  | 

### Return type

[**JumpcloudGoogleEmmConnectionStatus**](jumpcloud.google_emm.ConnectionStatus.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterprisesListEnterprises**
> JumpcloudGoogleEmmListEnterprisesResponse EnterprisesListEnterprises(ctx, optional)
List Google Enterprises

Lists all Google EMM enterprises. An empty list indicates that the Organization is not configured with a Google EMM enterprise yet.    Note: Currently only one Google Enterprise per Organization is supported.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/google-emm/enterprises \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryParametersLimit** | **string**|  | 
 **queryParametersSkip** | **string**|  | 
 **queryParametersFields** | [**[]string**](string.md)|  | 

### Return type

[**JumpcloudGoogleEmmListEnterprisesResponse**](jumpcloud.google_emm.ListEnterprisesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterprisesPatchEnterprise**
> JumpcloudGoogleEmmEnterprise EnterprisesPatchEnterprise(ctx, enterpriseId, body)
Update a Google Enterprise

Updates a Google EMM enterprise details.  #### Sample Request ``` curl -X PATCH https://console.jumpcloud.com/api/v2/google-emm/enterprises \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ 'allowDeviceEnrollment': true, 'deviceGroupId': 'string' }' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enterpriseId** | **string**|  | 
  **body** | [**Body4**](Body4.md)|  | 

### Return type

[**JumpcloudGoogleEmmEnterprise**](jumpcloud.google_emm.Enterprise.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignupURLsCreate**
> JumpcloudGoogleEmmSignupUrl SignupURLsCreate(ctx, )
Get a Signup URL to enroll Google enterprise

Creates a Google EMM enterprise signup URL.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/signup-urls \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JumpcloudGoogleEmmSignupUrl**](jumpcloud.google_emm.SignupURL.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebTokensCreateWebToken**
> JumpcloudGoogleEmmWebToken WebTokensCreateWebToken(ctx, body)
Get a web token to render Google Play iFrame

Creates a web token to access an embeddable managed Google Play web UI for a given Google EMM enterprise.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/google-emm/web-tokens \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JumpcloudGoogleEmmCreateWebTokenRequest**](JumpcloudGoogleEmmCreateWebTokenRequest.md)|  | 

### Return type

[**JumpcloudGoogleEmmWebToken**](jumpcloud.google_emm.WebToken.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

