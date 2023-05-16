# \PasswordManagerApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceServiceGetDevice**](PasswordManagerApi.md#DeviceServiceGetDevice) | **Get** /passwordmanager/devices/{UUID} | 
[**DeviceServiceListDevices**](PasswordManagerApi.md#DeviceServiceListDevices) | **Get** /passwordmanager/devices | 


# **DeviceServiceGetDevice**
> DevicePackageV1Device DeviceServiceGetDevice(ctx, uUID)


Get Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**DevicePackageV1Device**](device_package.v1.Device.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeviceServiceListDevices**
> DevicePackageV1ListDevicesResponse DeviceServiceListDevices(ctx, optional)


List Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | 
 **skip** | **int32**|  | 
 **sort** | **string**|  | 
 **fields** | [**[]string**](string.md)|  | 
 **filter** | [**[]string**](string.md)|  | 

### Return type

[**DevicePackageV1ListDevicesResponse**](device_package.v1.ListDevicesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

