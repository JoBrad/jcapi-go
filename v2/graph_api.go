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

type GraphApiService service


/* GraphApiService List the associations of an Active Directory instance
 This endpoint returns the direct associations of this Active Directory instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET &#39;https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/associations?targets&#x3D;user \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param activedirectoryId 
 @param targets Targets which a \&quot;active_directory\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphActiveDirectoryAssociationsList(ctx context.Context, activedirectoryId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/activedirectories/{activedirectory_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", fmt.Sprintf("%v", activedirectoryId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of an Active Directory instance
 This endpoint allows you to manage the _direct_ associations of an Active Directory instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/activedirectories/{AD_Instance_ID}/associations \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user\&quot;,     \&quot;id\&quot;: \&quot;{User_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param activedirectoryId 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationActiveDirectory) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphActiveDirectoryAssociationsPost(ctx context.Context, activedirectoryId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/activedirectories/{activedirectory_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", fmt.Sprintf("%v", activedirectoryId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationActiveDirectory); localVarOk {
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

/* GraphApiService List the Users bound to an Active Directory instance
 This endpoint will return all Users bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Active Directory instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Active Directory instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/users \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param activedirectoryId ObjectID of the Active Directory instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphActiveDirectoryTraverseUser(ctx context.Context, activedirectoryId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/activedirectories/{activedirectory_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", fmt.Sprintf("%v", activedirectoryId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to an Active Directory instance
 This endpoint will return all Users Groups bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Active Directory instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Active Directory instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/usergroups \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param activedirectoryId ObjectID of the Active Directory instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphActiveDirectoryTraverseUserGroup(ctx context.Context, activedirectoryId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/activedirectories/{activedirectory_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", fmt.Sprintf("%v", activedirectoryId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of an Application
 This endpoint will return the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Applications and User Groups.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET &#39;https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations?targets&#x3D;user_group \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param applicationId ObjectID of the Application.
 @param targets Targets which a \&quot;application\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphApplicationAssociationsList(ctx context.Context, applicationId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of an Application
 This endpoint allows you to manage the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Application and User Groups.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST &#39;https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations&#39; \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user_group\&quot;,     \&quot;id\&quot;: \&quot;{Group_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param applicationId ObjectID of the Application.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationApplication) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphApplicationAssociationsPost(ctx context.Context, applicationId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationApplication); localVarOk {
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

/* GraphApiService List the Users bound to an Application
 This endpoint will return all Users bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Application to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Application.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/users \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param applicationId ObjectID of the Application.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphApplicationTraverseUser(ctx context.Context, applicationId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to an Application
 This endpoint will return all Users Groups bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates  each path from this Application to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Application.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/usergroups \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param applicationId ObjectID of the Application.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphApplicationTraverseUserGroup(ctx context.Context, applicationId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a Command
 This endpoint will return the _direct_ associations of this Command.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Commands and User Groups.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/associations?targets&#x3D;system_group \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param commandId ObjectID of the Command.
 @param targets Targets which a \&quot;command\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphCommandAssociationsList(ctx context.Context, commandId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/commands/{command_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", fmt.Sprintf("%v", commandId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a Command
 This endpoint will allow you to manage the _direct_ associations of this Command.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Commands and User Groups.   #### Sample Request &#x60;&#x60;&#x60;  curl -X POST https://console.jumpcloud.com/api/v2/commands/{Command_ID}/associations \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system_group\&quot;,     \&quot;id\&quot;: \&quot;Group_ID\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param commandId ObjectID of the Command.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationCommand) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphCommandAssociationsPost(ctx context.Context, commandId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/commands/{command_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", fmt.Sprintf("%v", commandId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationCommand); localVarOk {
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

/* GraphApiService List the Systems bound to a Command
 This endpoint will return all Systems bound to a Command, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Command to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Command.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/systems \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param commandId ObjectID of the Command.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphCommandTraverseSystem(ctx context.Context, commandId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/commands/{command_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", fmt.Sprintf("%v", commandId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to a Command
 This endpoint will return all System Groups bound to a Command, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Command to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Command.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/systemgroups \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param commandId ObjectID of the Command.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphCommandTraverseSystemGroup(ctx context.Context, commandId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/commands/{command_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", fmt.Sprintf("%v", commandId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a G Suite instance
 This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET &#39;https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations?targets&#x3D;user_group \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param gsuiteId ObjectID of the G Suite instance.
 @param targets Targets which a \&quot;g_suite\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphGSuiteAssociationsList(ctx context.Context, gsuiteId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gsuites/{gsuite_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", fmt.Sprintf("%v", gsuiteId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a G Suite instance
 This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user_group\&quot;,     \&quot;id\&quot;: \&quot;{Group_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param gsuiteId ObjectID of the G Suite instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationGSuite) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphGSuiteAssociationsPost(ctx context.Context, gsuiteId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gsuites/{gsuite_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", fmt.Sprintf("%v", gsuiteId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationGSuite); localVarOk {
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

/* GraphApiService List the Users bound to a G Suite instance
 This endpoint will return all Users bound to a G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this G Suite instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this G Suite instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/users \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param gsuiteId ObjectID of the G Suite instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphGSuiteTraverseUser(ctx context.Context, gsuiteId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gsuites/{gsuite_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", fmt.Sprintf("%v", gsuiteId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to a G Suite instance
 This endpoint will return all User Groups bound to an G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this G Suite instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this G Suite instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{GSuite_ID}/usergroups \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param gsuiteId ObjectID of the G Suite instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphGSuiteTraverseUserGroup(ctx context.Context, gsuiteId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gsuites/{gsuite_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", fmt.Sprintf("%v", gsuiteId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a LDAP Server
 This endpoint returns the _direct_ associations of this LDAP Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.  #### Sample Request  &#x60;&#x60;&#x60;  curl -X GET &#39;https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations?targets&#x3D;user_group \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param ldapserverId ObjectID of the LDAP Server.
 @param targets Targets which a \&quot;ldap_server\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphLdapServerAssociationsList(ctx context.Context, ldapserverId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldapservers/{ldapserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a LDAP Server
 This endpoint allows you to manage the _direct_ associations of a LDAP Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user\&quot;,     \&quot;id\&quot;: \&quot;{User_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param ldapserverId ObjectID of the LDAP Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationLdapServer) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphLdapServerAssociationsPost(ctx context.Context, ldapserverId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldapservers/{ldapserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationLdapServer); localVarOk {
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

/* GraphApiService List the Users bound to a LDAP Server
 This endpoint will return all Users bound to an LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this LDAP server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this LDAP server instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param ldapserverId ObjectID of the LDAP Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphLdapServerTraverseUser(ctx context.Context, ldapserverId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldapservers/{ldapserver_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to a LDAP Server
 This endpoint will return all Users Groups bound to a LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this LDAP server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this LDAP server instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/usergroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param ldapserverId ObjectID of the LDAP Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphLdapServerTraverseUserGroup(ctx context.Context, ldapserverId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ldapservers/{ldapserver_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of an Office 365 instance
 This endpoint returns _direct_ associations of an Office 365 instance.   A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET &#39;https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations?targets&#x3D;user_group&#39; \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param office365Id ObjectID of the Office 365 instance.
 @param targets Targets which a \&quot;office_365\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphOffice365AssociationsList(ctx context.Context, office365Id string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/office365s/{office365_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", fmt.Sprintf("%v", office365Id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of an Office 365 instance
 This endpoint allows you to manage the _direct_ associations of a Office 365 instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user_group\&quot;,     \&quot;id\&quot;: \&quot;{Group_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param office365Id ObjectID of the Office 365 instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationOffice365) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphOffice365AssociationsPost(ctx context.Context, office365Id string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/office365s/{office365_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", fmt.Sprintf("%v", office365Id), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationOffice365); localVarOk {
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

/* GraphApiService List the Users bound to an Office 365 instance
 This endpoint will return all Users bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Office 365 instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Office 365 instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/users \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param office365Id ObjectID of the Office 365 suite.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphOffice365TraverseUser(ctx context.Context, office365Id string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/office365s/{office365_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", fmt.Sprintf("%v", office365Id), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to an Office 365 instance
 This endpoint will return all Users Groups bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Office 365 instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Office 365 instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60;   curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID/usergroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param office365Id ObjectID of the Office 365 suite.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphOffice365TraverseUserGroup(ctx context.Context, office365Id string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/office365s/{office365_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", fmt.Sprintf("%v", office365Id), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a Policy
 This endpoint returns the _direct_ associations of a Policy.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policies and Systems.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET &#39;https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/associations?targets&#x3D;system_group \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policyId ObjectID of the Policy.
 @param targets Targets which a \&quot;policy\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphPolicyAssociationsList(ctx context.Context, policyId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policies/{policy_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", fmt.Sprintf("%v", policyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a Policy
 This endpoint allows you to manage the _direct_ associations of a Policy.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policies and Systems.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/associations/ \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system_group\&quot;,     \&quot;id\&quot;: \&quot;{Group_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policyId ObjectID of the Policy.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationPolicy) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphPolicyAssociationsPost(ctx context.Context, policyId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policies/{policy_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", fmt.Sprintf("%v", policyId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationPolicy); localVarOk {
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

/* GraphApiService List the associations of a Policy Group.
 This endpoint returns the _direct_ associations of this Policy Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations?targets&#x3D;system \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param targets Targets which a \&quot;policy_group\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphPolicyGroupAssociationsList(ctx context.Context, groupId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a Policy Group
 This endpoint manages the _direct_ associations of this Policy Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system\&quot;,     \&quot;id\&quot;: \&quot;{SystemID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationPolicyGroup) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphPolicyGroupAssociationsPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationPolicyGroup); localVarOk {
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

/* GraphApiService List the members of a Policy Group
 This endpoint returns the Policy members of a Policy Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphPolicyGroupMembersList(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
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

/* GraphApiService Manage the members of a Policy Group
 This endpoint allows you to manage the Policy members of a Policy Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;policy\&quot;,     \&quot;id\&quot;: \&quot;{Policy_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationPolicyGroupMember) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphPolicyGroupMembersPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationPolicyGroupMember); localVarOk {
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

/* GraphApiService List the Policy Group&#39;s membership
 This endpoint returns all Policy members that are a member of this Policy Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/membership \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyGroupMembership(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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

/* GraphApiService List the Systems bound to a Policy Group
 This endpoint will return all Systems bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Policy Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyGroupTraverseSystem(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to Policy Groups
 This endpoint will return all System Groups bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Policy Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systemgroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the Policy Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyGroupTraverseSystemGroup(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policygroups/{group_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the parent Groups of a Policy
 This endpoint returns all the Policy Groups a Policy is a member of.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/memberof \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policyId ObjectID of the Policy.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyMemberOf(ctx context.Context, policyId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policies/{policy_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", fmt.Sprintf("%v", policyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService List the Systems bound to a Policy
 This endpoint will return all Systems bound to a Policy, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Policy to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policyId ObjectID of the Command.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyTraverseSystem(ctx context.Context, policyId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policies/{policy_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", fmt.Sprintf("%v", policyId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to a Policy
 This endpoint will return all Systems Groups bound to a Policy, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Policy to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET  https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/systemgroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param policyId ObjectID of the Command.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphPolicyTraverseSystemGroup(ctx context.Context, policyId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/policies/{policy_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", fmt.Sprintf("%v", policyId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a RADIUS  Server
 This endpoint returns the _direct_ associations of a Radius Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations?targets&#x3D;user_group \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param radiusserverId ObjectID of the Radius Server.
 @param targets Targets which a \&quot;radius_server\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphRadiusServerAssociationsList(ctx context.Context, radiusserverId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/radiusservers/{radiusserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", fmt.Sprintf("%v", radiusserverId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a RADIUS Server
 This endpoint allows you to manage the _direct_ associations of a Radius Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;type\&quot;:\&quot;user\&quot;,  \&quot;id\&quot;:\&quot;{USER_ID}\&quot;,  \&quot;op\&quot;:\&quot;add\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param radiusserverId ObjectID of the Radius Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationRadiusServer) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphRadiusServerAssociationsPost(ctx context.Context, radiusserverId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/radiusservers/{radiusserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", fmt.Sprintf("%v", radiusserverId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationRadiusServer); localVarOk {
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

/* GraphApiService List the Users bound to a RADIUS  Server
 This endpoint will return all Users bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this RADIUS server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this RADIUS server instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;   &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param radiusserverId ObjectID of the Radius Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphRadiusServerTraverseUser(ctx context.Context, radiusserverId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/radiusservers/{radiusserver_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", fmt.Sprintf("%v", radiusserverId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to a RADIUS  Server
 This endpoint will return all Users Groups bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this RADIUS server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this RADIUS server instance.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/usergroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param radiusserverId ObjectID of the Radius Server.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphRadiusServerTraverseUserGroup(ctx context.Context, radiusserverId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/radiusservers/{radiusserver_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", fmt.Sprintf("%v", radiusserverId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a Software Application
 This endpoint will return the _direct_ associations of a Software Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Software Application and System Groups.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/associations?targets&#x3D;system_group \\   -H &#39;accept: application/json&#39; \\   -H &#39;content-type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param softwareAppId ObjectID of the Software App.
 @param targets Targets which a \&quot;software_app\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphSoftwareappsAssociationsList(ctx context.Context, softwareAppId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/softwareapps/{software_app_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", fmt.Sprintf("%v", softwareAppId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a software application.
 This endpoint allows you to associate or disassociate a software application to a system or system group.  #### Sample Request &#x60;&#x60;&#x60; $ curl -X POST https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/associations \\ -H &#39;Accept: application/json&#39; \\ -H &#39;Content-Type: application/json&#39; \\ -H &#39;x-api-key: {API_KEY}&#39; \\ -d &#39;{   \&quot;id\&quot;: \&quot;&lt;object_id&gt;\&quot;,   \&quot;op\&quot;: \&quot;add\&quot;,   \&quot;type\&quot;: \&quot;system\&quot;  }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param softwareAppId ObjectID of the Software App.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationSoftwareApp) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphSoftwareappsAssociationsPost(ctx context.Context, softwareAppId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/softwareapps/{software_app_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", fmt.Sprintf("%v", softwareAppId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationSoftwareApp); localVarOk {
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

/* GraphApiService List the Systems bound to a Software App.
 This endpoint will return all Systems bound to a Software App, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Software App to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Software App.  See &#x60;/associations&#x60; endpoint to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param softwareAppId ObjectID of the Software App.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSoftwareappsTraverseSystem(ctx context.Context, softwareAppId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/softwareapps/{software_app_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", fmt.Sprintf("%v", softwareAppId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to a Software App.
 This endpoint will return all Systems Groups bound to a Software App, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this Software App to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Software App.  See &#x60;/associations&#x60; endpoint to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET  https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/systemgroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param softwareAppId ObjectID of the Software App.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSoftwareappsTraverseSystemGroup(ctx context.Context, softwareAppId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/softwareapps/{software_app_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", fmt.Sprintf("%v", softwareAppId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the associations of a System
 This endpoint returns the _direct_ associations of a System.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations?targets&#x3D;user \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param targets Targets which a \&quot;system\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphSystemAssociationsList(ctx context.Context, systemId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService Manage associations of a System
 This endpoint allows you to manage the _direct_ associations of a System.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;attributes\&quot;: {       \&quot;sudo\&quot;: {         \&quot;enabled\&quot;: true,         \&quot;withoutPassword\&quot;: false       }     },     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user\&quot;,     \&quot;id\&quot;: \&quot;UserID\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationSystem) 
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphSystemAssociationsPost(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
		return nil, err
	}
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationSystem); localVarOk {
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

/* GraphApiService List the associations of a System Group
 This endpoint returns the _direct_ associations of a System Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example System Groups and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations?targets&#x3D;user \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param targets Targets which a \&quot;system_group\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphSystemGroupAssociationsList(ctx context.Context, groupId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a System Group
 This endpoint allows you to manage the _direct_ associations of a System Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example System Groups and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user\&quot;,     \&quot;id\&quot;: \&quot;{UserID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationSystemGroup) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphSystemGroupAssociationsPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationSystemGroup); localVarOk {
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

/* GraphApiService List the members of a System Group
 This endpoint returns the system members of a System Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphSystemGroupMembersList(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
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

/* GraphApiService Manage the members of a System Group
 This endpoint allows you to manage the system members of a System Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system\&quot;,     \&quot;id\&quot;: \&quot;{System_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationSystemGroupMember) 
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphSystemGroupMembersPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
		return nil, err
	}
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationSystemGroupMember); localVarOk {
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

/* GraphApiService List the System Group&#39;s membership
 This endpoint returns all Systems that are a member of this System Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID/membership \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupMembership(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sort"].([]string); localVarOk {
		localVarQueryParams.Add("sort", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Commands bound to a System Group
 This endpoint will return all Commands bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group&#39;s type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System Group to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/commands \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "details" (string) This will provide detail descriptive response for the request.
 @return []CommandsGraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupTraverseCommand(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]CommandsGraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []CommandsGraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if err := typeCheckParameter(localVarOptionals["details"], "string", "details"); err != nil {
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
	if localVarTempParam, localVarOk := localVarOptionals["details"].(string); localVarOk {
		localVarQueryParams.Add("details", parameterToString(localVarTempParam, ""))
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

/* GraphApiService List the Policies bound to a System Group
 This endpoint will return all Policies bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System Group to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  This endpoint is not public yet as we haven&#39;t finished the code.  ##### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/policies \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupTraversePolicy(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Policy Groups bound to a System Group
 This endpoint will return all Policy Groups bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System Group to the corresponding Policy Group; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy Group from this System Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/policygroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupTraversePolicyGroup(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/policygroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Users bound to a System Group
 This endpoint will return all Users bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System Group to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/users \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupTraverseUser(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the User Groups bound to a System Group
 This endpoint will return all User Groups bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System Group to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/usergroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the System Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemGroupTraverseUserGroup(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systemgroups/{group_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the parent Groups of a System
 This endpoint returns all the System Groups a System is a member of.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/memberof \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemMemberOf(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService List the Commands bound to a System
 This endpoint will return all Commands bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/commands \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "details" (string) This will provide detail descriptive response for the request.
 @return []CommandsGraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemTraverseCommand(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]CommandsGraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []CommandsGraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

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
	if err := typeCheckParameter(localVarOptionals["details"], "string", "details"); err != nil {
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
	if localVarTempParam, localVarOk := localVarOptionals["details"].(string); localVarOk {
		localVarQueryParams.Add("details", parameterToString(localVarTempParam, ""))
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

/* GraphApiService List the Policies bound to a System
 This endpoint will return all Policies bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  This endpoint is not yet public as we have finish the code.  ##### Sample Request  &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/{System_ID}/policies \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemTraversePolicy(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Policy Groups bound to a System
 This endpoint will return all Policy Groups bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System to the corresponding Policy Group; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy Group from this System.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/policygroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemTraversePolicyGroup(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/policygroups"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

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
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService List the Users bound to a System
 This endpoint will return all Users bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/users \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemTraverseUser(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

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
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService List the User Groups bound to a System
 This endpoint will return all User Groups bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this System to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/usergroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "date" (string) Current date header for the System Context API
     @param "authorization" (string) Authorization header for the System Context API
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphSystemTraverseUserGroup(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

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
	if err := typeCheckParameter(localVarOptionals["date"], "string", "date"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["authorization"], "string", "authorization"); err != nil {
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
	if localVarTempParam, localVarOk := localVarOptionals["date"].(string); localVarOk {
		localVarHeaderParams["Date"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["authorization"].(string); localVarOk {
		localVarHeaderParams["Authorization"] = parameterToString(localVarTempParam, "")
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

/* GraphApiService List the associations of a User
 This endpoint returns the _direct_ associations of a User.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Users and Systems.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/associations?targets&#x3D;system_group \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param targets Targets which a \&quot;user\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphUserAssociationsList(ctx context.Context, userId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a User
 This endpoint allows you to manage the _direct_ associations of a User.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Users and Systems.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/users/{UserID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;attributes\&quot;: {       \&quot;sudo\&quot;: {       \&quot;enabled\&quot;: true,         \&quot;withoutPassword\&quot;: false       }     },     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system_group\&quot;,     \&quot;id\&quot;: \&quot;{GroupID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationUser) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphUserAssociationsPost(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationUser); localVarOk {
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

/* GraphApiService List the associations of a User Group.
 This endpoint returns the _direct_ associations of this User Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations?targets&#x3D;system \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param targets Targets which a \&quot;user_group\&quot; can be associated to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphUserGroupAssociationsList(ctx context.Context, groupId string, targets []string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("targets", parameterToString(targets, "csv"))
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

/* GraphApiService Manage the associations of a User Group
 This endpoint manages the _direct_ associations of this User Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;system\&quot;,     \&quot;id\&quot;: \&quot;{SystemID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationUserGroup) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphUserGroupAssociationsPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationUserGroup); localVarOk {
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

/* GraphApiService List the members of a User Group
 This endpoint returns the user members of a User Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphConnection*/
