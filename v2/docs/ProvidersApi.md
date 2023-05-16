# \ProvidersApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutotaskCreateConfiguration**](ProvidersApi.md#AutotaskCreateConfiguration) | **Post** /providers/{provider_id}/integrations/autotask | Creates a new Autotask integration for the provider
[**AutotaskDeleteConfiguration**](ProvidersApi.md#AutotaskDeleteConfiguration) | **Delete** /integrations/autotask/{UUID} | Delete Autotask Integration
[**AutotaskGetConfiguration**](ProvidersApi.md#AutotaskGetConfiguration) | **Get** /integrations/autotask/{UUID} | Retrieve Autotask Integration Configuration
[**AutotaskPatchMappings**](ProvidersApi.md#AutotaskPatchMappings) | **Patch** /integrations/autotask/{UUID}/mappings | Create, edit, and/or delete Autotask Mappings
[**AutotaskPatchSettings**](ProvidersApi.md#AutotaskPatchSettings) | **Patch** /integrations/autotask/{UUID}/settings | Create, edit, and/or delete Autotask Integration settings
[**AutotaskRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#AutotaskRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration/options | Get all Autotask ticketing alert configuration options for a provider
[**AutotaskRetrieveAllAlertConfigurations**](ProvidersApi.md#AutotaskRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration | Get all Autotask ticketing alert configurations for a provider
[**AutotaskRetrieveCompanies**](ProvidersApi.md#AutotaskRetrieveCompanies) | **Get** /integrations/autotask/{UUID}/companies | Retrieve Autotask Companies
[**AutotaskRetrieveCompanyTypes**](ProvidersApi.md#AutotaskRetrieveCompanyTypes) | **Get** /integrations/autotask/{UUID}/companytypes | Retrieve Autotask Company Types
[**AutotaskRetrieveContracts**](ProvidersApi.md#AutotaskRetrieveContracts) | **Get** /integrations/autotask/{UUID}/contracts | Retrieve Autotask Contracts
[**AutotaskRetrieveContractsFields**](ProvidersApi.md#AutotaskRetrieveContractsFields) | **Get** /integrations/autotask/{UUID}/contracts/fields | Retrieve Autotask Contract Fields
[**AutotaskRetrieveMappings**](ProvidersApi.md#AutotaskRetrieveMappings) | **Get** /integrations/autotask/{UUID}/mappings | Retrieve Autotask mappings
[**AutotaskRetrieveServices**](ProvidersApi.md#AutotaskRetrieveServices) | **Get** /integrations/autotask/{UUID}/contracts/services | Retrieve Autotask Contract Services
[**AutotaskRetrieveSettings**](ProvidersApi.md#AutotaskRetrieveSettings) | **Get** /integrations/autotask/{UUID}/settings | Retrieve Autotask Integration settings
[**AutotaskUpdateAlertConfiguration**](ProvidersApi.md#AutotaskUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/autotask/alerts/{alert_UUID}/configuration | Update an Autotask ticketing alert&#39;s configuration
[**AutotaskUpdateConfiguration**](ProvidersApi.md#AutotaskUpdateConfiguration) | **Patch** /integrations/autotask/{UUID} | Update Autotask Integration configuration
[**ConnectwiseCreateConfiguration**](ProvidersApi.md#ConnectwiseCreateConfiguration) | **Post** /providers/{provider_id}/integrations/connectwise | Creates a new ConnectWise integration for the provider
[**ConnectwiseDeleteConfiguration**](ProvidersApi.md#ConnectwiseDeleteConfiguration) | **Delete** /integrations/connectwise/{UUID} | Delete ConnectWise Integration
[**ConnectwiseGetConfiguration**](ProvidersApi.md#ConnectwiseGetConfiguration) | **Get** /integrations/connectwise/{UUID} | Retrieve ConnectWise Integration Configuration
[**ConnectwisePatchMappings**](ProvidersApi.md#ConnectwisePatchMappings) | **Patch** /integrations/connectwise/{UUID}/mappings | Create, edit, and/or delete ConnectWise Mappings
[**ConnectwisePatchSettings**](ProvidersApi.md#ConnectwisePatchSettings) | **Patch** /integrations/connectwise/{UUID}/settings | Create, edit, and/or delete ConnectWise Integration settings
[**ConnectwiseRetrieveAdditions**](ProvidersApi.md#ConnectwiseRetrieveAdditions) | **Get** /integrations/connectwise/{UUID}/agreements/{agreement_ID}/additions | Retrieve ConnectWise Additions
[**ConnectwiseRetrieveAgreements**](ProvidersApi.md#ConnectwiseRetrieveAgreements) | **Get** /integrations/connectwise/{UUID}/agreements | Retrieve ConnectWise Agreements
[**ConnectwiseRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#ConnectwiseRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration/options | Get all ConnectWise ticketing alert configuration options for a provider
[**ConnectwiseRetrieveAllAlertConfigurations**](ProvidersApi.md#ConnectwiseRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration | Get all ConnectWise ticketing alert configurations for a provider
[**ConnectwiseRetrieveCompanies**](ProvidersApi.md#ConnectwiseRetrieveCompanies) | **Get** /integrations/connectwise/{UUID}/companies | Retrieve ConnectWise Companies
[**ConnectwiseRetrieveCompanyTypes**](ProvidersApi.md#ConnectwiseRetrieveCompanyTypes) | **Get** /integrations/connectwise/{UUID}/companytypes | Retrieve ConnectWise Company Types
[**ConnectwiseRetrieveMappings**](ProvidersApi.md#ConnectwiseRetrieveMappings) | **Get** /integrations/connectwise/{UUID}/mappings | Retrieve ConnectWise mappings
[**ConnectwiseRetrieveSettings**](ProvidersApi.md#ConnectwiseRetrieveSettings) | **Get** /integrations/connectwise/{UUID}/settings | Retrieve ConnectWise Integration settings
[**ConnectwiseUpdateAlertConfiguration**](ProvidersApi.md#ConnectwiseUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/connectwise/alerts/{alert_UUID}/configuration | Update a ConnectWise ticketing alert&#39;s configuration
[**ConnectwiseUpdateConfiguration**](ProvidersApi.md#ConnectwiseUpdateConfiguration) | **Patch** /integrations/connectwise/{UUID} | Update ConnectWise Integration configuration
[**MtpIntegrationRetrieveAlerts**](ProvidersApi.md#MtpIntegrationRetrieveAlerts) | **Get** /providers/{provider_id}/integrations/ticketing/alerts | Get all ticketing alerts available for a provider&#39;s ticketing integration.
[**MtpIntegrationRetrieveSyncErrors**](ProvidersApi.md#MtpIntegrationRetrieveSyncErrors) | **Get** /integrations/{integration_type}/{UUID}/errors | Retrieve Recent Integration Sync Errors
[**PolicyGroupTemplatesDelete**](ProvidersApi.md#PolicyGroupTemplatesDelete) | **Delete** /providers/{provider_id}/policygrouptemplates/{id} | Deletes policy group template.
[**PolicyGroupTemplatesGet**](ProvidersApi.md#PolicyGroupTemplatesGet) | **Get** /providers/{provider_id}/policygrouptemplates/{id} | Gets a provider&#39;s policy group template.
[**PolicyGroupTemplatesGetConfiguredPolicyTemplate**](ProvidersApi.md#PolicyGroupTemplatesGetConfiguredPolicyTemplate) | **Get** /providers/{provider_id}/configuredpolicytemplates/{id} | Retrieve a configured policy template by id.
[**PolicyGroupTemplatesList**](ProvidersApi.md#PolicyGroupTemplatesList) | **Get** /providers/{provider_id}/policygrouptemplates | List a provider&#39;s policy group templates.
[**PolicyGroupTemplatesListConfiguredPolicyTemplates**](ProvidersApi.md#PolicyGroupTemplatesListConfiguredPolicyTemplates) | **Get** /providers/{provider_id}/configuredpolicytemplates | List a provider&#39;s configured policy templates.
[**PolicyGroupTemplatesListMembers**](ProvidersApi.md#PolicyGroupTemplatesListMembers) | **Get** /providers/{provider_id}/policygrouptemplates/{id}/members | Gets the list of members from a policy group template.
[**ProviderOrganizationsCreateOrg**](ProvidersApi.md#ProviderOrganizationsCreateOrg) | **Post** /providers/{provider_id}/organizations | Create Provider Organization
[**ProviderOrganizationsUpdateOrg**](ProvidersApi.md#ProviderOrganizationsUpdateOrg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
[**ProvidersGetProvider**](ProvidersApi.md#ProvidersGetProvider) | **Get** /providers/{provider_id} | Retrieve Provider
[**ProvidersListAdministrators**](ProvidersApi.md#ProvidersListAdministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
[**ProvidersListOrganizations**](ProvidersApi.md#ProvidersListOrganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
[**ProvidersPostAdmins**](ProvidersApi.md#ProvidersPostAdmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
[**ProvidersRemoveAdministrator**](ProvidersApi.md#ProvidersRemoveAdministrator) | **Delete** /providers/{provider_id}/administrators/{id} | Delete Provider Administrator
[**ProvidersRetrieveIntegrations**](ProvidersApi.md#ProvidersRetrieveIntegrations) | **Get** /providers/{provider_id}/integrations | Retrieve Integrations for Provider
[**ProvidersRetrieveInvoice**](ProvidersApi.md#ProvidersRetrieveInvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#39;s invoice.
[**ProvidersRetrieveInvoices**](ProvidersApi.md#ProvidersRetrieveInvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#39;s invoices.
[**SyncroCreateConfiguration**](ProvidersApi.md#SyncroCreateConfiguration) | **Post** /providers/{provider_id}/integrations/syncro | Creates a new Syncro integration for the provider
[**SyncroDeleteConfiguration**](ProvidersApi.md#SyncroDeleteConfiguration) | **Delete** /integrations/syncro/{UUID} | Delete Syncro Integration
[**SyncroGetConfiguration**](ProvidersApi.md#SyncroGetConfiguration) | **Get** /integrations/syncro/{UUID} | Retrieve Syncro Integration Configuration
[**SyncroPatchMappings**](ProvidersApi.md#SyncroPatchMappings) | **Patch** /integrations/syncro/{UUID}/mappings | Create, edit, and/or delete Syncro Mappings
[**SyncroPatchSettings**](ProvidersApi.md#SyncroPatchSettings) | **Patch** /integrations/syncro/{UUID}/settings | Create, edit, and/or delete Syncro Integration settings
[**SyncroRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#SyncroRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/syncro/alerts/configuration/options | Get all Syncro ticketing alert configuration options for a provider
[**SyncroRetrieveAllAlertConfigurations**](ProvidersApi.md#SyncroRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/syncro/alerts/configuration | Get all Syncro ticketing alert configurations for a provider
[**SyncroRetrieveBillingMappingConfigurationOptions**](ProvidersApi.md#SyncroRetrieveBillingMappingConfigurationOptions) | **Get** /integrations/syncro/{UUID}/billing_mapping_configuration_options | Retrieve Syncro billing mappings dependencies
[**SyncroRetrieveCompanies**](ProvidersApi.md#SyncroRetrieveCompanies) | **Get** /integrations/syncro/{UUID}/companies | Retrieve Syncro Companies
[**SyncroRetrieveMappings**](ProvidersApi.md#SyncroRetrieveMappings) | **Get** /integrations/syncro/{UUID}/mappings | Retrieve Syncro mappings
[**SyncroRetrieveSettings**](ProvidersApi.md#SyncroRetrieveSettings) | **Get** /integrations/syncro/{UUID}/settings | Retrieve Syncro Integration settings
[**SyncroUpdateAlertConfiguration**](ProvidersApi.md#SyncroUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/syncro/alerts/{alert_UUID}/configuration | Update a Syncro ticketing alert&#39;s configuration
[**SyncroUpdateConfiguration**](ProvidersApi.md#SyncroUpdateConfiguration) | **Patch** /integrations/syncro/{UUID} | Update Syncro Integration configuration


# **AutotaskCreateConfiguration**
> InlineResponse201 AutotaskCreateConfiguration(ctx, providerId, optional)
Creates a new Autotask integration for the provider

Creates a new Autotask integration for the provider. You must be associated with the provider to use this route. A 422 Unprocessable Entity response means the server failed to validate with Autotask.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **body** | [**AutotaskIntegrationReq**](AutotaskIntegrationReq.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskDeleteConfiguration**
> AutotaskDeleteConfiguration(ctx, uUID)
Delete Autotask Integration

Removes a Autotask integration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskGetConfiguration**
> AutotaskIntegration AutotaskGetConfiguration(ctx, uUID)
Retrieve Autotask Integration Configuration

Retrieves configuration for given Autotask integration id. You must be associated to the provider the integration is tied to in order to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**AutotaskIntegration**](AutotaskIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskPatchMappings**
> AutotaskMappingResponse AutotaskPatchMappings(ctx, uUID, optional)
Create, edit, and/or delete Autotask Mappings

Create, edit, and/or delete mappings between Jumpcloud organizations and Autotask companies/contracts/services. You must be associated to the same provider as the Autotask integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**AutotaskMappingRequest**](AutotaskMappingRequest.md)|  | 

### Return type

[**AutotaskMappingResponse**](AutotaskMappingResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskPatchSettings**
> AutotaskSettings AutotaskPatchSettings(ctx, uUID, optional)
Create, edit, and/or delete Autotask Integration settings

Create, edit, and/or delete Autotask settings. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**AutotaskSettingsPatchReq**](AutotaskSettingsPatchReq.md)|  | 

### Return type

[**AutotaskSettings**](AutotaskSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveAllAlertConfigurationOptions**
> AutotaskTicketingAlertConfigurationOptions AutotaskRetrieveAllAlertConfigurationOptions(ctx, providerId)
Get all Autotask ticketing alert configuration options for a provider

Get all Autotask ticketing alert configuration options for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**AutotaskTicketingAlertConfigurationOptions**](AutotaskTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveAllAlertConfigurations**
> AutotaskTicketingAlertConfigurationList AutotaskRetrieveAllAlertConfigurations(ctx, providerId)
Get all Autotask ticketing alert configurations for a provider

Get all Autotask ticketing alert configurations for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**AutotaskTicketingAlertConfigurationList**](AutotaskTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveCompanies**
> AutotaskCompanyResp AutotaskRetrieveCompanies(ctx, uUID, optional)
Retrieve Autotask Companies

Retrieves a list of Autotask companies for the given Autotask id. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**AutotaskCompanyResp**](AutotaskCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveCompanyTypes**
> AutotaskCompanyTypeResp AutotaskRetrieveCompanyTypes(ctx, uUID)
Retrieve Autotask Company Types

Retrieves a list of user defined company types from Autotask for the given Autotask id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**AutotaskCompanyTypeResp**](AutotaskCompanyTypeResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveContracts**
> InlineResponse2003 AutotaskRetrieveContracts(ctx, uUID, optional)
Retrieve Autotask Contracts

Retrieves a list of Autotask contracts for the given Autotask integration id. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveContractsFields**
> InlineResponse2004 AutotaskRetrieveContractsFields(ctx, uUID)
Retrieve Autotask Contract Fields

Retrieves a list of Autotask contract fields for the given Autotask integration id. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveMappings**
> InlineResponse2006 AutotaskRetrieveMappings(ctx, uUID, optional)
Retrieve Autotask mappings

Retrieves the list of mappings for this Autotask integration. You must be associated to the same provider as the Autotask integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveServices**
> InlineResponse2005 AutotaskRetrieveServices(ctx, uUID, optional)
Retrieve Autotask Contract Services

Retrieves a list of Autotask contract services for the given Autotask integration id. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskRetrieveSettings**
> AutotaskSettings AutotaskRetrieveSettings(ctx, uUID)
Retrieve Autotask Integration settings

Retrieve the Autotask integration settings. You must be associated to the same provider as the Autotask integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**AutotaskSettings**](AutotaskSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskUpdateAlertConfiguration**
> AutotaskTicketingAlertConfiguration AutotaskUpdateAlertConfiguration(ctx, providerId, alertUUID, optional)
Update an Autotask ticketing alert's configuration

Update an Autotask ticketing alert's configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **alertUUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **alertUUID** | **string**|  | 
 **body** | [**AutotaskTicketingAlertConfigurationRequest**](AutotaskTicketingAlertConfigurationRequest.md)|  | 

### Return type

[**AutotaskTicketingAlertConfiguration**](AutotaskTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutotaskUpdateConfiguration**
> AutotaskIntegration AutotaskUpdateConfiguration(ctx, uUID, optional)
Update Autotask Integration configuration

Update the Autotask integration configuration. A 422 Unprocessable Entity response means the server failed to validate with Autotask.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**AutotaskIntegrationPatchReq**](AutotaskIntegrationPatchReq.md)|  | 

### Return type

[**AutotaskIntegration**](AutotaskIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseCreateConfiguration**
> InlineResponse201 ConnectwiseCreateConfiguration(ctx, providerId, optional)
Creates a new ConnectWise integration for the provider

Creates a new ConnectWise integration for the provider. You must be associated with the provider to use this route. A 422 Unprocessable Entity response means the server failed to validate with ConnectWise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **body** | [**ConnectwiseIntegrationReq**](ConnectwiseIntegrationReq.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseDeleteConfiguration**
> ConnectwiseDeleteConfiguration(ctx, uUID)
Delete ConnectWise Integration

Removes a ConnectWise integration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseGetConfiguration**
> ConnectwiseIntegration ConnectwiseGetConfiguration(ctx, uUID)
Retrieve ConnectWise Integration Configuration

Retrieves configuration for given ConnectWise integration id. You must be associated to the provider the integration is tied to in order to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**ConnectwiseIntegration**](ConnectwiseIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwisePatchMappings**
> ConnectWiseMappingRequest ConnectwisePatchMappings(ctx, uUID, optional)
Create, edit, and/or delete ConnectWise Mappings

Create, edit, and/or delete mappings between Jumpcloud organizations and ConnectWise companies/agreements/additions. You must be associated to the same provider as the ConnectWise integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**ConnectWiseMappingRequest**](ConnectWiseMappingRequest.md)|  | 

### Return type

[**ConnectWiseMappingRequest**](ConnectWiseMappingRequest.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwisePatchSettings**
> ConnectWiseSettings ConnectwisePatchSettings(ctx, uUID, optional)
Create, edit, and/or delete ConnectWise Integration settings

Create, edit, and/or delete ConnectWiseIntegration settings. You must be associated to the same provider as the ConnectWise integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**ConnectWiseSettingsPatchReq**](ConnectWiseSettingsPatchReq.md)|  | 

### Return type

[**ConnectWiseSettings**](ConnectWiseSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveAdditions**
> InlineResponse2008 ConnectwiseRetrieveAdditions(ctx, uUID, agreementID, optional)
Retrieve ConnectWise Additions

Retrieves a list of ConnectWise additions for the given ConnectWise id and Agreement id. You must be associated to the same provider as the ConnectWise integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
  **agreementID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **agreementID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveAgreements**
> InlineResponse2007 ConnectwiseRetrieveAgreements(ctx, uUID, optional)
Retrieve ConnectWise Agreements

Retrieves a list of ConnectWise agreements for the given ConnectWise id. You must be associated to the same provider as the ConnectWise integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveAllAlertConfigurationOptions**
> ConnectWiseTicketingAlertConfigurationOptions ConnectwiseRetrieveAllAlertConfigurationOptions(ctx, providerId)
Get all ConnectWise ticketing alert configuration options for a provider

Get all ConnectWise ticketing alert configuration options for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**ConnectWiseTicketingAlertConfigurationOptions**](ConnectWiseTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveAllAlertConfigurations**
> ConnectWiseTicketingAlertConfigurationList ConnectwiseRetrieveAllAlertConfigurations(ctx, providerId)
Get all ConnectWise ticketing alert configurations for a provider

Get all ConnectWise ticketing alert configurations for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**ConnectWiseTicketingAlertConfigurationList**](ConnectWiseTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveCompanies**
> ConnectwiseCompanyResp ConnectwiseRetrieveCompanies(ctx, uUID, optional)
Retrieve ConnectWise Companies

Retrieves a list of ConnectWise companies for the given ConnectWise id. You must be associated to the same provider as the ConnectWise integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**ConnectwiseCompanyResp**](ConnectwiseCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveCompanyTypes**
> ConnectwiseCompanyTypeResp ConnectwiseRetrieveCompanyTypes(ctx, uUID)
Retrieve ConnectWise Company Types

Retrieves a list of user defined company types from ConnectWise for the given ConnectWise id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**ConnectwiseCompanyTypeResp**](ConnectwiseCompanyTypeResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveMappings**
> InlineResponse2009 ConnectwiseRetrieveMappings(ctx, uUID, optional)
Retrieve ConnectWise mappings

Retrieves the list of mappings for this ConnectWise integration. You must be associated to the same provider as the ConnectWise integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseRetrieveSettings**
> ConnectWiseSettings ConnectwiseRetrieveSettings(ctx, uUID)
Retrieve ConnectWise Integration settings

Retrieve the ConnectWise integration settings. You must be associated to the same provider as the ConnectWise integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**ConnectWiseSettings**](ConnectWiseSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseUpdateAlertConfiguration**
> ConnectWiseTicketingAlertConfiguration ConnectwiseUpdateAlertConfiguration(ctx, providerId, alertUUID, optional)
Update a ConnectWise ticketing alert's configuration

Update a ConnectWise ticketing alert's configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **alertUUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **alertUUID** | **string**|  | 
 **body** | [**ConnectWiseTicketingAlertConfigurationRequest**](ConnectWiseTicketingAlertConfigurationRequest.md)|  | 

### Return type

[**ConnectWiseTicketingAlertConfiguration**](ConnectWiseTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectwiseUpdateConfiguration**
> ConnectwiseIntegration ConnectwiseUpdateConfiguration(ctx, uUID, optional)
Update ConnectWise Integration configuration

Update the ConnectWise integration configuration. A 422 Unprocessable Entity response means the server failed to validate with ConnectWise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**ConnectwiseIntegrationPatchReq**](ConnectwiseIntegrationPatchReq.md)|  | 

### Return type

[**ConnectwiseIntegration**](ConnectwiseIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MtpIntegrationRetrieveAlerts**
> TicketingIntegrationAlertsResp MtpIntegrationRetrieveAlerts(ctx, providerId)
Get all ticketing alerts available for a provider's ticketing integration.

Get all ticketing alerts available for a provider's ticketing integration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**TicketingIntegrationAlertsResp**](TicketingIntegrationAlertsResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MtpIntegrationRetrieveSyncErrors**
> IntegrationSyncErrorResp MtpIntegrationRetrieveSyncErrors(ctx, uUID, integrationType)
Retrieve Recent Integration Sync Errors

Retrieves recent sync errors for given integration type and integration id. You must be associated to the provider the integration is tied to in order to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
  **integrationType** | **string**|  | 

### Return type

[**IntegrationSyncErrorResp**](IntegrationSyncErrorResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesDelete**
> PolicyGroupTemplatesDelete(ctx, providerId, id)
Deletes policy group template.

Deletes a Policy Group Template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesGet**
> PolicyGroupTemplate PolicyGroupTemplatesGet(ctx, providerId, id)
Gets a provider's policy group template.

Retrieves a Policy Group Template for this provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**PolicyGroupTemplate**](PolicyGroupTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesGetConfiguredPolicyTemplate**
> ConfiguredPolicyTemplate PolicyGroupTemplatesGetConfiguredPolicyTemplate(ctx, providerId, id)
Retrieve a configured policy template by id.

Retrieves a Configured Policy Templates for this provider and Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**ConfiguredPolicyTemplate**](ConfiguredPolicyTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesList**
> PolicyGroupTemplates PolicyGroupTemplatesList(ctx, providerId, optional)
List a provider's policy group templates.

Retrieves a list of Policy Group Templates for this provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**PolicyGroupTemplates**](PolicyGroupTemplates.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesListConfiguredPolicyTemplates**
> InlineResponse20014 PolicyGroupTemplatesListConfiguredPolicyTemplates(ctx, providerId, optional)
List a provider's configured policy templates.

Retrieves a list of Configured Policy Templates for this provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyGroupTemplatesListMembers**
> PolicyGroupTemplateMembers PolicyGroupTemplatesListMembers(ctx, providerId, id, optional)
Gets the list of members from a policy group template.

Retrieves a Policy Group Template's Members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **id** | **string**|  | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**PolicyGroupTemplateMembers**](PolicyGroupTemplateMembers.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProviderOrganizationsCreateOrg**
> Organization ProviderOrganizationsCreateOrg(ctx, providerId, optional)
Create Provider Organization

This endpoint creates a new organization under the provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **body** | [**CreateOrganization**](CreateOrganization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProviderOrganizationsUpdateOrg**
> Organization ProviderOrganizationsUpdateOrg(ctx, providerId, id, optional)
Update Provider Organization

This endpoint updates a provider's organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **id** | **string**|  | 
 **body** | [**Organization**](Organization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersGetProvider**
> Provider ProvidersGetProvider(ctx, providerId, optional)
Retrieve Provider

This endpoint returns details about a provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersListAdministrators**
> InlineResponse20013 ProvidersListAdministrators(ctx, providerId, optional)
List Provider Administrators

This endpoint returns a list of the Administrators associated with the Provider. You must be associated with the provider to use this route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **sortIgnoreCase** | [**[]string**](string.md)| The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersListOrganizations**
> InlineResponse20015 ProvidersListOrganizations(ctx, providerId, optional)
List Provider Organizations

This endpoint returns a list of the Organizations associated with the Provider. You must be associated with the provider to use this route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **sortIgnoreCase** | [**[]string**](string.md)| The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersPostAdmins**
> Administrator ProvidersPostAdmins(ctx, providerId, optional)
Create a new Provider Administrator

This endpoint allows you to create a provider administrator. You must be associated with the provider to use this route. You must provide either `role` or `roleName`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **body** | [**ProviderAdminReq**](ProviderAdminReq.md)|  | 

### Return type

[**Administrator**](Administrator.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRemoveAdministrator**
> ProvidersRemoveAdministrator(ctx, providerId, id)
Delete Provider Administrator

This endpoint removes an Administrator associated with the Provider. You must be associated with the provider to use this route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRetrieveIntegrations**
> IntegrationsResponse ProvidersRetrieveIntegrations(ctx, providerId, optional)
Retrieve Integrations for Provider

Retrieves a list of integrations this provider has configured. You must be associated to the provider to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**IntegrationsResponse**](IntegrationsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRetrieveInvoice**
> *os.File ProvidersRetrieveInvoice(ctx, providerId, iD)
Download a provider's invoice.

Retrieves an invoice for this provider. You must be associated to the provider to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **iD** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRetrieveInvoices**
> ProviderInvoiceResponse ProvidersRetrieveInvoices(ctx, providerId, optional)
List a provider's invoices.

Retrieves a list of invoices for this provider. You must be associated to the provider to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]

### Return type

[**ProviderInvoiceResponse**](ProviderInvoiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroCreateConfiguration**
> InlineResponse201 SyncroCreateConfiguration(ctx, providerId, optional)
Creates a new Syncro integration for the provider

Creates a new Syncro integration for the provider. You must be associated with the provider to use this route. A 422 Unprocessable Entity response means the server failed to validate with Syncro.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **body** | [**SyncroIntegrationReq**](SyncroIntegrationReq.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroDeleteConfiguration**
> SyncroDeleteConfiguration(ctx, uUID)
Delete Syncro Integration

Removes a Syncro integration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroGetConfiguration**
> SyncroIntegration SyncroGetConfiguration(ctx, uUID)
Retrieve Syncro Integration Configuration

Retrieves configuration for given Syncro integration id. You must be associated to the provider the integration is tied to in order to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**SyncroIntegration**](SyncroIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroPatchMappings**
> SyncroMappingRequest SyncroPatchMappings(ctx, uUID, optional)
Create, edit, and/or delete Syncro Mappings

Create, edit, and/or delete mappings between Jumpcloud organizations and Syncro companies. You must be associated to the same provider as the Syncro integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**SyncroMappingRequest**](SyncroMappingRequest.md)|  | 

### Return type

[**SyncroMappingRequest**](SyncroMappingRequest.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroPatchSettings**
> SyncroSettings SyncroPatchSettings(ctx, uUID, optional)
Create, edit, and/or delete Syncro Integration settings

Create, edit, and/or delete SyncroIntegration settings. You must be associated to the same provider as the Syncro integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**SyncroSettingsPatchReq**](SyncroSettingsPatchReq.md)|  | 

### Return type

[**SyncroSettings**](SyncroSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveAllAlertConfigurationOptions**
> SyncroTicketingAlertConfigurationOptions SyncroRetrieveAllAlertConfigurationOptions(ctx, providerId)
Get all Syncro ticketing alert configuration options for a provider

Get all Syncro ticketing alert configuration options for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**SyncroTicketingAlertConfigurationOptions**](SyncroTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveAllAlertConfigurations**
> SyncroTicketingAlertConfigurationList SyncroRetrieveAllAlertConfigurations(ctx, providerId)
Get all Syncro ticketing alert configurations for a provider

Get all Syncro ticketing alert configurations for a provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 

### Return type

[**SyncroTicketingAlertConfigurationList**](SyncroTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveBillingMappingConfigurationOptions**
> SyncroBillingMappingConfigurationOptionsResp SyncroRetrieveBillingMappingConfigurationOptions(ctx, uUID, optional)
Retrieve Syncro billing mappings dependencies

Retrieves a list of dependencies for Syncro billing mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**SyncroBillingMappingConfigurationOptionsResp**](SyncroBillingMappingConfigurationOptionsResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveCompanies**
> SyncroCompanyResp SyncroRetrieveCompanies(ctx, uUID, optional)
Retrieve Syncro Companies

Retrieves a list of Syncro companies for the given Syncro id. You must be associated to the same provider as the Syncro integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**SyncroCompanyResp**](SyncroCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveMappings**
> InlineResponse20010 SyncroRetrieveMappings(ctx, uUID, optional)
Retrieve Syncro mappings

Retrieves the list of mappings for this Syncro integration. You must be associated to the same provider as the Syncro integration to use this api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroRetrieveSettings**
> SyncroSettings SyncroRetrieveSettings(ctx, uUID)
Retrieve Syncro Integration settings

Retrieve the Syncro integration settings. You must be associated to the same provider as the Syncro integration to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 

### Return type

[**SyncroSettings**](SyncroSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroUpdateAlertConfiguration**
> SyncroTicketingAlertConfiguration SyncroUpdateAlertConfiguration(ctx, providerId, alertUUID, optional)
Update a Syncro ticketing alert's configuration

Update a Syncro ticketing alert's configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **alertUUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **alertUUID** | **string**|  | 
 **body** | [**SyncroTicketingAlertConfigurationRequest**](SyncroTicketingAlertConfigurationRequest.md)|  | 

### Return type

[**SyncroTicketingAlertConfiguration**](SyncroTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncroUpdateConfiguration**
> SyncroIntegration SyncroUpdateConfiguration(ctx, uUID, optional)
Update Syncro Integration configuration

Update the Syncro integration configuration. A 422 Unprocessable Entity response means the server failed to validate with Syncro.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uUID** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uUID** | **string**|  | 
 **body** | [**SyncroIntegrationPatchReq**](SyncroIntegrationPatchReq.md)|  | 

### Return type

[**SyncroIntegration**](SyncroIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

