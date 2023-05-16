/*
 * JumpCloud API
 *
 * # Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 
 *
 * API version: 2.0
 * Contact: support@jumpcloud.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type AppleMDMApiService service


/* AppleMDMApiService Get Apple MDM CSR Plist
 Retrieves an Apple MDM signed CSR Plist for an organization.  The user must supply the returned plist to Apple for signing, and then provide the certificate provided by Apple back into the PUT API.  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/csr \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdmSignedCsrPlist*/
func (a *AppleMDMApiService) ApplemdmsCsrget(ctx context.Context, appleMdmId string, localVarOptionals map[string]interface{}) (AppleMdmSignedCsrPlist,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdmSignedCsrPlist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/csr"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/octet-stream",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Delete an Apple MDM
 Removes an Apple MDM configuration.  Warning: This is a destructive operation and will remove your Apple Push Certificates.  We will no longer be able to manage your devices and the only recovery option is to re-register all devices into MDM.  #### Sample Request &#x60;&#x60;&#x60; curl -X DELETE https://console.jumpcloud.com/api/v2/applemdms/{id} \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdm*/
func (a *AppleMDMApiService) ApplemdmsDelete(ctx context.Context, id string, localVarOptionals map[string]interface{}) (AppleMdm,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdm
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Remove an Apple MDM Device&#39;s Enrollment
 Remove a single Apple MDM device from MDM enrollment.  #### Sample Request &#x60;&#x60;&#x60;   curl -X DELETE https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id} \\   -H &#39;accept: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdmDevice*/
func (a *AppleMDMApiService) ApplemdmsDeletedevice(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) (AppleMdmDevice,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdmDevice
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Get Apple MDM DEP Public Key
 Retrieves an Apple MDM DEP Public Key.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdmPublicKeyCert*/
func (a *AppleMDMApiService) ApplemdmsDepkeyget(ctx context.Context, appleMdmId string, localVarOptionals map[string]interface{}) (AppleMdmPublicKeyCert,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdmPublicKeyCert
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/depkey"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/x-pem-file",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Clears the Activation Lock for a Device
 Clears the activation lock on the specified device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesClearActivationLock(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Request the status of an OS update for a device
 Pass through to request the status of an OS update #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/osUpdateStatus \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesOSUpdateStatus(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/osUpdateStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Refresh activation lock information for a device
 Refreshes the activation lock information for a device  #### Sample Request  &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesRefreshActivationLockInformation(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Schedule an OS update for a device
 Schedules an OS update for a device  #### Sample Request  &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/scheduleOSUpdate \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{\&quot;install_action\&quot;: \&quot;INSTALL_ASAP\&quot;, \&quot;product_key\&quot;: \&quot;key\&quot;}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "body" (ScheduleOsUpdate) 
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesScheduleOSUpdate(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/scheduleOSUpdate"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(ScheduleOsUpdate); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Erase Device
 Erases a DEP-enrolled device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/erase \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "body" (Body) 
 @return */
func (a *AppleMDMApiService) ApplemdmsDeviceserase(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/erase"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Body); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService List AppleMDM Devices
 Lists all Apple MDM devices.  The filter and sort queries will allow the following fields: &#x60;createdAt&#x60; &#x60;depRegistered&#x60; &#x60;enrolled&#x60; &#x60;id&#x60; &#x60;osVersion&#x60; &#x60;serialNumber&#x60; &#x60;udid&#x60;  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices \\   -H &#39;accept: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xTotalCount" (int32) 
 @return []AppleMdmDevice*/
func (a *AppleMDMApiService) ApplemdmsDeviceslist(ctx context.Context, appleMdmId string, localVarOptionals map[string]interface{}) ([]AppleMdmDevice,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []AppleMdmDevice
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xTotalCount"], "int32", "xTotalCount"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sort"].([]string); localVarOk {
		localVarQueryParams.Add("sort", parameterToString(localVarTempParam, "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["xTotalCount"].(int32); localVarOk {
		localVarHeaderParams["x-total-count"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Lock Device
 Locks a DEP-enrolled device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/lock \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "body" (Body1) 
 @return */
func (a *AppleMDMApiService) ApplemdmsDeviceslock(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Body1); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Restart Device
 Restarts a DEP-enrolled device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/restart \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{\&quot;kextPaths\&quot;: [\&quot;Path1\&quot;, \&quot;Path2\&quot;]}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "body" (Body2) 
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesrestart(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Body2); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Shut Down Device
 Shuts down a DEP-enrolled device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/shutdown \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *AppleMDMApiService) ApplemdmsDevicesshutdown(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}/shutdown"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* AppleMDMApiService Get an Apple MDM Enrollment Profile
 Get an enrollment profile  Currently only requesting the mobileconfig is supported.  #### Sample Request  &#x60;&#x60;&#x60; curl https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/enrollmentprofiles/{ID} \\   -H &#39;accept: application/x-apple-aspen-config&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param id 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return Mobileconfig*/
func (a *AppleMDMApiService) ApplemdmsEnrollmentprofilesget(ctx context.Context, appleMdmId string, id string, localVarOptionals map[string]interface{}) (Mobileconfig,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mobileconfig
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/enrollmentprofiles/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/x-apple-aspen-config",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService List Apple MDM Enrollment Profiles
 Get a list of enrollment profiles for an apple mdm.  Note: currently only one enrollment profile is supported.  #### Sample Request &#x60;&#x60;&#x60;  curl https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/enrollmentprofiles \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []AppleMdm*/
func (a *AppleMDMApiService) ApplemdmsEnrollmentprofileslist(ctx context.Context, appleMdmId string, localVarOptionals map[string]interface{}) ([]AppleMdm,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []AppleMdm
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/enrollmentprofiles"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Details of an AppleMDM Device
 Gets a single Apple MDM device.  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id} \\   -H &#39;accept: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param deviceId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdmDevice*/
func (a *AppleMDMApiService) ApplemdmsGetdevice(ctx context.Context, appleMdmId string, deviceId string, localVarOptionals map[string]interface{}) (AppleMdmDevice,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdmDevice
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/devices/{device_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", fmt.Sprintf("%v", deviceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService List Apple MDMs
 Get a list of all Apple MDM configurations.  An empty topic indicates that a signed certificate from Apple has not been provided to the PUT endpoint yet.  Note: currently only one MDM configuration per organization is supported.  #### Sample Request &#x60;&#x60;&#x60; curl https://console.jumpcloud.com/api/v2/applemdms \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "limit" (int32) 
     @param "skip" (int32) The offset into the records to return.
 @return []AppleMdm*/
func (a *AppleMDMApiService) ApplemdmsList(ctx context.Context, localVarOptionals map[string]interface{}) ([]AppleMdm,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []AppleMdm
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Update an Apple MDM
 Updates an Apple MDM configuration.  This endpoint is used to supply JumpCloud with a signed certificate from Apple in order to finalize the setup and allow JumpCloud to manage your devices.  It may also be used to update the DEP Settings.  #### Sample Request &#x60;&#x60;&#x60;   curl -X PUT https://console.jumpcloud.com/api/v2/applemdms/{ID} \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;name\&quot;: \&quot;MDM name\&quot;,     \&quot;appleSignedCert\&quot;: \&quot;{CERTIFICATE}\&quot;,     \&quot;encryptedDepServerToken\&quot;: \&quot;{SERVER_TOKEN}\&quot;,     \&quot;dep\&quot;: {       \&quot;welcomeScreen\&quot;: {         \&quot;title\&quot;: \&quot;Welcome\&quot;,         \&quot;paragraph\&quot;: \&quot;In just a few steps, you will be working securely from your Mac.\&quot;,         \&quot;button\&quot;: \&quot;continue\&quot;,       },     },   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (AppleMdmPatch) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return AppleMdm*/
func (a *AppleMDMApiService) ApplemdmsPut(ctx context.Context, id string, localVarOptionals map[string]interface{}) (AppleMdm,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AppleMdm
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(AppleMdmPatch); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* AppleMDMApiService Refresh DEP Devices
 Refreshes the list of devices that a JumpCloud admin has added to their virtual MDM in Apple Business Manager - ABM so that they can be DEP enrolled with JumpCloud.  #### Sample Request &#x60;&#x60;&#x60;   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/refreshdepdevices \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param appleMdmId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *AppleMDMApiService) ApplemdmsRefreshdepdevices(ctx context.Context, appleMdmId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applemdms/{apple_mdm_id}/refreshdepdevices"
	localVarPath = strings.Replace(localVarPath, "{"+"apple_mdm_id"+"}", fmt.Sprintf("%v", appleMdmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