func (a *GraphApiService) GraphUserGroupMembersList(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphConnection,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphConnection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
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

/* GraphApiService Manage the members of a User Group
 This endpoint allows you to manage the user members of a User Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/members \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{     \&quot;op\&quot;: \&quot;add\&quot;,     \&quot;type\&quot;: \&quot;user\&quot;,     \&quot;id\&quot;: \&quot;{User_ID}\&quot;   }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (GraphOperationUserGroupMember) 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return */
func (a *GraphApiService) GraphUserGroupMembersPost(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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
	if localVarTempParam, localVarOk := localVarOptionals["body"].(GraphOperationUserGroupMember); localVarOk {
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

/* GraphApiService List the User Group&#39;s membership
 This endpoint returns all users members that are a member of this User Group.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/membership \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupMembership(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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

/* GraphApiService List the Active Directories bound to a User Group
 This endpoint will return all Active Directory Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding Active Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Active Directory from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/activedirectories \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseActiveDirectory(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/activedirectories"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Applications bound to a User Group
 This endpoint will return all Applications bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/applications \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseApplication(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Directories bound to a User Group
 This endpoint will return all Directories bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directories from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/directories \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseDirectory(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/directories"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the G Suite instances bound to a User Group
 This endpoint will return all G Suite Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID/gsuites \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseGSuite(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/gsuites"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the LDAP Servers bound to a User Group
 This endpoint will return all LDAP Servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/ldapservers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseLdapServer(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/ldapservers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Office 365 instances bound to a User Group
 This endpoint will return all Office 365 instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/office365s \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseOffice365(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/office365s"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the RADIUS Servers bound to a User Group
 This endpoint will return all RADIUS servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/radiusservers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseRadiusServer(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/radiusservers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Systems bound to a User Group
 This endpoint will return all Systems bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseSystem(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to User Groups
 This endpoint will return all System Groups bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this User Group.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systemgroups \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param groupId ObjectID of the User Group.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserGroupTraverseSystemGroup(ctx context.Context, groupId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/usergroups/{group_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", fmt.Sprintf("%v", groupId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the parent Groups of a User
 This endpoint returns all the User Groups a User is a member of.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/memberof \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserMemberOf(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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

/* GraphApiService List the Active Directory instances bound to a User
 This endpoint will return all Active Directory Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Active Directory instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Active Directory instance from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/activedirectories \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseActiveDirectory(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/activedirectories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Applications bound to a User
 This endpoint will return all Applications bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/applications \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseApplication(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Directories bound to a User
 This endpoint will return all Directories bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directory from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/directories \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseDirectory(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/directories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the G Suite instances bound to a User
 This endpoint will return all G-Suite Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/gsuites \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseGSuite(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/gsuites"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the LDAP servers bound to a User
 This endpoint will return all LDAP Servers bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/ldapservers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseLdapServer(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/ldapservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Office 365 instances bound to a User
 This endpoint will return all Office 365 Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/office365s \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseOffice365(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/office365s"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the RADIUS Servers bound to a User
 This endpoint will return all RADIUS Servers bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/radiusservers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseRadiusServer(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/radiusservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the Systems bound to a User
 This endpoint will return all Systems bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/systems\\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseSystem(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the System Groups bound to a User
 This endpoint will return all System Groups bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of compiled graph attributes for all paths followed.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/systemgroups\\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param userId ObjectID of the User.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
     @param "skip" (int32) The offset into the records to return.
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
 @return []GraphObjectWithPaths*/
func (a *GraphApiService) GraphUserTraverseSystemGroup(ctx context.Context, userId string, localVarOptionals map[string]interface{}) ([]GraphObjectWithPaths,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GraphObjectWithPaths
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{user_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
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

/* GraphApiService List the policy statuses for a system
 This endpoint returns the policy results for a particular system.  ##### Sample Request  &#x60;&#x60;&#x60; curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/policystatuses \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39;  &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param systemId ObjectID of the System.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "fields" ([]string) The comma separated fields included in the returned records. If omitted, the default list of fields will be returned. 
     @param "filter" ([]string) A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "sort" ([]string) The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
     @param "xOrgId" (string) Organization identifier that can be obtained from console settings.
 @return []PolicyResult*/
func (a *GraphApiService) PolicystatusesSystemsList(ctx context.Context, systemId string, localVarOptionals map[string]interface{}) ([]PolicyResult,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []PolicyResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/systems/{system_id}/policystatuses"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", fmt.Sprintf("%v", systemId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["fields"].([]string); localVarOk {
		localVarQueryParams.Add("fields", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].([]string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, "csv"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
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

